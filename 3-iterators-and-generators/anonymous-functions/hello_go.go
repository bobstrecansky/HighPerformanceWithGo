package main

import "fmt"

func main() {
	// Note the trailing () for this anonymous function invocation
	func() {
		fmt.Println("Hello Go")
	}()
}
