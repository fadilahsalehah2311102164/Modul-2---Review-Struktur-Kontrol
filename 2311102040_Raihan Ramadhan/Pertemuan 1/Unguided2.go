package main

import (
	"fmt"
	"math"
)

func main() {
	var beratKiri, beratKanan float64
	var totalBerat float64

	for {
		fmt.Print("\nMasukkan berat belanjaan di kedua kantong (pisahkan dengan spasi): ")

		// Menggunakan Scanln untuk membaca input dengan spasi
		_, err := fmt.Scanln(&beratKiri, &beratKanan)
		if err != nil {
			fmt.Println("Input tidak valid. Silakan masukkan dua angka.")
			continue
		}

		// Hitung selisih antara berat kiri dan kanan
		selisih := math.Abs(beratKiri - beratKanan)
		oleng := selisih >= 9

		// Tampilkan apakah motor Pak Andi akan oleng
		fmt.Printf("Sepeda motor Pak Andi akan oleng: %v\n", oleng)

		// Hitung total berat
		totalBerat = beratKiri + beratKanan

		// Jika total berat melebihi 150 kg atau salah satu berat negatif, hentikan program
		if totalBerat >= 150 || beratKiri < 0 || beratKanan < 0 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}
