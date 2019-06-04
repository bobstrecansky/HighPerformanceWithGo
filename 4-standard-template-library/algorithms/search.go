package main

import (
	"fmt"
	"sort"
)

func main() {
	data := []int{1, 2, 3, 4, 5, 6}
	findInt := 2
	out := sort.Search(len(data), func(i int) bool { return data[i] >= findInt })
	fmt.Printf("Integer %d was found in %d at position %d\n", findInt, data, out)
}
