package main

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/", promhttp.Handler())
	port := ":2112"
	fmt.Println("Prometheus Handler listening on port ", port)
	http.ListenAndServe(port, nil)
}
