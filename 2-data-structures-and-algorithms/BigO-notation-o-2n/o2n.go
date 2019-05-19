package o2n

func Fibonacci(i int) int {
	if i <= 1 {
		return 1
	}
	return Fibonacci(i-1) + Fibonacci(i-2)
}
