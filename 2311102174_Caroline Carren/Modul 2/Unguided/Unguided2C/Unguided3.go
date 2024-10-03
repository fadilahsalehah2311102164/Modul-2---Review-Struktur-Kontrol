package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel untuk menyimpan input bilangan
	var b int

	// Meminta pengguna untuk memasukkan bilangan
	fmt.Print("Bilangan: ")
	fmt.Scanln(&b)

	// Mencetak faktor-faktor dari bilangan
	fmt.Print("Faktor: ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
		}
	}

	// Mengecek apakah bilangan adalah bilangan prima
	prima := true
	if b <= 1 {
		prima = false
	} else {
		// Loop untuk mencari pembagi selain 1 dan b
		for i := 2; i*i <= b; i++ {
			if b%i == 0 {
				prima = false
				break
			}
		}
	}

	// Menampilkan apakah bilangan prima atau bukan
	fmt.Println("\nPrima:", prima)
}
