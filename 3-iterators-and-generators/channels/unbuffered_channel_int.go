package main

import (
	"fmt"
	"sort"
)

func checkForSortedInts(intArray []int, done chan bool) {
	result := sort.IntsAreSorted(intArray)
	fmt.Println(intArray, "sorted: ", result)
	done <- false
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
		intArray := []int{2, 11, 3, 34, 5, 0, 16} // unsorted
		searchNumber := 16
		fmt.Printf("Unsorted Array: %v\n", intArray)
		checkForSortedInts(intArray, ch)
		fmt.Printf("Sorted Array %v\n", intArray)
		searchInts(intArray, searchNumber, ch)
	}()
	<-ch
}
