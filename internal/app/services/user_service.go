package services

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/assembler"
	"hmdp/internal/app/dto"
	"hmdp/internal/domain/aggregate"
	"hmdp/internal/domain/entity"
)

type IUserService interface {
	SendCode(ctx *gin.Context, req *dto.UserSendCodeReq) (*dto.UserSendCodeRsp, error)
	LoginByCode(ctx *gin.Context, req *dto.UserLoginByCodeReq) (*dto.UserLoginByCodeRsp, error)
	Info(ctx *gin.Context, req *dto.UserInfoReq) (*dto.UserInfoRsp, error)
	Me(ctx *gin.Context, req *dto.UserMeReq) (*dto.UserMeRsp, error)
}

type UserService struct {
	UserAgg aggregate.IUserAggregate //调用领域聚合层|领域服务层
	UserReq *assembler.UserReq       //处理请求
	UserRsp *assembler.UserRsp       //处理响应

}

func NewUserService(aggr aggregate.IUserAggregate) IUserService {
	return &UserService{
		aggr,
		&assembler.UserReq{},
		&assembler.UserRsp{},
	}
}

func (s *UserService) SendCode(ctx *gin.Context, req *dto.UserSendCodeReq) (*dto.UserSendCodeRsp, error) {
	// dto2entity
	user := s.UserReq.D2ESendCode(req)
	session := sessions.Default(ctx)
	err := s.UserAgg.SendCode(ctx, session, user)
	if err != nil {
		return nil, err
	}
	return s.UserRsp.E2DSendCode(), nil
}

func (s *UserService) LoginByCode(ctx *gin.Context, req *dto.UserLoginByCodeReq) (*dto.UserLoginByCodeRsp, error) {
	user := s.UserReq.D2ELoginByCode(req)
	session := sessions.Default(ctx)
	err := s.UserAgg.LoginByCode(ctx, session, user, req.Code)
	if err != nil {
		return nil, err
	}
	return s.UserRsp.E2DLoginByCode(user), nil
}

// Info 登录情况下才能访问
func (s *UserService) Info(ctx *gin.Context, req *dto.UserInfoReq) (*dto.UserInfoRsp, error) {
	if user, ok := ctx.Get("user"); ok {
		return s.UserRsp.E2DInfo(user.(*entity.User)), nil
	}
	return nil, fmt.Errorf("查看用户详细信息失败")
}

// Me 登录情况下才能访问,返回用户的基本信息
func (s *UserService) Me(ctx *gin.Context, req *dto.UserMeReq) (*dto.UserMeRsp, error) {
	if user, ok := ctx.Get("user"); ok {
		return s.UserRsp.E2DMe(user.(*entity.User)), nil
	}
	return nil, fmt.Errorf("查看用户详细信息失败")
}
