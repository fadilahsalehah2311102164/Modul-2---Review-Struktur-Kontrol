package main

import (
	"fmt"
	"math"
)

func main() {
	var kantong1, kantong2 float64
	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scanln(&kantong1, &kantong2)

		total := kantong1 + kantong2

		if total > 150 {
			fmt.Println("Total berat tidak boleh melebihi 150 kg.")
			fmt.Println("Program Selesai")
			break // Keluar dari loop dan program selesai
		} else if kantong1 < 0 || kantong2 < 0 {
			fmt.Println("Berat tidak boleh negatif.")
		} else if math.Abs(kantong1-kantong2) < 9 {
			fmt.Println("Sepeda motor akan oleng: false")
		} else {
			fmt.Println("Sepeda motor akan oleng: true")
		}
	}
}
