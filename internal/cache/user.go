package cache

import (
	"context"
	"github.com/spf13/viper"
	"hmdp/internal/model"
	"hmdp/pkg/utils"
	"time"
)

// SaveUser saves a hash to redis
func SaveUser(ctx context.Context, token string, user *model.User) error {
	expire := time.Duration(viper.GetInt("server.session_expire")) * time.Minute
	return RedisStore.Set(ctx, token, utils.ToJSON(user), expire).Err()
}

func GetUser(ctx context.Context, token string, user *model.User) error {
	js, err := RedisStore.Get(ctx, token).Result()
	// 刷新过期时间
	RedisStore.Expire(ctx, token, time.Duration(viper.GetInt("server.session_expire"))*time.Minute)
	if err != nil {
		return err
	}
	err = utils.FromJSON(js, &user)
	if err != nil {
		return err
	}
	return nil
}
