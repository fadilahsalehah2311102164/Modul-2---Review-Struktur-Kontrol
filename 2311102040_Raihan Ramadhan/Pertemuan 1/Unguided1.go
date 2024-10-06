package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N int
	var pita []string

	// Membaca input jumlah bunga
	fmt.Print("Masukkan jumlah bunga: ")
	fmt.Scan(&N)

	reader := bufio.NewReader(os.Stdin)

	// Membersihkan newline yang tertinggal setelah input N
	reader.ReadString('\n')

	bungaKe := 1

	for bungaKe <= N {
		fmt.Printf("Bunga %d: ", bungaKe)
		bunga, _ := reader.ReadString('\n')
		bunga = strings.TrimSpace(bunga)

		// Jika input adalah "SELESAI", hentikan proses input
		if strings.ToUpper(bunga) == "SELESAI" {
			break
		}

		// Tambahkan bunga ke pita
		pita = append(pita, bunga)
		bungaKe++
	}

	// Tampilkan pita (gabungan bunga dengan separator " - ")
	fmt.Println("Pita:", strings.Join(pita, " - "))

	// Tampilkan jumlah bunga dalam pita
	fmt.Printf("Jumlah bunga dalam pita: %d\n", len(pita))
}
