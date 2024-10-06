package main

import "fmt"

func f(k int) float64 {
	return float64((4*k+2)*(4*k+2)) / float64((4*k+1)*(4*k+3))
}

func main() {
	var k int

	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&k)

	hasil := f(k)
	fmt.Printf("Nilai f(k) = %.10f\n", hasil)
}
