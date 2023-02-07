package cache

import (
	"context"
	"fmt"
	"github.com/Imomali1/todo-app/config"
	"github.com/redis/go-redis/v9"
	"time"
)

type ICache struct {
	Client *redis.Client
}

func NewClient(conf config.Configs) *ICache {
	client := &ICache{
		Client: redis.NewClient(&redis.Options{
			Addr:     fmt.Sprintf("%s:%d", conf.Redis.Host, conf.Redis.Port),
			Password: conf.Redis.Password,
			DB:       conf.Redis.DB,
		}),
	}

	return client
}

func (c *ICache) Set(ctx context.Context, key string, value interface{}, ttl int) error {
	return c.Client.Set(ctx, key, value, time.Duration(ttl)*time.Second).Err()
}

func (c *ICache) Keys(ctx context.Context, pattern string) ([]string, error) {
	return c.Client.Keys(ctx, pattern).Result()
}

func (c *ICache) Get(ctx context.Context, key string) (string, error) {
	return c.Client.Get(ctx, key).Result()
}
