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
GO adalah bahasa pemrograman yang digunakan untuk pembuatan aplikasi berat, seperti sistem operasi dan aplikasi web, sehingga programmer yang memahami GO dapat memanfaatkan peluang karir sebagai pembuat aplikasi berat. 

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

## Break
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
### 1. Menampilkan output yang diinputkan user

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
 <img src="https://github.com/user-attachments/assets/f944cd6d-089c-4e13-9a1f-f79788ffe3b9">

#### Screenshoot Output
 <img src="https://github.com/user-attachments/assets/e26f44de-25ec-4a77-886b-66c2087cc198">

#### Deskripsi Program
##### Program di atas adalah program sederhana berbasis bahasa Go yang berfungsi untuk membaca input berupa string dari pengguna, kemudian mencetak string tersebut kembali. Program ini hanya menggunakan satu variabel untuk menyimpan input dari pengguna.

##### Algoritma Program
- Deklarasi variabel nama sebagai string.
- Baca input dari pengguna menggunakan fmt.Scanln, dan simpan ke dalam variabel nama.
- Cetak kembali nilai dari variabel nama menggunakan fmt.Println.

##### Cara Kerja Program

*Input Pengguna:*


Program menunggu pengguna untuk memasukkan sebuah teks (string). Misalnya, pengguna memasukkan John Doe dan menekan "Enter".

*Menyimpan Input:*


Input dari pengguna, misalnya John Doe, akan disimpan dalam variabel nama.

*Mencetak Input:*


Program kemudian mencetak kembali nilai dari variabel nama, yang berisi string yang dimasukkan oleh pengguna. Jadi outputnya akan menjadi:
`John Doe`


### 2. Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah 'merah', 'kuning', 'hijau', dan 'ungu' selama 5 kali percobaan berulang.
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
 <img src="https://github.com/user-attachments/assets/875c415a-e9ea-48b0-96c8-1342a0163337">

#### Screenshoot Output



#### Deskripsi Program
##### Program ini bertujuan untuk mengecek apakah pengguna bisa memasukkan urutan warna yang benar (seperti yang telah ditentukan) dalam lima kali percobaan. Program akan menghentikan percobaan lebih awal jika urutan warna yang dimasukkan salah, dan memberikan hasil akhir apakah pengguna berhasil atau tidak.



##### Algoritma Program
- Inisialisasi Urutan Warna yang Benar: Tentukan urutan warna yang benar yaitu merah, kuning, hijau, ungu.

- Memulai Loop dengan 5 Percobaan: Jalankan loop sebanyak 5 kali untuk memberikan kesempatan kepada pengguna memasukkan urutan warna.

- Baca Input Pengguna:
Setiap percobaan, program membaca input dari pengguna berupa urutan warna yang dipisahkan oleh spasi.

- Bandingkan Urutan Input dengan Urutan yang Benar:

Pisahkan input menjadi array string.
Bandingkan urutan warna yang dimasukkan pengguna dengan urutan warna yang benar.

- Evaluasi Hasil:

Jika urutan salah, ubah variabel success menjadi false dan keluar dari loop.
Jika tidak ada kesalahan, lanjutkan ke percobaan berikutnya.

- Cetak Hasil:

Jika semua input benar, cetak "BERHASIL: true".
Jika ada kesalahan, cetak "BERHASIL: false".

##### Cara Kerja Program

- Input Pengguna: Program meminta pengguna untuk memasukkan urutan warna dalam satu baris. Pengguna harus memasukkan empat warna, dipisahkan oleh spasi, misalnya:
`merah kuning hijau ungu`

- Memeriksa Kesesuaian:
`Program membandingkan input pengguna dengan urutan warna yang benar,
Jika input sesuai, program akan melanjutkan ke percobaan berikutnya,
Jika input tidak sesuai, program menghentikan percobaan dan langsung mencetak hasil "BERHASIL: false".`

- Pengulangan:
` Program memberikan hingga 5 kesempatan (percobaan) kepada pengguna. Jika pengguna berhasil memasukkan urutan yang benar dalam semua percobaan, program mencetak hasil "BERHASIL: true". Namun, jika ada satu kali kesalahan, hasil langsung dinyatakan gagal.`


### 3. Menambahkan bilangan yang diinputkan user

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
<img src="https://github.com/user-attachments/assets/89fc8989-9c99-4b11-a275-aa4a47eff2d4">

