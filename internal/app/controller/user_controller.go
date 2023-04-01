package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"hmdp/internal/app/vo"
	"hmdp/internal/entity"
	"hmdp/internal/infrastructure/mysql"
	"hmdp/pkg/logger"
	"hmdp/pkg/serializer"
	"hmdp/pkg/utils"
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
	return &UserControllerImp{}
}

type UserControllerImp struct {
}

func (u *UserControllerImp) SendCode(ctx *gin.Context) {
	var req struct {
		Phone entity.Phone `json:"phone" form:"phone" binding:"required"`
	}
	// 参数绑定
	if err := ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	// 校验手机号码
	if !req.Phone.VerifyMobileFormat() {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "手机号码错误"})
	}
	// 下发验证码
	code := utils.GenValidateCode(6)
	session := sessions.Default(ctx)
	session.Set("code", code)
	err := session.Save()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"err": "手机号码错误"})
		return
	}
	logger.Logger.Info("code" + code)

	ctx.JSON(http.StatusOK, serializer.Response{
		Success:  true,
		ErrorMsg: "",
		Data:     code,
		Total:    0,
	})
}

func (u *UserControllerImp) Info(ctx *gin.Context) {
	// 查看用户详细信息
	if user, ok := ctx.Get("user"); ok {
		if u, ok := user.(*mysql.User); ok {
			ctx.JSON(http.StatusOK, vo.BuildUserResponse(u))
			return
		}
	}
}

func (u *UserControllerImp) Register(ctx *gin.Context) {
	//TODO implement me
	panic("implement me")
}

func (u *UserControllerImp) Login(ctx *gin.Context) {
	var err error
	var req struct {
		Phone    entity.Phone `json:"phone" form:"phone" binding:"required"`
		Code     string       `json:"code" form:"code"`
		Password string       `json:"password" form:"code"`
	}
	// 参数绑定
	if err = ctx.ShouldBind(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	// 校验手机号码
	if !req.Phone.VerifyMobileFormat() {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "手机号码错误"})
	}
	// 校验验证码是否一致
	session := sessions.Default(ctx)
	originalCode := session.Get("code")
	if code, ok := originalCode.(string); !ok || code != req.Code {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "验证码错误"})
	}
	// 数据库查询
	var user *mysql.User
	if user, err = mysql.GetUserOrCreate(string(req.Phone)); err != nil {
		ctx.JSON(http.StatusInternalServerError, serializer.Response{
			Success:  false,
			ErrorMsg: "数据库错误",
		})
	}
	session.Set("user_id", user.ID)
	err = session.Save()
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"err": "session报错失败"})
		return
	}
	ctx.JSON(http.StatusOK, vo.BuildUserResponse(user))
}

func (u *UserControllerImp) Me(ctx *gin.Context) {
	if user, ok := ctx.Get("user"); ok {
		if u, ok := user.(*mysql.User); ok {
			ctx.JSON(http.StatusOK, vo.BuildUserResponse(u))
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

//type LoginParam struct {
//	Phone    string `json:"phone" form:"phone" binding:"required"`
//	Code     string `json:"code" form:"code"`
//	Password string `json:"password" form:"code"`
//}
//
//func Login(c *gin.Context) {
//	params := LoginParam{}
//	if err := c.ShouldBind(&params); err != nil {
//		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
//		return
//	}
//	// 通过session校验验证码
//	session := sessions.Default(c)
//	uid := session.Get("code")
//	if uid == nil {
//		c.JSON(http.StatusBadRequest, serializer.Response{
//			Success:  false,
//			ErrorMsg: "验证码校验失败",
//		})
//		return
//	}
//	s := uid.(string)
//	if s != params.Code {
//		c.JSON(http.StatusBadRequest, serializer.Response{
//			Success:  false,
//			ErrorMsg: "验证码校验失败",
//		})
//		return
//	}
//	// 校验手机号码
//	if !utils.VerifyMobileFormat(params.Phone) {
//		c.JSON(http.StatusBadRequest, serializer.Response{
//			Success:  false,
//			ErrorMsg: "手机号校验失败",
//		})
//		return
//	}
//	// 数据库查询
//	user, _ := model.GetUserByPhone(params.Phone)
//	if user.Phone != "" { //用户不存在
//		session.Set("user_id", user.ID)
//		session.Save()
//		c.JSON(http.StatusOK, dto.BuildUserResponse(user))
//		return
//	}
//	user, err := model.CreateUserWithPhone(params.Phone)
//	if err != nil {
//		c.JSON(http.StatusOK, serializer.Response{
//			Success:  false,
//			ErrorMsg: "手机号校验失败",
//		})
//		return
//	}
//	session.Set("user_id", user.ID)
//	session.Save()
//	c.JSON(http.StatusOK, dto.BuildUserResponse(user))
//}

//func Me(c *gin.Context) {
//	if user, _ := c.Get("user"); user != nil {
//		if u, ok := user.(*mysql.User); ok {
//			c.JSON(http.StatusOK, dto.BuildUserResponse(*u))
//			return
//		}
//	}
//
//	return
//}
