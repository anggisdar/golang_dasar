package main

import (
	"errors"
	"fmt"
)

func main() {
	//array
	value := []int{80, 70, 60, 50, 40}
	total := sum(value)
	fmt.Println(total)

	//error handling
	result, err := calculate(10, 2, "*")
	if err != nil {
		fmt.Println("Error", err)
		fmt.Println(err.Error())
	}

	fmt.Println(result)
}

// for - range
func sum(numbers []int) int {
	var total int
	for _, number := range numbers {
		total = total + number
	}
	//return
	return total
}

// fungsi error
func calculate(number, numberTwo int, operation string) (int, error) {
	var result int
	var errorResult error

	//switch
	switch operation {
	case "+":
		result = number + numberTwo
	case "-":
		result = number - numberTwo
	case "/":
		result = number / numberTwo
	case "*":
		result = number * numberTwo
	default:
		errorResult = errors.New("data tidak ditemukan")
	}

	return result, errorResult

}
