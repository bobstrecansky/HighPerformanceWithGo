package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"contrib.go.opencensus.io/exporter/stackdriver"
	"go.opencensus.io/trace"
)

const server = ":1234"

func init() {
	exporter, err := stackdriver.NewExporter(stackdriver.Options{
		ProjectID: os.Getenv("GOOGLE_CLOUD_PROJECT"),
	})
	if err != nil {
		log.Fatal("Can't initialize GOOGLE_CLOUD_PROJECT enviornment variable", err)
	}
	trace.RegisterExporter(exporter)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}

func sleep(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	_, span := trace.StartSpan(ctx, "sleep")
	defer span.End()
	time.Sleep(1 * time.Second)
	fmt.Fprintln(w, "Done Sleeping")
}

func githubRequest(ctx context.Context, w http.ResponseWriter, r *http.Request) {
	_, span := trace.StartSpan(ctx, "githubRequest")
	defer span.End()
	res, err := http.Get("https://github.com")
	if err != nil {
		log.Fatal(err)
	}
	res.Body.Close()
	fmt.Fprintln(w, "Request to https://github.com completed with a status of: ", res.Status)
	span.End()
}

func main() {
	h := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		ctx, span := trace.StartSpan(context.Background(), "function/main")
		defer span.End()
		githubRequest(ctx, w, r)
		sleep(ctx, w, r)
	})

	http.Handle("/", h)

	log.Printf("serving at : %s", server)
	err := http.ListenAndServe(server, nil)
	if err != nil {
		log.Fatal("Couldn't start HTTP server: %s", err)
	}

}
