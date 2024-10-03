package main

import (
	"bufio"   // Mengimpor paket bufio untuk membaca input pengguna secara efisien
	"fmt"     // Mengimpor paket fmt untuk fungsi input/output seperti Println
	"os"      // Mengimpor paket os untuk menggunakan fungsi input dari sistem operasi
	"strings" // Mengimpor paket strings untuk manipulasi string, seperti split dan trim
)

func main() {
	// Urutan warna yang benar yang harus dimasukkan oleh pengguna
	correctOrder := []string{"merah", "kuning", "hijau", "ungu"}

	// Membaca input pengguna menggunakan bufio
	reader := bufio.NewReader(os.Stdin)

	// Variabel untuk melacak apakah semua percobaan berhasil
	success := true

	// Melakukan loop untuk 5 percobaan
	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i) // Menampilkan nomor percobaan ke layar

		// Membaca input dari pengguna dan memotong karakter newline di akhir
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input) // Menghapus spasi dan newline dari input

		// Memisahkan input berdasarkan spasi, menghasilkan array warna
		colors := strings.Split(input, " ")

		// Mengecek apakah setiap warna yang dimasukkan sesuai dengan urutan yang benar
		for j := 0; j < 4; j++ {
			if colors[j] != correctOrder[j] { // Membandingkan dengan urutan warna yang benar
				success = false // Jika salah satu warna tidak sesuai, set flag success menjadi false
				break           // Keluar dari loop jika ada kesalahan
			}
		}

		// Jika ada satu percobaan yang gagal, keluar dari loop utama
		if !success {
			break
		}
	}

	// Menampilkan hasil akhir, apakah semua percobaan berhasil
	if success {
		fmt.Println("BERHASIL: true") // Jika semua percobaan benar
	} else {
		fmt.Println("BERHASIL: false") // Jika ada percobaan yang salah
	}
}
