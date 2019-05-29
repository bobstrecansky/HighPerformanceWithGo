package main

import (
	"fmt"
	"sort"
)

func sortInts(intArray []int, done chan bool) {
	sort.Ints(intArray)
	fmt.Printf("Sorted Array: %v\n", intArray)
	done <- true
}

func searchInts(intArray []int, search int, done chan bool) {
	sorted := sort.Search(len(intArray), func(sorted int) bool { return intArray[sorted] >= search })
	if sorted < len(intArray) && intArray[sorted] == search {
		fmt.Printf("Search number: %d\nFound in position: %d\n", search, sorted)
	}
	done <- true
}
func main() {
	ch := make(chan bool)
	go func() {
		s := []int{2, 11, 3, 34, 5, 0, 16} // unsorted
		fmt.Printf("Unsorted Array: %v\n", s)
		sortInts(s, ch)
		searchInts(s, 5, ch)
	}()
	<-ch
}
