package main

import (
	"fmt"
)

func BiayaPengiriman(beratGram int) int {
	// Konversi berat ke kg dan sisa gram
	BeratKg := beratGram / 1000
	SisaBeratGram := beratGram % 1000

	// Hitung biaya dasar untuk kg
	biayaDasar := BeratKg * 10000

	// Hitung biaya tambahan untuk sisa gram
	biayaTambahan := 0
	if SisaBeratGram > 0 {
		if BeratKg >= 10 {
			// Jika berat lebih dari 10 kg, sisa gram digratiskan
			biayaTambahan = 0
		} else if SisaBeratGram <= 500 {
			// Jika sisa gram <= 500, biaya tambahan Rp. 5,- per gram
			biayaTambahan = SisaBeratGram * 5
		} else {
			// Jika sisa gram > 500, biaya tambahan Rp. 15,- per gram
			biayaTambahan = SisaBeratGram * 15
		}
	}

	// Total biaya adalah biaya dasar + biaya tambahan
	totalBiaya := biayaDasar + biayaTambahan

	// Cetak detail
	fmt.Printf("Detail berat: %d kg + %d gr\n", BeratKg, SisaBeratGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaDasar, biayaTambahan)
	return totalBiaya
}

func main() {
	var berat int
	// Minta input berat dari pengguna
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&berat)

	// Hitung biaya pengiriman
	totalBiaya := BiayaPengiriman(berat)

	// Tampilkan total biaya
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
