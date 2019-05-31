package main

import (
	"context"
	"fmt"
	"runtime"
	"time"

	"golang.org/x/sync/semaphore"
)

func main() {
	ctx := context.Background()
	var (
		sem    = semaphore.NewWeighted(int64(runtime.GOMAXPROCS(0)))
		result = make([]string, 5)
	)
	for i := range result {
		if err := sem.Acquire(ctx, 1); err != nil {
			break
		}
		go func(i int) {
			defer sem.Release(1)
			time.Sleep(100 * time.Millisecond)
			result[i] = "Semaphores are Cool \n"
		}(i)
	}
	if err := sem.Acquire(ctx, int64(runtime.GOMAXPROCS(0))); err != nil {
		fmt.Println("Error acquiring semaphore")
	}
	fmt.Println(result)
}
