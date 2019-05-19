package main

import (
	"fmt"
	"math/rand"
)

const NumItems int = 3

var int_data []int = make([]int, NumItems)

func InitInts() {
	for i := 0; i < NumItems; i++ {
		int_data[i] = rand.Intn(2)
	}
}

func IntCallbackIterator(cb func(int)) {
	for _, val := range int_data {
		fmt.Print(int_data, val)
		cb(val)
	}
}

func main() {
	InitInts()
	var sum int = 0
	cb := func(val int) {
		sum += val
	}
	IntCallbackIterator(cb)
	fmt.Print(sum)
}
