package main

import "fmt"

func main() {
	var bilangan153 int
	fmt.Print("Bilangan: ")
	fmt.Scan(&bilangan153)

	if bilangan153 <= 1 {
		fmt.Println("LEBIH DARI 1 WOY.")
		return
	}

	fmt.Print("Faktor: ")
	var faktor int = 0
	for i := 1; i <= bilangan153; i++ {
		if bilangan153%i == 0 {
			fmt.Print(i, " ")
			faktor++
		}
	}
	fmt.Println()

	if faktor == 2 {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
