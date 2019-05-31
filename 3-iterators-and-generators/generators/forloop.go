package main

import "fmt"

func count(n int) {
	for i := 0; i < n; i++ {
		fmt.Println("Counted", i)
	}
}

func main() {
	count(99999)
}
