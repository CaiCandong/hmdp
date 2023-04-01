package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"log"
	"time"
)

type RedisStore struct {
	client *redis.Client
}

var RediStore RedisStore

func InitRedisStore() *RedisStore {
	address := viper.GetString("redis.address")
	password := viper.GetString("redis.password")
	db := viper.GetInt(`server.db`)
	client := redis.NewClient(&redis.Options{
		Addr:     address,
		Password: password,
		DB:       db,
	})

	RediStore = RedisStore{client: client}
	return &RediStore
}

// set
func (rs RedisStore) Set(id string, value string) {
	ctx := context.Background()
	err := rs.client.Set(ctx, id, value, time.Minute*10).Err()
	if err != nil {
		log.Println(err)
	}
}

// get
func (rs RedisStore) Get(id string, clear bool) string {
	ctx := context.Background()
	val, err := rs.client.Get(ctx, id).Result()
	if err != nil {
		log.Println(err)
		return ""
	}
	if clear {
		err := rs.client.Del(ctx, id).Err()
		if err != nil {
			log.Println(err)
			return ""
		}
	}
	return val
}

// verify a capt
func (r RedisStore) Verify(id, answer string, clear bool) bool {
	v := r.Get(id, clear)
	return v == answer
}
