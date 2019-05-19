package on2

import "testing"

func benchmarkBubbleSort(x int, b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		BubbleSort(x)
	}
}

func BenchmarkBubbleSort10(b *testing.B)     { benchmarkBubbleSort(10, b) }
func BenchmarkBubbleSort100(b *testing.B)    { benchmarkBubbleSort(100, b) }
func BenchmarkBubbleSort1000(b *testing.B)   { benchmarkBubbleSort(1000, b) }
func BenchmarkBubbleSort10000(b *testing.B)  { benchmarkBubbleSort(10000, b) }
func BenchmarkBubbleSort100000(b *testing.B) { benchmarkBubbleSort(100000, b) }
