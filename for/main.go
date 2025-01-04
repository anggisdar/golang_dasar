package main

import "fmt"

func main() {

	for i := 1; i >= 100; i++ {
		fmt.Println("saya adalah backend golang", i)
	}

	// title := "Golang the best language"
	// for index, letter := range title {
	// 	fmt.Println("index:", index, "letter:", string(letter))
	// }

	// // tanpa index
	// title := "Golang the best language"

	// for _, letter := range title {
	// 	fmt.Println("letter:", string(letter))
	// }

	title := "Golang the best language"

	for index, letter := range title {
		letterString := string(letter)

		if letterString == "a" || letterString == "i" || letterString == "u" || letterString == "e" || letterString == "0" {
			fmt.Println("index:", index, "letter:", string(letter))
		}

		switch letterString {
		case "a", "i", "u", "e", "o":
			fmt.Println("index:", index, "letter:", string(letter))
		}
	}
}
