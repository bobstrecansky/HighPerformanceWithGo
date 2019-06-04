package main

import (
	"fmt"
)

func reverse(s []string) []string {
	for x, y := 0, len(s)-1; x < y; x, y = x+1, y-1 {
		s[x], s[y] = s[y], s[x]
	}
	return s
}
func main() {
	s := []string{"foo", "bar", "baz", "go", "stop"}
	reversedS := reverse(s)
	fmt.Println(reversedS)
}
