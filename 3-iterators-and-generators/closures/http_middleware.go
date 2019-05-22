package main

import (
	"fmt"
	"log"
	"net/http"
)

// Checks for proper creds for admin
func adminCheck(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("user") != "admin" {
			fmt.Fprintln(w, "Not an Admin")
			http.NotFound(w, r)
		}
	})
}

// Sets a HTTP 418 (I'm a Teapot) status code for the response
func newStatusCode(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Status Code Set")
		w.WriteHeader(http.StatusTeapot)
		h.ServeHTTP(w, r)
	})
}

// Adds a header, Foo:Bar
func addHeader(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Foo", "Bar")
		log.Println("Adding Foo Header")
		h.ServeHTTP(w, r)
	})
}

// Writes a HTTP Response
func writeResponse(w http.ResponseWriter, r *http.Request) {
	log.Println("Writing HTTP Response")
	fmt.Fprintln(w, "Hello PerfGo!")
}

// Wrap the middleware together
func main() {
	handler := http.HandlerFunc(writeResponse)
	http.Handle("/", newStatusCode(addHeader(handler)))
	http.Handle("/onlyHeader", addHeader(handler))
	http.Handle("/onlyStatus", newStatusCode(handler))
	http.Handle("/admin", adminCheck(handler))
	http.ListenAndServe(":1234", nil)
}
