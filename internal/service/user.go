package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hmdp/internal/assembler"
	"hmdp/internal/cache"
	"hmdp/internal/dto"
	"hmdp/internal/model"
	"hmdp/pkg/utils"
)

type IUserService interface {
	SendCode(ctx *gin.Context, phone string) (any, error) // 发送验证码
	LoginByCode(ctx *gin.Context, req *dto.LoginByCodeReq) (*dto.UserLoginByCodeRsp, error)
	Me(ctx *gin.Context) (*dto.UserMeRsp, error)
	Info(ctx *gin.Context, req *dto.UserInfoReq) (*dto.UserInfoRsp, error)
	FindUserById(ctx *gin.Context, req *dto.FindUserByIdReq) (*dto.FindUserByIdRsp, error)
	FindUserByPhone(ctx *gin.Context, phone string) (u *model.User, err error)
	IsFollowed(ctx *gin.Context, req *dto.IsFollowedReq) (*dto.IsFollowedRsp, error)
}

type UserService struct {
	DB  *gorm.DB
	Req *assembler.UserReq
	Rsp *assembler.UserRsp
}

func NewUserService(db *gorm.DB) *UserService {
	return &UserService{
		DB:  db,
		Req: &assembler.UserReq{},
		Rsp: &assembler.UserRsp{},
	}
}

func (s *UserService) SendCode(ctx *gin.Context, phone string) (any, error) {
	// 校验手机号码
	if !utils.VerifyMobileFormat(phone) {
		return nil, fmt.Errorf("%v:手机号码格式错误", phone)
	}
	// 生成验证码
	code := utils.GenValidateCode(6)
	// 下发验证码
	if err := utils.SendCode(code); err != nil {
		return nil, err
	}
	// 调用redis保存验证码
	key := fmt.Sprintf("code:%v", phone)
	err := cache.RedisStore.Set(ctx, key, code, 0).Err()
	return nil, err
}

func (s *UserService) Info(ctx *gin.Context, req *dto.UserInfoReq) (*dto.UserInfoRsp, error) {
	var userinfo model.UserInfo
	err := s.DB.Preload("User").Where("user_id = ?", req.ID).First(&userinfo).Error
	if err == nil && userinfo.User != nil {
		return s.Rsp.E2DInfo(&userinfo), nil
	}
	return nil, err
}

func (s *UserService) LoginByCode(ctx *gin.Context, req *dto.LoginByCodeReq) (*dto.UserLoginByCodeRsp, error) {
	//校验手机号码
	if !utils.VerifyMobileFormat(req.Phone) {
		return nil, fmt.Errorf("%v:手机号码格式错误", req.Phone)
	}
	// 校验验证码是否一致
	originalCode, err := cache.RedisStore.Get(ctx, fmt.Sprintf("code:%v", req.Phone)).Result()
	if err != nil {

		return nil, err
	}

	if req.Code != originalCode {
		return nil, fmt.Errorf("%v:验证码错误", req.Phone)
	}
	// 校验通过，删除验证码
	err = cache.RedisStore.Del(ctx, fmt.Sprintf("code:%v", req.Phone)).Err()

	// 数据库查询
	var user model.User
	err = s.DB.Where(model.User{Phone: req.Phone}).FirstOrCreate(&user).Error
	if err != nil {
		return nil, err
	}
	// 将用户的信息存放到redis中
	token := user.GenToken()
	err = cache.SaveUser(ctx, token, &user)

	return s.Rsp.E2DLoginByCode(&user), err
}

func (s *UserService) Me(ctx *gin.Context) (*dto.UserMeRsp, error) {
	user := ctx.MustGet("user").(*model.User)
	return s.Rsp.E2DMe(user), nil
}
func (s *UserService) FollowUser(ctx *gin.Context, req *dto.FollowUserReq) (*dto.FollowUserRsp, error) {
	follow := model.Follow{
		UserId:   req.UserId,
		FollowId: req.ID,
	}
	// UserId 关注 FollowId
	key := model.FollowRedisKey(req.UserId) // redis的key
	//关注
	if *req.Follow {
		err := s.DB.Where("user_id = ? AND follow_id = ?", follow.UserId, follow.FollowId).FirstOrCreate(&follow).Error
		if err != nil {
			return nil, err
		}
		// 保存到redis的set中
		cache.RedisStore.SAdd(ctx, key, req.ID)
		return &dto.FollowUserRsp{}, nil
	}
	// 取消关注
	err := s.DB.Where("user_id = ? AND follow_id = ?", follow.UserId, follow.FollowId).Delete(&follow).Error
	cache.RedisStore.SRem(ctx, key, req.ID)
	// 从redis的set中删除
	if err != nil {
		return nil, err
	}
	return &dto.FollowUserRsp{}, nil
}
func (s *UserService) FindUserById(ctx *gin.Context, req *dto.FindUserByIdReq) (*dto.FindUserByIdRsp, error) {
	var user model.User
	err := s.DB.Where("id = ?", req.ID).First(&user).Error
	return s.Rsp.E2DFindUserById(&user), err
}

func (s *UserService) FindUserByPhone(ctx *gin.Context, phone string) (u *model.User, err error) {
	var user model.User
	err = s.DB.Where("phone = ?", phone).First(&user).Error
	return &user, err
}

func (s *UserService) IsFollowed(ctx *gin.Context, req *dto.IsFollowedReq) (*dto.IsFollowedRsp, error) {
	var follow model.Follow
	err := s.DB.Where("user_id = ? AND follow_id = ?", req.CurrentUserId, req.FollowUserId).First(&follow).Error
	if err == gorm.ErrRecordNotFound {
		return &dto.IsFollowedRsp{Followed: false}, nil
	}
	return &dto.IsFollowedRsp{Followed: true}, nil
}

func (s *UserService) CommonFollow(ctx *gin.Context, req *dto.CommonFollowReq) ([]*dto.CommonFollowRsp, error) {
	UserIds, err := cache.RedisStore.SInter(ctx, model.FollowRedisKey(req.UserId), model.FollowRedisKey(req.CurrentUserId)).Result()
	if err != nil {
		return nil, err
	}
	var users []*model.User
	err = s.DB.Select("id, nick_name, icon").Where("id in (?)", UserIds).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return s.Rsp.E2DCommonFollow(users), nil
}
