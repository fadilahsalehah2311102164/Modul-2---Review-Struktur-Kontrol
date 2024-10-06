![unguided4](https://github.com/user-attachments/assets/b413a444-070d-419f-a26b-7e27a66ef3f3)<h2 align="center"><strong>LAPORAN PRAKTIKUM</strong></h2>
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
![guided3](https://github.com/user-attachments/assets/2a107fb4-26d5-48bc-a2b4-eaa7a0ff266c)

### Screenshot Output
![output3](https://github.com/user-attachments/assets/09f863f7-4bc8-4dda-bba0-64ecf9b37b03)

### Deskripsi Program 
### Deskripsi Program

Program ini dibuat untuk menerima lima buah angka dari pengguna, menjumlahkannya, dan menampilkan hasil dari penjumlahan tersebut. Pengguna memasukkan lima angka secara berurutan, dan program akan menjumlahkan angka-angka tersebut dan mencetak hasilnya di layar.

### Algoritma yang Digunakan

Algoritma yang digunakan adalah algoritma penjumlahan sederhana:
1. Meminta input dari pengguna sebanyak lima angka.
2. Menyimpan angka-angka tersebut dalam variabel.
3. Menjumlahkan semua angka yang diinputkan.
4. Menampilkan hasil penjumlahan.

### Cara Kerja Program

1. **Deklarasi Variabel**:
   - Program mendeklarasikan lima variabel `a`, `b`, `c`, `d`, `e` yang bertipe integer untuk menyimpan input dari pengguna.
   - Selain itu, variabel `hasil` digunakan untuk menyimpan hasil penjumlahan dari kelima angka tersebut.

2. **Membaca Input**:
   - Program menggunakan fungsi `fmt.Scanln(&a, &b, &c, &d, &e)` untuk menerima lima angka sekaligus dari pengguna. Pengguna akan diminta untuk memasukkan lima angka dipisah dengan spasi.
   
3. **Proses Penjumlahan**:
   - Setelah input dari pengguna diterima, program menjumlahkan kelima angka tersebut dengan formula `hasil = a + b + c + d + e`.

4. **Output**:
   - Program menampilkan hasil penjumlahan menggunakan fungsi `fmt.Println()`. Di sini, output yang ditampilkan berupa format:
     ```
     Hasil Penjumlahan a b c d e adalah = hasil
     ```
   - Misalnya, jika pengguna memasukkan angka `1 2 3 4 5`, program akan menampilkan:
     ```
     Hasil Penjumlahan 1 2 3 4 5 adalah = 14
     ```

### Input dan Output

Jika pengguna memasukkan:
```
2 3 4 9 8
```
Maka program akan mencetak:
```
Hasil Penjumlahan 2 3 4 9 8 adalah = 26
```

## Guided - 4
### Study Case I
Diberikan sebuah nilai akhir mata kuliah (NAM) [0..100] dan standar penilaian nilai mata kuliah (NMK) sebagai berikut:
![soal2c-2 1](https://github.com/user-attachments/assets/10b38d38-db72-4004-b88e-fa3530c5cd96)
![soal2c-2 2](https://github.com/user-attachments/assets/3c7bf1c1-216b-44ce-af1e-aa710571fa20)

### Source Code
![guided4](https://github.com/user-attachments/assets/0a2f9d0f-f24e-48a8-998f-2858bdbd631b)

### Screenshot Output
![outputg4 1](https://github.com/user-attachments/assets/49facde5-a9be-4ccc-812d-c5750ce2717b)

### Study Case II
#### A. Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal?

##### Output
![outputg4 2](https://github.com/user-attachments/assets/7fdfb04a-75ab-4e57-addd-3878b778de31)

Program menampilkan seksekusi sesuai dengan spesifikasi soal, dimana nila 80.1 merupakan lebih besar dari 80 dan mendapatkan nilai A

#### B. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya!
##### Penjelasan
Berikut ini adalah beberapa kesalahan yang terdapat pada program tersebut beserta penjelasannya:

##### 1. **Salah Ketik (Typo) pada Variabel dan Teks**
   - **Variabel `nam`**: Pada bagian komentar dan kode, variabel `nam` seharusnya `nilai`, atau minimal lebih deskriptif (contohnya `nilai` atau `angka`) agar mudah dipahami.
   - **Teks Output**: Pada teks `Nilai Indeks unutk nilai %.2f adalah %s\n`, terdapat kesalahan ketik pada kata *"unutk"*, yang seharusnya *"untuk"*.

   **Perbaikan**:
   ```go
   fmt.Printf("Nilai Indeks untuk nilai %.2f adalah %s\n", nam, nmk)
   ```

##### 2. **Pembagian Nilai Batas yang Tidak Jelas**
   - Pada kondisi `else if nam > 72.5`, perbandingan seperti ini tidak melibatkan batas yang tepat antara nilai yang masuk kategori `A`, `B`, `C`, dan seterusnya. Harus jelas kapan nilai **<= 80** misalnya, masuk ke dalam kategori `B` atau `A`.

   **Perbaikan**:
   Gunakan operator **>=** (lebih besar sama dengan) untuk memberikan batas yang jelas.
   - Misalnya, ketika nilai 80 masuk ke kategori A, maka logika menjadi:
     ```go
     if nam >= 80 {
         nmk = "A"
     } else if nam >= 72.5 {
         nmk = "B"
     } 
     ```

##### 3. **Rentang Nilai yang Tidak Jelas**
   - Pada rentang nilai terakhir `else if nam > 40`, nilai 40 tidak masuk ke rentang `F` atau kategori lainnya. Rentang seharusnya konsisten dan terdefinisi dengan baik.

   **Perbaikan**:
   Tambahkan `<=` di setiap perbandingan rentang nilai, dan pastikan setiap nilai bisa dimasukkan ke dalam kategori yang tepat.
   ```go
   if nam >= 80 {
       nmk = "A"
   } else if nam >= 72.5 {
       nmk = "B"
   } else if nam >= 65 {
       nmk = "C"
   } else if nam >= 50 {
       nmk = "D"
   } else if nam >= 40 {
       nmk = "E"
   } else {
       nmk = "F"
   }
   ```

##### Alur Program yang Benar:
1. **Meminta Input Nilai**:
   - Program meminta input nilai numerik dari pengguna.
   
2. **Menentukan Kategori Nilai**:
   - Program memeriksa apakah nilai yang diinputkan berada dalam rentang tertentu dan mengelompokkannya ke dalam kategori (A, B, C, D, E, F).
   - Rentang harus mencakup semua kemungkinan nilai dari yang tertinggi ke terendah, dan tidak boleh ada nilai yang tidak terdefinisi.

3. **Menampilkan Output**:
   - Program akan menampilkan hasil berupa indeks nilai huruf sesuai dengan nilai yang dimasukkan.

##### Versi Program yang Diperbaiki:
![guidedv 1](https://github.com/user-attachments/assets/2a55d742-1b59-4edc-b9cd-acc4f32b4019)

### Kesimpulan:
- Program memiliki beberapa kesalahan penulisan dan pembagian rentang nilai yang tidak jelas.
- Setelah diperbaiki, program akan bekerja dengan benar dengan menerima input nilai dari pengguna, mengelompokkannya berdasarkan rentang yang telah ditentukan, dan mencetak indeks huruf yang sesuai.

#### C. Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; dan 49.5. Seharusnya keluaran yang diperoleh adalah ‘A’, ‘B’, dan ‘D’.
##### Source Code
![guided4-rev](https://github.com/user-attachments/assets/b2667f6a-5668-47e5-b0a2-7a4aca2038a3)

##### Output
![outputg4 3](https://github.com/user-attachments/assets/b25172b3-426b-4a15-8133-61b6ae42255f)

### Deskripsi Program 
Program ini merupakan aplikasi sederhana yang mengonversi nilai numerik (nilai ujian atau tugas) yang dimasukkan oleh pengguna ke dalam **nilai indeks huruf** (A, B, C, D, E, F) berdasarkan rentang nilai tertentu.

#### Algoritma Program
1. **Input**: 
   - Program meminta pengguna untuk memasukkan nilai berupa angka desimal (float32).
2. **Proses Pengelompokan Nilai**:
   - Berdasarkan nilai yang dimasukkan, program akan memeriksa apakah nilai tersebut memenuhi syarat untuk mendapatkan indeks huruf sesuai rentang nilai yang telah ditentukan.
   - Rentang nilai yang digunakan dalam program adalah:
     - Nilai **≥ 80**: indeks **A**
     - Nilai **≥ 70 dan < 80**: indeks **B**
     - Nilai **≥ 65 dan < 70**: indeks **C**
     - Nilai **≥ 45 dan < 65**: indeks **D**
     - Nilai **≥ 40 dan < 45**: indeks **E**
     - Nilai **< 40**: indeks **F**
3. **Output**:
   - Setelah pengelompokan, program akan menampilkan indeks huruf yang sesuai dengan nilai yang telah dimasukkan.

#### Cara Kerja Program
1. **Deklarasi Variabel**:
   - Program mendeklarasikan dua variabel: `nilai` (tipe float32) untuk menyimpan nilai numerik dari input pengguna, dan `indeks` (tipe string) untuk menyimpan nilai huruf yang akan dihitung berdasarkan input nilai.

2. **Meminta Input**:
   - Program menggunakan fungsi `fmt.Scan()` untuk menerima input nilai dari pengguna. Input yang dimasukkan harus berupa angka desimal (float32), seperti 75.5 atau 90.

3. **Mengecek Kondisi**:
   - Program kemudian mengecek nilai yang dimasukkan menggunakan serangkaian pernyataan `if-else`. Setiap pernyataan memeriksa apakah nilai berada dalam rentang tertentu:
     - Jika nilai **≥ 80**, variabel `indeks` diisi dengan huruf **A**.
     - Jika nilai **≥ 70**, maka nilai tersebut diberi indeks **B**.
     - Dan seterusnya hingga indeks huruf **F** untuk nilai di bawah 40.

4. **Mencetak Output**:
   - Setelah proses pengecekan selesai, program menggunakan `fmt.Printf()` untuk mencetak hasil berupa nilai indeks huruf beserta nilai numerik yang dimasukkan oleh pengguna.
   - Misalnya, jika pengguna memasukkan nilai **75**, maka output program akan menjadi: 
     ```
     Nilai Indeks untuk nilai 75.00 adalah B
     ```

#### Cara Kerja Output
- Misalnya, pengguna memasukkan nilai **83**:
  - Program memeriksa apakah nilai tersebut lebih besar atau sama dengan 80. Karena 83 memenuhi syarat, maka `indeks = "A"`.
  - Program kemudian mencetak hasilnya sebagai **Nilai Indeks untuk nilai 83.00 adalah A**.
  
- Jika pengguna memasukkan nilai **72**:
  - Program memeriksa setiap rentang. Nilai 72 tidak memenuhi syarat untuk indeks A, tetapi memenuhi syarat untuk indeks B (karena lebih besar dari 70), sehingga `indeks = "B"`.
  - Program mencetak hasilnya sebagai **Nilai Indeks untuk nilai 72.00 adalah B**.

------

# UNGUIDED

## Unguided 1 - 2B - 2
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

### Source Code
![unguided1](https://github.com/user-attachments/assets/e11881ed-4bc9-4cbe-bb93-aad7afc11c64)

### Screenshot Output
![outputun1](https://github.com/user-attachments/assets/17f8a25d-5de4-4d03-b1dc-b45fe86bdcbb)

### Deskripsi Program 
Program ini adalah aplikasi sederhana yang meminta pengguna untuk memasukkan nama-nama bunga, lalu menyimpannya dalam sebuah daftar. Pengguna dapat terus menambahkan bunga hingga memasukkan perintah "SELESAI". Setelah selesai, program akan menampilkan daftar bunga yang telah dimasukkan, disatukan dengan tanda pemisah berupa " - ", dan juga menampilkan jumlah bunga yang dimasukkan.

#### Algoritma Program
1. **Deklarasi Variabel**:
   - Program menggunakan variabel `bunga` berupa slice string untuk menyimpan nama-nama bunga yang dimasukkan oleh pengguna.
   - Variabel `scanner` digunakan untuk membaca input pengguna.

2. **Looping untuk Input**:
   - Program menampilkan perintah kepada pengguna untuk memasukkan nama bunga. Pengguna dapat memasukkan satu atau lebih nama bunga.
   - Setiap kali pengguna mengetikkan nama bunga dan menekan Enter, program akan menyimpan input tersebut dalam slice `bunga`.
   - Program akan terus meminta input hingga pengguna mengetikkan "SELESAI" untuk mengakhiri proses input.

3. **Proses dan Pengecekan Input**:
   - Setiap input bunga akan dimasukkan ke dalam slice `bunga` menggunakan fungsi `append()`.
   - Jika pengguna memasukkan kata "SELESAI", program akan keluar dari loop dan berhenti meminta input tambahan.

4. **Penggabungan Bunga dan Penghitungan Jumlah**:
   - Setelah selesai menerima input, program akan menampilkan daftar bunga yang sudah dimasukkan dengan menyatukan setiap elemen dari slice `bunga` menggunakan tanda pemisah " - ". Fungsi `strings.Join()` digunakan untuk melakukan ini.
   - Program juga menghitung berapa banyak bunga yang dimasukkan oleh pengguna menggunakan fungsi `len()` dan menampilkannya.

#### Cara Kerja Program

1. **Input Data**:
   - Program dimulai dengan meminta pengguna untuk memasukkan nama bunga satu per satu.
   - Misalnya, pengguna memasukkan:
     - Bunga 1: "Mawar"
     - Bunga 2: "Melati"
     - Bunga 3: "Matahari"
     - Bunga 4: "Tulip"
   - Ketika pengguna mengetik "SELESAI", program akan berhenti meminta input.

2. **Pemrosesan Data**:
   - Program menyimpan setiap input bunga ke dalam slice `bunga`:
     - Setelah memasukkan semua bunga, `bunga` berisi: `["Mawar", "Melati", "Matahari", "Tulip"]`.

3. **Penggabungan dan Penghitungan**:
   - Program kemudian menggabungkan semua nama bunga menjadi satu string dengan tanda pemisah " - ":
     ```
     "Mawar - Melati - Matahari - Tulip"
     ```
   - Setelah itu, program menghitung jumlah bunga yang dimasukkan. Pada contoh ini, jumlahnya adalah 4.

### Kesimpulan
Program ini menggunakan loop untuk menerima input dari pengguna sampai input "SELESAI" diterima. Setelah itu, program menggabungkan semua nama bunga yang dimasukkan menggunakan tanda " - " dan menampilkan jumlah bunga yang diinputkan. Program ini sederhana, namun cukup efektif untuk melatih pemrosesan input dari pengguna, penggunaan slice, dan cara menggabungkan serta menampilkan data dalam Go.

## Unguided 2 - 2B - 3
### Study Case
Setiap hari Pak Andi membawa banyak barang belanjaan dari pasar dengan mengendarai sepeda motor. Barang belanjaan tersebut dibawa dalam kantong terpal di kiri-kanan motor. Sepeda motor tidak akan oleng jika selisih berat barang di kedua kantong sisi tidak lebih dari 9 kg.
<br>
Buatlah program Pak Andi yang menerima input dua buah bilangan real positif yang menyatakan berat total masing-masing isi kantong terpal. Program akan terus meminta input bilangan tersebut hingga salah satu kantong terpal berisi 9 kg atau lebih.
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):

![soalun2 1](https://github.com/user-attachments/assets/1c99aa30-a7bd-4909-87af-a2c9c6647ffd)

Pada modifikasi program tersebut, program akan menampilkan true jika selisih kedua isi kantong lebih dari atau sama dengan 9 kg. Program berhenti memproses apabila total berat isi kedua kantong melebihi 150 kg atau salah satu kantong beratnya negatif.
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):
![soalun2 2](https://github.com/user-attachments/assets/d7cf3d5b-db90-4d05-8c73-53e41ebb8fa9)

### Source Code
![unguided2](https://github.com/user-attachments/assets/d7611e65-5774-4f76-beed-f371197dda48)

### Screenshot Output
![outputun2](https://github.com/user-attachments/assets/624830d8-b751-46c3-a580-ab6e10008e64)

### Deskripsi Program 
Program ini merupakan aplikasi sederhana yang digunakan untuk menginput berat belanjaan dalam dua kantong dan menentukan apakah total berat tersebut memenuhi kriteria tertentu. Program ini memberikan umpan balik tentang kondisi berat dan stabilitas sepeda motor yang membawa kantong tersebut.

#### Algoritma

1. **Inisialisasi Variabel**:
   - `kantong1` dan `kantong2` untuk menyimpan berat dari masing-masing kantong belanja.
   
2. **Input Berat**:
   - Program meminta pengguna untuk memasukkan berat belanjaan di kedua kantong.
   
3. **Perhitungan Total**:
   - Menghitung total berat dari kedua kantong dengan menjumlahkan `kantong1` dan `kantong2`.
   
4. **Kondisi dan Umpan Balik**:
   - Jika total berat lebih dari 150 kg:
     - Tampilkan pesan bahwa total berat tidak boleh melebihi 150 kg.
     - Tampilkan pesan "Program Selesai".
     - Hentikan eksekusi program dengan `break`.
   - Jika salah satu dari berat kantong negatif:
     - Tampilkan pesan bahwa berat tidak boleh negatif.
   - Jika perbedaan berat antara dua kantong kurang dari 9 kg:
     - Tampilkan pesan bahwa sepeda motor tidak akan oleng (stabil).
   - Jika tidak memenuhi kondisi di atas:
     - Tampilkan pesan bahwa sepeda motor akan oleng (tidak stabil).

#### Cara Kerja Program
1. Program mulai dengan mengimpor paket yang diperlukan (`fmt` dan `math`).
2. Menggunakan loop `for`, program meminta pengguna untuk memasukkan berat dari dua kantong.
3. Program membaca input dari pengguna dan menyimpan nilai ke dalam variabel `kantong1` dan `kantong2`.
4. Total berat dihitung dengan menjumlahkan kedua variabel tersebut.
5. Program kemudian memeriksa apakah total berat melebihi 150 kg, jika iya, program menampilkan pesan dan keluar.
6. Jika berat salah satu kantong negatif, program memberikan peringatan.
7. Jika perbedaan antara berat kedua kantong kurang dari 9 kg, program menyatakan bahwa sepeda motor akan stabil. Jika tidak, program menyatakan bahwa sepeda motor akan oleng.
8. Proses ini diulang hingga total berat mencapai 150 kg, di mana program akan berhenti.

### Input dan Output

- Jika pengguna memasukkan total berat lebih dari 150 kg, akan ditampilkan:
  ```
  Total berat tidak boleh melebihi 150 kg.
  Program Selesai
  ```

- Jika pengguna memasukkan berat negatif, akan ditampilkan:
  ```
  Berat tidak boleh negatif.
  ```

- Jika berat di kedua kantong stabil (perbedaan < 9 kg), akan ditampilkan:
  ```
  Sepeda motor akan oleng: false
  ```

- Jika berat tidak stabil (perbedaan >= 9 kg), akan ditampilkan:
  ```
  Sepeda motor akan oleng: true
  ```

Program ini memberikan cara yang efektif untuk memantau dan mengontrol berat barang yang dibawa, memastikan keselamatan saat mengangkut barang menggunakan sepeda motor.

## Unguided 3 - 2B - 4
### Study Case
Diberikan sebuah persamaan sebagai berikut ini.

![soalun3 1](https://github.com/user-attachments/assets/1ce588db-5b16-44e1-8fa1-52f7d7d6768e)

Buatlah sebuah program yang menerima input sebuah bilangan sebagai K, kemudian menghitung dan menampilkan nilai f(K) sesuai persamaan di atas.
Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):

![soalun3 2](https://github.com/user-attachments/assets/23f0b46e-a4a9-4852-9e63-752e59f9b395)

### Source Code
![unguided3](https://github.com/user-attachments/assets/66d5d22a-6571-4f82-8578-af5a38f94426)

### Screenshot Output
![ouputun3](https://github.com/user-attachments/assets/ab96f150-3f2b-47a3-87b9-f5e03d1556a5)

### Deskripsi Program 
Program ini merupakan aplikasi sederhana yang bertujuan untuk menampilkan dua nilai: satu integer (`k`) dan satu floating-point (`f`) dalam format tabel sederhana. Program ini mengilustrasikan bagaimana cara mencetak nilai dengan format yang rapi ke layar.

#### Algoritma

1. **Inisialisasi Variabel**:
   - Mendeklarasikan variabel `k` bertipe `int` dan menginisialisasinya dengan nilai `100`.
   - Mendeklarasikan variabel `f` bertipe `float64` dan menginisialisasinya dengan nilai `1.0000061880`.

2. **Mencetak Header Tabel**:
   - Mencetak garis horizontal sebagai pemisah atas tabel.

3. **Mencetak Nilai**:
   - Mencetak nilai dari variabel `k` dan `f` dengan label yang sesuai dan format yang rapi.

4. **Mencetak Footer Tabel**:
   - Mencetak garis horizontal sebagai penutup tabel.

#### Cara Kerja Program

1. Program dimulai dengan mengimpor paket `fmt`, yang digunakan untuk mencetak output ke konsol.
2. Dua variabel diinisialisasi:
   - `k` diatur ke nilai `100`.
   - `f` diatur ke nilai `1.0000061880`.
3. Program mencetak garis pembuka tabel dengan menggunakan `fmt.Println()`, yang mencetak garis horizontal.
4. Nilai variabel `k` dicetak dalam format tabel dengan label "Nilai K".
5. Nilai variabel `f` dicetak dalam format tabel dengan label "Nilai f(K)".
6. Program mencetak garis penutup tabel dengan menggunakan `fmt.Println()`.

## Unguided 4 - 2C- 1
### Study Case
PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut!

Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.

Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):
![soalun4](https://github.com/user-attachments/assets/4ebfd1cb-90da-49e1-b274-7ef6a27ef552)

### Source Code
![unguided4](https://github.com/user-attachments/assets/9f283330-364a-4d38-b3ab-144b0f860e8c)

### Screenshot Output
![outputun4](https://github.com/user-attachments/assets/e545aa4f-9dbc-41ef-b803-66cc42df7a88)

### Deskripsi Program
Program ini merupakan aplikasi yang menghitung dan menampilkan biaya pengiriman untuk beberapa parsel berdasarkan berat masing-masing parsel. Biaya dihitung berdasarkan tarif per kilogram dan tambahan biaya per gram. Program menggunakan struktur data peta (map) dan slice untuk menyimpan informasi tentang parsel.

#### Algoritma

1. **Mendefinisikan Data Parsel**:
   - Buat slice berisi map yang menyimpan informasi tentang berat parsel dan detail berat dalam kilogram dan gram.

2. **Menghitung Biaya untuk Setiap Parsel**:
   - Iterasi melalui setiap parsel dalam slice `parcels`.
   - Untuk setiap parsel, panggil fungsi `hitungBiayaParcel` untuk menghitung total biaya berdasarkan berat parsel.
   - Simpan hasil total biaya kembali ke dalam map yang bersangkutan.
   - Cetak informasi detail untuk setiap parsel.

3. **Fungsi `hitungBiayaParcel`**:
   - Terima berat parsel sebagai argumen.
   - Hitung biaya total berdasarkan berat dalam kilogram dan gram.
   - Kembalikan total biaya ke fungsi pemanggil.

#### Cara Kerja Program

1. Program dimulai dengan mendeklarasikan slice `parcels` yang berisi beberapa map, masing-masing menyimpan informasi berat dan detail berat untuk tiga parsel.
   
2. Dengan menggunakan loop `for`, program iterasi melalui setiap parsel yang ada di dalam slice `parcels`:
   - Menghitung biaya dengan memanggil fungsi `hitungBiayaParcel`, dengan parameter berat dari parsel.
   - Menyimpan total biaya kembali ke dalam map parsel.
   - Mencetak informasi dari setiap parsel, termasuk:
     - Nomor contoh parsel.
     - Berat parsel dalam gram.
     - Detail berat dalam format kilogram dan gram.
     - Total biaya pengiriman.

3. Fungsi `hitungBiayaParcel` melakukan langkah berikut:
   - Mengatur tarif biaya per kilogram dan biaya tambahan per gram.
   - Menghitung jumlah kilogram dan sisa gram dari berat parsel.
   - Menghitung total biaya berdasarkan tarif yang telah ditentukan.
   - Mengembalikan nilai total biaya.

## Unguided - 2C - 3
### Study Case
Sebuah bilangan bulat b memiliki faktor bilangan f > 0 jika f habis membagi b. Contoh: 2 merupakan faktor dari bilangan 6 karena 6 habis dibagi 2.

Buatlah program yang menerima input sebuah bilangan bulat b dan b > 1. Program harus dapat mencari dan menampilkan semua faktor dari bilangan tersebut!

Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):
![soalun5](https://github.com/user-attachments/assets/37014a1f-8cd0-4f30-aa01-fc88c99a26b6)

Bilangan bulat b > 0 merupakan bilangan prima p jika dan hanya jika memiliki persis dua faktor bilangan saja, yaitu 1 dan dirinya sendiri.

Lanjutkan program sebelumnya. Setelah menerima masukan sebuah bilangan bulat b > 0. Program tersebut mencari dan menampilkan semua faktor bilangan tersebut. Kemudian, program menentukan apakah b merupakan bilangan prima.

Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):
![soalun5 2](https://github.com/user-attachments/assets/ccfa27ea-67c4-47a8-88c4-d496d36656b1)

### Source Code
![unguided5](https://github.com/user-attachments/assets/a62c74e7-34b3-4f01-9b03-a6843540e8e7)

### Screenshot Output
![outputun5](https://github.com/user-attachments/assets/0bd95a67-b0e6-41ee-8e4a-e75ef1f09978)

### Deskripsi Program
Program ini dibuat untuk menghitung dan menampilkan faktor dari sebuah bilangan bulat yang dimasukkan oleh pengguna serta menentukan apakah bilangan tersebut adalah bilangan prima atau tidak. Program menggunakan input dari pengguna dan memprosesnya menggunakan loop untuk mencari faktor dan melakukan pengecekan bilangan prima.

#### Algoritma

1. **Input Bilangan**:
   - Program meminta pengguna untuk memasukkan bilangan bulat (`b`).

2. **Mencetak Faktor**:
   - Menggunakan loop `for` dari 1 hingga `b`, program memeriksa setiap angka (`i`) untuk menentukan apakah `i` adalah faktor dari `b` dengan memeriksa jika `b % i == 0`.
   - Jika kondisi tersebut benar, program mencetak `i` sebagai faktor.

3. **Pengecekan Bilangan Prima**:
   - Inisialisasi variabel `prima` sebagai `true`.
   - Jika `b` kurang dari atau sama dengan 1, ubah `prima` menjadi `false` karena bilangan 0 dan 1 bukanlah bilangan prima.
   - Jika `b` lebih besar dari 1, program menggunakan loop untuk memeriksa kemungkinan faktor dari 2 hingga akar kuadrat dari `b` (dengan `i*i <= b`).
   - Jika ditemukan suatu `i` yang membagi `b` (yaitu `b % i == 0`), set `prima` ke `false` dan keluar dari loop.

4. **Menampilkan Hasil**:
   - Program mencetak apakah bilangan yang dimasukkan adalah bilangan prima atau tidak berdasarkan nilai dari variabel `prima`.

#### Cara Kerja Program

1. Program dimulai dengan mendeklarasikan variabel `b` untuk menyimpan input bilangan bulat dari pengguna.
   
2. Pengguna diminta untuk memasukkan sebuah bilangan, yang kemudian disimpan dalam variabel `b` menggunakan `fmt.Scanln(&b)`.

3. Selanjutnya, program mencetak "Faktor: " dan menggunakan loop `for` untuk mencari semua faktor dari bilangan `b`:
   - Untuk setiap bilangan `i` dari 1 hingga `b`, program memeriksa apakah `b` dapat dibagi habis oleh `i`.
   - Jika bisa, `i` dicetak sebagai faktor.

4. Setelah mencetak semua faktor, program melanjutkan untuk menentukan apakah `b` adalah bilangan prima:
   - Jika `b` kurang dari atau sama dengan 1, program langsung menetapkan `prima` ke `false`.
   - Jika `b` lebih dari 1, program memeriksa semua angka dari 2 hingga akar kuadrat dari `b` untuk mencari faktor.
   - Jika ditemukan faktor, `prima` diubah menjadi `false`.

5. Program akhirnya mencetak hasil pengecekan bilangan prima dalam format "Prima: true" atau "Prima: false".
