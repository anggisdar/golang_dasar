package main

import (
	"fmt"
)

func main() {
	// "function string"
	// printMyResult("saya sedang")
	// printMyResult("belajar golang")

	// "function int"
	result := add(10, 20)
	fmt.Println(result)

	// sentence := printMyResult("saya sedang")
	// fmt.Println(sentence)
}

// pemanggilan yang berada di satu package main dengan awalan huruf kecil
func add(number int, numberTwo int) int {
	result := number + numberTwo
	return result
}

// func printMyResult(sentence string) {
// 	fmt.Println(sentence)

// func printMyResult(sentence string) string {
// 	newSentence := sentence + "belajar golang"
// 	return newSentence // mengembalikan nilai dengan "return" yang menghasilkan output
// }
