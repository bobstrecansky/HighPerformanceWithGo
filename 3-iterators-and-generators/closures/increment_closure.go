package main

import "fmt"

func incrementCounter() func() int {
	initializedNumber := 0
	return func() int {
		initializedNumber++
		return initializedNumber
	}
}

func main() {
	n1 := incrementCounter()
	fmt.Println("n1 increment counter #1: ", n1()) // First invocation of n1
	fmt.Println("n1 increment counter #2: ", n1()) // Notice the second invocation; n1 is called twice, so n1 == 2
	n2 := incrementCounter()                       // New instance of initializedNumber
	fmt.Println("n2 increment counter #1: ", n2()) // n2 is only called once, so n2 == 1
	fmt.Println("n1 increment counter #3: ", n1()) // // state of n1 is not changed with the n2 calls
}
