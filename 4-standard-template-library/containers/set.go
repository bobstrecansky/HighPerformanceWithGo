package main

import "fmt"

func main() {
	s := make(map[int]bool)

	for i := 0; i < 5; i++ {
		s[i] = true
	}

	delete(s, 4)

	if s[2] {
		fmt.Println("s[2] is set")
	}
	if !s[4] {
		fmt.Println("s[4] was deleted")
	}
}
