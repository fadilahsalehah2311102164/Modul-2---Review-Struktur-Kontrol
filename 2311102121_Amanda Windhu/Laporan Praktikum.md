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

### 2. Suatu pita (string) berisi kumpulan nama-nama bunga yang dipisahkan oleh spasi dan '–', contoh pita diilustrasikan seperti berikut ini.<br/>Pita: mawar – melati – tulip – teratai – kamboja – anggrek<br/>Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita.<br/> (Petunjuk: gunakan operasi penggabungan string dengan operator “+”.)<br/>Tampilkan isi pita setelah proses input selesai.<br/>Modifikasi program sebelumnya, proses input akan berhenti apabila user mengetikkan `SELESAI`. Kemudian tampilkan isi pita beserta banyaknya bunga yang ada di dalam pita<br/>

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

func main() {
	// Membuat reader untuk membaca input dari pengguna
	reader := bufio.NewReader(os.Stdin)

	// Meminta input jumlah bunga yang akan dimasukkan (bilangan bulat positif N)
	fmt.Print("N: ")
	var N int
	for {
		// Baca input dari pengguna
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error membaca input: %v", err)
			return
		}

		// Konversi input ke integer
		N, err = strconv.Atoi(strings.TrimSpace(input))
		if err != nil || N <= 0 {
			fmt.Println("Harap masukkan bilangan bulat positif.")
		} else {
			break
		}
	}

	// Inisialisasi variabel pita (string) untuk menyimpan nama bunga
	var pita string
	var count int // Menyimpan jumlah bunga yang dimasukkan

	// Loop untuk menerima input nama bunga sebanyak N kali
	for i := 1; i <= N; i++ {
		fmt.Printf("Bunga %d: ", i) // Menambahkan instruksi

		// Membaca input dari pengguna
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Printf("Error: %v", err)
			return // Keluar dari program jika ada kesalahan
		}

		// Menghapus spasi dan karakter newline dari input
		input = strings.TrimSpace(input)

		// Cek jika pengguna mengetik "SELESAI"
		if strings.ToUpper(input) == "SELESAI" {
			break // Menghentikan input jika "SELESAI" dimasukkan
		}

		// Menggabungkan nama bunga dengan pita menggunakan " – " sebagai pemisah
		if pita == "" {
			pita = input // Jika pita masih kosong, langsung masukkan nama bunga
		} else {
			pita = pita + " – " + input // Jika sudah ada isinya, tambahkan dengan pemisah " – "
		}

		count++ // Menambah jumlah bunga yang dimasukkan
	}

	// Menampilkan isi pita dan bunga setelah semua input dimasukkan
	fmt.Println("Pita:", pita)
	fmt.Printf("Bunga: %d", count)
	
}
```
### Output: ![image](https://github.com/user-attachments/assets/9f77d530-d6d2-4c78-93e9-b117793bba7f)<br/>
![image](https://github.com/user-attachments/assets/adcb8be8-6eb8-431e-9d93-3d4f3e6fee06)

Kode di atas meminta pengguna untuk memasukkan jumlah bunga yang akan dimasukkan, kemudian mengumpulkan nama-nama bunga tersebut sebanyak jumlah yang diminta. Pengguna dapat mengakhiri proses penginputan lebih awal dengan mengetikkan "SELESAI". Setelah itu, program akan menyimpan nama-nama bunga dalam sebuah string yang dipisahkan oleh " – " dan menampilkan hasilnya berserta jumlah total bunga yang dimasukkan.<br/>

### 3. Setiap hari Pak Andi membawa banyak barang belanjaan dari pasar dengan mengendarai sepeda motor. Barang belanjaan tersebut dibawa dalam kantong terpal di kiri-kanan motor. Sepeda motor tidak akan oleng jika selisih berat barang di kedua kantong sisi tidak lebih dari 9 kg. Buatlah program Pak Andi yang menerima input dua buah bilangan real positif yang menyatakan berat total masing-masing isi kantong terpal. Program akan terus meminta input bilangan tersebut hingga salah satu kantong terpal berisi 9 kg atau lebih.<br/>Pada modifikasi program tersebut, program akan menampilkan true jika selisih kedua isi kantong lebih dari atau sama dengan 9 kg. Program berhenti memproses apabila total berat isi kedua kantong melebihi 150 kg atau salah satu kantong beratnya negatif.

# Sebelum di modifikasi

```go
package main

