package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var N int
	fmt.Print("N: ")
	fmt.Scan(&N)

	pita_153 := ""
	hitung_153 := 0

	reader := bufio.NewReader(os.Stdin)

	for y := 1; y <= N; y++ {
		fmt.Printf("bunga%d: ",y)
		input, _ :=reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Jika user mengetik "SELESAI", hentikan proses input
		if strings.ToLower(input) == "SELESAI" {
			break
		}

		// Tambahkan nama bunga ke pita
		if pita_153 == "" {
			pita_153 = input
		} else {
			pita_153 += " - " + input
		}
		hitung_153++
	}

	fmt.Printf("\nPita: " + pita_153)
	fmt.Printf("\nbunga : %d\n", hitung_153)
}
