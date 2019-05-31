package main

import "fmt"

const top = 100

func main() {
	err := callbackLoop(top, func(n int) error {
		fmt.Println(n)
		return nil
	})
	if err != nil {
		panic(err)
	}
}

func callbackLoop(top int, callback func(n int) error) error {
	for i := 0; i < top; i++ {
		err := callback(i)
		if err != nil {
			return err
		}
	}
	return nil
}
