package singleton

var counter *int

func SetupCounter() {
	if counter == nil {
		counter = new(int)
	}
}

func Increment() {
	if counter == nil {
		SetupCounter()
	}
	*counter++
}

func GetValue() int {
	if counter == nil {
		SetupCounter()
	}
	return *counter
}
