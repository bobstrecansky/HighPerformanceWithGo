package main

import (
	"fmt"
	"net/http"

	_ "net/http/pprof"
	"runtime"
	"time"
)

func main() {
	http.HandleFunc("/leak", leakyAbstraction)
	http.ListenAndServe("localhost:6060", nil)
}

func leakyAbstraction(w http.ResponseWriter, r *http.Request) {
	ch := make(chan string, 15)
	for {
		//for i := 0; i < 10; i++ {
		fmt.Fprintln(w, "Number of Goroutines: ", runtime.NumGoroutine())
		go func() { ch <- wait() }()
	}
}

func wait() string {
	time.Sleep(5 * time.Microsecond)
	return "Hello Gophers!"
}
