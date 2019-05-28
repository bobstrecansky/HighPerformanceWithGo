package main

import "fmt"

func incrementCounter() func() int {
	var initializedNumber = 0
	return func() int {
		initializedNumber++
		return initializedNumber
	}
}

func main() {
	o := incrementCounter()
	fmt.Print(o())

	fmt.Print(o())
	fmt.Print(o())
	fmt.Print(o())
	fmt.Print(o())
}
