package main

import "fmt"

func salutation() func() string {
	return func() string { return "Hello Go" }
}

func main() {
	hello := salutation()
	fmt.Println(hello())
}
