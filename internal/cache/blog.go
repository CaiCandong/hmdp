package cache

import (
	"context"
	"fmt"
)

func IsLike(ctx context.Context, blogId uint, userId uint) bool {
	key := fmt.Sprintf("like:%d", blogId)
	result, _ := RedisStore.SIsMember(ctx, key, userId).Result()
	return result
}
