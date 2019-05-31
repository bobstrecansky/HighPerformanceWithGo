package main

import (
	"fmt"
	"sort"
)

func main() {
	intArray := []int{2, 11, 3, 34, 5, 0, 16} // unsorted
	searchNumber := 34
	sorted := sort.SearchInts(intArray, searchNumber)
	if sorted < len(intArray) {
		fmt.Printf("Found element %d at array position %d\n", searchNumber, sorted)
	} else {
		fmt.Printf("Element %d not found in array %v\n", searchNumber, intArray)
	}
}
