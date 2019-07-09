package main

import (
	"fmt"
)

var SerialNumber = "unlicenced"

func main() {
	if SerialNumber == "ABC123" {
		fmt.Println("Valid Serial Number!")
	} else {
		fmt.Println("Invalid Serial Number")
	}
}
