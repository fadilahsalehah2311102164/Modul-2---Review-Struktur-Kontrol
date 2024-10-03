package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan nilai numerik dan nilai huruf
	var nam float32 // Variabel 'nam' untuk menyimpan nilai numerik dengan tipe float32
	var nmk string  // Variabel 'nmk' untuk menyimpan nilai huruf (indeks)

	// Meminta input nilai dari pengguna
	fmt.Print("Masukkan nilai: ")
	fmt.Scan(&nam) // Membaca input nilai numerik dari pengguna

	// Logika penentuan nilai huruf berdasarkan nilai numerik
	if nam >= 80 {
		nmk = "A" // Jika nilai 80 atau lebih, indeksnya A
	} else if nam >= 70 {
		nmk = "B" // Jika nilai 70 atau lebih, indeksnya B
	} else if nam >= 60 {
		nmk = "C" // Jika nilai 60 atau lebih, indeksnya C
	} else if nam >= 50 {
		nmk = "D" // Jika nilai 50 atau lebih, indeksnya D
	} else if nam >= 40 {
		nmk = "E" // Jika nilai 40 atau lebih, indeksnya E
	} else {
		nmk = "F" // Jika nilai di bawah 40, indeksnya F
	}

	// Menampilkan hasil penentuan nilai huruf (indeks)
	fmt.Printf("Nilai Indeks untuk nilai %.2f adalah %s\n", nam, nmk)
}
