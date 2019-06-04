package main

import "fmt"

func findMinInt(a []int) int {
	var minInt int = a[0]
	for _, i := range a {
		if minInt > i {
			minInt = i
		}
	}
	return minInt

}

func findMaxInt(b []int) int {
	var max int = b[0]
	for _, i := range b {
		if max < i {
			max = i
		}
	}
	return max
}

func main() {
	intData := []int{3, 1, 2, 5, 6, 4}
	minResult := findMinInt(intData)
	maxResult := findMaxInt(intData)
	fmt.Println("Minimum value in array: ", minResult)
	fmt.Println("Maximum value in array: ", maxResult)
}
