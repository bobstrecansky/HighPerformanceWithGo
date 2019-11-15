package main

import (
	"math/rand"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {

	saturation := prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "saturation",
		Help: "A gauge of the saturation golden singnal",
	})

	requests := prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "requests",
			Help: "A counter for the requests golden signal",
		},
		[]string{"code", "method"},
	)

	latency := prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name:    "latency",
			Help:    "A histogram of latencies for the latency golden signal",
			Buckets: []float64{.25, .5, 1, 2.5, 5, 10},
		},
		[]string{"handler", "method"},
	)

	goldenSignalHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		sleepTime := rand.Intn(100)
		//sleepTime := rand.Intn(10000)
		time.Sleep(time.Duration(sleepTime) * time.Millisecond)

		if sleepTime%4 == 0 {
			http.Error(w, "Not Found", http.StatusNotFound)
		} else if sleepTime%5 == 0 {
			http.Error(w, "Server Error", http.StatusInternalServerError)
		} else {
			w.Write([]byte("Golden Signal Handler"))
		}
	})

	goldenSignalChain := promhttp.InstrumentHandlerInFlight(saturation,
		promhttp.InstrumentHandlerDuration(latency.MustCurryWith(prometheus.Labels{"handler": "signals"}),
			promhttp.InstrumentHandlerCounter(requests, goldenSignalHandler),
		),
	)

	prometheus.MustRegister(saturation, requests, latency)

	http.Handle("/metrics", promhttp.Handler())
	http.Handle("/signals", goldenSignalChain)
	http.ListenAndServe(":1234", nil)
}
