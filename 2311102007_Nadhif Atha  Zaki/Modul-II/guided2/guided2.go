package main

import (
	"fmt"
)

func main() {
	// Warna yang diharapkan untuk setiap percobaan
	expected := []string{"merah", "kuning", "hijau", "ungu"}

	// Menampung status keberhasilan dari setiap percobaan
	success := true

	// Loop untuk 5 kali percobaan
	for i := 1; i <= 5; i++ {
		var warna1, warna2, warna3, warna4 string
		fmt.Printf("Percobaan %d: ", i)
		fmt.Scan(&warna1, &warna2, &warna3, &warna4)

		// Memeriksa apakah input sama dengan yang diharapkan
		if warna1 != expected[0] || warna2 != expected[1] || warna3 != expected[2] || warna4 != expected[3] {
			success = false
		}
	}

	// Menampilkan hasil akhir
	fmt.Printf("BERHASIL: %v\n", success)
}
