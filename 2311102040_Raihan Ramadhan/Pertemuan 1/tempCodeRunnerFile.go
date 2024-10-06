package main

import (
	"fmt"
)

func hitungBiaya(berat int) (int, string, string) {
	// Konversi berat dari gram ke kg dan sisa gram
	kg := berat / 1000
	sisaGram := berat % 1000

	// Biaya dasar per kilogram
	biayaKg := kg * 10000

	// Hitung biaya untuk sisa gram
	biayaSisaGram := 0
	if sisaGram > 0 {
		if sisaGram <= 500 {
			biayaSisaGram = sisaGram * 5
		} else {
			biayaSisaGram = sisaGram * 15
		}
	}

	// Jika total berat lebih dari 10 kg, biaya sisa gram tidak dikenakan
	if berat > 10000 {
		biayaSisaGram = 0
	}

	// Total biaya
	totalBiaya := biayaKg + biayaSisaGram

	// Format string untuk detail berat dan rincian biaya
	detailBerat := fmt.Sprintf("%d kg + %d gr", kg, sisaGram)
	detailBiaya := fmt.Sprintf("Rp. %d + Rp. %d", biayaKg, biayaSisaGram)
	return totalBiaya, detailBerat, detailBiaya
}

func main() {
	// Input berat parsel dalam gram
	var berat int

	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scan(&berat)

	// Hitung biaya
	totalBiaya, detailBerat, detailBiaya := hitungBiaya(berat)

	// Tampilkan hasil
	fmt.Println("Detail berat:", detailBerat)
	fmt.Println("Rincian biaya:", detailBiaya)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
