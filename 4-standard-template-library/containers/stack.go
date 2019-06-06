package main

import (
	"fmt"

	stack "github.com/golang-collections/collections/stack"
)

func main() {
	// Create a new stack
	fmt.Println("Creating New Stack")
	exstack := stack.New()
	fmt.Println("Pushing 1 to stack")
	exstack.Push(1) // push 1 to stack
	fmt.Println("Top of Stack is : ", exstack.Peek())
	fmt.Println("Popping 1 from stack")
	exstack.Pop() // remove 1 from stack
	fmt.Println("Stack length is : ", exstack.Len())
}
