package main

import (
	"fmt"
)

var SerialNumber = "unlicensed"

func main() {
	if SerialNumber == "ABC123" {
		fmt.Println("Valid Serial Number!")
	} else {
		fmt.Println("Invalid Serial Number")
	}
}
