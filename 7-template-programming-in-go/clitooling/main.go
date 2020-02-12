package main

import (
	"fmt"
	"os"

	"github.com/HighPerformanceWithGo/7-metaprogramming-in-go/clitooling/cmd"
)

func main() {

	if err := cmd.DateCommand.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
