package main

import (
	"fmt"
)

func printToChannel(s string, ch chan string) {
	for index := range s {
		ch <- string(s[index])
	}
	close(ch)
}

func printBool(b bool, ch chan bool) {
	ch <- b
	close(ch)
}

func main() {
	ch1 := make(chan string)
	ch2 := make(chan bool)
	go printToChannel("HELLO", ch1)
	go printToChannel("GOLANG", ch1)
	go printBool(true, ch2)

	// Print all the results from printToChannel calls
	for out := range ch1 {
		fmt.Println(out)
	}

	// Print result from printBool
	fmt.Println(<-ch2)

}
