package singleton

import (
	"sync"
	"testing"
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
	if GetValue() != numOfGoroutines {
		t.Errorf("Counter value is not correct, expected: %d, actual: %d", numOfGoroutines, GetValue())
	}
}
