package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan nilai numerik dan nilai huruf
	var nam float32 // Variabel 'nam' untuk menyimpan nilai numerik dengan tipe float32
	var nmk string  // Variabel 'nmk' untuk menyimpan nilai huruf (indeks)

	// Meminta input nilai dari pengguna
	fmt.Print("Masukkan nilai : ")
	fmt.Scan(&nam) // Membaca input nilai numerik dari pengguna

	// Logika penentuan nilai huruf berdasarkan nilai numerik
	if nam > 80 {
		nmk = "A" // Jika nilai lebih dari 80, indeksnya A
	} else if nam > 72.5 {
		nmk = "B" // Jika nilai lebih dari 72.5, indeksnya B
	} else if nam > 65 {
		nmk = "C" // Jika nilai lebih dari 65, indeksnya C
	} else if nam > 50 {
		nmk = "D" // Jika nilai lebih dari 50, indeksnya D
	} else if nam > 40 {
		nmk = "E" // Jika nilai lebih dari 40, indeksnya E
	} else {
		nmk = "F" // Jika nilai di bawah atau sama dengan 40, indeksnya F
	}

	// Menampilkan hasil penentuan nilai huruf (indeks)
	fmt.Printf("Nilai Indeks untuk nilai %.2f adalah %s\n", nam, nmk)
}
