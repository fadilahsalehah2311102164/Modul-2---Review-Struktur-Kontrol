package main

import (
	"fmt"
)


func hitungbiaya(berat int) int {
	
	const biayaperkg = 10000
	const biayalebih = 5
	const biayakurang = 15

	
	kg := berat / 1000
	sisa153 := berat % 1000

	
	biaya := kg * biayaperkg

	
	if sisa153 >= 500 {
		biaya += sisa153 * biayalebih
	} else {
		biaya += sisa153 * biayakurang
	}

	if berat >= 10000 {
		biaya = kg * biayaperkg
	}

	return biaya
}

func main() {
 var berat int
	fmt.Printf("PT POS COY\n")
	fmt.Printf("Berat parsel (gram): ")
	fmt.Scan(&berat)
	biaya := hitungbiaya(berat)
	fmt.Printf("Total Biaya: Rp. %d\n\n", biaya)
	
}
