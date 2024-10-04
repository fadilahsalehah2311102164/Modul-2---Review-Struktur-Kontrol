<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>

<br>

<h2 align="center"><strong>MODUL II</strong></h2>
<h2 align="center"><strong> REVIEW STRUKTUR KONTROL </strong></h2>

<br>

<p align="center">

  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Rangga Pradarrell Fathi / 2311102200<br>
  S1 IF 11 05
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu:</strong><br>
  Arif Amrulloh, S.Kom., M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong><br>
  <strong>FAKULTAS INFORMATIKA</strong><br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong><br>
  <strong>2024</strong>
</p>

------

## I. Dasar Teori
### Pengertian Go
GO adalah bahasa pemrograman yang digunakan untuk pembuatan aplikasi berat, seperti sistem operasi dan aplikasi web, sehingga programmer yang memahami GO dapat memanfaatkan peluang karir sebagai pembuat aplikasi berat[1].

## 1. Tipe Data

### Integer (Bilangan Bulat)
Go menyediakan beberapa tipe integer:
- `int`: ukuran tergantung platform (32 atau 64 bit)
- `int8`: -128 hingga 127
- `int16`: -32768 hingga 32767
- `int32`: -2147483648 hingga 2147483647
- `int64`: -9223372036854775808 hingga 9223372036854775807
- `uint`: unsigned integer, tergantung platform
- `uint8`: 0 hingga 255
- `uint16`: 0 hingga 65535
- `uint32`: 0 hingga 4294967295
- `uint64`: 0 hingga 18446744073709551615

Contoh penggunaan:
```go
var angka int = 42
var umur uint8 = 25
```

### Float (Bilangan Desimal)
Go memiliki dua tipe float:
- `float32`: ±1.18×10⁻³⁸ hingga ±3.4×10³⁸
- `float64`: ±2.23×10⁻³⁰⁸ hingga ±1.80×10³⁰⁸

Contoh penggunaan:
```go
var pi float64 = 3.14159265359
var suhu float32 = 27.5
```

### String
String di Go adalah urutan karakter UTF-8 yang tidak dapat diubah (immutable).

Fitur string di Go:
1. Dapat menggunakan tanda kutip ganda (`"`) untuk string biasa
2. Dapat menggunakan tanda kutip balik (`` ` ``) untuk raw string

Contoh penggunaan:
```go
var nama string = "Budi"
var alamat string = `Jalan Merdeka No. 17
Kota Jakarta`  // Raw string bisa multi-baris
```

### Boolean
Tipe bool memiliki dua nilai: `true` atau `false`

Contoh penggunaan:
```go
var lulus bool = true
var menikah bool = false
```

## 2. Percabangan

### If-Else
Struktur if-else di Go memiliki beberapa fitur unik:
1. Tidak memerlukan tanda kurung `()` untuk kondisi
2. Kurung kurawal `{}` wajib digunakan
3. Dapat memiliki statement singkat sebelum kondisi

Contoh:
```go
// If-else biasa
if nilai >= 75 {
    fmt.Println("Lulus")
} else {
    fmt.Println("Tidak lulus")
}

// If dengan statement singkat
if skor := hitungSkor(); skor > 100 {
    fmt.Println("Skor maksimum!")
}
```

### Switch
Switch di Go lebih fleksibel dibanding bahasa lain:
1. Tidak perlu `break` di setiap case
2. Case bisa memiliki multiple values
3. Bisa menggunakan kondisi di setiap case

Contoh:
```go
// Switch biasa
switch buah {
case "apel":
    fmt.Println("Ini apel")
case "jeruk", "lemon":  // Multiple values
    fmt.Println("Ini jeruk atau lemon")
default:
    fmt.Println("Buah tidak dikenal")
}

// Switch dengan kondisi
switch {
case nilai >= 90:
    fmt.Println("A")
case nilai >= 80:
    fmt.Println("B")
default:
    fmt.Println("C")
}
```

## 3. Perulangan

Go hanya memiliki satu konstruksi perulangan: `for`. Namun, ini bisa digunakan dalam beberapa cara:

### For Tradisional
```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```
Komponen:
1. Statement inisialisasi: `i := 0`
2. Kondisi: `i < 5`
3. Statement post: `i++`

### For sebagai While
```go
counter := 0
for counter < 5 {
    fmt.Println(counter)
    counter++
}
```
Hanya menggunakan kondisi, mirip `while` di bahasa lain.

### For Range
Digunakan untuk iterasi melalui struktur data:
```go
slice := []string{"apel", "jeruk", "mangga"}
for index, nilai := range slice {
    fmt.Printf("Index: %d, Nilai: %s\n", index, nilai)
}
```
Dapat digunakan pada:
1. Array atau slice
2. String
3. Map
4. Channel

### Infinite Loop
```go
for {
    fmt.Println("Loop tanpa henti")
    // Gunakan 'break' untuk keluar dari loop
}
```

## Tips Tambahan
1. Gunakan `break` untuk keluar dari loop
2. Gunakan `continue` untuk melompat ke iterasi berikutnya
3. Label bisa digunakan dengan `break` dan `continue` untuk kontrol yang lebih spesifik

Contoh label:
```go
outerLoop:
for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
        if i*j > 10 {
            break outerLoop
        }
    }
}
```
## II. GUIDED
1. Menampilkan output yang diinputkan user

#### Source Code
```go
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
#### Screenshoot Source Code


#### Screenshoot Output


#### Deskripsi Program


2. Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah 'merah', 'kuning', 'hijau', dan 'ungu' selama 5 kali percobaan berulang.
Buatlah sebuah program yang menerima Input berupa warna dari ke 4 gelas reaks! sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuat dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.

