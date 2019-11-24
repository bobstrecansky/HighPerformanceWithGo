package benchmark

import "testing"

func BenchmarkCPrint(b *testing.B) {
	CgoPrint(b.N)
}

func BenchmarkGoPrint(b *testing.B) {
	GoPrint(b.N)
}