#### Screenshoot Output


#### Deskripsi Program

##### Program ini bertujuan untuk:

- Menerima lima input angka integer dari pengguna.
- Menjumlahkan kelima angka tersebut.
- Menampilkan hasil penjumlahan beserta input yang dimasukkan oleh pengguna.



##### Algoritma Program
- Deklarasi variabel a, b, c, d, dan e untuk menyimpan input angka, serta variabel hasil untuk menyimpan hasil penjumlahan.
- Membaca input: Program meminta pengguna memasukkan lima angka integer.
- Menjumlahkan angka: Setelah input diterima, program menjumlahkan kelima angka tersebut.
- Menampilkan hasil: Program menampilkan hasil penjumlahan beserta lima angka yang dimasukkan.

##### Cara Kerja Program
1 Input Pengguna:
- Program meminta pengguna untuk memasukkan lima angka integer yang dipisahkan oleh spasi. Contoh input:
`10 20 30 40 50`

2 Penjumlahan:
- Setelah pengguna memasukkan angka-angka tersebut, program menjumlahkan angka-angka tersebut. Pada contoh input di atas:
`hasil = 10 + 20 + 30 + 40 + 50 = 150`

3 Mencetak Hasil:
- Setelah penjumlahan selesai, program menampilkan hasil penjumlahan dengan format yang rapi:
`Hasil Penjumlahan 10 20 30 40 50 adalah = 150`


### 4.  Diberikan sebuah nilai akhir mata kullah (NAM) [0..100] dan standar penilaian nilai mata kullah (NMK) sebagai berikut:
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
<img src="https://github.com/user-attachments/assets/1b173449-87f7-4687-ae3e-101e34510a8d">


#### Screenshoot Output

#### Deskripsi Program
##### Program ini bertujuan untuk:
- Meminta pengguna untuk memasukkan sebuah nilai dalam bentuk angka desimal (float).
- Berdasarkan nilai tersebut, program akan memberikan nilai huruf (indeks) sesuai dengan kriteria yang sudah ditentukan.
- Menampilkan nilai indeks berdasarkan nilai yang dimasukkan pengguna.

##### Algoritma Program
- Deklarasi variabel: Program mendeklarasikan variabel nam (float) untuk menyimpan nilai yang diinputkan oleh pengguna, dan nmk (string) untuk menyimpan nilai indeks berdasarkan hasil pengkondisian.
- Meminta input dari pengguna: Program meminta pengguna memasukkan nilai numerik.
- Pengkondisian: Program mengevaluasi nilai yang dimasukkan dengan serangkaian pernyataan if-else untuk menentukan indeks (A, B, C, D, E, F) berdasarkan rentang nilai yang telah ditentukan.
- Menampilkan hasil: Setelah indeks ditentukan, program akan mencetak nilai dan indeks yang sesuai ke layar.
  
##### Cara Kerja Program
1 Input Pengguna:
- Program meminta pengguna untuk memasukkan sebuah nilai numerik. Misalnya, pengguna memasukkan nilai 78.5.
2 Pemeriksaan Rentang Nilai:
- Setelah nilai diinput, program memeriksa dalam rentang mana nilai tersebut berada menggunakan struktur kontrol if-else.
3 Jika nilai yang dimasukkan adalah 78.5, maka program akan mengecek:
- Apakah nilai lebih dari 80? Tidak, lanjut ke berikutnya.
- Apakah nilai lebih dari 72.5? Ya, maka variabel nmk diberi nilai "B".
4 Menampilkan Nilai dan Indeks:
- Setelah pengecekan selesai, program menampilkan hasil akhir dengan nilai yang dimasukkan pengguna dan nilai indeks yang sesuai. Pada contoh di atas, hasilnya adalah:
`Nilai indeks untuk nilai 78.50 adalah B`

## III. UNGUIDED

1. Suatu pita (string) berisi kumpulan nama-nama bunga yang dipisahkan oleh spasi dan '-', contoh pita diilustrasikan seperti berikut ini.

Pita: mawar - melati - tulip - teratai - kamboja - anggrek

Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita.

(Petunjuk: gunakan operasi penggabungan string dengan operator "+").

Tampilkan isi pita setelah proses input selesai.

