package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"hmdp/internal/domain/entity"
	"hmdp/internal/infrastructure/cache"
)

func EnableCookieSession() gin.HandlerFunc {
	address := viper.GetString("redis.address")
	password := viper.GetString("redis.password")
	cookieKey := viper.GetString(`server.cookie_key`)
	store, err := redis.NewStore(10, "tcp", address, password, []byte(cookieKey))
	if err != nil {
		panic(err)
	}

	return sessions.Sessions("go-gin-chat", store)
}

func CurrentUserByToken() gin.HandlerFunc {
	return func(c *gin.Context) {
		// Get the token from the request header or query string
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			token = c.Query("token")
		}

		// Check if the token is valid
		if token == "" {
			c.Next()
			return
		}

		// Validate the token using your authentication logic
		// In this example, we're just checking if the token is "123456"
		var user entity.User
		err := cache.GetUser(c, token, &user)
		if err != nil {
			c.Next()
			return
		}
		c.Set("user", &user)
		// Token is valid, so continue to the next middleware or handler
		c.Next()
	}
}
