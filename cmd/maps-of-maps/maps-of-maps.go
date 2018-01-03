package main

import (
	"fmt"
)

func main() {
	vehicules := map[string]map[string]int{
		"Car": {
			"Wheels":  4,
			"Seats":   5,
			"Airbags": 7,
		},
		"Bus": {
			"Wheels": 4,
			"Seats":  30,
		},
		"Bike": {
			"Wheels": 2,
			"Seats":  2,
		},
	}

	for vehicule, caracteristics := range vehicules {
		fmt.Printf("%v, ", vehicule)
		for caracteristic, value := range caracteristics {
			fmt.Printf("%v: %v, ", caracteristic, value)
		}
		fmt.Printf("\b\b\n")
	}
}
