package redis

import (
	"accounting-service/core/environment"
	"context"
	"github.com/go-redis/redis/v8"
	"time"
)

type Cache struct {
	env    *environment.Environment
	ctx    context.Context
	client *redis.Client
}

func New(env *environment.Environment, ctx context.Context) *Cache {
	rdb := redis.NewClient(&redis.Options{
		Addr:     env.RedisURL,
		Password: env.RedisPassword,
		DB:       0, // use default DB
	})
	return &Cache{env: env, ctx: ctx, client: rdb}
}

func (conf *Cache) SetValue(key string, value interface{}, duration time.Duration) error {
	err := conf.client.Set(conf.ctx, key, value, duration).Err()
	return err
}

func (conf *Cache) GetValue(key string) (string, error) {
	val, err := conf.client.Get(conf.ctx, key).Result()
	return val, err
}
