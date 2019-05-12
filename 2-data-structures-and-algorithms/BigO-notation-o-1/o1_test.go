package oone

import "testing"

func BenchmarkThree(b *testing.B) {
	for n := 0; n < b.N; n++ {
		ThreeWords()
	}
}

func BenchmarkTen(b *testing.B) {
	for n := 0; n < b.N; n++ {
		TenWords()
	}
}
