package main

import (
	"errors"
	"fmt"
)

func main() {
	var contohError error = errors.New("Terjadi error")
	fmt.Println(contohError.Error())
}
