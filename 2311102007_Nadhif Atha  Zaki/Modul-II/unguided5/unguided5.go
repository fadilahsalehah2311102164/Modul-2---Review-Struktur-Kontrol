package main

import "fmt"

func main() {
	var bilangan int

	fmt.Print("Bilangan: ")
	fmt.Scan(&bilangan)

	if bilangan <= 1 {
		fmt.Println("Bilangan harus lebih besar dari 1.")
		return
	}

	fmt.Print("Faktor: ")
	var jumlahFaktor int = 0
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Print(i, " ")
			jumlahFaktor++
		}
	}
	fmt.Println()

	if jumlahFaktor == 2 {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}
