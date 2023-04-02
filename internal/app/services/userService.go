package services

import (
	"fmt"
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/vo"
	"hmdp/internal/domain/entity"
	"hmdp/internal/domain/repository"
	"hmdp/internal/domain/valueobject"
	"hmdp/pkg/serializer"
	"hmdp/pkg/utils"
	"net/http"
)

type IUserService interface {
	SendCode(ctx *gin.Context, phone valueobject.Phone) (res serializer.Response, err error)
	Info(ctx *gin.Context) (res serializer.Response, err error)
	LoginByCode(ctx *gin.Context, phone valueobject.Phone, Code string) (res serializer.Response, err error)
}

type UserService struct {
	userRepo repository.IUserRepo
}

type UserConfiguration func(os *UserService) error

func NewUserService(cfgs ...UserConfiguration) IUserService {
	// Create the user service
	os := &UserService{}
	// Apply all Configurations passed in
	for _, cfg := range cfgs {
		// Pass the service into the configuration function
		err := cfg(os)
		if err != nil {
			return nil
		}
	}
	return os
}

func WithUserRepo(userRepo repository.IUserRepo) UserConfiguration {
	return func(os *UserService) error {
		os.userRepo = userRepo
		return nil
	}
}

// func WithUserID

func (s *UserService) LoginByCode(ctx *gin.Context, phone valueobject.Phone, Code string) (res serializer.Response, err error) {

	// 校验手机号码
	if !phone.VerifyMobileFormat() {
		return serializer.Response{Success: false, ErrorMsg: "手机号码错误"}, nil
	}
	// 校验验证码是否一致
	session := sessions.Default(ctx)
	originalCode := session.Get("code")
	if code, ok := originalCode.(string); !ok || code != Code {
		return res, err
	}
	// 数据库查询
	var user entity.User
	if user, err = s.userRepo.GetUserOrCreate(string(phone)); err != nil {
		return res, err
	}
	session.Set("user_id", user.ID)
	err = session.Save()
	if err != nil {
		return res, err
	}
	res.Data = user
	return vo.BuildUserResponse(user), nil
}

func (s *UserService) SendCode(ctx *gin.Context, phone valueobject.Phone) (res serializer.Response, err error) {
	// 校验手机号码
	if !phone.VerifyMobileFormat() {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "手机号码错误"})
	}
	// 生成验证码
	code := utils.GenValidateCode(6)
	session := sessions.Default(ctx)
	session.Set("code", code)
	err = session.Save()
	// 下发验证码
	err = phone.SendCode(code)
	res.Success = true
	return res, err
}

func (s *UserService) Info(ctx *gin.Context) (res serializer.Response, err error) {
	// 查看用户详细信息
	if user, ok := ctx.Get("user"); ok {
		if u, ok := user.(*entity.User); ok {
			return vo.BuildUserResponse(*u), err
		}
	}
	return res, fmt.Errorf("查看用户详细信息失败")
}
