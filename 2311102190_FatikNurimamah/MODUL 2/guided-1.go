package main

import (
	"fmt"
)
func main () {

	var nama string
	
	fmt.Print("Masukkan nama: ")
	fmt.Scanln(&nama)
	fmt.Println("Nama yang dimasukkan:", nama)
}
