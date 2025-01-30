package main

import (
	"fmt"
)

// type Luas interface { // kontrak yang harus dipenuhi
type BangunDatar interface { // kontrak yang harus dipenuhi
	HitungLuas() int // method
}

type Persegi struct { // struct
	Sisi int
}

func (persegi Persegi) HitungLuas() int { // fungsi
	return persegi.Sisi * persegi.Sisi //return
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (PersegiPanjang PersegiPanjang) HitungLuas() int {
	return PersegiPanjang.Panjang * PersegiPanjang.Lebar
}

// tes

type Tes struct {
	Angka int
}

func (tes Tes) HitungLuas() int {
	return 1000
}

// masalah
func main() {
	PersegiPanjang := PersegiPanjang{Panjang: 6, Lebar: 5}
	luas := SeberapaLuas(PersegiPanjang)
	fmt.Println("Luas Persegi Panjang :", luas)

	persegi := Persegi{Sisi: 4}
	luas = SeberapaLuas(persegi)
	fmt.Println("Luas Persegi:", luas)

	tes := Tes{Angka: 5}
	luas = SeberapaLuas(tes)
	fmt.Println("luas Tes", luas)
}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}
