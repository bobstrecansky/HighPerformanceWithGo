package main

import (
	"fmt"
	"time"
)

func main() {
	ch1 := make(chan string, 1)
	ch2 := make(chan string, 1)

	// RPC call example (completes successfully)
	go func() {
		time.Sleep(50 * time.Millisecond)
		ch1 <- "Completed RPC"
	}()

	select {
	case receivedRPC := <-ch1:
		fmt.Println("RPC call result: ", receivedRPC)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("RPC timed out")
	}

	// API call example (times out)
	go func() {
		time.Sleep(1000 * time.Millisecond)
		ch2 <- "Completed API Call"
	}()
	select {
	case receivedAPI := <-ch2:
		fmt.Println("API call result: ", receivedAPI)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("API timed out")
	}
}
