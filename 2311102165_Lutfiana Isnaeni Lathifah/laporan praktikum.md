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
  Lutfiana Isnaeni L /2311102165<br>
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
## Dasar Teori

Berikan penjelasan teori terkait materi modul ini dengan bahasa anda sendiri serta susunan yang terstruktur per topiknya.

## Guided 

### 1. [Nama Topik]

```C++
package main

import (
	"fmt"
)

func main() {

	var nama string

	fmt.Scanln(&nama)
	fmt.Println(nama)
}
```
#### Output:
![240302_00h00m06s_screenshot]()

Kode di atas digunakan untuk mengambil input dari pengguna dan mencetaknya. Pertama, variabel nama dideklarasikan sebagai string untuk menyimpan input. Kemudian, fungsi fmt.Scanln(&nama) digunakan untuk menerima input dari pengguna dan menyimpannya ke dalam variabel nama. Setelah itu, fungsi fmt.Println(nama) mencetak input yang telah dimasukkan oleh pengguna. Kode ini efektif untuk menerima dan menampilkan data teks yang dimasukkan selama program berjalan.

### 2. [Nama Topik]

```C++
package main

import "fmt"

func main() {

	var a, b, c, d, e int
	var hasil int
	fmt.Scanln(&a, &b, &c, &d, &e)

	hasil = a + b + c + d + e
	fmt.Println("Hasil Penjualan", a, b, c, d, e, "adalah =", hasil)
}
```
#### Output:
![240302_00h00m06s_screenshot]()

Kode di atas digunakan untuk menghitung total penjualan dari lima input angka yang dimasukkan oleh pengguna. Pertama, lima variabel a, b, c, d, dan e dideklarasikan untuk menyimpan nilai integer dari input. Pengguna kemudian memasukkan lima angka yang disimpan ke dalam variabel tersebut menggunakan fmt.Scanln(&a, &b, &c, &d, &e). Setelah itu, program menjumlahkan kelima angka tersebut dan menyimpannya di variabel hasil. Terakhir, program mencetak hasil penjumlahan beserta angka-angka input dengan format "Hasil Penjualan (angka-angka) adalah = (hasil)".

### 3. [Nama Topik]

```C++
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
		fmt.Println("BERHASIL : true")
	} else {
		fmt.Println("BERHASIL : false")
	}
}
```
#### Output:
![240302_00h00m06s_screenshot]()

Kode di atas digunakan untuk memeriksa apakah pengguna dapat memasukkan urutan warna yang benar sebanyak lima kali percobaan. Urutan warna yang benar adalah "merah", "kuning", "hijau", "ungu". Pada setiap percobaan, program meminta input dari pengguna berupa serangkaian warna yang dipisahkan oleh spasi. Program kemudian memisahkan input tersebut dan memeriksa apakah urutannya sesuai dengan urutan yang benar. Jika ada kesalahan pada salah satu percobaan, program akan keluar dari loop dan menampilkan "BERHASIL : false". Jika semua percobaan benar, program akan menampilkan "BERHASIL : true".

### 4. [Nama Topik]

```C++
package main

import "fmt"

func main() {
	var nam float32
	var nmk string

	// Meminta input nilai
	fmt.Print("Masukkan nilai : ")
	fmt.Scan(&nam)

	// Logika penentuan nilai huruf berdasarkan nilai numerik
	if nam > 80 {
		nmk = "A"
	} else if nam > 72.5 {
		nmk = "B"
	} else if nam > 65 {
		nmk = "C"
	} else if nam > 50 {
		nmk = "D"
	} else if nam > 40 {
		nmk = "E"
	} else {
		nmk = "F"
	}

	// Menampilkan hasil
	fmt.Printf("Nilai Indeks untuk nilai %.2f adalah %s\n", nam, nmk)
}
```
#### Output:
![240302_00h00m06s_screenshot]()

Kode di atas digunakan untuk mengonversi nilai numerik yang diinput oleh pengguna menjadi nilai huruf berdasarkan rentang tertentu. Pengguna diminta memasukkan nilai numerik (tipe `float32`), kemudian program menentukan nilai huruf (`A`, `B`, `C`, `D`, `E`, atau `F`) berdasarkan logika kondisional `if-else`. Setiap rentang nilai memiliki indeks huruf yang berbeda, misalnya nilai di atas 80 mendapatkan "A", sedangkan nilai di bawah 40 mendapatkan "F". Setelah itu, hasilnya ditampilkan dengan format "Nilai Indeks untuk nilai (nilai numerik) adalah (nilai huruf)".

## Unguided 

### 1. [Soal]

C++
#include <iostream>
using namespace std;

int main() {
    cout << "ini adalah file code unguided praktikan" << endl;
    return 0;
}

#### Output:
![240302_00h00m06s_screenshot]()

Kode di atas digunakan untuk mencetak teks "ini adalah file code guided praktikan" ke layar menggunakan function cout untuk mengeksekusi nya.

#### Full code Screenshot:
![240309_10h21m35s_screenshot]()