import (
	"fmt"
)

func main() {
	var beratKiri, beratKanan float64

	for {
		// Menerima input berat kedua kantong dalam satu baris
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&beratKiri, &beratKanan)

		// Cek jika salah satu kantong mencapai atau lebih dari 9 kg
		if beratKiri >= 9 || beratKanan >= 9 {
			fmt.Println("Proses selesai.")
			break
		}

		// Hitung selisih berat antara kantong kiri dan kanan
		selisih := beratKiri - beratKanan
		if selisih < 0 {
			selisih = -selisih
		}

		// Cek apakah selisih berat melebihi 9 kg
		if selisih > 9 {
			fmt.Println("Selisih berat antara kantong kiri dan kanan melebihi 9 kg.")
		}
	}
}

```
## Output:![image](https://github.com/user-attachments/assets/e28b4ab7-67eb-499e-b445-28104b997a8f)<br/>

# Setelah di modifikasi

```go
package main

import (
	"fmt"
)

func main() {
	var beratKiri, beratKanan float64

	for {
		// Menerima input berat kedua kantong dalam satu baris
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		fmt.Scan(&beratKiri, &beratKanan)

		// Cek jika salah satu kantong memiliki berat negatif
		if beratKiri < 0 || beratKanan < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		// Cek jika total berat kedua kantong melebihi 150 kg
		totalBerat := beratKiri + beratKanan
		if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		// Hitung selisih berat antara kantong kiri dan kanan
		selisih := beratKiri - beratKanan
		if selisih < 0 {
			selisih = -selisih
		}

		// Menampilkan hasil apakah sepeda motor akan oleng atau tidak
		if selisih >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}
	}
}

```
### Output: ![image](https://github.com/user-attachments/assets/e38d04e0-dec9-4866-85e4-353a2937c2d2)<br/>

Kode di atas bertujuan untuk membantu Pak Andi menjaga keseimbangan beban sepeda motornya dengan memeriksa selisih berat antara dua kantong terpal yang diisi barang belanjaan. Program menerima input dari pengguna berupa dua nilai yang mewakili berat barang di kantong kiri dan kanan. Program ini akan terus meminta input dari pengguna hingga salah satu dari dua kondisi terpenuhi: salah satu kantong memiliki berat lebih dari atau sama dengan 9 kg, atau selisih berat antara kedua kantong melebihi batas aman, yaitu 9 kg. Jika salah satu kantong mencapai berat 9 kg atau lebih, program akan menampilkan pesan "Proses selesai." dan menghentikan eksekusi. Di setiap iterasi, setelah pengguna memasukkan berat kedua kantong, program akan menghitung selisih berat antara kantong kiri dan kanan. Jika hasil perhitungan menunjukkan bahwa selisih tersebut lebih dari 9 kg, program akan menampilkan pesan: "Selisih berat antara kantong kiri dan kanan melebihi 9 kg." Jika tidak, program akan melanjutkan meminta input baru dari pengguna.<br/>
Proses ini akan terus berulang selama kedua kondisi penghentian belum terpenuhi. Sebagai contoh, jika pengguna memasukkan berat kantong kiri sebesar 5.5 kg dan berat kantong kanan sebesar 3.2 kg, program akan menghitung selisihnya, yaitu 5.5 - 3.2 = 2.3 kg, yang masih kurang dari 9 kg. Karena itu, program tidak akan menampilkan pesan dan akan melanjutkan meminta input. Selanjutnya, jika pengguna memasukkan berat kantong kiri sebesar 10 kg dan kantong kanan sebesar 0 kg, selisihnya adalah 10 kg, yang melebihi batas aman. Maka, program akan menampilkan pesan: "Selisih berat antara kantong kiri dan kanan melebihi 9 kg." Jika kemudian pengguna memasukkan berat kantong kiri sebesar 9 kg dan kantong kanan sebesar 2 kg, program akan menghentikan eksekusi dengan menampilkan pesan "Proses selesai." karena berat kantong kiri telah mencapai 9 kg.

### 4. Diberikan sebuah persamaan sebagai berikut ini.<br/>
![image](https://github.com/user-attachments/assets/ddea519e-dd44-437b-9403-a5d66b26462a)<br/>
Buatlah sebuah program yang menerima input sebuah bilangan sebagai K, kemudian menghitung dan menampilkan nilai f(K) sesuai persamaan di atas. Akar 2 merupakan bilangan Irasional. Meskipun demikian, nilai tersebut dapat dihampiri dengan rumus beriku:<br/>
![image](https://github.com/user-attachments/assets/f174cfd9-2bc7-4383-8cae-d87a9d80648a)<br/>
Modifikasi program sebelumnya yang menerima input integer K dan menghitung akar 2 untuk K tersebut. Hampiran akar 2 dituliskan dalam ketelitian 10 angka di belakang koma.

```go
package main

