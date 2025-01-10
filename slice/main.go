package main

import "fmt"

func main() {
	// ini slice
	var spareparts = []string{"Mainboard", "Subboard", "mainFpc", "Battery"}

	for _, sparepart := range spareparts {
		fmt.Println(sparepart)
	}

	//slice append
	var console []string
	console = append(console, "psp")

	fmt.Println(console)
}
