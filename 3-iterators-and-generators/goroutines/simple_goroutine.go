package main

import (
	"fmt"
	"time"
)

func printSleep(s string) {
	for index, stringVal := range s {
		fmt.Printf("%#U at index %d\n", stringVal, index)
		time.Sleep(1 * time.Millisecond)
	}
}

func main() {
	const t time.Duration = 9
	go printSleep("HELLO GOPHERS")
	time.Sleep(t * time.Millisecond)
	fmt.Println("sleep complete")
}
