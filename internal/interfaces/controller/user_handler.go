package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/dto"
	"hmdp/internal/app/services"
	"hmdp/pkg/serializer"
	"net/http"
)

type UserHandler struct {
	UserService services.IUserService
}

func NewUserHandler(UserService services.IUserService) *UserHandler {
	return &UserHandler{UserService}
}

func (u *UserHandler) SendCode(ctx *gin.Context) {
	req := dto.UserSendCodeReq{}
	// 参数解析
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
	}
	// 调用service层
	rsp, err := u.UserService.SendCode(ctx, &req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
	return
}

func (u *UserHandler) Info(ctx *gin.Context) {
	req := &dto.UserInfoReq{}
	// 参数解析
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
	}
	info, err := u.UserService.Info(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(info))
}

func (u *UserHandler) Login(ctx *gin.Context) {
	req := &dto.UserLoginByCodeReq{}
	// 参数绑定
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	user, err := u.UserService.LoginByCode(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(user))
	return
}

func (u *UserHandler) Me(ctx *gin.Context) {
	req := &dto.UserMeReq{}
	info, err := u.UserService.Me(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(info))
	return
}
