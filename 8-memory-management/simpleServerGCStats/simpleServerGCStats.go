package main

import (
	"fmt"
	"io"
	"net/http"
	"runtime/debug"
	"time"
)

func main() {
	Handler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Memory Management Test")
	}

	go func() {
		for {
			gc := new(debug.GCStats)
			fmt.Println("\nTime: ", time.Now())
			fmt.Println("Last Garbage Collection: ", gc.LastGC)
			fmt.Println("Number of Garbage Collections: ", gc.NumGC)
			fmt.Println("Total Pause for all Collections: ", gc.PauseTotal)
			fmt.Println("Pause: ", gc.Pause)
			fmt.Println("PauseEnd: ", gc.PauseEnd)
			fmt.Println("PauseQuantiles: ", gc.PauseQuantiles)
			debug.FreeOSMemory()
			time.Sleep(5 * time.Second)
		}
	}()

	http.HandleFunc("/", Handler)
	http.ListenAndServe(":1234", nil)
}
