package comparison

import (
	"encoding/json"
	"fmt"
	"log"
)

type CarData struct {
	Cylinders int
	Brand     string
	Mpg       float64
}

func CarJSON() {
	honda := CarData{Cylinders: 4, Brand: "Toyota", Mpg: 42.6}
	var carDataJson []byte
	carDataJson, err := json.Marshal(honda)
	if err != nil {
		log.Println(err)
	}
	fmt.Sprintf(string(carDataJson))
}
