package main

import "fmt"

func main() {
	m := make(map[int]string)
	m[1] = "car"
	m[2] = "train"
	m[3] = "plane"
	fmt.Println("Full Map:\t ", m)
	fmt.Println("m[3] value:\t ", m[3])
	fmt.Println("Length of map:\t ", len(m))
}
