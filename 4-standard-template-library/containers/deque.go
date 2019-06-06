package main

import (
	"fmt"

	"gopkg.in/karalabe/cookiejar.v1/collections/deque"
)

func main() {
	d := deque.New()
	elements := []string{"foo", "bar", "baz"}
	for i := range elements {
		d.PushLeft(elements[i])
	}
	fmt.Println(d.PopLeft())  // queue => ["foo", "bar"]
	fmt.Println(d.PopRight()) // queue => ["bar"]
	fmt.Println(d.PopLeft())  // queue => empty
}
