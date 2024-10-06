<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
<h2 align="center"><strong>ALGORITMA DAN PEMROGRAMAN 2</strong></h2>
<h2 align="center"><strong>MODUL II</strong></h2>
<h2 align="center"><strong> REVIEW STRUKTUR KONTROL </strong></h2>

<p align="center">

  <img src="https://github.com/user-attachments/assets/0a03461e-7740-4661-9e83-9925031bd72c" alt="Logo" width="200"/>

</p>

<br>

<p align="center">
  <strong>Disusun Oleh:</strong><br>
  Yesika Widiyani<br>
  2311102195<br>
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

# DASAR TEORI

## Bahasa Go
Bahasa Go, yang juga dikenal dengan nama Golang, adalah bahasa pemrograman yang dikembangkan oleh Google pada tahun 2007 dan dirilis secara publik pada tahun 2009. Go dirancang untuk menjadi bahasa yang sederhana, efisien, dan mudah dipahami, dengan fitur-fitur yang mendukung pemrograman paralel dan konkuren. Go sering digunakan dalam pengembangan sistem, aplikasi web, dan infrastruktur cloud.

<br>

Filosofi desain Go bertujuan untuk:<br>
- Kesederhanaan: Sintaks yang mudah dipahami dan minimistik.
- Kecepatan: Kompilasi yang cepat dan performa runtime yang baik.
- Keterbacaan: Kode yang mudah dibaca dan dipelihara.

## Struktur Program GO 
Dalam pemrograman Go, struktur atau struct adalah tipe yang ditentukan pengguna untuk menyimpan kumpulan kolom yang berbeda menjadi satu kolom.
<br>

Dalam kerangka program yang ditulis dalam bahasa pemrograman Go, program utama selalu mempunyai dua komponen berikut:<br>
-	package main merupakan penanda bahwa file ini berisi program utama.
-	func main() berisi kode utama dari sebuah program Go.<br>

**Contoh Struktur Program**

```go
package main

import "fmt"

func main() {
    fmt.Println("Hello, Go!")
}
```

## Tipe Data Dalam Go
Go memiliki beberapa tipe data, yang dibagi menjadi dua kategori: tipe data dasar dan tipe data terstruktur.<br>
### 1. Tipe Data Dasar
- Integer: Tipe untuk bilangan bulat. Contoh: `int`, `int8`, `int16`, `int32`, `int64`.
- Float: Tipe untuk bilangan pecahan. Contoh: `float32`, `float64`.
- Boolean: Tipe yang hanya memiliki dua nilai: `true` dan `false`.
- String: Tipe untuk teks, ditandai dengan tanda kutip ganda.

### 2. Tipe Data Terstruktur
- Array: Sekumpulan elemen dengan tipe yang sama. Ukuran array tetap.<br>
```go
var angka [5]int
angka[0] = 1
```
- Slice: Versi dinamis dari array. Dapat berubah ukurannya.<br>
```go
buah := []string{"apel", "jeruk"}
buah = append(buah, "pisang") // Menambahkan elemen
```
- Struct: Kumpulan field dengan tipe data yang berbeda.<br>
```go
type Mahasiswa struct {
    Nama    string
    Umur    int
    Jurusan string
}
```

## Struktur Kontrol
Struktur kontrol adalah salah satu konsep dasar dalam pemrograman yang memungkinkan pengembang untuk mengontrol alur eksekusi program berdasarkan kondisi tertentu. Dalam bahasa pemrograman Go (Golang), struktur kontrol mirip dengan bahasa pemrograman lainnya, namun ada beberapa fitur dan karakteristik unik yang membedakan Go dari bahasa lain. Berikut adalah dasar teori tentang struktur kontrol dalam bahasa Go:
<br>

### 1. Statement If-Else
Struktur if digunakan untuk menjalankan blok kode tertentu jika suatu kondisi bernilai benar (true). Jika tidak, blok else dapat digunakan untuk menjalankan kode lain.<br>
**Sintaks:**
```go
if condition {
    // code if condition is true
} else {
    // code if condition is false
}
```
<br>

### 2. Statement switch
Statement switch di Go menawarkan cara yang lebih sederhana dan efisien dibandingkan dengan serangkaian statement if-else yang panjang. Struktur ini biasanya digunakan untuk melakukan seleksi berdasarkan nilai dari sebuah variabel.
<br>
**Sintaks:**
```go
switch expression {
case value1:
    // code block
case value2:
    // code block
default:
    // default code block
}
```

### 3. Statement for
Go hanya memiliki satu jenis loop, yaitu for. Namun, for di Go cukup fleksibel dan bisa digunakan untuk berbagai kasus iterasi, baik sebagai loop klasik, loop berbasis kondisi, atau loop tanpa batas.
<br>
**Sintaks:**
```go
for initialization; condition; post {
    // code block
}
```

