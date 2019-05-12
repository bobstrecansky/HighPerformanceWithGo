package onlogn

import "testing"

func BenchMarkBinarySearch5(b *testing.B) {
	for n := 0; n < b.N; n++ {
		intGroup := []int{5, 1, 2, 3, 4}
		intRequest := 3
		BinarySearch(intGroup, intRequest)
	}
}

func BenchMarkBinarySearch10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		intGroup := []int{6, 5, 7, 9, 8, 1, 2, 3, 4, 10}
		intRequest := 6
		BinarySearch(intGroup, intRequest)
	}
}
