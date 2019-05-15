package o2n

import "testing"

func benchmarkFibonacci(x int, b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		Fibonacci(x)
	}
}

func BenchmarkFibonacci1(b *testing.B)  { benchmarkFibonacci(1, b) }
func BenchmarkFibonacci2(b *testing.B)  { benchmarkFibonacci(2, b) }
func BenchmarkFibonacci5(b *testing.B)  { benchmarkFibonacci(5, b) }
func BenchmarkFibonacci10(b *testing.B) { benchmarkFibonacci(10, b) }
func BenchmarkFibonacci15(b *testing.B) { benchmarkFibonacci(15, b) }
func BenchmarkFibonacci20(b *testing.B) { benchmarkFibonacci(20, b) }
