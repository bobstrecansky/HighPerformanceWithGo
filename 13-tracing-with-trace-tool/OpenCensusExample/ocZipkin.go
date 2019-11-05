package main

import (
	"context"
	"log"
	"net/http"
	"time"

	"contrib.go.opencensus.io/exporter/zipkin"
	"go.opencensus.io/trace"

	"github.com/go-redis/redis"
	openzipkin "github.com/openzipkin/zipkin-go"
	zipkinHTTP "github.com/openzipkin/zipkin-go/reporter/http"
)

func tracingServer() {
	l, err := openzipkin.NewEndpoint("oc-zipkin", "192.168.1.5:5454")
	if err != nil {
		log.Fatalf("Couldn't create zipkinEndpoint: %v", err)
	}
	r := zipkinHTTP.NewReporter("http://0.0.0.0:9411/api/v2/spans")
	z := zipkin.NewExporter(r, l)
	trace.RegisterExporter(z)
	trace.ApplyConfig(trace.Config{DefaultSampler: trace.AlwaysSample()})
}

func makeRequest(ctx context.Context, url string) string {
	log.Printf("Retrieving URL")
	_, span := trace.StartSpan(ctx, "httpRequest")
	defer span.End()
	res, _ := http.Get(url)
	defer res.Body.Close()
	time.Sleep(100 * time.Millisecond)
	log.Printf("URL Response : %s", res.Status)

	span.Annotate([]trace.Attribute{
		trace.StringAttribute("URL Response Code", res.Status),
	}, "HTTP Response Status Code:"+res.Status)

	time.Sleep(50 * time.Millisecond)
	return res.Status
}

func writeToRedis(ctx context.Context, key string, value string) {
	log.Printf("Writing to Redis")
	_, span := trace.StartSpan(ctx, "redisWrite")
	defer span.End()
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "",
		DB:       0,
	})

	err := client.Set(key, value, 0).Err()
	if err != nil {
		panic(err)
	}
}

func main() {
	tracingServer()
	ctx, span := trace.StartSpan(context.Background(), "main")
	defer span.End()

	for i := 0; i < 10; i++ {
		url := "https://golang.org/"
		respStatus := makeRequest(ctx, url)
		writeToRedis(ctx, url, respStatus)
	}
}
