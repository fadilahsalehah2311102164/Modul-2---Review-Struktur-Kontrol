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
		fmt.Printf("bunga%d: ",i)
		input, _ :=reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Tambahkan nama bunga ke pita
		if pita_153 == "" {
			pita_153 = input
		} else {
			pita_153 += " - " + input
		}
		hitung_153++
	}

	fmt.Printf("\nPita: " + pita_153)
	
}
