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

func (s *UserService) Info(ctx *gin.Context, id uint) (interface{}, error) {
	return nil, nil
}

func (s *UserService) LoginByCode(ctx *gin.Context, req *dto.UserLoginByCodeReq) (*dto.UserLoginByCodeRsp, error) {
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

func (s *UserService) Me(ctx *gin.Context) (interface{}, error) {
	return nil, nil
}

//
//func (s *UserService) UpdateUser(ctx *gin.Context, id uint, data interface{}) (interface{}, error) {
//
//}

func (s *UserService) FindUserById(ctx *gin.Context, id uint) (*model.User, error) {
	var user model.User
	err := s.DB.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (s *UserService) FindUserByPhone(ctx *gin.Context, phone string) (u *model.User, err error) {
	var user model.User
	err = s.DB.Where("phone = ?", phone).First(&user).Error
	return &user, err
}
