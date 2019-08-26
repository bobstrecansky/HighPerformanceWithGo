package main

import (
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.Handle("/promMetrics", promhttp.Handler())
	http.ListenAndServe(":1234", nil)
}
