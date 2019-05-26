package main

import (
	"fmt"
	"net/http"
)

// Checks for a “user:admin” header, proper credentials for the admin path
func adminCheck(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if r.Header.Get("user") != "admin" {
			http.Error(w, "Not Authorized", 401)
			return
		}
		fmt.Fprintln(w, "Admin Portal")
		h.ServeHTTP(w, r)
	})
}

// Sets a HTTP 418 (I'm a Teapot) status code for the response
func newStatusCode(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusTeapot)
		h.ServeHTTP(w, r)
	})
}

// Adds a header, Foo:Bar
func addHeader(h http.Handler) http.HandlerFunc {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Foo", "Bar")
		h.ServeHTTP(w, r)
	})
}

// Writes a HTTP Response
func writeResponse(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello PerfGo!")
}

// Wrap the middleware together
func main() {
	handler := http.HandlerFunc(writeResponse)
	http.Handle("/", addHeader(newStatusCode(handler)))
	http.Handle("/onlyHeader", addHeader(handler))
	http.Handle("/onlyStatus", newStatusCode(handler))
	http.Handle("/admin", adminCheck(handler))
	http.ListenAndServe(":1234", nil)
}
