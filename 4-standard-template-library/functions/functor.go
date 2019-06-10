package main

import (
	"fmt"

	"github.com/go-functional/core/functor"
)

func main() {
	intSlice := []int{1, 3, 5, 7}
	fmt.Println("Int Slice:\t", intSlice)
	intFunctor := functor.LiftIntSlice(intSlice)
	fmt.Println("Lifted Slice:\t", intFunctor)

	// Apply a square to our given functor
	squareFunc := func(i int) int {
		return i * i
	}

	// Apply a mod 3 to our given functor
	modThreeFunc := func(i int) int {
		return i % 3
	}

	squared := intFunctor.Map(squareFunc)
	fmt.Println("Squared: \t", squared)

	modded := squared.Map(modThreeFunc)
	fmt.Println("Modded: \t", modded)
}
