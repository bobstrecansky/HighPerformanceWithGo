package main

import (
	"io"
	"net/http"
	_ "net/http/pprof"
	"time"
)

func main() {
	Handler := func(w http.ResponseWriter, req *http.Request) {
		time.Sleep(5 * time.Second)
		io.WriteString(w, "Network Trace Profile Test")
	}

	http.HandleFunc("/", Handler)
	http.ListenAndServe(":1234", nil)
}
