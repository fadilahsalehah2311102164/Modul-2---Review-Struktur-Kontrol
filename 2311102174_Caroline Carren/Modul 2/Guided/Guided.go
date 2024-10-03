package main

import (
	"fmt" // Mengimpor paket fmt yang menyediakan fungsi untuk input/output standar
)

func main() {
	var nama string // Deklarasi variabel nama dengan tipe string

	// Menggunakan fmt.Scanln untuk membaca input dari pengguna.
	// Harus menggunakan &nama karena Scanln mengharapkan pointer agar bisa mengisi nilai ke variabel nama
	fmt.Scanln(&nama)

	// Mencetak nilai variabel nama ke layar setelah input diambil dari pengguna
	fmt.Println(nama)
}
