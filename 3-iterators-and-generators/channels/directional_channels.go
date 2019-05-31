package main

import "fmt"

func main() {
	foo := make(chan<- string, 10)
	bar := make(chan<- string, 10)
	baz := make(<-chan string, 10)

	fmt.Println(foo)
	fmt.Println(bar)
	fmt.Println(baz)
}
