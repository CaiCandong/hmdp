package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"hmdp/internal/dto"
	"hmdp/internal/model"
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

func (c *UserHandler) IsFollowed(ctx *gin.Context) {
	req := &dto.IsFollowedReq{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
	}
	// 获取当前用户id
	user, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusBadRequest, fmt.Errorf("获取当前用户失败"))
		return
	}
	req.CurrentUserId = user.(*model.User).ID
	rsp, _ := c.svc.IsFollowed(ctx, req)
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
	return
}

func (c *UserHandler) CommonFollow(ctx *gin.Context) {

	// TODO 共同关注
}

// FollowUser 关注
func (c *UserHandler) FollowUser(ctx *gin.Context) {
	req := &dto.FollowUserReq{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
	}
	// 获取当前用户id
	user := ctx.MustGet("user")
	req.UserId = user.(*model.User).ID
	c.svc.FollowUser(ctx, req)
}

func (c *UserHandler) FindUserById(ctx *gin.Context) {
	req := &dto.FindUserByIdReq{}
	if err := ctx.ShouldBindUri(req); err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
	}
	rsp, err := c.svc.FindUserById(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(rsp))
	return
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
	u, ok := ctx.Get("user")
	if !ok {
		ctx.JSON(http.StatusBadRequest, serializer.ParamErr("未登录", nil))
	}
	req.ID = u.(*model.User).ID
	info, err := c.svc.Info(ctx, req)
	if err != nil {
		ctx.JSON(http.StatusOK, serializer.ParamErr("", err))
		return
	}
	ctx.JSON(http.StatusOK, serializer.Success(info))
}

func (c *UserHandler) UserLogin(ctx *gin.Context) {
	req := &dto.LoginByCodeReq{}
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
