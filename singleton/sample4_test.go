package singleton

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSingletonCounter2(t *testing.T) {
	// Arrange
	numOfGoroutines := 1000
	var wg sync.WaitGroup
	wg.Add(numOfGoroutines)

	// Act
	for i := 0; i < numOfGoroutines; i++ {
		go func() {
			defer wg.Done()
			Increment()
		}()
	}
	wg.Wait()

	// Assert
	assert.Equal(t, numOfGoroutines, GetValue())
}
