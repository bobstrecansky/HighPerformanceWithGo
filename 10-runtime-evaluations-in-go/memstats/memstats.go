package main

import (
	"fmt"
	"net/http"
	"runtime"
)

func memStats(w http.ResponseWriter, r *http.Request) {
	var memStats runtime.MemStats
	runtime.ReadMemStats(&memStats)
	fmt.Fprintln(w, "Alloc:", memStats.Alloc)
	fmt.Fprintln(w, "Total Alloc:", memStats.TotalAlloc)
	fmt.Fprintln(w, "Sys:", memStats.Sys)
	fmt.Fprintln(w, "Lookups:", memStats.Lookups)
	fmt.Fprintln(w, "Mallocs:", memStats.Mallocs)
	fmt.Fprintln(w, "Frees:", memStats.Frees)

	fmt.Fprintln(w, "Heap Alloc:", memStats.HeapAlloc)
	fmt.Fprintln(w, "Heap Sys:", memStats.HeapSys)
	fmt.Fprintln(w, "Heap Idle:", memStats.HeapIdle)
	fmt.Fprintln(w, "Heap In Use:", memStats.HeapInuse)
	fmt.Fprintln(w, "Heap Released:", memStats.HeapReleased)
	fmt.Fprintln(w, "Heap Objects:", memStats.HeapObjects)

	fmt.Fprintln(w, "Stack In Use:", memStats.StackInuse)
	fmt.Fprintln(w, "Stack Sys:", memStats.StackSys)
	fmt.Fprintln(w, "MSpanInuse:", memStats.MSpanInuse)
	fmt.Fprintln(w, "MSpan Sys:", memStats.MSpanSys)
	fmt.Fprintln(w, "MCache In Use:", memStats.MCacheInuse)
	fmt.Fprintln(w, "MCache Sys:", memStats.MCacheSys)
	fmt.Fprintln(w, "Buck Hash Sys:", memStats.BuckHashSys)

	fmt.Fprintln(w, "EnableGC:", memStats.EnableGC)
	fmt.Fprintln(w, "GCSys:", memStats.GCSys)
	fmt.Fprintln(w, "Other Sys:", memStats.OtherSys)
	fmt.Fprintln(w, "Next GC:", memStats.NextGC)
	fmt.Fprintln(w, "Last GC:", memStats.LastGC)
	fmt.Fprintln(w, "Num GC:", memStats.NumGC)
	fmt.Fprintln(w, "Num Forced GC:", memStats.NumForcedGC)

	fmt.Fprintln(w, "Pause Total NS:", memStats.PauseTotalNs)
	fmt.Fprintln(w, "Pause Ns:", memStats.PauseNs)
	fmt.Fprintln(w, "Pause End:", memStats.PauseEnd)
	fmt.Fprintln(w, "GCCPUFraction:", memStats.GCCPUFraction)
	fmt.Fprintln(w, "BySize Size:", memStats.BySize)
}

func main() {
	http.HandleFunc("/", memStats)
	http.ListenAndServe(":1234", nil)
}
