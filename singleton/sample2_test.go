package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingletonCache2(t *testing.T) {
	// Arrange
	SetupCache()

	// Act
	Add("key", "value")

	// Assert
	for i := 0; i < 5; i++ {
		SetupCache()
		res := Get("key")
		assert.Equal(t, "value", res)
	}
}
