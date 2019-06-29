package main

import "fmt"

func main() {
	var x interface{}
	x = "hello Go"
	fmt.Printf("(%v, %T)\n", x, x)
	x = 123
	fmt.Printf("(%v, %T)\n", x, x)
	x = true
	fmt.Printf("(%v, %T)\n", x, x)
}
