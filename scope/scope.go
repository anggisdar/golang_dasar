package scope

import "fmt"

//variable global yang bisa diakses publik
var GlobalVar = "Saya Publik"

//variable yang "hanya" bisa diakses lokal dalam fungsi
func main() {
	localVar := "Saya Lokal"
	fmt.Println(globalVar)
	fmt.Println(localVar)

	//ini blok if blok-scope
	if true {
		blockVar := "Saya Block"
		fmt.Println(blockVar)
	}

	// fmt.println(blokvar)blokvar hanya bisa di dalam blok if
}
