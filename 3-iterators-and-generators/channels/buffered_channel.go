package main

import "fmt"

func main() {

	buffered_channel := make(chan string, 2)
	buffered_channel <- "foo"
	buffered_channel <- "bar"

	// Length of channel is 2 because both elements added to channel
	fmt.Println("Channel Length After Add: ", len(buffered_channel))

	// Pop foo and bar off the stack
	fmt.Println(<-buffered_channel)
	fmt.Println(<-buffered_channel)

	// Length of channel is 0 because both elements removed from channel
	fmt.Println("Channel Length After Pop: ", len(buffered_channel))

	// Push baz to the stack
	buffered_channel <- "baz"

	// Store baz as a variable, out
	out := <-buffered_channel
	fmt.Println(out)
	close(buffered_channel)
}
