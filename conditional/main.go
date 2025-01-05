package main

import "fmt"

func main() {
	score := 80
	var grade string
	var narasi string

	// if score > 80 {
	// 	grade = "AMAT BAIK"
	// } else if score > 70 {
	// 	grade = "BAIK"
	// } else if score >= 60 {
	// 	grade = "CUKUP"
	// } else {
	// 	grade = "TIDAK LULUS"
	// }

	fmt.Println(grade)

	const address string = "nusantara"

	switch address {
	case "nusantara":
		fmt.Println("Address is Nusantara")
	case "jakarta":
		fmt.Println("Address iS Jakarta")
	default:
		fmt.Println("Address is not Nusantara or Jakarta")
	}

	switch {
	case score < 0 || score > 100:
		grade = "NULL"
		narasi = "nilai tidak ada didalam range"
	case score > 80:
		grade = "A"
		narasi = "Nilai Kamu sangat baik"
	case score > 70:
		grade = "B"
		narasi = "Nilai kamu baik"
	case score >= 60:
		grade = "C"
		narasi = "nilai kamu cukup"
	default:
		grade = "D"
		narasi = "maaf kamu belum bisa lanjut"
	}

	fmt.Println("value", score)
	fmt.Println("grade", grade)
	fmt.Println("feedback", narasi)
}