Modifikasi program sebelumnya, proses input akan berhenti apabila user mengetikkan 'SELESAI'. Kemudian tampilkan isi pita beserta banyaknya bunga yang ada di dalam pita.

#### Source Code
```go
//Rangga Pradarrell Fathi
//2311102200
//11-IF-05
package main

import (
	"fmt"
	"strings"
)

func main() {
	var rangkaianBunga string    // Variabel untuk menyimpan rangkaian nama bunga
	var jumlahBunga int = 0      // Variabel untuk menghitung jumlah bunga

	// Perulangan untuk meminta input nama bunga
	for i := 1; ; i++ {
		var namaBunga string
		fmt.Printf("Bunga %d: ", i)  
		fmt.Scan(&namaBunga)      

		// Mengecek apakah input adalah "SELESAI"
		if strings.ToUpper(namaBunga) == "SELESAI" {
			break                    
		}

		// Menambahkan nama bunga ke rangkaian
		if rangkaianBunga == "" {
			rangkaianBunga = namaBunga
		} else {
			rangkaianBunga += " - " + namaBunga 
		}
		jumlahBunga++  
	}

	// Menampilkan hasil
	fmt.Printf("Pita: %s\n", rangkaianBunga)    
	fmt.Printf("Bunga: %d\n", jumlahBunga)      
```
#### Screenshoot Source Code
<img src="https://github.com/user-attachments/assets/5567ae9d-c1c4-4d1d-a1e4-4d479d8f2313">


#### Screenshoot Output


#### Deskripsi Program
##### Program ini bertujuan untuk:
- Mengumpulkan rangkaian nama bunga yang diinputkan oleh pengguna.
- Menghitung jumlah bunga yang dimasukkan.
- Berhenti menerima input jika pengguna memasukkan kata "SELESAI" dan menampilkan rangkaian nama bunga serta jumlah bunga yang telah dimasukkan.

##### Algoritma Program
1 Inisialisasi variabel:
rangkaianBunga: Menyimpan rangkaian nama bunga.
2 jumlahBunga: 
Menghitung jumlah bunga yang dimasukkan.
3 Pengulangan untuk input nama bunga:
- Program meminta pengguna memasukkan nama bunga dalam perulangan tanpa batas.
- Setiap nama bunga yang diinputkan akan ditambahkan ke dalam rangkaian nama bunga dengan pemisah " - ".
- Jika pengguna memasukkan "SELESAI", perulangan berhenti.
4 Menyusun rangkaian nama bunga:
-Jika nama bunga adalah input pertama, disimpan langsung.
- Jika bukan yang pertama, tambahkan dengan pemisah " - ".
5 Menghitung jumlah bunga:
Setiap kali pengguna memasukkan nama bunga (selain "SELESAI"), variabel jumlahBunga akan bertambah.
6 Menampilkan hasil:
Setelah pengguna selesai menginput, program menampilkan rangkaian bunga dan jumlah bunga yang diinput.
  
##### Cara Kerja Program
1 Input Nama Bunga:
- Program meminta pengguna untuk memasukkan nama bunga secara berulang. Misalnya:
`Bunga 1: Mawar
Bunga 2: Melati
Bunga 3: Anggrek
Bunga 4: SELESAI
Input berakhir ketika pengguna memasukkan kata SELESAI (tidak sensitif terhadap huruf besar/kecil).`

2 Penggabungan Nama Bunga:
- Setiap nama bunga akan disusun ke dalam rangkaian dengan format "Bunga 1 - Bunga 2 - Bunga 3". Untuk contoh input di atas, hasil rangkaian menjadi:
`Mawar - Melati - Anggrek`

3 Menghitung Jumlah Bunga:
Program menghitung jumlah bunga yang dimasukkan. Jika pengguna memasukkan tiga bunga, variabel jumlahBunga akan menyimpan angka 3.

4 Menampilkan Hasil:
Setelah perulangan selesai, program akan menampilkan hasil rangkaian bunga dan jumlah bunga yang telah diinputkan. Hasilnya adalah:
`Pita: Mawar - Melati - Anggrek
Bunga: 3`


