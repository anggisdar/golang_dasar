package main

type Luas interface { // kontrak yang harus dipenuhi
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

func main() {
}
