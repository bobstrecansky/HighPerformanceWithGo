### Go Standard Library Sorting Algorithm Benchmark

This example is a benchmark of the Go standard library search algorithm.  It downloads the Go source code, sets the GORROT_BOOTSTRAP as a toolchain to run the benchmarking tests, and executes the TestSearch test, which includes a grouping of benchmarks for different sorting algorithms.  An example output from this command is shown below.

```
$ bash benchmarksort.bash 
Cloning into 'go'...
remote: Enumerating objects: 156, done.
remote: Counting objects: 100% (156/156), done.
remote: Compressing objects: 100% (124/124), done.
remote: Total 381095 (delta 45), reused 46 (delta 29), pack-reused 380939
Receiving objects: 100% (381095/381095), 183.19 MiB | 33.72 MiB/s, done.
Resolving deltas: 100% (302878/302878), done.
Building Go cmd/dist using /usr/lib/golang.
Building Go toolchain1 using /usr/lib/golang.
Building Go bootstrap cmd/go (go_bootstrap) using Go toolchain1.
Building Go toolchain2 using go_bootstrap and Go toolchain1.
Building Go toolchain3 using go_bootstrap and Go toolchain2.
Building packages and commands for linux/amd64.
---
Installed Go for linux/amd64 in /home/bob/git/HighPerformanceWithGo/1-introduction/go
Installed commands in /home/bob/git/HighPerformanceWithGo/1-introduction/go/bin
=== RUN   TestSearch
--- PASS: TestSearch (0.00s)
=== RUN   TestSearchEfficiency
--- PASS: TestSearchEfficiency (0.00s)
=== RUN   TestSearchWrappers
--- PASS: TestSearchWrappers (0.00s)
=== RUN   TestSearchWrappersDontAlloc
--- SKIP: TestSearchWrappersDontAlloc (0.00s)
    search_test.go:135: skipping; GOMAXPROCS>1
=== RUN   TestSearchExhaustive
--- PASS: TestSearchExhaustive (0.00s)
goos: linux
goarch: amd64
pkg: sort
BenchmarkSearchWrappers-8       	 8813284	       128 ns/op
BenchmarkSortString1K-8         	    7394	    155458 ns/op
BenchmarkSortString1K_Slice-8   	    7021	    147020 ns/op
BenchmarkStableString1K-8       	    5936	    202265 ns/op
BenchmarkSortInt1K-8            	   15049	     80777 ns/op
BenchmarkStableInt1K-8          	   12108	     92240 ns/op
BenchmarkStableInt1K_Slice-8    	   14799	     88723 ns/op
BenchmarkSortInt64K-8           	     152	   7538665 ns/op
BenchmarkSortInt64K_Slice-8     	     170	   7135752 ns/op
BenchmarkStableInt64K-8         	     146	   8137775 ns/op
BenchmarkSort1e2-8              	   25747	     48348 ns/op
BenchmarkStable1e2-8            	   12775	     94729 ns/op
BenchmarkSort1e4-8              	     132	   8910633 ns/op
BenchmarkStable1e4-8            	      39	  26609879 ns/op
BenchmarkSort1e6-8              	       1	1372163197 ns/op
BenchmarkStable1e6-8            	       1	5366714895 ns/op
PASS
ok  	sort	36.883s
```
