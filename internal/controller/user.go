package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/dto"
	"hmdp/internal/service"
	"hmdp/pkg/serializer"
	"net/http"
)

type UserHandler struct {
	svc *service.UserService
}

func NewUserHandler(svc *service.UserService) *UserHandler {
	return &UserHandler{svc}
}

func (c *UserHandler) UserSendCode(ctx *gin.Context) {
	req := dto.UserSendCodeReq{}
	// 参数解析
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
	}
	// 调用service层
	rsp, err := c.svc.SendCode(ctx, req.Phone)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
	return
}

func (c *UserHandler) UserInfo(ctx *gin.Context) {
	req := &dto.UserInfoReq{}
	// 参数解析
	if err := ctx.ShouldBindUri(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
	}
	info, err := c.svc.Info(ctx, req.ID)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(info))
}

func (c *UserHandler) UserLogin(ctx *gin.Context) {
	req := &dto.UserLoginByCodeReq{}
	// 参数绑定
	if err := ctx.ShouldBind(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	user, err := c.svc.LoginByCode(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(user))
	return
}

func (c *UserHandler) UserMe(ctx *gin.Context) {
	info, err := c.svc.Me(ctx)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(info))
	return
}
