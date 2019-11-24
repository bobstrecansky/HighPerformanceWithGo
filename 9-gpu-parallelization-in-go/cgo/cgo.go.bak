package main

/*
 #include <stdio.h>

 const char* hello_gophers() {
	 return "Hello Gophers!";
 }
*/
import "C"
import "fmt"

func main() {
	fmt.Println(C.GoString(C.hello_gophers()))
}
