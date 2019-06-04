package main

import (
	"container/list"
	"fmt"
)

func main() {
	// Our goal is to create a linked list with a value of 1,2,3,4
	// Create a new doubly linked list
	ll := list.New()
	three := ll.PushBack(3)           // stack representation -> [3]
	four := ll.InsertBefore(4, three) // stack representation -> [4 3]
	ll.InsertBefore(2, three)         // stack representation -> [4 2 3]
	ll.MoveToBack(four)               // stack representation -> [2 3 4]
	ll.PushFront(1)                   // stack representation -> [1 2 3 4]
	listLength := ll.Len()

	// Show the linkedList type and length
	fmt.Printf("ll type: %T\n", ll)
	fmt.Println("ll length: :", listLength)

	// Iterate through ll to print the values
	for e := ll.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
