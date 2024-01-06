package singleton

import (
	"sync"
	"testing"
)

func TestCounter_Increment_Concurrency(t *testing.T) {
	numOfGoroutines := 1000
	counter := NewCounter()
	var wg sync.WaitGroup
	wg.Add(numOfGoroutines)

	for i := 0; i < numOfGoroutines; i++ {
		go func() {
			defer wg.Done()
			counter.Increment()
		}()
	}

	wg.Wait()

	if counter.GetValue() != numOfGoroutines {
		t.Errorf("Counter value is not correct, expected: %d, actual: %d", numOfGoroutines, counter.GetValue())
	}

}
