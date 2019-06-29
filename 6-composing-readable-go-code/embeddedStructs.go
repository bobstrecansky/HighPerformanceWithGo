package main

import "fmt"

func main() {

	type Utencils struct {
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
		Utencils
		Appliances
	}

	bobKitchen := new(Kitchen)
	bobKitchen.Utencils.fork = "3 prong"
	bobKitchen.Utencils.knife = "dull"
	bobKitchen.Utencils.spoon = "deep"
	bobKitchen.Appliances.stove = "6 burner"
	bobKitchen.Appliances.dishwasher = "3 rack"
	bobKitchen.Appliances.oven = "self cleaning"

	fmt.Printf("%+v\n", bobKitchen)
}
