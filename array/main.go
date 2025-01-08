package main

import "fmt"

// ini array
func main() {
	var fruits [4]string
	fruits[0] = "Leci"
	fruits[1] = "Manggo"
	fruits[2] = "Grep"
	fruits[3] = "Banana"

	for _, fruit := range fruits {
		fmt.Println(fruit)
	}

	// ini slice
	var spareparts = []string{"Mainboard", "Subboard", "mainFpc", "Battery"}

	for _, sparepart := range spareparts {
		fmt.Println(sparepart)
	}

	//slice append
	var console []string
	console = append(console, "psp")

	fmt.Println(console)

	// ini byte
	byteArray := []byte("ini byte")
	for _, byteArr := range byteArray {
		fmt.Println(byteArr)
	}

}
