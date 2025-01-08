package main

import (
	"fmt"
)

func main() {
	var myMap map[string]int
	myMap = map[string]int{}

	myMap["go"] = 2
	fmt.Println(myMap)

}
