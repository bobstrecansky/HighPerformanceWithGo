package on2

import "testing"

func benchmarkQuadraticTimings(x, y int, b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		quadraticTimings(x, y)
	}
}

func BenchmarkQuadraticTimings10(b *testing.B)   { benchmarkQuadraticTimings(10, 10, b) }
func BenchmarkQuadraticTimings100(b *testing.B)  { benchmarkQuadraticTimings(100, 100, b) }
func BenchmarkQuadraticTimings1000(b *testing.B) { benchmarkQuadraticTimings(1000, 1000, b) }