import (
	"fmt"
)

// Fungsi untuk menghitung nilai hampiran akar 2 berdasarkan iterasi K
func calculateSqrt2(k int) float64 {
	product := 1.0 // Mulai dengan nilai produk 1

	// Iterasi dari 0 hingga K
	for i := 0; i <= k; i++ {
		numerator := (4*float64(i) + 2) * (4*float64(i) + 2) // (4k + 2)^2
		denominator := (4*float64(i) + 1) * (4*float64(i) + 3) // (4k + 1)(4k + 3)
		product *= numerator / denominator
	}

	return product // Mengembalikan nilai hasil perkalian
}

func main() {
	var k int

	// Membaca input nilai K
	fmt.Print("Masukkan nilai K: ")
	fmt.Scan(&k)

	// Menghitung nilai hampiran sqrt(2)
	result := calculateSqrt2(k)

	// Menampilkan hasil dengan 10 angka di belakang koma
	fmt.Printf("Nilai akar 2 = %.10f\n", result)
}
```
### Output: ![image](https://github.com/user-attachments/assets/f8d88a43-db9f-445d-b5ff-ea3a6e90038c)

Kode di atas untuk menghitung nilai hampiran akar 2 menggunakan metode iterasi. Dalam program ini, terdapat fungsi `calculateSqrt2` yang menerima parameter integer k, yang menentukan jumlah iterasi yang dilakukan untuk mendekati nilai akar 2. Fungsi ini menghitung hasil perkalian berdasarkan rumus matematis tertentu, di mana numerator dihitung dengan rumus 
(4k+2) pangkat 2 dan denominator dengan (4k+1)(4k+3). Hasil dari setiap iterasi dikalikan secara berurutan untuk menghasilkan nilai hampiran yang semakin akurat. Di dalam fungsi `main`, program meminta pengguna untuk memasukkan nilai k, kemudian memanggil fungsi `calculateSqrt2` untuk menghitung hasilnya. Akhirnya, program menampilkan hasil tersebut dengan format 10 angka di belakang koma. Dengan demikian, pengguna dapat memperoleh nilai hampiran dari akar 2 berdasarkan jumlah iterasi yang ditentukan, memberikan gambaran tentang sifat irasional dari angka tersebut.<br/>

### 5. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut!<br/> Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.<br/>

```go
package main

import (
	"fmt"
)

