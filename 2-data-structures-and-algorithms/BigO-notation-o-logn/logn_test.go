package logn

import "testing"
var result int
func benchmarkBinarySearchTimings(i int, b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		binarySearchTimings(i)
	}
}

func BenchmarkBinarySearchTimings10(b *testing.B)     { benchmarkBinarySearchTimings(10, b) }
func BenchmarkBinarySearchTimings100(b *testing.B)    { benchmarkBinarySearchTimings(100, b) }
func BenchmarkBinarySearchTimings200(b *testing.B)    { benchmarkBinarySearchTimings(200, b) }
func BenchmarkBinarySearchTimings300(b *testing.B)    { benchmarkBinarySearchTimings(300, b) }
func BenchmarkBinarySearchTimings1000(b *testing.B)   { benchmarkBinarySearchTimings(1000, b) }
func BenchmarkBinarySearchTimings2000(b *testing.B)   { benchmarkBinarySearchTimings(2000, b) }
func BenchmarkBinarySearchTimings3000(b *testing.B)   { benchmarkBinarySearchTimings(3000, b) }
func BenchmarkBinarySearchTimings5000(b *testing.B)   { benchmarkBinarySearchTimings(5000, b) }
func BenchmarkBinarySearchTimings10000(b *testing.B)  { benchmarkBinarySearchTimings(10000, b) }
func BenchmarkBinarySearchTimings100000(b *testing.B) { benchmarkBinarySearchTimings(100000, b) }
