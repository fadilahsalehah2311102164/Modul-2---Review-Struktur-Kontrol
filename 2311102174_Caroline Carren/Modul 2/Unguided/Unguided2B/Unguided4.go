package main

import (
	"fmt"
)

func main() {
	// Deklarasi variabel k dengan tipe float64
	var k float64
	// Menampilkan pesan untuk memasukkan nilai K
	fmt.Print("Nilai K = ")
	// Membaca input dari pengguna dan menyimpannya ke variabel k
	fmt.Scanln(&k)

	// Hitung f(K) dengan rumus: f(k) = ((4*k + 2)^2) / ((4*k + 1) * (4*k + 3))
	f_k := (4*k + 2) * (4*k + 2) / ((4*k + 1) * (4*k + 3))

	// Menampilkan hasil perhitungan f(K) dengan presisi 10 angka di belakang koma
	fmt.Printf("Nilai f(K) = %.10f\n", f_k)
}
