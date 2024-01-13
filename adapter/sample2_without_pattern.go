package adapter

import (
	"context"

	"github.com/redis/go-redis/v9"
)

type CacheLocal2 struct {
	cache map[string]string
}

func (c *CacheLocal2) Get(ctx context.Context, key string) (string, error) {
	return c.cache[key], nil
}

func (c *CacheLocal2) Set(ctx context.Context, key, value string) (string, error) {
	c.cache[key] = value
	return value, nil
}

// O que acontece se eu quiser trocar o redis por outro cache?
// Ou se eu quiser executar outro cache quando em modo de teste? em localhost?
type UserService2 struct {
	cache redis.Client // interface própria do redis
}

func (u *UserService2) GetUserName(ctx context.Context, userID string) (string, error) {
	return u.cache.Get(ctx, userID).Result()
}

func (u *UserService2) SetUserName(ctx context.Context, userID, userName string) (string, error) {
	return u.cache.Set(ctx, userID, userName, 0).Result()
}
