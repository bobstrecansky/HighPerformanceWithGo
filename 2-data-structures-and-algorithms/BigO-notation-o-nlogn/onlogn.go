package onlogn

import "sort"

func BinarySearch(intGroup []int, intRequest int) int {
	out := sort.SearchInts(intGroup, intRequest)
	return out
}
