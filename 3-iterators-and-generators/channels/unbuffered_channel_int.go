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

func searchInts(intArray []int, searchNumber int, done chan bool) {
	sorted := sort.SearchInts(intArray, searchNumber)
	if sorted < len(intArray) {
		fmt.Printf("Found element %d at array position %d\n", searchNumber, sorted)
	} else {
		fmt.Printf("Element %d not found in array %v\n", searchNumber, intArray)
	}
	done <- true
}
func main() {
	ch := make(chan bool)
	go func() {
		s := []int{2, 11, 3, 34, 5, 0, 16} // unsorted
		fmt.Println("Unsorted Array: ", s)
		searchNumber := 16
		sortInts(s, ch)
		searchInts(s, searchNumber, ch)
	}()
	<-ch
}