func hitungBiayaKirim(berat int) int {
	// Menghitung berat dalam kilogram dan gram
	kg := berat / 1000
	gram := berat % 1000

	// Biaya pengiriman per kilogram
	biayaPerKg := 10000
	biayaTotal := kg * biayaPerKg

	// Biaya tambahan untuk sisa gram
	biayaTambahan := 0
	if kg >= 10 {
		biayaTambahan = 0 // Gratis biaya tambahan jika berat lebih dari 10 kg
	} else {
		if gram >= 500 {
			biayaTambahan = gram * 5 // Rp. 5 per gram jika sisa >= 500 gram
		} else {
			biayaTambahan = gram * 15 // Rp. 15 per gram jika sisa < 500 gram
		}
	}

	// Total biaya
	return biayaTotal + biayaTambahan
}

func main() {
	var berat int

	// Meminta input berat dari pengguna
	fmt.Print("Berat parsel (gram): ")
	fmt.Scan(&berat)

	// Menghitung berat dalam kg dan gram
	kg := berat / 1000
	gram := berat % 1000

	// Menghitung total biaya pengiriman
	biayaPerKg := 10000 * kg
	biayaTambahan := 0

	// Kondisi untuk biaya tambahan berdasarkan sisa berat gram
	if kg >= 10 {
		biayaTambahan = 0 // Gratis biaya tambahan jika lebih dari 10 kg
	} else {
		if gram >= 500 {
			biayaTambahan = gram * 5 // Rp. 5 per gram untuk sisa >= 500 gram
		} else {
			biayaTambahan = gram * 15 // Rp. 15 per gram untuk sisa < 500 gram
		}
	}

	// Menghitung total biaya
	totalBiaya := biayaPerKg + biayaTambahan

	// Menampilkan hasil
	fmt.Printf("Detail berat: %d kg + %d gr\n", kg, gram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaPerKg, biayaTambahan)
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}
```

## Output: ![image](https://github.com/user-attachments/assets/59368d3d-ff15-4887-a01a-50e366329ddf)

Kode di atas berfungsi untuk menghitung biaya pengiriman parsel berdasarkan berat dalam gram dengan ketentuan yang diberikan. Program dimulai dengan meminta pengguna untuk memasukkan berat parsel dalam satuan gram. Setelah input diberikan, berat tersebut dipecah menjadi dua bagian: berat dalam kilogram (kg) dan sisa berat dalam gram. Biaya pengiriman dasar adalah Rp. 10.000,- per kilogram, dan biaya tambahan dikenakan untuk sisa berat dalam gram. Jika berat sisa adalah 500 gram atau lebih, tambahan biaya dikenakan Rp. 5,- per gram, sementara jika kurang dari 500 gram, biaya tambahan adalah Rp. 15,- per gram. Namun, apabila berat total parsel lebih dari 10 kilogram, maka biaya tambahan untuk sisa gram digratiskan. Setelah menghitung biaya untuk kilogram dan gram, program menampilkan rincian berat dan biaya dalam format yang terperinci, serta menampilkan total biaya pengiriman.

### 6. Diberikan sebuah nilai akhir mata kuliah (NAM) [0..100] dan standar penilaian nilai mata kuliah (NMK) sebagai berikut:<br/> ![image](https://github.com/user-attachments/assets/05fa8371-e853-40d9-8ee4-193873afcc1c)<br/> Program berikut menerima input sebuah bilangan rill yang menyatakan NAM. Program menghitung NMK dan menampilkanny.<br/>

```go
package main

import "fmt"

func main() {
	var nam float64
	var nmk string

	// Meminta input nilai
	fmt.Print("Nilai akhir mata kuliah: ")
	fmt.Scan(&nam)

	// Logika penentuan nilai huruf berdasarkan nilai numerik
	if nam > 80 {
		nmk = "A"
	} if nam > 72.5 {
		nmk = "AB"
	} if nam > 65 {
		nam = "B"
	} if nam > 57.5 {
		nmk = "BC"
	} if nam > 50 {
		nmk = "C"
	} if nam > 40 {
		nmk = "D"
	} else if nam <= 40 {
		nam = "E"
	}

	// Menampilkan hasil
	fmt.Printf("Nilai mata kuliah:" ,nmk)
}
```

## Jawablah pertanyaan-pertanyaan berikut:<br/>a. Jika nam diberikan adalah 80.1, apakah keluaran dari 











