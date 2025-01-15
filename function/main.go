package main

import (
	"fmt"
)

func main() {
	// "function string"
	/* printMyResult("saya sedang")
	printMyResult("belajar golang")*/

	// "function int"
	result := add(10, 20)
	fmt.Println(result)

	//mengembalikan nilai lebih dari satu
	luas, keliling := calculate(10, 20)
	fmt.Println(luas)
	fmt.Println(keliling)

	/* sentence := printMyResult("saya sedang")
	fmt.Println(sentence)*/
}

// pemanggilan yang berada di satu package main dengan awalan huruf kecil
/* func add(number int, numberTwo int) int {
result := number + numberTwo
return result */

// lebih clean code langsung dilakukan pemangilan "return" tanpa harus dideklarasikan ulang
func add(number int, numberTwo int) int {
	return number + numberTwo
}

// mengembalikan nilai lebih dari satu
/*func calculate(panjang int, lebar int) (int, int) {
luas := panjang * lebar
keliling := 2 * (panjang * lebar)

return luas, keliling*/

// function dengan predefined return value
func calculate(panjang int, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)

	return
}

// func printMyResult(sentence string) {
// 	fmt.Println(sentence)

// func printMyResult(sentence string) string {
// 	newSentence := sentence + "belajar golang"
// 	return newSentence // mengembalikan nilai dengan "return" yang menghasilkan output
// }
