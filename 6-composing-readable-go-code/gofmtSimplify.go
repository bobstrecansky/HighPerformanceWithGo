package main

import "fmt"

func main() {

	var tmp = []int{1, 2, 3}
	b := tmp[1:len(tmp)]
	fmt.Println(b)
	for i, _ := range tmp {
		fmt.Println(tmp[i])
	}

}
