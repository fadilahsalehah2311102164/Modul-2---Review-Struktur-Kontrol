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
  Aditya Sulistiawan / 2311102193<br>
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
### 1. Pengenalan Go
Go (atau Golang) adalah bahasa pemrograman open source yang dikembangkan oleh Google pada tahun 2007 oleh Robert Griesemer, Rob Pike, dan Ken Thompson. Go dirancang dengan tujuan untuk menggabungkan kemudahan pemrograman seperti Python dengan kinerja dan keamanan dari bahasa yang dikompilasi seperti C++. Bahasa Go memiliki beberapa karakteristik utama:

- Kompilasi cepat
- Dukungan konkurensi bawaan melalui goroutines dan channels
- Sistem garbage collection otomatis
- Sintaks yang sederhana dan mudah dipahami
- Mendukung pemrograman berorientasi objek meskipun tidak memiliki inheritance
## 2. Tipe Data dan Variabel

### 2.1 Deklarasi Variabel
Go menggunakan tipe data statis dan inferensi tipe. Variabel dapat dideklarasikan dengan beberapa cara:

Contoh penggunaan:
```go
var nama string = "John"     // Deklarasi eksplisit
umur := 25                   // Deklarasi singkat dengan inferensi tipe
var tinggi float64           // Deklarasi tanpa inisialisasi
```

### 2.2 Tipe Data Dasar
Go memiliki beberapa tipe data dasar:

#### 1. Numerik

- Integer: int, int8, int16, int32, int64
- Unsigned Integer: uint, uint8, uint16, uint32, uint64
- Float: float32, float64
- Complex: complex64, complex128

#### 2. Boolean

