package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"hmdp/internal/infrastructure/mysql"
)

// CurrentUser 获取登录用户
func CurrentUser() gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			user, err := mysql.GetUser(uid)
			if err == nil {
				c.Set("user", &user)
			}
		}
		c.Next()
	}
}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, ok := c.Get("user"); ok && user != nil {
			if _, ok := user.(*mysql.User); ok {
				c.Next()
				return
			}
		}

		//c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}
