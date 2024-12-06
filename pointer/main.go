package main

import "fmt"

func main() {
	var parts = [5]string{"mb", "subboard", "mainFpc", "receiver", "speaker"}
	partsPointer := &parts
	fmt.Println("array as pointer", partsPointer)
}
