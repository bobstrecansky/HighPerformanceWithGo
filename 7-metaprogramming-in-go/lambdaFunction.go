package main

import "fmt"

func modulus(x, y int) int {
	return x % y
}

func main() {
	z := modulus(3, 2)
	fmt.Println(z)
}