## Fungsi
Fungsi adalah blok kode yang melakukan tugas tertentu. Fungsi dapat memiliki parameter dan nilai kembali.
<br>
**Contoh Fungsi**
```go
func tambah(a int, b int) int {
    return a + b
}
```

## Error Handling
Go menggunakan pendekatan sederhana untuk menangani error dengan mengembalikan nilai error dari fungsi.
<br>
**Contoh Penanganan Error**
```go
func bagi(a, b int) (int, error) {
    if b == 0 {
        return 0, fmt.Errorf("Tidak bisa membagi dengan nol")
    }
    return a / b, nil
}
```

------

# GUIDED

## Guided 1 - Soal Nama
### Study Case
Membuat sebuah program sederhana yang meminta pengguna memasukkan nama mereka, lalu mencetak nama tersebut ke layar. Program ini dirancang menggunakan bahasa pemrograman Go dan bertujuan untuk memahami dasar penggunaan input/output serta variabel.

### Source Code
![guided1](https://github.com/user-attachments/assets/92144e4a-0c63-48cd-ac6c-297bf862a942)

### Screenshot Output
![outputg1](https://github.com/user-attachments/assets/4903b775-5221-4fb3-8e9e-427527b71b46)

### Deskripsi Program
Program ini merupakan sebuah program sederhana yang dibuat dengan menggunakan bahasa pemrograman Go (Golang). Tujuannya adalah untuk membaca input teks dari pengguna (seperti nama), lalu menampilkan kembali input tersebut sebagai output. Program ini memanfaatkan fungsi-fungsi standar dari paket `fmt` untuk menangani proses input dan output.

#### Algoritma Program
1. **Deklarasi variabel:** Program memulai dengan mendeklarasikan variabel `nama` yang memiliki tipe data `string`. Variabel ini akan menyimpan input dari pengguna.
2. **Input dari pengguna:** Program meminta input dari pengguna menggunakan fungsi `fmt.Scanln`. Input ini diambil dari baris perintah (command line) dan disimpan dalam variabel `nama`.
3. **Menampilkan output:** Setelah mendapatkan input dari pengguna, program kemudian mencetak nilai dari variabel `nama` menggunakan fungsi `fmt.Println` untuk ditampilkan sebagai output.

#### Cara Kerja Program
1. **Inisialisasi Program:** Saat program dijalankan, ia pertama kali mendeklarasikan variabel `nama` dengan tipe string, tetapi belum diisi dengan nilai apa pun.
   
2. **Menerima Input Pengguna:**
   - Program kemudian akan berhenti sejenak, menunggu pengguna untuk memasukkan teks.
   - Setelah pengguna mengetikkan nama atau teks apa pun dan menekan tombol **Enter**, fungsi `fmt.Scanln(&nama)` akan mengambil input tersebut dan menyimpannya dalam variabel `nama`.
   - Tanda `&` di depan `nama` berarti program bekerja dengan "alamat" dari variabel tersebut, yang dikenal sebagai pointer dalam Go. Ini memungkinkan fungsi `Scanln` untuk mengisi variabel `nama` dengan input pengguna.

3. **Menampilkan Output:**
   - Setelah input diambil, program menggunakan fungsi `fmt.Println(nama)` untuk mencetak teks yang telah diinput oleh pengguna.
   - Teks tersebut akan ditampilkan di layar sebagai output, yang berupa nilai dari variabel `nama`.

#### Alur Kerja Program
**Input dan Output:**
- Saat program dijalankan, ia akan menunggu pengguna mengetikkan nama:  
  **Input:** `yesika`

- Setelah pengguna menekan **Enter**, program akan mencetak:<br>
  **Output:** `yesika`

Jadi, intinya, program ini membaca teks yang dimasukkan oleh pengguna, kemudian menampilkannya kembali sebagai output.

## Guided 2 - 2B - 1
### Study Case
Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah ‘merah’, ‘kuning’, ‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.
<br>

Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):

### Source Code
![guided2](https://github.com/user-attachments/assets/07dc276b-9f06-44d9-808a-cfd9afbf4fa3)

### Screenshot Output
![outputg2](https://github.com/user-attachments/assets/c3e8cb59-525e-481a-a7eb-78c982ea9ccf)

### Deskripsi Program
Program ini merupakan program berbasis bahasa Go yang dirancang untuk memeriksa apakah pengguna dapat memasukkan urutan warna yang benar dalam lima percobaan. Urutan warna yang benar adalah: **"merah", "kuning", "hijau", dan "ungu"**. Jika pengguna berhasil memasukkan urutan yang tepat di setiap percobaan, program akan menampilkan hasil **"BERHASIL : true"**, sedangkan jika pengguna gagal, program akan menampilkan **"BERHASIL : false"**.

#### Algoritma Program
1. **Inisialisasi Urutan Warna:** Program dimulai dengan mendefinisikan urutan warna yang benar dalam sebuah slice bernama `correctOrder`.
2. **Loop untuk 5 Percobaan:** Program menggunakan loop untuk mengizinkan pengguna mencoba sebanyak 5 kali.
3. **Membaca Input Pengguna:** Pada setiap percobaan, program meminta pengguna untuk memasukkan urutan warna dalam satu baris, memisahkan warna dengan spasi.
4. **Pengecekan Input:** Program membandingkan input pengguna dengan urutan warna yang benar.
5. **Evaluasi Hasil:** Jika pengguna gagal di salah satu percobaan, program segera keluar dari loop dan menampilkan hasil **false**. Jika berhasil di semua percobaan, program menampilkan hasil **true**.

#### Cara Kerja Program
1. **Inisialisasi:**
   - Program pertama kali mendeklarasikan slice `correctOrder` yang menyimpan urutan warna yang benar: `{"merah", "kuning", "hijau", "ungu"}`.
   
2. **Input dari Pengguna:**
   - Program menggunakan `bufio.NewReader` untuk membaca input dari pengguna di setiap percobaan.
   - Input tersebut kemudian dibersihkan dari karakter newline (`\n`) menggunakan `strings.TrimSpace`.

3. **Proses Looping:**
   - Program menggunakan loop for yang berjalan sebanyak 5 kali (karena ada 5 percobaan).
   - Pada setiap iterasi, pengguna diminta untuk memasukkan urutan warna, yang kemudian dipisah menjadi elemen-elemen terpisah menggunakan `strings.Split` dengan pemisah spasi (`" "`).

4. **Pengecekan Urutan Warna:**
   - Setelah input pengguna dipecah menjadi slice warna, program membandingkan setiap warna yang diinputkan dengan urutan warna yang benar.
   - Jika ditemukan perbedaan antara input pengguna dan urutan yang benar, variabel `success` diubah menjadi **false** dan program keluar dari loop.

5. **Hasil Akhir:**
   - Setelah selesai mengecek semua percobaan, jika tidak ada kesalahan di setiap percobaan, program mencetak **"BERHASIL : true"**.
   - Jika ditemukan kesalahan di salah satu percobaan, program langsung mencetak **"BERHASIL : false"** dan keluar dari loop lebih awal.

#### Input dan Output
**Input:**  
Pengguna memasukkan urutan warna pada setiap percobaan, misalnya:

```
Percobaan 1: merah kuning hijau ungu
Percobaan 2: merah kuning hijau ungu
Percobaan 3: merah kuning hijau ungu
Percobaan 4: merah kuning hijau ungu
Percobaan 5: merah kuning hijau ungu
```

**Output:**  
Karena semua percobaan benar, program akan menghasilkan:

```
BERHASIL : true
```

Jika di salah satu percobaan pengguna salah memasukkan urutan warna, misalnya:

```
Percobaan 3: merah biru hijau ungu
```

**Output:**  
Program akan langsung menghasilkan:

```
BERHASIL : false
```

## Guided - 3 - Soal abcde
### Study Case
Membuat program sederhana yang memungkinkan pengguna memasukkan 5 angka integer, kemudian program menjumlahkan kelima angka tersebut dan menampilkan hasilnya.

### Source Code
![Uploading guided3.png…]()

### Screenshot Output


### Deskripsi Program 

## 2B - 3
### Study Case
### Source Code
### Screenshot Output
### Deskripsi Program 


# UNGUIDED

## 1. 2B - 2
### Study Case
Suatu pita (string) berisi kumpulan nama-nama bunga yang dipisahkan oleh spasi dan ‘–‘, contoh pita diilustrasikan seperti berikut ini.
Pita: mawar – melati – tulip – teratai – kamboja – anggrek
<br>
Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita.
<br>
(Petunjuk: gunakan operasi penggabungan string dengan operator “+” ). Tampilkan isi pita setelah proses input selesai.
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):

![soal2b-2 1](https://github.com/user-attachments/assets/11fcd8f8-0835-4186-ad0f-e516d5b2aaaa)

Modifikasi program sebelumnya, proses input akan berhenti apabila user mengetikkan ‘SELESAI’. Kemudian tampilkan isi pita beserta banyaknya bunga yang ada di dalam pita. Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):

![soal2b-2 2](https://github.com/user-attachments/assets/85ad6097-1e95-42ec-9a70-0e5530533ad4)


## 2B - 3
### Study Case
### Source Code
### Screenshot Output
### Deskripsi Program 

## 2B - 3
### Study Case
### Source Code
### Screenshot Output
### Deskripsi Program 


# UNGUIDED
## 2B - 2
### Study Case
### Source Code
### Screenshot Output
### Deskripsi Program 

## 2B - 3
### Study Case
### Source Code
### Screenshot Output
### Deskripsi Program 

## 2C - 1
### Study Case
### Source Code
### Screenshot Output
### Deskripsi Program 

## 2C - 2
### Study Case
### Source Code
### Screenshot Output
### Deskripsi Program 

## 2C - 3
### Study Case
### Source Code
### Screenshot Output
### Deskripsi Program 
