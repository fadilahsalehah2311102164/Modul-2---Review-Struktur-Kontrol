package main

import (
	"fmt" // Mengimpor paket fmt untuk fungsi input/output
)

func main() {
	// Deklarasi variabel a, b, c, d, e untuk menyimpan input dari pengguna
	var a, b, c, d, e int
	var hasil int // Variabel untuk menyimpan hasil penjumlahan

	// Menggunakan fmt.Scanln untuk membaca 5 angka integer dari input pengguna
	fmt.Scanln(&a, &b, &c, &d, &e)

	// Melakukan penjumlahan dari semua variabel a, b, c, d, e
	hasil = a + b + c + d + e

	// Mencetak hasil penjumlahan ke layar
	fmt.Println("Hasil Penjumlahan ", a, b, c, d, e, "adalah = ", hasil)
}
