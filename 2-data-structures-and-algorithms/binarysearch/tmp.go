package main

import (
	"fmt"
	"sort"
)

func main() {

	intArray := []int{3, 2, 5, 4, 1, 6}
	sortedIntArray := []int{1, 2, 3, 4, 5, 6}
	searchNumber := 3
	//	sort.Ints(intArray)
	x := sort.IntsAreSorted(intArray)
	y := sort.SearchInts(sortedIntArray, searchNumber)
	fmt.Println(x)
	fmt.Print(y)
}
