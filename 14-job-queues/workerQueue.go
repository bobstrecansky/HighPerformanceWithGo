package main

import (
	"log"
	"net/http"
)

const queueSize = 50
const workers = 10
const port = "1234"

type job struct {
	name    string
	payload string
}

func runJob(id int, individualJob job) {
	log.Printf("Worker %d: Completed: %s with payload %s", id, individualJob.name, individualJob.payload)
}

func main() {
	jobQueue := make(chan job, queueSize)
	for i := 1; i <= workers; i++ {
		go func(i int) {
			for j := range jobQueue {
				runJob(i, j)
			}
		}(i)

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		submittedJob := job{r.FormValue("name"), r.FormValue("payload")}
		jobQueue <- submittedJob
	})

	http.ListenAndServe(":"+port, nil)
}
