package main

import (
	_ "expvar"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const maxWorkers = 10

type job struct {
	name     string
	duration time.Duration
}

func doWork(id int, j job) {
	fmt.Printf("worker%d: started %s, working for %fs\n", id, j.name, j.duration.Seconds())
	time.Sleep(j.duration)
	fmt.Printf("worker%d: completed %s!\n", id, j.name)
}

func main() {
	// channel for jobs
	jobs := make(chan job)

	// start workers
	wg := &sync.WaitGroup{}
	wg.Add(maxWorkers)
	for i := 1; i <= maxWorkers; i++ {
		go func(i int) {
			defer wg.Done()

			for j := range jobs {
				doWork(i, j)
			}
		}(i)
	}

	// add jobs
	for i := 0; i < 100; i++ {
		name := fmt.Sprintf("job-%d", i)
		duration := time.Duration(rand.Intn(1000)) * time.Millisecond
		fmt.Printf("adding: %s %s\n", name, duration)
		jobs <- job{name, duration}
	}
	close(jobs)

	// wait for workers to complete
	wg.Wait()
}
