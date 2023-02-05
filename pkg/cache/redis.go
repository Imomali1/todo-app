package cache

import (
	"context"
	"github.com/Imomali1/todo-app/config"
	"github.com/redis/go-redis/v9"
	"time"
)

type Cache struct {
	Client *redis.Client
}

func NewClient(conf config.Configs) *Cache {
	client := &Cache{
		Client: redis.NewClient(&redis.Options{
			Addr:     conf.Redis.Host + conf.Redis.Port,
			Password: conf.Redis.Password,
			DB:       conf.Redis.DB,
		}),
	}

	return client
}

func (c *Cache) Set(ctx context.Context, key string, value interface{}, ttl int) error {
	return c.Client.Set(ctx, key, value, time.Duration(ttl)*time.Second).Err()
}

func (c *Cache) Keys(ctx context.Context, pattern string) ([]string, error) {
	return c.Client.Keys(ctx, pattern).Result()
}

func (c *Cache) Get(ctx context.Context, key string) (string, error) {
	return c.Client.Get(ctx, key).Result()
}
