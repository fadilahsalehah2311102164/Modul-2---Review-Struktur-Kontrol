package main

import (
	"fmt"
)

func main() {
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scanln(&b)

	fmt.Print("Faktor: ")
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			fmt.Print(i, " ")
		}
	}
	fmt.Println()

	prima := true
	if b <= 1 {
		prima = false
	} else {
		for i := 2; i*i <= b; i++ {
			if b%i == 0 {
				prima = false
				break
			}
		}
	}

	fmt.Println("Prima:", prima)
}
