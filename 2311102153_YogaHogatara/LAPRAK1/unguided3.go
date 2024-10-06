package main

import (
	"fmt"
)

func f(k int) float64 {
	
	atas:= (4*float64(k) + 2) * (4*float64(k) + 2)
	bawah := (4*float64(k) + 1) * (4*float64(k) + 3)
	return atas / bawah
}



func main() {
	
	var k int
	fmt.Print("Nilai K = ")
	fmt.Scan(&k)

	hasil := f(k)
	fmt.Printf("Nilai f(k) = %.10f\n", hasil)
}
