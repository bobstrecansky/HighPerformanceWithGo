package main

import (
	"fmt"
	"sort"
)

func main() {
	input := []string{"foo", "bar", "baz"}
	var result []string
	// closure callback
	func() {
		result = append(input, "abc")       // Append to the array
		result = append(result, "def")      // Append to the array again
		sort.Sort(sort.StringSlice(result)) // Sort the larger array
	}()
	fmt.Print(result)
}
