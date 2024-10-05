package main

import "fmt"

func main() {
    // Mendeklarasikan variabel untuk menyimpan nilai akhir dan nama nilai
	var nam float32  // Variabel untuk menyimpan nilai akhir mata kuliah 
	var nmk string   // Variabel untuk menyimpan huruf nilai

	// Meminta user untuk menginput nilai akhir mata kuliah
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nam)   // Membaca input nilai dari user

	// Logika penentuan nilai huruf berdasarkan nilai akhir yang diinputkan 
	if nam >= 80 {
		nmk = "A"
	} else if nam >= 70 {
		nmk = "B"
	} else if nam >= 60 {
		nmk = "C"
	} else if nam >= 50 {
		nmk = "D"
	} else if nam >= 40 {
		nmk = "E"
	} else {
		nmk = "F"
	}

	// Menampilkan hasil
	fmt.Printf("Nilai Indeks untuk nilai %.2f adalah %s\n", nam, nmk)
}
