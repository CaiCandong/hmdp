package middleware

import (
	"github.com/gin-gonic/gin"
	"hmdp/internal/model"
)

// CurrentUser 获取登录用户
//func CurrentUser(repo repository.IUserRepo) gin.HandlerFunc {
//	return func(c *gin.Context) {
//		session := sessions.Default(c)
//		uid := session.Get("user_id")
//		if uid != nil {
//			user := &entity.User{Model: gorm.Model{ID: uid.(uint)}}
//			err := repo.GetUser(user)
//			if err == nil {
//				c.Set("user", user)
//			}
//		}
//		c.Next()
//	}
//}

func AuthRequired() gin.HandlerFunc {
	return func(c *gin.Context) {
		if user, ok := c.Get("user"); ok && user != nil {
			_, ok := user.(*model.User)
			if ok {
				c.Next()
				return
			}
		}

		//c.JSON(200, serializer.CheckLogin())
		c.Abort()
	}
}
