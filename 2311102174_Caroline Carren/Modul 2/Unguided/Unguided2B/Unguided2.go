package main

import "fmt"

func main() {
	// Deklarasi variabel untuk menyimpan pita dan banyaknya bunga
	var pita string     // Variabel 'pita' digunakan untuk menggabungkan nama bunga
	var banyakBunga int // Variabel 'banyakBunga' untuk menghitung jumlah bunga yang dimasukkan

	// Menggunakan loop untuk meminta input bunga dari pengguna secara berulang
	for {
		var bunga string
		// Meminta input nama bunga
		fmt.Print("Bunga (ketik 'SELESAI' untuk berhenti): ")
		fmt.Scanln(&bunga) // Membaca input dari pengguna

		// Jika pengguna mengetik "SELESAI", maka keluar dari loop
		if bunga == "SELESAI" {
			break
		}

		// Jika pita sudah memiliki isi, tambahkan pemisah " – "
		if pita != "" {
			pita += " – "
		}
		// Menambahkan nama bunga ke pita
		pita += bunga

		// Meningkatkan jumlah bunga yang dimasukkan
		banyakBunga++
	}

	// Menampilkan hasil akhir
	fmt.Println("Pita:", pita)                   // Menampilkan pita yang berisi nama-nama bunga
	fmt.Println("Banyaknya bunga:", banyakBunga) // Menampilkan jumlah bunga yang dimasukkan
}
