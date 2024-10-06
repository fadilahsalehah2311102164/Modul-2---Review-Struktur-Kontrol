package main

import "fmt"

func hitungBiayaKirim(beratParsel int) int {
	// Menghitung berat dalam kg dan sisa gram
	beratKg := beratParsel / 1000
	sisaGram := beratParsel % 1000

	// Menghitung biaya dasar
	biayaDasar := beratKg * 10000

	// Menghitung biaya tambahan berdasarkan sisa gram
	var biayaTambahan int
	if sisaGram >= 500 {
		biayaTambahan = sisaGram * 5
	} else if sisaGram > 0 && beratKg > 10 {
		// Jika sisa gram kurang dari 500 gram dan total berat lebih dari 10 kg, maka gratis
		biayaTambahan = 0
	} else {
		biayaTambahan = sisaGram * 15
	}

	// Menghitung total biaya
	totalBiaya := biayaDasar + biayaTambahan
	return totalBiaya
}

func main() {
	fmt.Println("PT POS - BlayaPos")

	for {
		var beratParsel int

		fmt.Print("Berat parsel (gram): ")
		fmt.Scan(&beratParsel)

		if beratParsel <= 0 {
			fmt.Println("Berat parsel harus positif.")
			continue
		}

		totalBiaya := hitungBiayaKirim(beratParsel)
		fmt.Println("Total biaya: Rp.", totalBiaya)
	}
}
