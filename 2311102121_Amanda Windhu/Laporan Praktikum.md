### UNGUIDED ###

### 1. Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturut-turut adalah ‘merah’, ‘kuning’, ‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.<br/> Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.<br/>

```go
package main

import (
	"bufio"    // Package untuk membaca input dari pengguna melalui terminal
	"fmt"      // Package untuk menampilkan output ke terminal
	"os"       // Package untuk interaksi dengan sistem operasi
	"strings"  // Package untuk manipulasi string
)

func main() {
	// Array yang menyimpan urutan warna yang benar
	correctOrder := []string{"merah", "kuning", "hijau", "ungu"}

	// Membuat reader untuk membaca input pengguna dari terminal
	reader := bufio.NewReader(os.Stdin)
	success := true // Flag untuk menandai apakah input sesuai urutan warna yang benar

	// Loop untuk percobaan sebanyak 5 kali
	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)   // Menampilkan percobaan ke berapa
		input, _ := reader.ReadString('\n') // Membaca input dari pengguna hingga baris baru ('\n')
		input = strings.TrimSpace(input)    // Menghapus spasi atau karakter newline dari input

		colors := strings.Split(input, " ") // Memecah string input berdasarkan spasi menjadi slice
		for j := 0; j < 4; j++ {            // Mengecek apakah setiap warna sesuai dengan urutan yang benar
			if colors[j] != correctOrder[j] {  // Jika ada warna yang salah urutan
				success = false  // Tandai success sebagai false
				break            // Hentikan pengecekan lebih lanjut
			}
		}
		if !success {  // Jika success sudah false, hentikan percobaan lebih lanjut
			break
		}
	}

	// Menampilkan hasil akhir
	if success {
		fmt.Println("BERHASIL : true")  // Jika semua input sesuai urutan
	} else {
		fmt.Println("BERHASIL : false") // Jika ada urutan warna yang salah
	}
}
```
### Output: 
![image](https://github.com/user-attachments/assets/c27ee543-0db9-44ec-bad1-d6dd4dfb006e)

Kode di atas untuk meminta pengguna memasukkan urutan warna dan memeriksa apakah input tersebut sesuai dengan urutan yang benar yang telah ditentukan. Pengguna memiliki maksimal lima percobaan untuk memasukkan urutan yang benar. Jika pengguna memasukkan urutan yang benar pada suatu saat, program akan berhenti dan menunjukkan keberhasilan; jika tidak, program akan menunjukkan kegagalan setelah lima percobaan.
Program ini menggunakan paket `bufio` untuk membaca input pengguna secara efisien, `fmt` untuk mencetak output ke konsol, `os` untuk berinteraksi dengan sistem operasi, dan `strings` untuk manipulasi string. Urutan warna yang benar ditentukan dalam sebuah slice sebagai `["merah", "kuning", "hijau", "ungu"]`.<br/>
Dalam program, sebuah reader buffered diinisialisasi untuk membaca input dari terminal, dan sebuah flag boolean `success` diatur ke `true` untuk menandai apakah input pengguna benar. Program kemudian menjalankan loop lima kali, meminta pengguna untuk memasukkan urutan warna pada setiap percobaan.<br/>
Setiap input dibaca hingga karakter newline, dibersihkan dari spasi, dan dipisahkan menjadi string warna individual. Kemudian, program memeriksa setiap warna dalam input pengguna terhadap urutan yang benar. Jika ada warna yang tidak cocok, flag `success` diatur ke `false` dan pengecekan lebih lanjut dihentikan.
Jika flag `success` masih `false` setelah lima percobaan, program akan menampilkan output "BERHASIL : false"; jika tidak, program akan menampilkan "BERHASIL : true".<br/>
Secara keseluruhan, program ini menunjukkan bagaimana menangani input pengguna, melakukan manipulasi string, dan menerapkan alur kontrol dasar dalam Go, memberikan cara sederhana namun menarik bagi pengguna untuk berinteraksi dengan urutan warna sambil belajar tentang konsep pemrograman seperti loop, kondisi, dan slice.<br/>

### 2. Suatu pita (string) berisi kumpulan nama-nama bunga yang dipisahkan oleh spasi dan '–', contoh pita diilustrasikan seperti berikut ini.<br/>Pita: mawar – melati – tulip – teratai – kamboja – anggrek<br/> Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita.<br/> (Petunjuk: gunakan operasi penggabungan string dengan operator “+”.)<br/>Tampilkan isi pita setelah proses input selesai.<br/>
