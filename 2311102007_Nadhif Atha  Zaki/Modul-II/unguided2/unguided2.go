package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("Modul Praktikum Algoritma dan Pemrograman 2")

	for {
		var beratKantong1, beratKantong2 float64

		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&beratKantong1, &beratKantong2)

		// Validasi input
		if beratKantong1 < 0 || beratKantong2 < 0 {
			fmt.Println("Berat tidak boleh negatif. Masukkan ulang.")
			continue
		}

		// Hitung total berat
		totalBerat := beratKantong1 + beratKantong2

		// Hentikan program jika total berat melebihi 150 kg
		if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		// Hitung selisih berat
		selisihBerat := math.Abs(beratKantong1 - beratKantong2)

		// Cek apakah sepeda motor akan oleng
		if selisihBerat >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}
	}
}
