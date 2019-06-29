package main

import (
	"fmt"
	"reflect"
)

func main() {
	var foo string = "Hi Go!"
	fooType := reflect.TypeOf(foo)
	fmt.Println("Foo type: ", fooType)

}
