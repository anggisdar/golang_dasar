package main

import (
	"fmt"
)

func main() {

	food := []map[string]string{

		{"name": "bakso", "daging": "Ayam"},
		{"name": "sate", "daging": "Sapi"},
		{"name": "siomay", "daging": "ikan"},
	}

	for _, foods := range food {
		fmt.Println(foods)
	}

}
