package cache

import (
	"context"
	"fmt"
	"github.com/spf13/viper"
	"hmdp/internal/model"
	"hmdp/pkg/utils"
	"time"
)

func SaveShopById(ctx context.Context, shop *model.Shop) error {
	expire := time.Duration(viper.GetInt("server.session_expire")) * time.Second
	return RedisStore.Set(ctx, fmt.Sprintf("shop:%v", shop.ID), utils.ToJSON(shop), expire).Err()
}

func SaveShopType(ctx context.Context, shopType []*model.ShopType) error {
	expire := time.Duration(viper.GetInt("server.session_expire")) * time.Second
	return RedisStore.Set(ctx, fmt.Sprintf("shop_type"), utils.ToJSON(shopType), expire).Err()
}

func GetShopType(ctx context.Context) ([]*model.ShopType, error) {
	var shopType []*model.ShopType
	js, err := RedisStore.Get(ctx, fmt.Sprintf("shop_type")).Result()
	if err != nil {
		return nil, err
	}
	// 刷新过期时间
	RedisStore.Expire(ctx, fmt.Sprintf("shop_type"), time.Duration(viper.GetInt("server.session_expire"))*time.Second)
	err = utils.FromJSON(js, &shopType)
	return shopType, err
}
func DeleteShopById(ctx context.Context, shop *model.Shop) error {
	return RedisStore.Del(ctx, fmt.Sprintf("shop:%v", shop.ID)).Err()
}
