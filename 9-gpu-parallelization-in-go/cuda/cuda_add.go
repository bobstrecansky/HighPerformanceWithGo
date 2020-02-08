package main

import "fmt"

//#cgo CFLAGS: -I.
//#cgo LDFLAGS: -L. -ltest
//#cgo LDFLAGS: -lcudart
//#include <cuda_add.h>
import "C"

func main() {
	fmt.Printf("Invoking cuda library...\n")
	fmt.Println("Done ", C.test_add())
}
