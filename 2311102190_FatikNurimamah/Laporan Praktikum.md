<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL II</strong></h2>
<h2 align="center"><strong> REVIEW STRUKTUR KONTROL </strong></h2>

<br>

<p align="center">
  
  <img src="https://github.com/user-attachments/assets/741cb565-774a-4298-b1fb-22ebf35822f1" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Fatik Nurimamah / 2311102190<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------


## Daftar Isi
1. [Dasar Teori](#dasar-teori)
2. [Guided](#guided)
3. [Unguided](#unguided)
4. [Kesimpulan](#kesimpulan)
5. [Daftar Pustaka](#daftar-pustaka)

## Dasar Teori
Struktur kontrol dari Go berkaitan dengan bahasa C namun berbeda dalam hal-hal tertentu. Tidak ada pengulangan do atau while, hanya for; switch yang lebih fleksibel; if dan switch bisa menggunakan perintah inisialisasi seperti halnya pada for; perintah break dan continue memiliki label identifikasi yang opsional; dan ada beberapa kontrol struktur baru termasuk switch pada tipe dan komunikasi multiplexer, select. Sintaksnya juga sedikit berbeda: tidak ada tanda kurung dan bagian badan dari kontrol harus selalu dibatasi oleh kurung kurawal.

**1. `if`**
   
Dalam Go, bentuk if yang paling sederhana adalah sebagai berikut:

```go
if x > 0 {
	return y
}
```












## Guided 

### 1. Program Sederhana untuk Membaca dan Menampilkan Nama

### Source Code :
```go
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
```
### Output:
![Screenshot 2024-10-04 225448](https://github.com/user-attachments/assets/b72455a1-8a1d-46da-b35f-8a0275f0f7a8)

### Full code Screenshot:
![Screenshot 2024-10-04 225654](https://github.com/user-attachments/assets/5bfcd0c0-e51f-45ac-b1d8-9fe3c1f264ee)

### Deskripsi Program : 

### Algoritma Program

### Cara Kerja Program


### 2. Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. 
**Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah 'merah', 'kuning', 'hijau', dan 'ungu' selama 5 kali percobaan berulang.Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya**
 ![Screenshot 2024-10-04 230120](https://github.com/user-attachments/assets/bda3c6e1-fb3c-4a49-a36c-b0f90bb1903d)

 ### Source Code :
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Urutan warna yang benar
	correctOrder := []string{"merah", "kuning", "hijau", "ungu"}

	// Membaca input untuk 5 percobaan
	reader := bufio.NewReader(os.Stdin)
	success := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)

		// Membaca input dari pengguna
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		// Memisahkan input berdasarkan spasi
		colors := strings.Split(input, " ")

		// Mengecek apakah urutan warna sesuai
		for j := 0; j < 4; j++ {
			if colors[j] != correctOrder[j] {
				success = false
				break
			}
		}

		// Jika ada percobaan yang tidak sesuai, keluar dari loop
		if !success {
			break
		}
	}

	// Menampilkan hasil
	if success {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}
}
```
## Output:
![Screenshot 2024-10-04 232908](https://github.com/user-attachments/assets/1d18fa8a-bde5-412b-84fe-116ab4b62667)

### Full code Screenshot:
![Screenshot 2024-10-04 232952](https://github.com/user-attachments/assets/d6f86650-497a-4dfb-91b8-b4dc27bf7ae8)

### Deskripsi Program : 

### Algoritma Program

### Cara Kerja Program


### 3. Penjumlahan 5 Angka dari Input Pengguna

### Source Code :
```go
package main 

import "fmt"

func main() {
	var a, b, c, d, e int
	var hasil int

	// Menambahkan prompt untuk input
	fmt.Print("Masukkan penjualan (a, b, c, d, e): ")
	fmt.Scanln(&a, &b, &c, &d, &e)

	hasil = a + b + c + d + e
	fmt.Println("Hasil Penjualan", a, b, c, d, e, "adalah =", hasil)
}

```
### Output:


### Full code Screenshot:


### Deskripsi Program : 

### Algoritma Program

### Cara Kerja Program

### 4.Diberikan sebuah nilai akhir mata kuliah (NAM) [0..100] dan standar penilaian nilai mata kuliah (NMK) sebagai berikut:

### Source Code :
```go

```
### Output:


### Full code Screenshot:


### Deskripsi Program : 

### Algoritma Program

### Cara Kerja Program



## Unguided 

### 1.Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita.


### Source Code :
```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Variabel pita digunakan untuk menyimpan daftar nama bunga
	var pita string
	var TotalBunga int

	// Menggunakan scanner untuk membaca input dari pengguna
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Masukkan nama bunga dan ketik 'SELESAI' untuk mengakhiri. :\n")

	for {
		TotalBunga++
		fmt.Printf("Bunga %d: ", TotalBunga)

		// Membaca input dari pengguna
		scanner.Scan()
		input := scanner.Text()

		// Cek jika pengguna mengetikkan "SELESAI" maka program akan berhenti
		if strings.ToUpper(input) == "SELESAI" {
			TotalBunga-- // Mengurangi jumlah bunga karena input "SELESAI" tidak dihitung
			break
		}

		// Menambahkan nama bunga ke pita 
		if pita == "" {
			pita = input
		} else {
			pita += " - " + input
		}
	}

	// Menampilkan hasil akhir pita dan jumlah bunga yang dimasukkan
	fmt.Println("Pita:", pita)
	fmt.Println("Total bunga:", TotalBunga )
}

```
### Output:
![Screenshot 2024-10-02 205457](https://github.com/user-attachments/assets/a9c25f50-1b96-432a-8a24-0ce28162d27f)

### Full code Screenshot:
![Screenshot 2024-10-04 231610](https://github.com/user-attachments/assets/aadc38e3-1a0e-4b22-96b0-d8ca9963825a)

### Deskripsi Program : 
Program ini digunakan untuk menerima input dari pengguna berupa nama-nama bunga, yang disimpan dalam variabel pita. Pengguna dapat terus memasukkan nama bunga sampai mengetikkan kata "SELESAI", yang akan menghentikan proses input. Setelah selesai, program akan menampilkan daftar lengkap nama bunga yang dimasukkan serta total bunga yang dihitung (tidak termasuk input "SELESAI").

### Algoritma Program
1. Inisialisasi variabel pita sebagai string kosong dan TotalBunga sebagai penghitung jumlah bunga.
2. Gunakan loop untuk terus menerima input dari pengguna.
3. Di dalam loop: Tambahkan nama bunga yang dimasukkan ke dalam variabel pita. Dan jika pengguna mengetik "SELESAI", akhiri loop dan kurangi penghitung jumlah bunga karena input ini tidak dianggap sebagai bunga.
4. Setelah keluar dari loop, tampilkan daftar lengkap bunga yang dimasukkan serta jumlah total bunga yang dihitung.

### Cara Kerja Program
1. Input Nama Bunga: Program meminta pengguna untuk memasukkan nama bunga satu per satu hingga pengguna mengetikkan kata "SELESAI".
2. Penggabungan Nama: Setiap nama bunga yang diberikan akan ditambahkan ke dalam variabel pita dan dipisahkan dengan tanda strip (" - ") untuk memudahkan pembacaan.
3. Penghitungan Jumlah Bunga: Setiap kali pengguna memasukkan nama bunga, penghitung TotalBunga bertambah satu, namun jika pengguna mengetik "SELESAI", penghitung ini dikurangi satu agar input "SELESAI" tidak dihitung sebagai bunga.
4. Output: Program menampilkan daftar nama-nama bunga yang dimasukkan pengguna dan jumlah total bunga yang valid setelah proses input berakhir.


### 2. Buatlah program Pak Andi yang menerima input dua buah bilangan real positif yang menyatakan berat total masing-masing isi kantong terpal. Program akan terus meminta input bilangan tersebut hingga salah satu kantong terpal berisi 9 kg atau lebih.


### Source Code :
```go
package main

import (
	"fmt"
	"math"
)

func main() {
	var BeratKantongKiri, BeratKantongKanan float64

	for {
		// Meminta pengguna untuk memasukkan berat kedua kantong
		fmt.Print("Masukkan berat barang di kedua kantong: ")
		fmt.Scan(&BeratKantongKiri, &BeratKantongKanan)

		// Mengecek apakah berat salah satu kantong bernilai negatif
		if BeratKantongKiri < 0 || BeratKantongKanan < 0 {
			fmt.Println("Proses selesai. Salah satu kantong memiliki berat negatif.")
			break
		}

		// Mengecek apakah total berat kedua kantong melebihi 150 kg
		TotalBeratKantong := BeratKantongKiri + BeratKantongKanan
		if TotalBeratKantong > 150 {
			fmt.Println("Proses selesai. Total berat kedua kantong melebihi 150 kg.")
			break
		}

		// Menghitung selisih berat antara kedua kantong
		SelisihBeratKantong := math.Abs(BeratKantongKiri -BeratKantongKanan)

		// Mengecek apakah selisih berat menyebabkan sepeda motor oleng
		if SelisihBeratKantong >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: True")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: False")
		}
	}
}

```
### Output:
![Screenshot 2024-10-02 214136](https://github.com/user-attachments/assets/4229d529-f363-492e-a72a-1129bb8546b6)

### Full code Screenshot:
![Screenshot 2024-10-04 231930](https://github.com/user-attachments/assets/09287ed6-180a-4da8-8d39-48588b456c58)

### Deskripsi Program : 

### Algoritma Program

### Cara Kerja Program


### 3. Diberikan sebuah persamaan sebagai berikut ini.


### Source Code :
```go
package main

import (
    "fmt"
)

// Fungsi untuk menghitung nilai f(k)
func f(k float64) float64 {
    pembilang := (4*k + 2) * (4*k + 2) 
    penyebut := (4*k + 1) * (4*k + 3)  
    return pembilang / penyebut
}

func main() {
    var k float64
    // Minta input nilai K dari user
    fmt.Print("Masukkan nilai K: ")
    fmt.Scanln(&k)

    // Hitung f(K) dan tampilkan hasilnya
    hasil := f(k)
    fmt.Printf("Nilai f(K) = %.10f\n", hasil)
}
```
### Output:
![Screenshot 2024-10-04 205308](https://github.com/user-attachments/assets/48ff47ca-61c1-4899-863d-e6cce4830006)

### Full code Screenshot:
![Screenshot 2024-10-05 103402](https://github.com/user-attachments/assets/b9d51301-9995-4786-9011-5ba4d16bf035)

### Deskripsi Program : 

### Algoritma Program

### Cara Kerja Program


### 4. 


### Source Code :
```go
package main

import (
	"fmt"
)

func BiayaPengiriman(beratGram int) int {
	// Konversi berat ke kg dan sisa gram
	BeratKg := beratGram / 1000
	SisaBeratGram := beratGram % 1000

	// Hitung biaya dasar untuk kg
	biayaDasar := BeratKg * 10000

	// Hitung biaya tambahan untuk sisa gram
	biayaTambahan := 0
	if SisaBeratGram > 0 {
		if BeratKg >= 10 {
			// Jika berat lebih dari 10 kg, sisa gram digratiskan
			biayaTambahan = 0
		} else if SisaBeratGram <= 500 {
			// Jika sisa gram <= 500, biaya tambahan Rp. 5,- per gram
			biayaTambahan = SisaBeratGram * 5
		} else {
			// Jika sisa gram > 500, biaya tambahan Rp. 15,- per gram
			biayaTambahan = SisaBeratGram * 15
		}
	}

	// Total biaya adalah biaya dasar + biaya tambahan
	totalBiaya := biayaDasar + biayaTambahan

	// Cetak detail
	fmt.Printf("Detail berat: %d kg + %d gr\n", BeratKg, SisaBeratGram)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biayaDasar, biayaTambahan)
	return totalBiaya
}

func main() {
	var berat int
	// Minta input berat dari pengguna
	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&berat)

	// Hitung biaya pengiriman
	totalBiaya := BiayaPengiriman(berat)

	// Tampilkan total biaya
	fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
}

