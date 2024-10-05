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
  Wahyu Hidayat / 2311102178<br>
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
Go, atau Golang, adalah bahasa pemrograman yang dirancang untuk kesederhanaan dan efisiensi dalam pengembangan perangkat lunak. Menurut Donovan dan Kernighan (2015), Go menawarkan sintaks yang bersih, kinerja tinggi melalui kompilasi langsung, dan dukungan untuk pemrograman paralel. Dengan sistem garbage collection dan ekosistem yang kaya, Go cocok untuk aplikasi yang membutuhkan skalabilitas dan responsivitas tinggi[1].

## 1. Struktur Program Go
Struktur program dalam bahasa Go memiliki beberapa komponen kunci:

Package Declaration: Setiap file Go dimulai dengan mendeklarasikan paketnya. Misalnya, package main berarti file ini adalah bagian dari paket utama yang bisa dieksekusi[1].

Import Statement: Digunakan untuk mengimpor paket lain yang dibutuhkan. Contoh, import "fmt" digunakan untuk mengimpor paket format yang memungkinkan kita menggunakan fungsi-fungsi seperti fmt.Println[1].

Main Function: Setiap program Go memerlukan fungsi main(), yang merupakan titik awal eksekusi program. Tanpa fungsi ini, program tidak akan dapat berjalan[1].

Contoh Struktur Program:
```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}

```

## 2. Tipe Data dan Instruksi Dasar
Go memiliki beberapa tipe data dasar, antara lain:

Integer: Tipe data untuk angka bulat, seperti int, int8, int16, int32, int64.
Floating Point: Tipe data untuk angka desimal, seperti float32 dan float64.
Boolean: Tipe data yang hanya memiliki dua nilai, yaitu true dan false.
String: Tipe data untuk teks.
Array: Koleksi elemen dengan tipe yang sama.
Slice: Mirip dengan array tetapi lebih fleksibel, dapat berubah ukuran.
Map: Koleksi pasangan kunci-nilai.
Contoh penggunaan tipe data dan instruksi dasar:
```go
package main

import "fmt"

func main() {
    var nama string = "John"
    var umur int = 30
    var tinggi float64 = 175.5
    var aktif bool = true

    fmt.Println("Nama:", nama)
    fmt.Println("Umur:", umur)
    fmt.Println("Tinggi:", tinggi)
    fmt.Println("Aktif:", aktif)
}
```
## 3. Struktur Kontrol Perulangan
Struktur kontrol perulangan memungkinkan kita untuk mengulangi eksekusi blok kode. Dalam Go, kita hanya memiliki satu jenis perulangan, yaitu for. for dapat digunakan dalam beberapa cara[1]:

For dengan Inisialisasi dan Kondisi: Ini mirip dengan loop di banyak bahasa lain.
For Tanpa Kondisi: Ini akan menciptakan loop tak terbatas, jika tidak dihentikan dengan break[1].
For dengan Range: Digunakan untuk iterasi melalui elemen dalam slice, array, atau map[1].
Contoh perulangan menggunakan for:
```go
package main

import "fmt"

func main() {
    // Perulangan dengan inisialisasi dan kondisi
    for i := 1; i <= 5; i++ {
        fmt.Println("Iterasi ke", i)
    }

    // Perulangan tanpa kondisi (loop tak terbatas)
    j := 1
    for {
        if j > 5 {
            break // Menghentikan perulangan
        }
        fmt.Println("Iterasi ke", j)
        j++
    }

    // Perulangan dengan range
    angka := []int{1, 2, 3, 4, 5}
    for index, value := range angka {
        fmt.Printf("Index: %d, Value: %d\n", index, value)
    }
}

```

## 4. Struktur Kontrol Percabangan
Percabangan memungkinkan program untuk membuat keputusan berdasarkan kondisi tertentu. Dalam Go, kita memiliki beberapa cara untuk melakukan percabangan:

If Statement: Memeriksa kondisi dan menjalankan blok kode jika kondisi terpenuhi.
Else If dan Else: Menambahkan lebih banyak kondisi yang dapat diperiksa.
Switch Statement: Memeriksa satu nilai terhadap beberapa kemungkinan[1].
Contoh penggunaan percabangan:
```go
package main

import "fmt"

func main() {
    angka := 10

    // Menggunakan if-else
    if angka > 10 {
        fmt.Println("Angka lebih besar dari 10")
    } else if angka == 10 {
        fmt.Println("Angka sama dengan 10")
    } else {
        fmt.Println("Angka lebih kecil dari 10")
    }

    // Menggunakan switch
    hari := 3
    switch hari {
    case 1:
        fmt.Println("Hari Senin")
    case 2:
        fmt.Println("Hari Selasa")
    case 3:
        fmt.Println("Hari Rabu")
    default:
        fmt.Println("Hari lain")
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


2. Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah 'merah', 'kuning', 'hijau', dan 'ungu' selama 5 kali percobaan berulang[1].
Buatlah sebuah program yang menerima Input berupa warna dari ke 4 gelas reaks! sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuat dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya[1].

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
[1]Donovan, A. A., & Kernighan, B. W. (2015). The Go Programming Language. Addison-Wesley.
------
