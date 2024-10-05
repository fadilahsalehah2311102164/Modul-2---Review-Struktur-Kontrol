package main

import (
	"fmt"
)

// Fungsi untuk menemukan semua faktor dari bilangan b
func cariFaktor(b int) []int {
	var daftarfaktor []int
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			daftarfaktor = append(daftarfaktor, i)
		}
	}
	return daftarfaktor
}

// Fungsi untuk memeriksa apakah bilangan b adalah bilangan prima.
func cekBilPrima(b int) bool {
	if b <= 1 {
		return false
	}
	for i := 2; i*i <= b; i++ {
		if b%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	// untuk mencari faktor dan menampilkannya
	faktor := cariFaktor(b)
	fmt.Print("Faktor: ")
	for _, f := range faktor {
		fmt.Print(f, " ")
	}
	fmt.Println()

	// untuk mengecek dan menampilkan apakah bilangan tersebut prima
	if cekBilPrima(b) {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
