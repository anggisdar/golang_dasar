package main

import "fmt"

func main() {
	score := 80
	var grade string

	if score <= 50 {
		grade = "E"
	} else if score <= 60 {
		grade = "c"
	} else if score < 70 {
		grade = "C"
	} else {
		grade = "NULL"
	}

	fmt.Println(grade)
}
