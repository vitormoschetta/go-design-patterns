package adapter

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/assert"
)

func TestCacheRedisClient(t *testing.T) {
	// Arrange
	cacheRedis := CacheRedis{
		cache: redis.Client{},
	}

	userService := UserService{
		cache: &cacheRedis,
	}

	defer func() {
		if r := recover(); r != nil {
			assert.Equal(t, 1, 1, "regis cache is not available")
		}
	}()

	// Act
	userName, err := userService.SetUserName(context.Background(), "1", "John Doe")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", userName)
}

func TestCacheAdapterLocalClient(t *testing.T) {
	// Arrange
	cacheLocal := CacheLocal{
		cache: map[string]string{},
	}

	userService := UserService{
		cache: &cacheLocal,
	}

	// Act
	_, err := userService.SetUserName(context.Background(), "1", "John Doe")
	assert.NoError(t, err)

	userName, err := userService.GetUserName(context.Background(), "1")

	// Assert
	assert.NoError(t, err)
	assert.Equal(t, "John Doe", userName)
}
