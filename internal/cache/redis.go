package cache

import (
	"context"
	"time"
	"github.com/redis/go-redis/v9"
)

type RedisCache struct {
	client *redis.Client
}

func NewRedisCache(addr, password string, db int) (*RedisCache, error) {
	client := redis.NewClient(&redis.Options{
    Addr:     addr,
    Password: password, 
    DB:       db,      
  })

  _, err := client.Ping(context.Background()).Result()
  if err!= nil {
    return nil, err
  }

  return &RedisCache{client: client}, nil

}

func (c *RedisCache) Get(ctx context.Context, key string) (string, error) {
  return c.client.Get(ctx, key).Result()
}

func (c *RedisCache) Set(ctx context.Context, key string, value string, expiration time.Duration) error {
  return c.client.Set(ctx, key, value, expiration).Err()
}

func (c *RedisCache) Delete(ctx context.Context, key string) error {
  return c.client.Del(ctx, key).Err()
}