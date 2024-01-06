package singleton

import "sync"

var (
	counter     int
	counterLock sync.Mutex
)

func Increment() {
	counterLock.Lock()
	defer counterLock.Unlock()
	counter++
}

func GetValue() int {
	counterLock.Lock()
	defer counterLock.Unlock()
	return counter
}
