package iterators

import "testing"

func benchmarkLoop(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		simpleLoop(i)
	}
}

func benchmarkCallback(i int, b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		CallbackLoop(i)
	}
}

func benchmarkNext(i int, b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		NextLoop(i)
	}
}

func benchmarkBufferedChan(i int, b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		BufferedChanLoop(i)
	}
}

func benchmarkUnbufferedChan(i int, b *testing.B) {
	b.ResetTimer()
	for n := 0; n < b.N; n++ {
		UnbufferedChanLoop(i)
	}
}

func BenchmarkLoop10000000(b *testing.B)           { benchmarkLoop(1000000, b) }
func BenchmarkCallback10000000(b *testing.B)       { benchmarkCallback(1000000, b) }
func BenchmarkNext10000000(b *testing.B)           { benchmarkNext(1000000, b) }
func BenchmarkBufferedChan10000000(b *testing.B)   { benchmarkBufferedChan(1000000, b) }
func BenchmarkUnbufferedChan10000000(b *testing.B) { benchmarkUnbufferedChan(1000000, b) }
