package onlogn

import (
	"testing"
)

//func benchmarkBinarySearch(i int, b *testing.B) {
//	for n := 0; n < b.N; n++ {
//		rand.Seed(time.Now().Unix())
//		intGroup := generateIntSlice(i)
//		intRequest := intGroup[rand.Intn(len(intGroup))]
//		b.ResetTimer()
//		BinarySearch(intGroup, intRequest)
//	}
//}
//
//func BenchmarkBinarySearch10(b *testing.B) { benchmarkBinarySearch(2, b) }

//func BenchmarkBinarySearch100(b *testing.B)  { benchmarkBinarySearch(100, b) }
//func BenchmarkBinarySearch1000(b *testing.B) { benchmarkBinarySearch(1000, b) }

func benchmarkOnlognLoop(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		OnlognLoop(i)
	}
}

func BenchmarkOnlognLoop10(b *testing.B)   { benchmarkOnlognLoop(10, b) }
func BenchmarkOnlognLoop100(b *testing.B)  { benchmarkOnlognLoop(100, b) }
func BenchmarkOnlognLoop1000(b *testing.B) { benchmarkOnlognLoop(1000, b) }
