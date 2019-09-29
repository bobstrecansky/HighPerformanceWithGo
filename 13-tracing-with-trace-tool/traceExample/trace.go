package main

import (
	"log"
	"os"
	"runtime/trace"
)

func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	ch := make(chan string)
	go func() {
		ch <- "Hi Gophers"
	}()
	<-ch
	log.Printf("Trace Completed")
}
