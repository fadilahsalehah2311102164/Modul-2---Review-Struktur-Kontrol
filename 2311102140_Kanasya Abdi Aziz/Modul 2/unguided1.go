package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var bunga []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Masukkan bunga, ketik 'SELESAI' untuk selesai:")
	for {
		fmt.Print("Bunga: ")
		scanner.Scan()
		input := scanner.Text()

		if input == "SELESAI" {
			break
		}

		bunga = append(bunga, input)
	}

	fmt.Println("Pita:", strings.Join(bunga, " - "))
	fmt.Println("Bunga:", len(bunga))
}
