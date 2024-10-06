package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var jumlahBunga int
	var pita []string
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("N: ")
	fmt.Scanln(&jumlahBunga)

	for i := 1; i <= jumlahBunga; i++ {
		fmt.Printf("Bunga %d: ", i)
		scanner.Scan()
		bunga := scanner.Text()
		pita = append(pita, bunga)
	}

	result := strings.Join(pita, " - ")
	fmt.Printf("Pita: %s\n", result)
}
