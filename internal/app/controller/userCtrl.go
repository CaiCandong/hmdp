package controller

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/services"
	"hmdp/internal/app/vo"
	"hmdp/internal/domain/entity"
	"hmdp/internal/domain/valueobject"
	"hmdp/internal/infrastructure/mysql"
	"net/http"
)

type UserController interface {
	Register(ctx *gin.Context)
	Login(ctx *gin.Context)
	Me(ctx *gin.Context)
	SendCode(ctx *gin.Context)
	Logout(ctx *gin.Context)
	Update(ctx *gin.Context)
	Info(ctx *gin.Context)
}

func NewUserController() UserController {
	// 进行依赖注入
	userRepo := mysql.NewUserRepo()
	userService := services.NewUserService(services.WithUserRepo(userRepo))
	return &UserControllerImp{userService}
}

type UserControllerImp struct {
	userService services.IUserService
}

func (u *UserControllerImp) SendCode(ctx *gin.Context) {
	var req struct {
		Phone valueobject.Phone `json:"phone" form:"phone" binding:"required"`
	}
	// 参数解析
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	// 调用service层
	rep, err := u.userService.SendCode(ctx, req.Phone)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, rep)
		return
	}
	ctx.JSON(http.StatusOK, rep)
	return
}

func (u *UserControllerImp) Info(ctx *gin.Context) {
	info, err := u.userService.Info(ctx)
	if err != nil {
		info.Success = false
		ctx.JSON(http.StatusBadRequest, info)
		return
	}
	ctx.JSON(http.StatusOK, info)
}

func (u *UserControllerImp) Register(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControllerImp) Login(ctx *gin.Context) {
	var err error
	var req struct {
		Phone    valueobject.Phone `json:"phone" form:"phone" binding:"required"`
		Code     string            `json:"code" form:"code"`
		Password string            `json:"password" form:"code"`
	}
	// 参数绑定
	if err = ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	resp, err := u.userService.LoginByCode(ctx, req.Phone, req.Code)
	if err != nil {
		resp.Success = false
		ctx.JSON(http.StatusBadRequest, resp)
		return
	}
	ctx.JSON(http.StatusOK, resp)
	return
}

func (u *UserControllerImp) Me(ctx *gin.Context) {
	if user, ok := ctx.Get("user"); ok {
		if u, ok := user.(*entity.User); ok {
			ctx.JSON(http.StatusOK, vo.BuildUserResponse(*u))
			return
		}
	}
}

func (u *UserControllerImp) Logout(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControllerImp) Update(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}
