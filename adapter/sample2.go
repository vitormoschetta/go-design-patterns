package adapter

import (
	"context"

	"github.com/redis/go-redis/v9"
)

// Definimos um contrato para o cache, para que, independente do cache que for usado, ele tenha que implementar esses métodos
type ICache interface {
	Get(ctx context.Context, key string) (string, error)
	Set(ctx context.Context, key, value string) (string, error)
}

type CacheLocal struct {
	cache map[string]string
}

func (c *CacheLocal) Get(ctx context.Context, key string) (string, error) {
	return c.cache[key], nil
}

func (c *CacheLocal) Set(ctx context.Context, key, value string) (string, error) {
	c.cache[key] = value
	return value, nil
}

// ADAPTER
// Aqui criamos um adapter para o redis, que implementa a interface ICache
type CacheRedis struct {
	cache redis.Client
}

func (c *CacheRedis) Get(ctx context.Context, key string) (string, error) {
	return c.cache.Get(ctx, key).Result()
}

func (c *CacheRedis) Set(ctx context.Context, key, value string) (string, error) {
	return c.cache.Set(ctx, key, value, 0).Result()
}

type UserService struct {
	cache ICache // usamos a interface ICache, para que possamos usar qualquer cache que implemente essa interface
}

func (u *UserService) GetUserName(ctx context.Context, userID string) (string, error) {
	return u.cache.Get(ctx, userID)
}

func (u *UserService) SetUserName(ctx context.Context, userID, userName string) (string, error) {
	return u.cache.Set(ctx, userID, userName)
}
