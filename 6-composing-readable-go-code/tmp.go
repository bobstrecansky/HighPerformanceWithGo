package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x int = 1
	v := reflect.ValueOf(x)
	fmt.Println(v.Type())
	fmt.Println(reflect.ValueOf(v))
}
