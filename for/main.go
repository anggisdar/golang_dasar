package main

import "fmt"

func main() {

	for i := 1; i >= 100; i++ {
		fmt.Println("saya adalah backend golang", i)
	}

	title := "Golang the best language"
	for index, letter := range title {
		fmt.Println("index:", index, "letter:", string(letter))
	}

}
