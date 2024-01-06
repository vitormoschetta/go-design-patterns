package singleton

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingletonCounter2(t *testing.T) {
	// Arrange
	// Act
	for i := 0; i < 5; i++ {
		SetupCounter()
		Increment()
	}

	// Assert
	assert.Equal(t, 5, GetValue())
}
