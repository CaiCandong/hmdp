package cache

import (
	"context"
	"time"
)

func Lock(ctx context.Context, key string) (bool, error) {
	return RedisStore.SetNX(ctx, key, "1", time.Second*10).Result()
}

func Unlock(ctx context.Context, key string) error {
	return RedisStore.Del(ctx, key).Err()
}
