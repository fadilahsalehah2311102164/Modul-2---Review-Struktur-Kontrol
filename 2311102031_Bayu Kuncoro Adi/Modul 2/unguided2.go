package main

import (
	"bufio"   // Mengimpor paket bufio untuk membaca input dari pengguna
	"fmt"     // Mengimpor paket fmt untuk menampilkan output
	"os"      // Mengimpor paket os untuk menggunakan input dari sistem operasi
	"strconv" // Mengimpor paket strconv untuk konversi string ke tipe data numerik
	"strings" // Mengimpor paket strings untuk manipulasi string
)

func main() {
	reader := bufio.NewReader(os.Stdin) // Membuat reader untuk membaca input pengguna

	// Loop untuk meminta input pengguna secara berulang
	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		input, _ := reader.ReadString('\n') // Membaca input pengguna
		input = strings.TrimSpace(input)    // Menghapus spasi dan karakter newline di awal dan akhir

		// Memisahkan input menjadi dua bagian berdasarkan spasi
		berat := strings.Split(input, " ")

		// Mengecek apakah input terdiri dari dua angka
		if len(berat) != 2 {
			fmt.Println("Input tidak valid. Masukkan dua angka.") // Menampilkan pesan kesalahan jika input tidak valid
			continue                                              // Melanjutkan ke iterasi berikutnya untuk meminta input ulang
		}

		// Mengonversi input dari string ke float64
		berat1, _ := strconv.ParseFloat(berat[0], 64)
		berat2, _ := strconv.ParseFloat(berat[1], 64)

		// Jika salah satu berat negatif, sepeda tidak oleng dan keluar dari loop
		if berat1 < 0 || berat2 < 0 {
			fmt.Println("Sepeda motor pak Andi akan oleng: false") // Menampilkan bahwa sepeda tidak oleng
			break                                                  // Keluar dari loop
		}

		// Jika total berat lebih dari 150 kg, sepeda tidak oleng dan keluar dari loop
		if berat1+berat2 > 150 {
			fmt.Println("Sepeda motor pak Andi akan oleng: false") // Sepeda tidak oleng karena berat berlebihan
			break                                                  // Keluar dari loop
		}

		// Menghitung selisih berat antara kedua kantong
		selisih := berat1 - berat2
		if selisih < 0 {
			selisih = -selisih // Mengubah selisih menjadi nilai absolut (selisih selalu positif)
		}

		// Jika selisih berat lebih dari atau sama dengan 9 kg, sepeda akan oleng
		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}

	// Pesan yang ditampilkan setelah proses selesai
	fmt.Println("Proses selesai.")
}