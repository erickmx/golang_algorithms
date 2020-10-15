package fibonacci

func Fibonacci(number int64) int64 {
	if number < 2 {
		return number
	}

	return Fibonacci(number-1) + Fibonacci(number-2)
}
