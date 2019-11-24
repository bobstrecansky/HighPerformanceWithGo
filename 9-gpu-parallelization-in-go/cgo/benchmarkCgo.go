package benchmark

/*
 #include <stdio.h>

 const char* hello_gophers() {
	 return "Hello Gophers!";
 }
*/
import "C"
import "fmt"

func CgoPrint(n int) {
	for i := 0; i < n; i++ {
		fmt.Sprintf(C.GoString(C.hello_gophers()))
	}
}

func GoPrint(n int) {
	for i := 0; i < n; i++ {
		fmt.Sprintf("Hello Gophers!")
	}
}
