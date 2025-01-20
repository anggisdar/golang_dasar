package main

import "fmt"

func main() {

	/* numberA := 5
	numberB := &numberA // & alamat memory dalam bentuk heksadesimal

	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB) // * mengembalikan nilai "proses derefrencing"

	*numberB = 10 // mendeklarasikan ulang * mengembalikan nilai "proses derefrencing"

	 fmt.Println(numberA)*/

	// refence & direference
	var numberA int = 5
	var numberB *int = &numberA // reference

	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB) // dereference

	numberA = 20

	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB)

	// dengan array
	var parts = [5]string{"mb", "subboard", "mainFpc", "receiver", "speaker"}
	partsPointer := &parts
	fmt.Println("array as pointer", partsPointer)

}
