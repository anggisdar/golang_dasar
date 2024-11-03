package main

import "fmt"

func main() {
	var a = 10

	fmt.Println(a > 10 && a < 5)
	fmt.Println(a > 10 || a < 5)
	fmt.Println(!(a > 10 && a < 5))

}