```
### Output:
![Screenshot 2024-10-05 105040](https://github.com/user-attachments/assets/dbe90925-f7db-431e-b38f-d8bafb0a680b)

### Full code Screenshot:
![Screenshot 2024-10-05 104036](https://github.com/user-attachments/assets/f5512033-3c8b-40a3-ac13-7408d7983d83)

### Deskripsi Program : 

### Algoritma Program

### Cara Kerja Program


### 5. 

### Source Code :
```go
package main

import (
	"fmt"
)

// Fungsi untuk menemukan semua faktor dari bilangan b
func cariFaktor(b int) []int {
	var daftarfaktor []int
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			daftarfaktor = append(daftarfaktor, i)
		}
	}
	return daftarfaktor
}

// Fungsi untuk memeriksa apakah bilangan b adalah bilangan prima.
func cekBilPrima(b int) bool {
	if b <= 1 {
		return false
	}
	for i := 2; i*i <= b; i++ {
		if b%i == 0 {
			return false
		}
	}
	return true
}

func main() {
	var b int
	fmt.Print("Bilangan: ")
	fmt.Scan(&b)

	// untuk mencari faktor dan menampilkannya
	faktor := cariFaktor(b)
	fmt.Print("Faktor: ")
	for _, f := range faktor {
		fmt.Print(f, " ")
	}
	fmt.Println()

	// untuk mengecek dan menampilkan apakah bilangan tersebut prima
	if cekBilPrima(b) {
		fmt.Println("Prima: true")
	} else {
		fmt.Println("Prima: false")
	}
}

```
### Output:
![Screenshot 2024-10-04 211706](https://github.com/user-attachments/assets/776fbe4f-bc4d-4220-afb8-e79f17793768)

### Full code Screenshot:
![Screenshot 2024-10-05 105707](https://github.com/user-attachments/assets/933b7e72-f02b-433e-a1e6-cf4ac91fd4c3)

### Deskripsi Program : 

### Algoritma Program

### Cara Kerja Program




## Daftar Pustaka
[1] Effective Go. (n.d.). Retrieved from Golang Indonesia: https://golang-id.org/doc/effective_go.html 










 
