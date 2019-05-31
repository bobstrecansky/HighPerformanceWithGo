package main

import (
	"fmt"
	"time"
)

func main() {
	var out = make([]string, 5)
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		out[i] = "This loop is slow\n"
	}
	fmt.Println(out)
}
