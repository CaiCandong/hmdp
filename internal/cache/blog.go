package cache

import (
	"context"
	"github.com/redis/go-redis/v9"
	"hmdp/internal/model"
	"strconv"
)

func IsLike(ctx context.Context, blogId uint, userId uint) bool {
	key := model.BlogLikeKey(blogId)
	_, err := RedisStore.ZScore(ctx, key, strconv.Itoa(int(userId))).Result()
	if err == redis.Nil { // 未点赞
		return false
	}
	return true
}
