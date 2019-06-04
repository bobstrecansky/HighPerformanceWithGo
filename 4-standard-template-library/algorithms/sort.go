package main

import (
	"fmt"
	"sort"
)

func main() {

	intData := []int{3, 1, 2, 5, 6, 4}
	stringData := []string{"foo", "bar", "baz"}
	floatData := []float64{1.5, 3.6, 2.5, 10.6}

	sort.Ints(intData)
	sort.Strings(stringData)
	sort.Float64s(floatData)

	fmt.Println("Sorted Integers: ", intData, "\nSorted Strings: ", stringData, "\nSorted Floats: ", floatData)
}
