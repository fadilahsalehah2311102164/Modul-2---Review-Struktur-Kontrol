package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Variabel pita digunakan untuk menyimpan daftar nama bunga
	var pita string
	var TotalBunga int

	// Menggunakan scanner untuk membaca input dari pengguna
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan nama bunga dan ketik 'SELESAI' untuk mengakhiri. :\n")

	for {
		TotalBunga++
		fmt.Printf("Bunga %d: ", TotalBunga)

		// Membaca input dari pengguna
		scanner.Scan()
		input := scanner.Text()

		// Cek jika pengguna mengetikkan "SELESAI" maka program akan berhenti
		if strings.ToUpper(input) == "SELESAI" {
			TotalBunga-- // Mengurangi jumlah bunga karena input "SELESAI" tidak dihitung
			break
		}

		// Menambahkan nama bunga ke pita 
		if pita == "" {
			pita = input
		} else {
			pita += " - " + input
		}
	}

	// Menampilkan hasil akhir pita dan jumlah bunga yang dimasukkan
	fmt.Println("Pita:", pita)
	fmt.Println("Total bunga:", TotalBunga )
}
