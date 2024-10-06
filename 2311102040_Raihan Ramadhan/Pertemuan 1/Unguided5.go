package main

import (
	"fmt"
	"strings"
)

// Fungsi untuk mencari faktor dari bilangan
func cariFaktor(b int) []int {
	var faktor []int
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			faktor = append(faktor, i)
		}
	}
	return faktor
}

// Fungsi untuk memeriksa apakah bilangan prima
func apakahPrima(b int) bool {
	if b < 2 {
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

	// Meminta input dari user
	fmt.Print("Masukkan bilangan: ")
	fmt.Scan(&b)

	// Mencari faktor bilangan
	faktor := cariFaktor(b)

	// Mengubah slice faktor menjadi string dengan spasi sebagai pemisah
	faktorStr := strings.Trim(strings.Join(strings.Fields(fmt.Sprint(faktor)), " "), "[]")

	// Mencetak faktor tanpa tanda kurung
	fmt.Printf("Faktor: %s\n", faktorStr)

	// Menentukan apakah bilangan prima
	if apakahPrima(b) {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
