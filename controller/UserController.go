package controller

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"hmdp/model"
	"hmdp/serializer"
	"hmdp/utils"
	"net/http"
)

type SendCodeParam struct {
	Phone string `json:"phone" form:"phone" binding:"required"`
}

func SendCode(c *gin.Context) {
	params := SendCodeParam{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	}
	// 校验手机号码
	if !utils.VerifyMobileFormat(params.Phone) {
		c.JSON(http.StatusBadRequest, gin.H{"err": "手机号码错误"})
	}
	// 生成随机的验证码
	code := utils.GenValidateCode(6)
	// 将验证码存放到session中
	session := sessions.Default(c)
	session.Set("code", code)
	session.Save()
	//session.
	// TODO:下发验证码
	utils.Logger.Info("code" + code)
	// 返回短信发送成功
	c.JSON(http.StatusOK, serializer.Response{
		Success:  true,
		ErrorMsg: "",
		Data:     code,
		Total:    0,
	})
}

type LoginParam struct {
	Phone    string `json:"phone" form:"phone" binding:"required"`
	Code     string `json:"code" form:"code"`
	Password string `json:"password" form:"code"`
}

func Login(c *gin.Context) {
	params := LoginParam{}
	if err := c.ShouldBind(&params); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
		return
	}
	// 通过session校验验证码
	session := sessions.Default(c)
	uid := session.Get("code")
	if uid == nil {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Success:  false,
			ErrorMsg: "验证码校验失败",
		})
		return
	}
	s := uid.(string)
	if s != params.Code {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Success:  false,
			ErrorMsg: "验证码校验失败",
		})
		return
	}
	// 校验手机号码
	if !utils.VerifyMobileFormat(params.Phone) {
		c.JSON(http.StatusBadRequest, serializer.Response{
			Success:  false,
			ErrorMsg: "手机号校验失败",
		})
		return
	}
	// 数据库查询
	user, _ := model.GetUserByPhone(params.Phone)
	if user.Phone == "" { //用户不存在
		user.Phone = params.Phone
		model.DB.Create(&user)
	}
	session.Set("user_id", user)
	session.Save()
	c.JSON(http.StatusBadRequest, serializer.Response{
		Success:  true,
		ErrorMsg: "登录成功",
	})
	return
}

func Me(c *gin.Context) {
	//params := LoginParam{}
	//if err := c.ShouldBind(params); err != nil {
	//	c.JSON(http.StatusBadRequest, gin.H{"err": err.Error()})
	//}
	////// 生成随机的验证码
	//GenValidateCode(6)
}
