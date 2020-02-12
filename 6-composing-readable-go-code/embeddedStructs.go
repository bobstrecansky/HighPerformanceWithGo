package main

import "fmt"

func main() {

	type Utensils struct {
		fork  string
		spoon string
		knife string
	}

	type Appliances struct {
		stove      string
		dishwasher string
		oven       string
	}

	type Kitchen struct {
		Utensils
		Appliances
	}

	bobKitchen := new(Kitchen)
	bobKitchen.Utensils.fork = "3 prong"
	bobKitchen.Utensils.knife = "dull"
	bobKitchen.Utensils.spoon = "deep"
	bobKitchen.Appliances.stove = "6 burner"
	bobKitchen.Appliances.dishwasher = "3 rack"
	bobKitchen.Appliances.oven = "self cleaning"

	fmt.Printf("%+v\n", bobKitchen)
}
