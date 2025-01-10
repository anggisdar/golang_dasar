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

	var mymap map[string]int
	mymap = map[string]int{}

	mymap["ruby"] = 9
	mymap["javascript"] = 8

	fmt.Println(mymap)
}
