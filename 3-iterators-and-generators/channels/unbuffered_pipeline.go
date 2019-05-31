package main

import (
	"fmt"
)

func main() {

	runeCh := make(chan int)
	runeConvCh := make(chan rune)

	go func() {
		for i := 33; i < 123; i++ {
			runeCh <- i
		}
	}()

	go func() {
		for i := 33; i < 123; i++ {
			j := <-runeCh
			runeConvCh <- rune(j)
		}
	}()
	for i := 33; i < 123; i++ {
		k := <-runeConvCh
		fmt.Println(string(k))
	}
}
