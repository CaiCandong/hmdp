package middleware

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/redis"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
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
