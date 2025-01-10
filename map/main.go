package main

import (
	"fmt"
)

func main() {
	// var myMap map[string]int
	// myMap = map[string]int{}

	// myMap["go"] = 2
	// fmt.Println(myMap)

	myMap := map[string]string{
		"golang":     "static language",
		"javascript": "dynamic language",
	}

	for key, value := range myMap {
		fmt.Println("key:", key, "value:", value)
	}

	fmt.Println("============")

	delete(myMap, "javascript")

	for key, value := range myMap {
		fmt.Println("key:", key, "value:", value)

	}

	// menampilkan nilai yang tidak ada dengan bool
	value, isAvailable := myMap["php"]
	fmt.Println(isAvailable)
	fmt.Println(value)

}
