package main

import "fmt"

func main() {

	bufferedChannel := make(chan int, 3)
	bufferedChannel <- 1
	bufferedChannel <- 3
	bufferedChannel <- 5
	close(bufferedChannel)
	for i := range bufferedChannel {
		fmt.Println(i)
	}
}
