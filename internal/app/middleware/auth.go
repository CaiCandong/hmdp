package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"hmdp/internal/domain/entity"
	"hmdp/internal/domain/repository"
)

// CurrentUser 获取登录用户
func CurrentUser(repo repository.IUserRepo) gin.HandlerFunc {
	return func(c *gin.Context) {
		session := sessions.Default(c)
		uid := session.Get("user_id")
		if uid != nil {
			user := &entity.User{Model: gorm.Model{ID: uid.(uint)}}
			err := repo.GetUser(user)
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
			if _, ok := user.(*entity.User); ok {
				c.Next()
				return
			}
		}

		//c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}
