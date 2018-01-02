package main

import (
	"fmt"
)

func main() {

	elements := make(map[string]string)

	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"
	elements["Be"] = "Beryllium"
	elements["B"] = "Boron"
	elements["C"] = "Carbon"
	elements["N"] = "Nitrogen"
	elements["O"] = "Oxygen"
	elements["F"] = "Fluorine"
	elements["Ne"] = "Neon"

	// Print elements
	fmt.Printf("%v\n", elements)
	// Iterate over elements
	for key, value := range elements {
		fmt.Printf("Key: %v, value: %v\n", key, value)
	}
	// Find the value of a key
	var key string
	key = "N"
	if name, ok := elements[key]; ok {
		fmt.Printf("Key: %v, value: %v\n", key, name)
	} else {
		fmt.Printf("Key: %v was not found\n", key)
	}
	key = "Z"
	if name, ok := elements[key]; ok {
		fmt.Printf("Key: %v, value: %v\n", key, name)
	} else {
		fmt.Printf("Key: %v was not found\n", key)
	}

}
