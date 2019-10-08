package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"sync"

	_ "github.com/go-sql-driver/mysql"
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
	db, err := sql.Open("mysql", "root:admin@0.0.0.0/mysql")
	if err != nil {
		panic(err)
	}
	defer db.Close()

}

func createTable() {
	db, err := sql.Open("mysql", "root:admin@0.0.0.0/mysql")
	if err != nil {
		panic(err)
	}
	defer db.Close()

}
func main() {
	jobQueue := make(chan job, queueSize)
	wg := &sync.WaitGroup{}
	wg.Add(workers)

	for i := 1; i <= workers; i++ {
		go func(i int) {
			defer wg.Done()
			for j := range jobQueue {
				runJob(i, j)
			}
		}(i)

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		submittedJob := job{r.FormValue("name"), r.FormValue("payload")}
		jobQueue <- submittedJob
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Worker %d: Completed: %s with payload %s", submittedJob.name, submittedJob.payload)
	})

	http.ListenAndServe(":"+port, nil)
	wg.Wait()
}
