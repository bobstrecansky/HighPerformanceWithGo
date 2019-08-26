package main

import (
	"time"
)

func main() {
	c := make(chan bool, 1)
	go panicRoutine(c)
	for i := 0; i < 2; i++ {
		<-c
	}
}

func panicRoutine(c chan bool) {
	time.Sleep(100 * time.Millisecond)
	panic("Goroutine Panic")
	c <- true
}
