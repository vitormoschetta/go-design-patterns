package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton(t *testing.T) {
	// Setup 1
	// Arrange
	cache := Cache{} // not using NewCache()

	// Act
	cache.Add("key", "value")
	res := cache.Get("key")

	// Assert
	assert.Equal(t, "value", res)

	// Setup 2
	// Arrange
	cache2 := NewCache()

	// Act
	res2 := cache2.Get("key")

	// Assert
	assert.Equal(t, "value", res2)

	// Setup 3
	// Arrange
	cache3 := NewCache()

	// Act
	res3 := cache3.Get("key")

	// Assert
	assert.Equal(t, "value", res3)
}
