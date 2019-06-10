package main

import "fmt"

func main() {
	stringExample := []string{"foo", "bar", "baz"}
	for i, out := range stringExample {
		fmt.Println(i, out)
	}
}
