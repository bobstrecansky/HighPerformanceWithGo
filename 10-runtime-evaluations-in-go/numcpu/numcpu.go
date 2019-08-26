package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Number of CPUs Available: ", runtime.NumCPU())
}
