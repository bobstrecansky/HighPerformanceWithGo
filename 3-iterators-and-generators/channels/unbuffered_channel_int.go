package main

import (
	"fmt"
	"sort"
)

func sortInts(intArray []int, done chan bool) {
	sort.Ints(intArray)
	fmt.Println(intArray)
	done <- true
}

func searchInts(intArray []int, done chan bool) {
	i := sort.Search(len(intArray), func(i int) bool { return intArray[i] >= 5 })
	if i < len(intArray) && intArray[i] == 5 {
		fmt.Println("Found in position :", i)
	}
	done <- true
}
func main() {
	ch := make(chan bool)
	go func() {
		s := []int{2, 11, 3, 34, 5, 0, 16} // unsorted
		sortInts(s, ch)
		searchInts(s, ch)
	}()
	<-ch
}
