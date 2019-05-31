package main

import "fmt"

const n = 100

func main() {
	simpleLoop(n)
}

func simpleLoop(n int) {
	for i := 0; i < 100; i++ {
		fmt.Println(i)
	}
}
