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
	fmt.Println("Done ", C.cuda_multiply())
	elapsed := time.Since(start)
	fmt.Println("C Execution took", elapsed)
}
