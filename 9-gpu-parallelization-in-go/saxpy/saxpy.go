package main

import "fmt"

//#cgo CFLAGS: -I.
//#cgo LDFLAGS: -L. -lsaxpy
//#cgo LDFLAGS: -lcudart
//#include <saxpy.h>
import "C"

func main() {
	fmt.Printf("Invoking cuda library...\n")
	fmt.Println("Done ", C.matrix_multiplication())
}
