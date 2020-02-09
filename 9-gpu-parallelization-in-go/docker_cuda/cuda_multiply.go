package main

import (
	"fmt"
	"time"
)

//#cgo CFLAGS: -I.
//#cgo LDFLAGS: -L. -lmultiply
//#cgo LDFLAGS: -lcudart
//#include <cuda_multiply.h>
import "C"

func main() {
	fmt.Printf("Invoking cuda library...\n")
	start := time.Now()
	C.cuda_multiply()
	elapsed := time.Since(start)
	fmt.Println("\nCuda Execution took", elapsed)
}
