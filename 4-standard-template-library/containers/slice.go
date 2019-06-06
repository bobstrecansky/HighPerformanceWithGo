package main

import "fmt"

// Remove i indexed item in slice
func remove(s []string, i int) []string {
	copy(s[i:], s[i+1:])
	return s[:len(s)-1]
}

func main() {
	slice := []string{"foo", "bar", "baz"}     // create a slice
	slice = append(slice, "tri")               // append a slice
	fmt.Println("Appended Slice: ", slice)     // print slice [foo, bar baz, tri]
	slice = remove(slice, 2)                   // remove slice item #2 (baz)
	fmt.Println("After Removed Item: ", slice) // print slice [foo, bar, tri]
}
