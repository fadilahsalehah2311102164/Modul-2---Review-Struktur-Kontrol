package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	//urutan warna yang benar
	correctOrder := []string{"merah", "kuning", "hijau", "ungu"}

	//membaca input untuk 5 percobaan
	reader := bufio.NewReader(os.Stdin)
	success := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)

		//membaca input dari pengguna
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		//memisahkan input berdasarkan spasi
		colors := strings.Split(input, " ")

		//mengecek apakah urutan warna sesuai
		for j := 0; j < 4; j++ {
			if colors[j] != correctOrder[j] {
				success = false
				break
			}
		}

		//jika ada percobaan yang tidak sesuai, keluar dari loop
		//if !success {
			//break
		//}
	}

	//menampilkan hasil
	if success {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}
