package main

import (
	"fmt"
)

func main() {
	var K int

	fmt.Print("Masukkan nilai K: ")
	fmt.Scanln(&K)

	prod := 1.0
	for k := 0; k <= K; k++ {
		atas := (4*float64(k) + 2) * (4*float64(k) + 2)  // (4k + 2)^2
		bawah := (4*float64(k) + 1) * (4*float64(k) + 3) // (4k + 1)(4k + 3)
		prod *= atas / bawah
	}

	fmt.Printf("Nilai akar 2 = %.10f\n", prod)
}
