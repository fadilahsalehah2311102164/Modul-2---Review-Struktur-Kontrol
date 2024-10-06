package main

import (
	"fmt"
)

func hitungF(k float64) float64 {
	// Menghitung f(K) sesuai dengan persamaan yang diberikan
	atas := (4*k + 2) * (4*k + 2)
	bawah := (4*k + 1) * (4*k + 3)
	return atas / bawah
}

func main() {
	var k float64
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	fK := hitungF(k)
	fmt.Printf("Nilai f(K) = %.10f\n", fK)
}
