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

## DASAR TEORI
<p style="text-align: justify; font-size: 12px;">

### Bahasa Go
Bahasa Go, yang juga dikenal dengan nama Golang, adalah bahasa pemrograman yang dikembangkan oleh Google pada tahun 2007 dan dirilis secara publik pada tahun 2009. Go dirancang untuk menjadi bahasa yang sederhana, efisien, dan mudah dipahami, dengan fitur-fitur yang mendukung pemrograman paralel dan konkuren. Go sering digunakan dalam pengembangan sistem, aplikasi web, dan infrastruktur cloud.

<br>

Filosofi desain Go bertujuan untuk:<br>
- Kesederhanaan: Sintaks yang mudah dipahami dan minimistik.
- Kecepatan: Kompilasi yang cepat dan performa runtime yang baik.
- Keterbacaan: Kode yang mudah dibaca dan dipelihara.

### Struktur Program GO 
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

### Tipe Data Dalam Go
Go memiliki beberapa tipe data, yang dibagi menjadi dua kategori: tipe data dasar dan tipe data terstruktur.<br>
#### 1. Tipe Data Dasar
- Integer: Tipe untuk bilangan bulat. Contoh: `int`, `int8`, `int16`, `int32`, `int64`.
- Float: Tipe untuk bilangan pecahan. Contoh: `float32`, `float64`.
- Boolean: Tipe yang hanya memiliki dua nilai: `true` dan `false`.
- String: Tipe untuk teks, ditandai dengan tanda kutip ganda.

#### 2. Tipe Data Terstruktur
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

### Struktur Kontrol
Struktur kontrol adalah salah satu konsep dasar dalam pemrograman yang memungkinkan pengembang untuk mengontrol alur eksekusi program berdasarkan kondisi tertentu. Dalam bahasa pemrograman Go (Golang), struktur kontrol mirip dengan bahasa pemrograman lainnya, namun ada beberapa fitur dan karakteristik unik yang membedakan Go dari bahasa lain. Berikut adalah dasar teori tentang struktur kontrol dalam bahasa Go:
<br>

#### 1. Statement If-Else
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

#### 2. Statement switch
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

#### 3. Statement for
Go hanya memiliki satu jenis loop, yaitu for. Namun, for di Go cukup fleksibel dan bisa digunakan untuk berbagai kasus iterasi, baik sebagai loop klasik, loop berbasis kondisi, atau loop tanpa batas.
<br>
**Sintaks:**
```go
for initialization; condition; post {
    // code block
}
```

### Fungsi
Fungsi adalah blok kode yang melakukan tugas tertentu. Fungsi dapat memiliki parameter dan nilai kembali.
<br>
**Contoh Fungsi**
```go
func tambah(a int, b int) int {
    return a + b
}
```

### Error Handling
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

# GUIDED



