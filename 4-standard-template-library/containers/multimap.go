package main

import (
	"fmt"

	"github.com/jwangsadinata/go-multimap/slicemultimap"
)

type cars []struct {
	year  int
	style string
}

func main() {

	newCars := cars{{2019, "convertible"}, {1966, "fastback"}, {2019, "SUV"}, {1920, "truck"}}
	multimap := slicemultimap.New()

	for _, car := range newCars {
		multimap.Put(car.year, car.style)
	}

	for _, style := range multimap.KeySet() {
		color, _ := multimap.Get(style)
		fmt.Printf("%v: %v\n", style, color)
	}
}
