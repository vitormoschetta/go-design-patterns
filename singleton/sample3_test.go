package singleton

import (
	"sync"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCounter_Increment_Concurrency(t *testing.T) {
	// Arrange
	numOfGoroutines := 1000
	counter := NewCounter()
	var wg sync.WaitGroup
	wg.Add(numOfGoroutines)

	// Act
	for i := 0; i < numOfGoroutines; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	// Assert
	assert.Equal(t, numOfGoroutines, counter.GetValue())
}
