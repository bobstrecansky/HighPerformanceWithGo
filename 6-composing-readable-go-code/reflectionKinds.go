package main

import (
	"fmt"
	"reflect"
)

func main() {
	i := []string{"foo", "bar", "baz"}
	ti := reflect.TypeOf(i)
	fmt.Println(ti.Kind())
}