2. Setiap hari Pak Andi membawa banyak barang belanjaan dari pasar dengan mengendarai sepeda motor. Barang belanjaan tersebut dibawa dalam kantong terpal di kiri-kanan motor. Sepeda motor tidak akan oleng jika selisih berat barang di kedua kantong sisi tidak lebih dari 9 kg. Buatlah program Pak Andi yang menerima input dua buah bilangan real positif yang menyatakan berat total masing-masing isi kantong terpal. Program akan terus meminta input bilangan tersebut hingga salah satu kantong terpal berisi 9 kg atau lebih. Pada modifikasi program tersebut, program akan menampilkan true Jika selisih kedua isi kantong lebih dari atau sama dengan 9 kg. Program berhenti memproses apabila total berat isi kedua kantong melebihi 150 kg atau salah satu kantong beratnya negatif.


#### Source Code
```go
//Rangga Pradarrell Fathi
//2311102200
//11-IF-05
package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		var beratKantong1, beratKantong2 float64

		// Meminta input berat kedua kantong belanjaan
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		_, err := fmt.Scanf("%f %f", &beratKantong1, &beratKantong2)
		
		// Pengecekan apakah input valid
		if err != nil {
			fmt.Println("Input tidak valid. Mohon masukkan dua angka.")
			continue
		}

		// Pengecekan berat negatif
		if beratKantong1 < 0 || beratKantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		// Pengecekan total berat
		totalBerat := beratKantong1 + beratKantong2
		if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}

		// Pengecekan selisih berat
		selisihBerat := math.Abs(beratKantong1 - beratKantong2)
		akanOleng := selisihBerat >= 9

		// Menampilkan hasil apakah motor akan oleng
		fmt.Printf("Sepeda motor pak Andi akan oleng: %t\n", akanOleng)

		// Pengecekan apakah salah satu kantong mencapai atau melebihi 9 kg
		if beratKantong1 >= 9 || beratKantong2 >= 9 {
			fmt.Println("Proses selesai.")
			break
		}
	}
}
```
#### Screenshoot Source Code
<img src="https://github.com/user-attachments/assets/d1873b28-1b33-44c2-9aa4-1a9e9d1b8b40">


#### Screenshoot Output


#### Deskripsi Program
##### Program ini bertujuan untuk:
- Meminta pengguna memasukkan berat dua kantong belanjaan secara terus-menerus.
- Mengecek apakah motor Pak Andi akan oleng berdasarkan perbedaan berat antara kedua kantong belanjaan.
- Program akan berhenti dalam kondisi tertentu, seperti jika total berat melebihi 150 kg, berat negatif, atau jika salah satu kantong mencapai atau melebihi 9 kg.

##### Algoritma Program
1 Inisialisasi variabel:
- beratKantong1, beratKantong2: Variabel untuk menyimpan berat kedua kantong belanjaan.
Pengulangan tak terbatas dimulai.
2 Meminta input:
- Program meminta input dua angka (berat dua kantong belanjaan) dari pengguna.
3 Pengecekan validitas input:
- Jika input tidak valid (bukan angka), program meminta input ulang.
4 Pengecekan kondisi berat negatif:
- Jika salah satu berat kantong negatif, program berhenti.
5 Pengecekan total berat:
- Jika total berat lebih dari 150 kg, program berhenti.
6 Menghitung selisih berat:
- Program menghitung selisih antara kedua berat kantong dan mengecek apakah selisihnya 9 kg atau lebih.
- Jika selisih berat ≥ 9, motor akan dianggap oleng dan hasilnya ditampilkan.
7 Pengecekan berat kantong ≥ 9 kg:
- Jika salah satu kantong beratnya mencapai atau melebihi 9 kg, program berhenti.
8 Mengulangi proses:
- Jika tidak ada kondisi berhenti yang terpenuhi, program kembali meminta input dari pengguna.
  
##### Cara Kerja Program
1 Memasukkan Berat Kantong:
- Program meminta pengguna untuk memasukkan berat dua kantong belanjaan. Misalnya:
- Masukan berat belanjaan di kedua kantong: 8 7
  
2 Validasi Input:
- Jika input tidak valid (misalnya pengguna memasukkan huruf atau karakter non-angka), program akan menampilkan pesan kesalahan dan meminta input ulang:
- Input tidak valid. Mohon masukkan dua angka.
  
3 Menghitung Selisih Berat:
Jika input valid, program menghitung selisih berat antara kedua kantong. Misalnya:
- Berat kantong pertama = 8 kg, berat kantong kedua = 7 kg
- Selisih = 1 kg, karena ini kurang dari 9 kg, motor tidak akan oleng.

