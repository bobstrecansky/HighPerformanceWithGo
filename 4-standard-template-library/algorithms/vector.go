package main

import vector "container/vector"
import "fmt"

func main() {
	vec := vector.New(0)
	buf := make([]byte, 10)
	vec.Push(buf)

	for i := 0; i < vec.Len(); i++ {
		el := vec.At(i).([]byte)
		fmt.Print(el, "\n")
	}
}
