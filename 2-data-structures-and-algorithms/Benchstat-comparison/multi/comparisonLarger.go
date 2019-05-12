package comparison_larger

import (
	"encoding/json"
	"fmt"
	"log"
)

type MoreCarData struct {
	Cylinders int
	Origin    string
	Mpg       float64
	Doors     int
	FuelType  string
}

func CarJSON() {
	sedan := MoreCarData{Cylinders: 4, Origin: "Japan", Mpg: 42.6, Doors: 4, FuelType: "Petrol"}
	convertible := MoreCarData{Cylinders: 6, Origin: " USA", Mpg: 28, Doors: 2, FuelType: "Diesel"}
	cars := []MoreCarData{sedan, convertible}
	for _, car := range cars {
		var carDataJson []byte
		carDataJson, err := json.Marshal(car)
		if err != nil {
			log.Println(err)
		}
		fmt.Sprintf(string(carDataJson))
	}
}