4 Program menampilkan hasil:
`Sepeda motor pak Andi akan oleng: false`

5 Pengecekan Total Berat:
Program memeriksa apakah total berat kedua kantong melebihi 150 kg. Jika iya, program menampilkan pesan dan berhenti:
`Proses selesai.`

6 Pengecekan Berat Kantong ≥ 9 kg:
Jika salah satu kantong memiliki berat 9 kg atau lebih, program menampilkan pesan dan berhenti.

7 Mengulangi Proses:
Jika tidak ada kondisi berhenti yang terpenuhi, program mengulangi proses, meminta input dari pengguna lagi.


3. Diberikan sebuah persamaan sebagai berikut ini.
   Buatlah sebuah program yang menerima input sebuah bilangan sebagai K, kemudian menghitung dan menampilkan nilai f(K) sesuai persamaan di atas.


#### Source Code
```go
//Rangga Pradarrell Fathi
//2311102200
//11-IF-05
package main

import (
	"fmt"  // Package untuk input/output
	"math" // Package untuk fungsi matematika
)

// hitungFungsiK menghitung nilai f(k) = (4k + 2)^2 / ((4k + 1)(4k + 3))
func hitungFungsiK(k float64) float64 {
	// Menghitung pembilang: (4k + 2)^2
	pembilang := math.Pow(4*k+2, 2)
	
	// Menghitung penyebut: (4k + 1)(4k + 3)
	penyebut := (4*k + 1) * (4*k + 3)
	
	// Mengembalikan hasil pembagian
	return pembilang / penyebut
}

func main() {
	var k float64
	
	// Input nilai k dari pengguna
	fmt.Print("Nilai K = ")
	_, err := fmt.Scanln(&k)
	
	// Kondisi error jika input tidak valid
	if err != nil {
		fmt.Println("Error: Input harus berupa angka")
		return
	}
	
	// Menghitung f(k) dengan memanggil fungsi hitungFungsiK
	hasil := hitungFungsiK(k)
	
	// Menampilkan hasil dengan format 10 angka desimal
	fmt.Printf("Nilai f(k) = %.10f\n", hasil)
}
```
#### Screenshoot Source Code
<img src="https://github.com/user-attachments/assets/dccde510-d39a-4c01-b9fa-d1d7acd9329d">


#### Screenshoot Output


#### Deskripsi Program
##### Program akan meminta pengguna untuk memasukkan nilai *k*, kemudian menghitung nilai F(k), berdasarkan rumus tersebut, dan menampilkan hasilnya dengan presisi hingga 10 angka desimal.

##### Algoritma Program
- Inisialisasi
- Input dari pengguna
- Perhitungan F(k)
- Menampilkan Hasil
  
##### Cara Kerja Program
1 Input nilai *k*
2 Perhitungan f(k)
3 Menampilkan Hasil
4 Penanganan kesalahan input

4. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut!

Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.


#### Source Code
```go
//Rangga Pradarrell Fathi
//2311102200
//11-IF-05
package main

import (
	"fmt"
)

// hitungBiayaPengiriman menghitung biaya pengiriman berdasarkan berat parsel
func hitungBiayaPengiriman(beratGram int) (int, string, string, int) {
	biayaDasar := 10000 // Biaya dasar per kg
	biayaTambahan := 0

	// Konversi ke kg dan sisa gram
	beratKg := beratGram / 1000
	sisaGram := beratGram % 1000

	// Hitung biaya dasar
	biayaTotal := biayaDasar * beratKg

	// Hitung biaya tambahan untuk sisa gram
	if sisaGram > 0 {
		if sisaGram <= 500 {
			biayaTambahan = sisaGram * 15 // Rp. 15 per gram untuk <= 500g
		} else {
			biayaTambahan = sisaGram * 5 // Rp. 5 per gram untuk > 500g
		}
		biayaTotal += biayaTambahan
	}

	// Format detail berat
	detailBerat := fmt.Sprintf("%d kg + %d gr", beratKg, sisaGram)

	// Format detail biaya
	detailBiaya := fmt.Sprintf("Rp. %d", biayaDasar*beratKg)
	if biayaTambahan > 0 {
		detailBiaya += fmt.Sprintf(" + Rp. %d", biayaTambahan)
	}

	return biayaTotal, detailBerat, detailBiaya, biayaTambahan
}

func main() {
	fmt.Println("Program Perhitungan Biaya Pengiriman POS")
	fmt.Print("Masukkan berat parsel (dalam gram): ")

	var beratParsel int
	_, err := fmt.Scan(&beratParsel)
	if err != nil {
		fmt.Println("Input tidak valid. Mohon masukkan angka.")
		return
	}

	// Contoh output format
	fmt.Println("\nBerat parsel (gram):", beratParsel)

	totalBiaya, detailBerat, detailBiaya, _ := hitungBiayaPengiriman(beratParsel)

	fmt.Println("Detail:", detailBerat)
	fmt.Println("Detail biaya:", detailBiaya)
	fmt.Println("Total biaya: Rp.", totalBiaya)
}
```
#### Screenshoot Source Code
<img src="https://github.com/user-attachments/assets/1e9e87dc-e682-4ac9-a9f3-52a6dfe08c1d">


