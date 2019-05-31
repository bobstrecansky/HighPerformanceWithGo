package main

import (
	"fmt"
	"net/http"
	"sync"
	"time"
)

func retrieve(url string, wg *sync.WaitGroup) {
	// WaitGroup Counter-- when Goroutine is finished
	defer wg.Done()
	start := time.Now()
	res, err := http.Get(url)
	end := time.Since(start)
	if err != nil {
		panic(err)
	}
	// Print the status code from the response
	fmt.Println(url, res.StatusCode, end)

}

func main() {
	var wg sync.WaitGroup
	var urls = []string{"https://godoc.org", "https://www.packtpub.com", "https://kubernetes.io/"}
	for i := range urls {
		// Increment WaitGroup Counter when new Goroutine is called
		wg.Add(1)
		go retrieve(urls[i], &wg)
	}
	// Wait for the collection of Goroutines to finish
	wg.Wait()
}