#### Source Code
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
		fmt.Println("BERHASIL : true")
	} else {
		fmt.Println("BERHASIL : false")
	}
}

```
#### Screenshoot Source Code


#### Screenshoot Output


#### Deskripsi Program


3. Menambahkan bilangan yang diinputkan user

#### Source Code
```go
package main

import (
	"fmt" 
)

func main() {
	
	var a, b, c, d, e int
	var hasil int 

	fmt.Scanln(&a, &b, &c, &d, &e)
	
	hasil = a + b + c + d + e

	fmt.Println("Hasil Penjumlahan ", a, b, c, d, e, "adalah = ", hasil)
}
```
#### Screenshoot Source Code


#### Screenshoot Output


#### Deskripsi Program


4. Diberikan sebuah nilai akhir mata kullah (NAM) [0..100] dan standar penilaian nilai mata kullah (NMK)
   sebagai berikut:



#### Source Code
```go
package main

import "fmt"

func main() {
	var nam float32
	var nmk string

	fmt.Print("Masukkan Nilai: ")
	fmt.Scan(&nam)

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

	fmt.Printf("Nilai indeks untuk nilai %.2f adalah %s\n", nam, nmk)
}
```
#### Screenshoot Source Code

#### Screenshoot Output

#### Deskripsi Program

```go
package main

import "fmt"

func main() {
	var nam float64 
	var nmk string

	fmt.Print("Masukkan Nilai: ")
	fmt.Scan(&nam)

	if nam >= 80 {
		nmk = "A"
	} else if nam >= 65 {
		nmk = "B"
	} else if nam >= 50 {
		nmk = "C"
	} else if nam >= 40 {
		nmk = "D"
	} else {
		nmk = "F"
	}

	fmt.Printf("Nilai indeks untuk nilai %.2f adalah %s\n", nam, nmk)
}

```


## III. UNGUIDED

1. Suatu pita (string) berisi kumpulan nama-nama bunga yang dipisahkan oleh spasi dan '-', contoh pita diilustrasikan seperti berikut ini.

Pita: mawar - melati - tulip - teratai - kamboja - anggrek

Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita.

(Petunjuk: gunakan operasi penggabungan string dengan operator "+").

Tampilkan isi pita setelah proses input selesai.

Modifikasi program sebelumnya, proses input akan berhenti apabila user mengetikkan 'SELESAI'. Kemudian tampilkan isi pita beserta banyaknya bunga yang ada di dalam pita.

#### Source Code
```go


```
#### Screenshoot Source Code


#### Screenshoot Output


#### Deskripsi Program


2. Setiap hari Pak Andi membawa banyak barang belanjaan dari pasar dengan mengendarai sepeda motor. Barang belanjaan tersebut dibawa dalam kantong terpal di kiri-kanan motor. Sepeda motor tidak akan oleng jika selisih berat barang di kedua kantong sisi tidak lebih dari 9 kg.

Buatlah program Pak Andi yang menerima input dua buah bilangan real positif yang menyatakan berat total masing-masing isi kantong terpal. Program akan terus meminta input bilangan tersebut hingga salah satu kantong terpal berisi 9 kg atau lebih.

Pada modifikasi program tersebut, program akan menampilkan true Jika selisih kedua isi kantong lebih dari atau sama dengan 9 kg. Program berhenti memproses apabila total berat isi kedua kantong melebihi 150 kg atau salah satu kantong beratnya negatif.


#### Source Code
```go

```
#### Screenshoot Source Code


#### Screenshoot Output


#### Deskripsi Program


3. Diberikan sebuah persamaan sebagai berikut ini.
   Buatlah sebuah program yang menerima input sebuah bilangan sebagai K, kemudian menghitung dan menampilkan nilai f(K) sesuai persamaan di atas.



#### Source Code
```go



```
#### Screenshoot Source Code


#### Screenshoot Output


#### Deskripsi Program
 

4. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut!

Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.


#### Source Code
```go

```
#### Screenshoot Source Code


#### Screenshoot Output


#### Deskripsi Program
Program ini berisi tentang menghitung dan mengonversikan berat yang dimasukkan oleh user dengan menggunakan gram dan kilogram. Program ini juga menghitung biaya pengiriman dari ketentuan yang ditentukan.


5. Sebuah bilangan bulat b memiliki faktor bilangan f > O jika f habis membagi b. Contoh: 2 merupakan faktor dari bilangan 6 karena 6 habis dibagi 2.

Buatlah program yang menerima input sebuah bilangan bulat b dan b> 1. Program harus dapat mencari dan menampilkan semua faktor dari bilangan tersebut!

Bilangan bulat b > 0 merupakan bilangan prima p jika dan hanya jika memiliki persis dua faktor bilangan saja, yaitu 1 dan dirinya sendiri.

Lanjutkan program sebelumnya. Setelah menerima masukan sebuah bilangan bulat b > 0. Program tersebut mencari dan menampilkan semua faktor bilangan tersebut. Kemudian, program menentukan apakah b merupakan bilangan prima.

#### Source Code
```go


```
#### Screenshoot Source Code


#### Screenshoot Output

#### Deskripsi Program
Program ini berisi tentang menentukan bilangan yang dimasukkan termasuk bilangan prima atau bukan. Pada awal program, user diminta untuk memasukkan bilangan bulat terlebih dahulu


## Referensi 
[1] Wali, M., Nengsih, TA, Hts, DIG, Choirina, P., Awaludin, AAR, Yusuf, M., ... & Baradja, A. (2023). Pengantar 15 Bahasa Pemrograman Terbaik Masa Depan (Referensi & Coding Untuk Pemula) . PT. Sonpedia Penerbitan Indonesia.
------
