package main

import "fmt"

// Fungsi untuk menghitung biaya pengiriman
func hitungBiaya(beratGram int) (int, int, int) {
	// Konversi berat dari gram ke kg
	beratKg := beratGram / 1000
	sisaGram := beratGram % 1000

	// Hitung biaya pengiriman per kg
	biayaKg := beratKg * 10000

	// Hitung biaya pengiriman untuk sisa gram
	var biayaSisa int
	if sisaGram >= 500 {
		biayaSisa = sisaGram * 5
	} else if sisaGram > 0 && beratKg <= 10 {
		biayaSisa = sisaGram * 15
	} else {
		biayaSisa = 0
	}

	return biayaKg, biayaSisa, beratKg
}

func main() {
	var beratGram int
	fmt.Print("Masukkan berat parsel (dalam gram): ")
	fmt.Scanln(&beratGram)

	// Panggil fungsi hitungBiaya untuk mendapatkan biaya pengiriman
	biayaKg, biayaSisa, beratKg := hitungBiaya(beratGram)

	// Tampilkan hasil perhitungan
	fmt.Printf("Total berat: %d kg %d gram\n", beratKg, beratGram%1000)
	fmt.Printf("Biaya pengiriman: Rp. %d\n", biayaKg+biayaSisa)
	fmt.Printf("  - Biaya per kg: Rp. %d\n", biayaKg)
	fmt.Printf("  - Biaya sisa gram: Rp. %d\n", biayaSisa)
}