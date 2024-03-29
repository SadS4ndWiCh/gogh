package store

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type RedisStore struct{
    redis *redis.Client
}

func NewRedisStore(url string) *RedisStore {
	redisOpts, err := redis.ParseURL(url)
	if err != nil {
		panic(err)
	}

	redisClient := redis.NewClient(redisOpts)

    return &RedisStore{
        redis: redisClient,
    }
}

func (st *RedisStore) Get(ctx context.Context, key string) (interface{}, error) {
    value, err := st.redis.Get(ctx, key).Result()
    if err != nil {
        return nil, err
    }

    return value, nil
}

func (st *RedisStore) Set(ctx context.Context, key string, value interface{}) error {
    return st.redis.Set(ctx, key, value, time.Hour).Err()
}
