package main

import (
	"fmt"
	"math"
)

func main() {
	var BeratKantongKiri, BeratKantongKanan float64

	for {
		// Meminta pengguna untuk memasukkan berat kedua kantong
		fmt.Print("Masukkan berat barang di kedua kantong: ")
		fmt.Scan(&BeratKantongKiri, &BeratKantongKanan)

		// Mengecek apakah berat salah satu kantong bernilai negatif
		if BeratKantongKiri < 0 || BeratKantongKanan < 0 {
			fmt.Println("Proses selesai. Salah satu kantong memiliki berat negatif.")
			break
		}

		// Mengecek apakah total berat kedua kantong melebihi 150 kg
		TotalBeratKantong := BeratKantongKiri + BeratKantongKanan
		if TotalBeratKantong > 150 {
			fmt.Println("Proses selesai. Total berat kedua kantong melebihi 150 kg.")
			break
		}

		// Menghitung selisih berat antara kedua kantong
		SelisihBeratKantong := math.Abs(BeratKantongKiri -BeratKantongKanan)

		// Mengecek apakah selisih berat menyebabkan sepeda motor oleng
		if SelisihBeratKantong >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: True")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: False")
		}
	}
}
