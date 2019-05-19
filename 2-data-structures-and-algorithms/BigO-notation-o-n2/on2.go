package on2

import (
	"math/rand"
	"time"
)

func generateRandomSlice(size int) []int {

	slice := make([]int, size, size)
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < size; i++ {
		slice[i] = rand.Intn(999) - rand.Intn(999)
	}
	return slice
}

func BubbleSort(i int) []int {
	k := generateRandomSlice(i)
	for i := 1; i < len(k); i++ {
		for j := 0; j < len(k)-i; j++ {
			if k[j] > k[j+1] {
				k[j], k[j+1] = k[j+1], k[j]
			}
		}
	}
	return k
}