#### Screenshoot Output


#### Deskripsi Program
##### Program ini menghitung biaya pengiriman berdasarkan berat parsel yang diinputkan oleh pengguna dalam satuan gram. 

##### Algoritma Program
- Input Berat Parsel
- Perhitungan Biaya Pengiriman
- Menghitung biaya tambahan berdasarkan sisa gram
- Menampilkan Hasil

##### Cara Kerja Program
1 Meminta Input dari Pengguna
2 Memanggil Fungsi hitungBiayaPengiriman
3 Di dalam fungsi ini, program melakukan konversi
4 Menampilkan Hasil

5. Sebuah bilangan bulat b memiliki faktor bilangan f > O jika f habis membagi b. Contoh: 2 merupakan faktor dari bilangan 6 karena 6 habis dibagi 2.

Buatlah program yang menerima input sebuah bilangan bulat b dan b> 1. Program harus dapat mencari dan menampilkan semua faktor dari bilangan tersebut!

Bilangan bulat b > 0 merupakan bilangan prima p jika dan hanya jika memiliki persis dua faktor bilangan saja, yaitu 1 dan dirinya sendiri.

Lanjutkan program sebelumnya. Setelah menerima masukan sebuah bilangan bulat b > 0. Program tersebut mencari dan menampilkan semua faktor bilangan tersebut. Kemudian, program menentukan apakah b merupakan bilangan prima.

#### Source Code
```go
//Rangga Pradarrell Fathi
//2311102200
//11-IF-05
package main

import (
	"fmt"
)

// fungsi faktorBawah menerima sebuah bilangan integer dan mengembalikan slice berisi faktor-faktornya
func faktorBawah(input int) []int {
	var faktor = make([]int, 0)
	for i := 1; i <= input; i++ {
		if input%i == 0 {
			faktor = append(faktor, i)
		}
	}
	return faktor
}

func main() {
	var input int
	// meminta input dari pengguna
	fmt.Print("Masukkan bilangan: ")
	fmt.Scanln(&input)

	// memeriksa apakah input adalah bilangan positif
	if input <= 0 {
		fmt.Println("Masukkan bilangan positif!")
	} else {
		faktor := faktorBawah(input)

		// menampilkan hasil
		fmt.Println("Faktor-faktor dari", input, ":")
		for _, value := range faktor {
			fmt.Println(value)
		}
	}
}

```
#### Screenshoot Source Code
<img src="https://github.com/user-attachments/assets/a2801502-6883-4047-8aa4-64759d8aa8f5">


#### Screenshoot Output

#### Deskripsi Program
##### Program ini bertujuan untuk mencari dan menampilkan faktor-faktor dari suatu bilangan bulat positif yang dimasukkan oleh pengguna. Faktor adalah bilangan-bilangan yang membagi habis bilangan tersebut.

##### Algoritma Program
- Input Bilangan:
- Pengecekan Validitas Input
- Menghitung Faktor-Faktor
- Menampilkan Faktor-Faktor

##### Cara Kerja Program
- Menerima Input dari Pengguna
- Memanggil Fungsi faktorBawah
- Menampilkan Hasil
- Pengecekan Bilangan Negatif


## Referensi 
[1] Donovan, A. A., & Kernighan, B. W. (2015). The Go Programming Language. Addison-Wesley Professional. 
[2] Bodner, J. (2021). Learning Go: An Idiomatic Approach to Real-World Go Programming. O'Reilly Media. 
------
