package main

import (
	"fmt"
)

func main() {

	title := "Golang the best language"

	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("index :", index, " letter :", string(letter))
		}
	}

	for index, letter := range title {
		letterString := string(letter)

		if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "i" {
			fmt.Println("index :", index, "letter :", string(letter))

		}
	}
}
