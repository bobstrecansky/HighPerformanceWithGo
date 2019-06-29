package main

import (
	"fmt"
	"reflect"
)

func main() {

	example := "foo"
	exampleVal := reflect.ValueOf(example)
	fmt.Println(exampleVal)

}