- bool (true atau false

#### 3. String

- string (urutan karakter Unicode)

#### 4. Tipe Data Khusus

- byte (alias untuk uint8)
- rune (alias untuk int32, merepresentasikan karakter Unicode)


## 3. Percabangan

Percabangan dalam Go memungkinkan program untuk membuat keputusan dan mengeksekusi kode yang berbeda berdasarkan kondisi tertentu.

### 3.1 If-Else
Struktur if-else di Go memiliki beberapa fitur unik:
1. Tidak memerlukan tanda kurung `()` untuk kondisi
2. Kurung kurawal `{}` wajib digunakan
3. Dapat memiliki statement singkat sebelum kondisi

Contoh:
```go
if kondisi {
    // kode jika kondisi true
} else if kondisiLain {
    // kode jika kondisiLain true
} else {
    // kode jika semua kondisi false
}
```

### 3.2 Switch
Switch di Go lebih fleksibel dibanding bahasa lain:
1. Tidak perlu `break` di setiap case
2. Case bisa memiliki multiple values
3. Bisa menggunakan kondisi di setiap case

Contoh:
```go
switch variabel {
case nilai1:
    // kode untuk nilai1
case nilai2:
    // kode untuk nilai2
default:
    // kode jika tidak ada case yang cocok
}
```

## 4. Perulangan

Go hanya memiliki satu konstruksi perulangan, yaitu for, namun dapat digunakan dengan berbagai cara.

### For Tradisional
```go
for inisialisasi; kondisi; increment {
    // kode yang diulang
}
```
Komponen:
1. Statement inisialisasi: `i := 0`
2. Kondisi: `i < 5`
3. Statement post: `i++`

### For sebagai While
```go
for kondisi {
    // kode yang diulang selama kondisi true
}
```
Hanya menggunakan kondisi, mirip `while` di bahasa lain.

### Infinite Loop
'''go
for {
    // kode yang diulang selamanya
    if kondisiKeluar {
        break
    }
}
'''

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

## II. GUIDED

### Guided 1
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
![sc](https://github.com/user-attachments/assets/0b484870-e638-4594-8480-090f6cf528ab)


#### Screenshoot Output
![gui1](https://github.com/user-attachments/assets/3be11c25-a917-43b7-bffb-d96411368f39)


#### Deskripsi Program
Program ini dibuat dengan bahasa pemrograman Go. Program ini akan menerima input berupa string yang merupakan nama dari user,
menyimpannya di variabel "nama", dan kemudian mencetak nama tersebut ke konsol.

#### Algoritma
** 1. Input:
- Terima input bilangan bulat b dari pengguna. Pastikan b > 1.

** 2. Mencari dan Menampilkan Faktor:

- Buat perulangan (loop) dari f = 1 hingga f <= b.
- Di dalam loop, periksa apakah b habis dibagi f (menggunakan operator modulo %).
- Jika ya, maka f adalah faktor dari b. Tampilkan nilai f.

** 3. Menentukan Prima:

- Inisialisasi variabel jumlahFaktor = 0.
- Ulangi langkah 2.
- Setiap kali menemukan faktor, tambahkan nilai jumlahFaktor dengan 1.
- Setelah perulangan selesai:
- Jika jumlahFaktor == 2 (hanya memiliki dua faktor: 1 dan dirinya sendiri), maka b adalah bilangan prima.
- Jika tidak, maka b bukan bilangan prima.

** 4. Output:

- Tampilkan semua faktor yang telah ditemukan.
- Tampilkan apakah b merupakan bilangan prima atau bukan.

#### Cara Kerja Program
 1. Program dimulai dari fungsi main()
 2. Variabel nama dideklarasikan
 3. Program menunggu input dari pengguna
 4. Setelah pengguna memasukkan teks dan menekan Enter:
   - Input disimpan dalam variabel nama
   - Program melanjutkan ke baris berikutnya
 5. Isi variabel nama ditampilkan ke layar
 6. Program selesai

### Guided 2

Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah 'merah', 'kuning', 'hijau', dan 'ungu' selama 5 kali percobaan berulang.
Buatlah sebuah program yang menerima Input berupa warna dari ke 4 gelas reaks! sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuat dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.

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

	correctOrder := []string{"merah", "kuning", "hijau","ungu"}

	reader := bufio.NewReader(os.Stdin)
	success := true

	for i := 1; i <= 5; i++{
		fmt.Printf("Percobaan %d:" , i)

		input, _:= reader.ReadString('\n')
		input = strings.TrimSpace(input)

		colors := strings.Split(input, " ")

		for j :=0; j < 4; j++ {
			if colors[j] !=correctOrder[j]{
				success = false 
				break
			}
		}

		if !success {
			break
		}
	}

	if success {
		fmt.Println("BERHASIL : true")
	} else {
		fmt.Println ("BERHASIL : false")
	}
}
```
#### Screenshoot Source Code
![SC](https://github.com/user-attachments/assets/0aca8dee-98f8-4e24-b648-988e02b0de29)

#### Screenshoot Output
![gui2](https://github.com/user-attachments/assets/32f981cb-51f9-4e5e-8b5c-a6d59285b548)

#### Deskripsi Program
Program ini memeriksa apakah urutan warna yang dimasukkan pengguna sesuai dengan urutan yang telah ditentukan. Program akan memberikan hingga 5 peluang, dan akan berhenti jika urutan salah atau semua percobaan berhasil.

#### Algoritma Program
 1.Inisialisasi:

- Tetapkan urutan warna yang benar: "merah", "kuning", "hijau", "ungu".
- Tetapkan variabel success menjadi true (asumsi awal berhasil).
- Atur percobaan maksimum (misalnya, 5 kali).

 2.Perulangan Percobaan:

- Untuk setiap percobaan (1 hingga percobaan maksimum):
- Tampilkan pesan "Percobaan [nomor percobaan]:".
- Baca input urutan warna dari pengguna.
- Pisahkan input menjadi array warna individual.
- Bandingkan setiap warna dalam input dengan urutan yang benar:
- Jika ada warna yang tidak cocok pada posisinya:
Set success menjadi false.
Hentikan perulangan (tidak perlu melanjutkan percobaan).
- Jika semua warna cocok, lanjutkan ke percobaan berikutnya.

 3.Evaluasi Hasil:

- Jika success bernilai true (semua percobaan berhasil):
Tampilkan pesan "BERHASIL : true".
- Jika success bernilai false (ada percobaan yang gagal):
Tampilkan pesan "BERHASIL : false".

#### Cara Kerja Program
 1.Input:
Program meminta pengguna untuk memasukkan sebuah bilangan bulat (b) melalui fmt.Scanln(&b).

 2.Validasi Input:
Memastikan bahwa bilangan yang dimasukkan lebih besar dari 1 (b <= 1). Jika tidak, program akan menampilkan pesan kesalahan dan berhenti.

 3.Mencari Faktor:
Program menggunakan loop for untuk mengiterasi dari 1 hingga b.
Di dalam loop, program memeriksa apakah b habis dibagi oleh i menggunakan operator modulo (%).
Jika habis dibagi (b%i == 0), maka i adalah faktor dari b dan dicetak ke layar.

 4.Menentukan Bilangan Prima:
Program menginisialisasi variabel prima dengan true, mengasumsikan bahwa b adalah prima.
Jika b kurang dari atau sama dengan 1, maka prima diset menjadi false karena 1 dan bilangan negatif bukan prima.
Loop for kedua digunakan untuk memeriksa apakah ada bilangan lain selain 1 dan dirinya sendiri yang dapat membagi b. Loop ini berjalan dari 2 hingga akar kuadrat dari b.
Jika b habis dibagi oleh i dalam loop ini, maka b bukan prima, dan prima diset menjadi false. Loop kemudian dihentikan dengan break.

5.Menampilkan Hasil:
Program mencetak "Prima: " diikuti dengan nilai prima (true jika prima, false jika tidak).

### Guided 3
3. Menambahkan bilangan yang diinputkan user

#### Source Code
```go
package main 

import "fmt"

func main(){

	var a, b, c, d, e int
	var hasil int
	fmt.Scanln(&a, &b, &c, &d, &e)

	hasil = a+b+c+d+e
	fmt.Println("Hasil Penjualan" , a,b,c,d,e, "adalah =", hasil)
}
```
#### Screenshoot Source Code
![SC](https://github.com/user-attachments/assets/4e2eaf01-ada6-45a2-8190-2a98ba83a2e1)


#### Screenshoot Output
![gui3](https://github.com/user-attachments/assets/028256ff-fac4-415e-a7cd-2c6c9ff4ba54)


#### Deskripsi Program
Program ini merupakan aplikasi sederhana berbasis konsol yang dikembangkan dengan menggunakan bahasa pemrograman Go. Program ini dirancang untuk menerima lima angka dari pengguna, menjumlahkan nilainya, dan menampilkan jumlah beserta nilai masukannya dalam format  terstruktur.

#### Algoritma Program
- Inisialisasi

Deklarasikan lima variabel integer untuk input (a, b, c, d, e)
Deklarasikan satu variabel integer untuk hasil penjumlahan

- Input Data

Tampilkan prompt untuk input (implisit)
Terima lima angka dari pengguna dalam satu baris
Simpan angka-angka tersebut ke dalam variabel a, b, c, d, dan e

- Proses

Jumlahkan nilai dari variabel a, b, c, d, dan e
Simpan hasil penjumlahan ke dalam variabel hasil

- Output

Tampilkan nilai-nilai input dan hasil penjumlahan dalam format yang ditentukan

#### Cara Kerja Kelompok
1. Program dimulai
2. Variabel-variabel dideklarasikan
3. Program menunggu input 5 angka dari pengguna
4. Setelah input diterima, program menjumlahkan kelima angka
5. Hasil penjumlahan ditampilkan ke layar
6. Program selesai

### Guided 4
4. Diberikan sebuah nilai akhir mata kullah (NAM) [0..100] dan standar penilaian nilai mata kullah (NMK)

#### Source Code
```go
package main

import "fmt"

func main() {
	var nam float32
	var nmk string

	fmt.Println("Masukan nilai: ")
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
		nmk = "TOLOL!!!"
	}

	fmt.Printf("Nilai Indeks untuk nilai %.2f adalah %s\n", nam, nmk)
}
```
#### Screenshoot Source Code
![SC](https://github.com/user-attachments/assets/d5b136de-ebc7-4bbb-8375-6ee2849f2320)

#### Screenshoot Output
![gui4](https://github.com/user-attachments/assets/5dbfd523-510c-4bcb-9a8a-fc268fef47c8)

#### Deskripsi Program
Program yang Anda berikan  meminta pengguna  memasukkan  nilai dalam format desimal (float32). Selanjutnya program  menentukan nilai indeks berdasarkan rentang nilai yang ditentukan dan terakhir menampilkan hasil nilai indeks.

#### Algoritma Program
1. Menerima Input:
Minta pengguna untuk memasukkan sebuah bilangan bulat (b).

2. Mencari dan Menampilkan Faktor:
Gunakan loop for untuk iterasi dari 1 hingga b (inklusif).
Untuk setiap angka i dalam loop:
Jika b habis dibagi oleh i (sisa bagi sama dengan 0), maka i adalah faktor dari b.
Tampilkan i sebagai faktor.

3. Menentukan Prima:
Inisialisasi variabel jumlah_faktor dengan 0.
Gunakan loop for lagi untuk iterasi dari 1 hingga b.
Untuk setiap angka i dalam loop:
Jika b habis dibagi oleh i, tambahkan 1 ke jumlah_faktor.
Setelah loop selesai:
Jika jumlah_faktor sama dengan 2 (hanya 1 dan dirinya sendiri yang merupakan faktor), maka bilangan tersebut prima.
Jika tidak, bilangan tersebut bukan prima.

4. Menampilkan Hasil:
Tampilkan apakah bilangan tersebut prima atau bukan.

#### Cara Kerja Program
1. Deklarasi Variabel:
var nam float32: Mendeklarasikan variabel nam dengan tipe data float32 untuk menyimpan nilai yang akan diinput oleh pengguna.
var nmk string: Mendeklarasikan variabel nmk dengan tipe data string untuk menyimpan nilai indeks yang akan ditentukan berdasarkan nilai nam.

2. Input Pengguna:
fmt.Println("Masukan nilai: "): Menampilkan pesan kepada pengguna untuk memasukkan nilai.
fmt.Scan(&nam): Menerima input nilai dari pengguna dan menyimpannya ke dalam variabel nam.

3. Struktur Kontrol if-else if-else:
Program menggunakan struktur kontrol if-else if-else untuk menentukan nilai indeks berdasarkan rentang nilai nam:
if nam > 80: Jika nilai nam lebih besar dari 80, maka nmk akan diset menjadi "A".
else if nam > 72.5: Jika nilai nam lebih besar dari 72.5 (dan tidak lebih besar dari 80), maka nmk akan diset menjadi "B".
else if nam > 65: Jika nilai nam lebih besar dari 65 (dan tidak lebih besar dari 72.5), maka nmk akan diset menjadi "C".
else if nam > 50: Jika nilai nam lebih besar dari 50 (dan tidak lebih besar dari 65), maka nmk akan diset menjadi "D".
else if nam > 40: Jika nilai nam lebih besar dari 40 (dan tidak lebih besar dari 50), maka nmk akan diset menjadi "E".
else: Jika semua kondisi di atas tidak terpenuhi (nilai nam kurang dari atau sama dengan 40), maka nmk akan diset menjadi "TOLOL!!!".

4. Output:
fmt.Printf("Nilai Indeks untuk nilai %.2f adalah %s\n", nam, nmk): Menampilkan hasil nilai indeks (nmk) yang telah ditentukan berdasarkan nilai input (nam) dengan format dua digit di belakang koma.

## III. UNGUIDED

1. Suatu pita (string) berisi kumpulan nama-nama bunga yang dipisahkan oleh spasi dan '-', contoh pita diilustrasikan seperti berikut ini.

Pita: mawar - melati - tulip - teratai - kamboja - anggrek

Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita.

(Petunjuk: gunakan operasi penggabungan string dengan operator "+").

Tampilkan isi pita setelah proses input selesai.

Modifikasi program sebelumnya, proses input akan berhenti apabila user mengetikkan 'SELESAI'. Kemudian tampilkan isi pita beserta banyaknya bunga yang ada di dalam pita.

#### Source Code
```go
package main

import (
	"fmt"
	"strings"
)

func main() {
	var bunga []string
	
	fmt.Println("Masukkan nama-nama bunga (ketik 'SELESAI' untuk mengakhiri):")
	
	for {
		var input string
		fmt.Printf("Masukkan nama bunga ke-%d: ", len(bunga)+1)
		fmt.Scanln(&input)
		
		if strings.ToUpper(input) == "SELESAI" {
			break
		}
		
		if input != "" {
			bunga = append(bunga, input)
		}
	}
	
	if len(bunga) == 0 {
		fmt.Println("\nTidak ada bunga yang dimasukkan.")
		return
	}
	

	pita := strings.Join(bunga, " - ")
	
	
	fmt.Println("\nHasil pita:")
	fmt.Println(pita)
	fmt.Printf("Banyak bunga dalam pita: %d\n", len(bunga))
}

```
#### Screenshoot Source Code
![SC](https://github.com/user-attachments/assets/170448bc-3cd9-40c9-a109-31002c20a5d9)


#### Screenshoot Output
![Unguided 1-1](https://github.com/user-attachments/assets/491b1173-6b96-48ee-bd0c-aca0dab73f2e)


#### Deskripsi Program
Program  meminta pengguna untuk memasukkan nama bunga  berulang kali hingga pengguna memasukkan kata "DONE". Nama bunga yang Anda masukkan akan disimpan dalam potongan bernama "Bunga". Saat pengguna mengetik "SELESAI", program akan memeriksa apakah  nama bunga telah dimasukkan. Jika tidak ada, program akan menampilkan pesan "Tidak ada minat yang dimasukkan". Dan tolong berhenti. Setelah bunga  dimasukkan, program menggunakan fungsi strings.Join() untuk menggabungkan nama bunga  menjadi satu string  yang dipisahkan oleh karakter '-'. String gabungan ini disimpan dalam variabel band. Terakhir, program menampilkan rangkaian pita (representasi pita bunga) dan jumlah bunga yang terdapat dalam pita.

#### Algoritma Program 
1. Menerima Input:
Minta pengguna memasukkan bilangan bulat b.

2. Mencari dan Menampilkan Faktor:
Gunakan loop for dari 1 hingga b.
Untuk setiap nilai i dalam loop:
Jika b habis dibagi i, maka i adalah faktor dari b.
Tampilkan i.
3. Menentukan Bilangan Prima:
Inisialisasi variabel jumlah_faktor dengan 0.
Gunakan loop for dari 1 hingga b.
Untuk setiap nilai i dalam loop:
Jika b habis dibagi i, maka tambah jumlah_faktor dengan 1.
Setelah loop selesai, periksa nilai jumlah_faktor:
Jika jumlah_faktor sama dengan 2, maka b adalah bilangan prima (prima = True).
Jika tidak, b bukan bilangan prima (prima = False).

4. Menampilkan Hasil:
Tampilkan apakah b merupakan bilangan prima atau bukan.

#### Cara Kerja Program

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
