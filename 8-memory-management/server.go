package main

import (
	"io"
	"net/http"
)

func main() {
	Handler := func(w http.ResponseWriter, req *http.Request) {
		io.WriteString(w, "Memory Management Test")
	}

	http.HandleFunc("/", Handler)
	http.ListenAndServe(":1234", nil)
}
