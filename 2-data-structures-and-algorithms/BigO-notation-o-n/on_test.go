package on

import "testing"

func benchmarkSimpleLoop(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		simpleLoopSum(i)
	}
}

func BenchmarkSimpleLoop10(b *testing.B)   { benchmarkSimpleLoop(10, b) }
func BenchmarkSimpleLoop100(b *testing.B)  { benchmarkSimpleLoop(100, b) }
func BenchmarkSimpleLoop1000(b *testing.B) { benchmarkSimpleLoop(1000, b) }
