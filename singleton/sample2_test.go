package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingleton2(t *testing.T) {
	// Setup 1
	// Arrange
	SetupCache()

	// Act
	Add("key", "value")
	res := Get("key")

	// Assert
	assert.Equal(t, "value", res)

	// Setup 2
	// Arrange
	SetupCache()

	// Act
	res2 := Get("key")

	// Assert
	assert.Equal(t, "value", res2)
}
