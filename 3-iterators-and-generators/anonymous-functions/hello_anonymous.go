package main

import "fmt"

func HelloGo() {
	fmt.Println("Hello Go from a Function")

}

func main() {

	HelloGo()
	func() { fmt.Println("Hello Go from an Anonymous Function") }()
	var hello func() = func() { fmt.Println("Hello Go from an Anonymous Function Assigned to a Variable") }
	hello()
}
