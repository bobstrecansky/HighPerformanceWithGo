package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	go func() {
		for {
			var r runtime.MemStats
			runtime.ReadMemStats(&r)
			fmt.Println("\nTime: ", time.Now())
			fmt.Println("Runtime MemStats Sys: ", r.Sys)
			fmt.Println("Runtime Heap Allocation: ", r.HeapAlloc)
			fmt.Println("Runtime Heap Idle: ", r.HeapIdle)
			fmt.Println("Runtime Heap In Use: ", r.HeapInuse)
			fmt.Println("Runtime Heap HeapObjects: ", r.HeapObjects)
			fmt.Println("Runtime Heap Released: ", r.HeapReleased)
			time.Sleep(5 * time.Second)
		}
	}()

	fmt.Println("Hello Gophers")
	time.Sleep(11 * time.Second)
}
