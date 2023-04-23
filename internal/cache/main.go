package cache

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"hmdp/internal/model"
	"hmdp/pkg/utils"
	"time"
)

var RedisStore *redis.Client
var (
	ErrEmptyRecord = fmt.Errorf("record not found")
)

func InitRedisStore() *redis.Client {
	address := viper.GetString("redis.address")
	password := viper.GetString("redis.password")
	db := viper.GetInt(`server.db`)
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	RedisStore = client
	return client
}

func GetShopById(ctx context.Context, shop *model.Shop) error {
	js, err := RedisStore.Get(ctx, fmt.Sprintf("shop:%v", shop.ID)).Result()
	// 未命中
	if err != nil {
		return err
	}
	// 命中空值
	if js == "" {
		return ErrEmptyRecord
	}
	// 命中真实值
	return utils.FromJSON(js, shop)
}

func SaveNotFind(ctx context.Context, shop *model.Shop) error {
	expire := time.Duration(1) * time.Minute
	return RedisStore.Set(ctx, fmt.Sprintf("shop:%v", shop.ID), "", expire).Err()
}
