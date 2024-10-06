package main

import (
	"fmt"
	"strings"
)

func main() {
	var pita string
	var bunga string
	var count int

	// Meminta input nama bunga hingga pengguna mengetik "SELESAI"
	for {
		fmt.Print("Masukkan nama bunga (ketik 'SELESAI' untuk selesai): ")
		fmt.Scanln(&bunga)

		// Jika user mengetik 'SELESAI', hentikan proses input
		if strings.ToUpper(bunga) == "SELESAI" {
			break
		}

		// Menambahkan bunga ke pita, dipisahkan dengan tanda " – "
		if pita == "" {
			pita = bunga
		} else {
			pita = pita + " – " + bunga
		}
		count++ // Menghitung jumlah bunga yang dimasukkan
	}

	// Menampilkan hasil pita dan jumlah bunga setelah proses selesai
	fmt.Println("Pita: ", pita)
	fmt.Printf("Jumlah bunga: %d\n", count)
}
