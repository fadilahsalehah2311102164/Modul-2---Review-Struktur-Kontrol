package main

import (
	"fmt"
)

func main() {
	var nilai float32
	var indeks string

	// Meminta input nilai
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nilai)

	// Logika penetuan nilai huruf berdasarkan nilai numerik
	if nilai >= 80 {
		indeks = "A"
	} else if nilai >= 70 {
		indeks = "B"
	} else if nilai >= 65 {
		indeks = "C"
	} else if nilai >= 45 {
		indeks = "D"
	} else if nilai >= 40 {
		indeks = "E"
	} else {
		indeks = "F"
	}

	// Menampilkan hasil
	fmt.Printf("Nilai Indeks untuk nilai %.2f adalah %s\n", nilai, indeks)
}
