package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	// Arrange
	cache := &Cache{}

	// Act
	cache.Add("key", "value")

	// Assert
	for i := 0; i < 5; i++ {
		cache = NewCache()
		res := cache.Get("key")
		assert.Equal(t, "value", res)
	}
}
