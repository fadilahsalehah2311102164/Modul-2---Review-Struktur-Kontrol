<p align="center">
  <strong>LAPORAN PRAKTIKUM</strong>
  <br>
  <strong>ALGORITMA DAN PEMROGRAMAN 2</strong>
  <br>
</p>

<br>

<p align="center">
  <strong>MODUL II</strong>
  <br>
  <strong>REVIEW STRUKTUR KONTROL</strong>
  <br>
</p>

<br>

<p align="center">
  <img src="https://github.com/user-attachments/assets/296eb3c2-bd6b-401f-8256-3661150ec39e" alt="Logo" width="200"/>
</p>

<br>

<p align="center">
  <strong>Disusun Oleh :</strong>
  <br>
  Mutia Rani Zahra Meilani
  <br>
  2311102182
  <br>
  S1IF1105
</p>

<br>

<p align="center">
  <strong>Dosen Pengampu :</strong><br>
  Arif Amrulloh, S.Kom., M.Kom.
</p>

<br>

<p align="center">
  <strong>PROGRAM STUDI S1 TEKNIK INFORMATIKA</strong>
  <br>
  <strong>FAKULTAS INFORMATIKA</strong>
  <br>
  <strong>TELKOM UNIVERSITY PURWOKERTO</strong>
  <br>
  <strong>2024</strong>
</p>

## <strong> <span style="color:#E5D9F2"> DASAR TEORI </span> </strong>

<span style="color:#C4D7FF; font-size:16px"><strong> ── PENGERTIAN GOLANG</strong></span>
<br>
Golang adalah bahasa pemrograman yang diciptakan dan dikembangkan oleh tim Engineer Google pada tahun 2009. Awalnya bahasa tersebut hanya digunakan untuk kepentingan internal. Kemudian bahasa ini dirilis untuk kepentingan public dan bersifat Open Source sehingga siapapun bisa mengembangkan bahasa Go atau Golang.

<span style="color:#C4D7FF; font-size:16px"><strong> ── TIPE DATA</strong></span>
<br>
Tipe data dalam Golang mendefinisikan jenis nilai yang dapat disimpan dalam sebuah variabel. Golang merupakan bahasa pemrograman yang strongly typed, artinya setiap variabel harus memiliki tipe data yang jelas.

**Tipe Data Dasar :**

- Integer : Menyimpan bilangan bulat (misal: int8, int16, int32, int64, uint8, uint16, uint32, uint64). Ukurannya bervariasi tergantung pada jumlah bit yang digunakan.
- Float : Menyimpan bilangan desimal (misal: float32, float64). Digunakan untuk perhitungan yang melibatkan bilangan pecahan.
- String : Kumpulan karakter (misal: "hello", "world"). String dalam Golang sebenarnya adalah slice of bytes, sehingga dapat dimanipulasi seperti array.
- Boolean : Menyimpan nilai benar atau salah (true atau false). Digunakan untuk membuat keputusan dalam program.

**Contoh Penggunaan Tipe Data Dasar**

```go
// Tipe Data Integer
var umur int = 25

// Tipe Data Float
var tinggi float64 = 175.5

// Tipe Data String
var nama string = "John Doe"

// Tipe Data Boolean
var isMarried bool = false
```

**Tipe Data Lainnya :**

- Array : Kumpulan nilai dengan tipe data yang sama dan ukuran yang tetap. Elemen array diakses menggunakan indeks yang dimulai dari 0.
- Slice : Mirip dengan array, tetapi ukurannya dinamis. Slice merupakan bagian dari array yang dapat berubah ukuran.
- Map : Kumpulan pasangan kunci-nilai, di mana kunci harus unik dan tipe datanya bisa berupa string, integer, atau jenis data dasar lainnya.
- Struct : Tipe data yang mendefinisikan kumpulan field dengan tipe data yang berbeda-beda. Struct digunakan untuk merepresentasikan data yang kompleks.

**Contoh Penggunaan Tipe Data Lainnya**

```go
// Tipe Data Array
var angka [5]int = [5]int{1, 2, 3, 4, 5}

// Tipe Data Slice
var buah []string = []string{"apel", "pisang", "mangga"}

// Tipe Data Map
var capital = map[string]string{"Indonesia": "Jakarta", "Inggris": "London"}

// Tipe Data Struct
type Person struct {
    Name  string
    Age   int
    City  string
}
var person Person
person.Name = "Alice"
person.Age = 30
person.City = "New York"

```

<span style="color:#C4D7FF; font-size:16px"><strong> ── PERCABANGAN</strong></span>
<br>
Percabangan digunakan untuk membuat keputusan dalam program, yaitu menjalankan blok kode tertentu berdasarkan kondisi yang diberikan. Golang menyediakan pernyataan if dan switch untuk melakukan percabangan.

**If Statement**
<br>
Digunakan untuk mengeksekusi blok kode jika kondisi yang diberikan bernilai benar (true). Ada fitur short statement yang memungkinkan kita mendeklarasikan variabel sebelum memeriksa kondisi, dan variabel tersebut hanya berlaku dalam lingkup if statement.

Contoh Penggunaan If Statement

```go
if nilai := 75; nilai >= 70 {
    fmt.Println("Anda lulus")
}
```

**If-Else Statement**
<br>
Ketika kita memiliki lebih dari satu kondisi yang ingin diuji, kita dapat menggunakan if-else untuk membuat percabangan yang lebih kompleks.

Contoh Penggunaan If Else Statement :

```go
if nilai >= 80 {
    fmt.Println("A")
} else if nilai >= 70 {
    fmt.Println("B")
} else if nilai >= 60 {
    fmt.Println("C")
} else {
    fmt.Println("D")
}
```

**Switch Statement**
<br>
Switch merupakan alternatif yang lebih bersih daripada if-else ketika kita memiliki banyak kondisi untuk diuji. Switch juga mendukung multiple cases dalam satu baris dan bahkan dapat digunakan tanpa ekspresi, yang akan bertindak seperti rangkaian if-else.

Contoh Penggunaan Switch Statement :

```go
switch hari {
case "Senin":
    fmt.Println("Hari kerja")
case "Sabtu", "Minggu":
    fmt.Println("Hari libur")
default:
    fmt.Println("Hari biasa")
}
```

<span style="color:#C4D7FF; font-size:16px"><strong> ── PERULANGAN</strong></span>
<br>
Perulangan digunakan untuk menjalankan blok kode secara berulang hingga kondisi tertentu terpenuhi. Golang menyediakan pernyataan for untuk melakukan perulangan.

**For Loop Standar**
<br>
Ini adalah bentuk perulangan yang paling umum, di mana kita mengatur nilai awal, kondisi, dan inkrementasi dalam satu baris. Perulangan ini sangat cocok untuk iterasi dengan jumlah yang diketahui.

Contoh Penggunaan For Loop :

```go
for i := 0; i < 5; i++ {
    fmt.Println(i)
}
```

**For Sebagai While**
<br>
Jika kita hanya memberikan kondisi tanpa inisialisasi atau inkrementasi, for loop dapat berfungsi seperti while loop, di mana iterasi terus berjalan selama kondisi terpenuhi.

Contoh Penggunaan For sebagai While:

```go
n := 1
for n < 5 {
    fmt.Println(n)
    n++
}
```

**Infinite Loop**
<br>
tanpa kondisi apapun menghasilkan perulangan tak terbatas. Ini sering digunakan bersamaan dengan break untuk menghentikan loop ketika kondisi tertentu terpenuhi.

Contoh Penggunaan Infinite Loop:

```go
for {
    fmt.Println(n)
    n++
    if n > 5 {
        break
    }
}
```

**For-Range Loop**
<br>
For-range digunakan untuk iterasi pada koleksi seperti slice, array, map, atau string. Pada slice, kita bisa mengakses indeks dan nilai, sementara pada map kita bisa mengakses kunci dan nilai.

Contoh Penggunaan For-Range Loop :

```go
slice := []string{"apple", "banana", "cherry"}
for i, nilai := range slice {
    fmt.Printf("index: %d, nilai: %s\n", i, nilai)
}
```

<span style="color:#C4D7FF; font-size:16px"><strong> ── CONTROL STATEMENT DALAM PERULANGAN</strong></span>
<br>

**Break**
<br>
Perintah ini digunakan untuk menghentikan perulangan lebih awal, biasanya digunakan untuk keluar dari *for loop* saat kondisi tertentu tercapai.
  
Contoh Penggunaan Break:
```go
  for i := 0; i < 10; i++ {
    if i == 5 {
        break
    }
    fmt.Println(i)
}
```

**Continue**
<br>
Continue digunakan untuk melewati iterasi saat ini dan melanjutkan ke iterasi berikutnya. Ini sering dipakai ketika kita ingin menghindari kondisi tertentu dalam loop.

Contoh Penggunaan Continue :
```go
  for i := 0; i < 5; i++ {
    if i == 2 {
        continue
    }
    fmt.Println(i)
}
```

## <strong> <span style="color:#E5D9F2"> GUIDED </span> </strong>

### <span style="color:#FFCFB3"> ── Guided 1 </span>

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

#### Output

![image](https://github.com/user-attachments/assets/07b9b52c-3681-4222-b24f-78c1d501bf86)

Deskripsi Program :
Program ini bertujuan untuk menerima input nama dari pengguna dan mencetak kembali nama yang telah diinputkan pengguna.

### <span style="color:#FFCFB3"> ── Guided 2 </span>

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
	correctOrder := []string{"merah", "kuning", "hijau", "ungu"}

	reader := bufio.NewReader(os.Stdin)
	succes := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d : ", i)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		colors := strings.Split(input, " ")

		for j := 0; j < 4; j++ {
			if colors[j] != correctOrder[j] {
				succes = false
				break
			}
		}

		if !succes {
			break
		}
	}

	if succes {
		fmt.Println("BERHASIL: TRUE")
	} else {
		fmt.Println("BERHASIL: FALSE")
	}
}
```

#### Output

Jika urutan warna benar :
<br>
![image](https://github.com/user-attachments/assets/33b834bc-9ffd-4de0-ae90-30151dc06061)

Jika urutan warna salah :
<br>
![image](https://github.com/user-attachments/assets/1155b127-44cd-44a8-94c1-d8d5813775c8)

Deskripsi Program :
Program ini bertujuan untuk verifikasi urutan warna yang telah ditentukan oleh program. Program ini menggunakan struktur perulangan dan struktur percabangan. Pada program ini pengguna harus memasukan urutan warna yang benar sebanyak 5 kali. Jika pengguna berhasil memasukan urutan warna dengan benar maka program akan mencetak "BERHASIL: TRUE", namun jika pengguna gagal memasukan urutan warna maka program akan mencetak "BERHASIL: FALSE".

### <span style="color:#FFCFB3"> ── Guided 3 </span>

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
	fmt.Println("Hasil Penjumlahan ", a, b, c, d, e, "adalah ", hasil)
}
```

#### Output

![image](https://github.com/user-attachments/assets/e402a026-ba0a-4e3c-8e55-4d4a2f86941e)

Deskripsi Program :
Program ini bertujuan untuk menghitung jumlah bilangan yang diinput oleh pengguna. Program ini hanya menggunakan operasi aritmatika sederhana. Program akan meminta pengguna untuk memasukan 5 bilang bulat, kemudian program akan menjumlahkan bilangan tersebut dan mecetak hasilnya.

### <span style="color:#FFCFB3"> ── Guided 4 </span>

#### Source Code

```go
package main

import "fmt"

func main() {
	var nam float32
	var nmk string

	fmt.Print("Masukan nilai : ")
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
	fmt.Printf("Nilai Indeks untuk Nilai %.2f adalah : %s\n ", nam, nmk)
}
```

#### Output

![image](https://github.com/user-attachments/assets/02c961a6-be9b-432e-a20f-a8109ef504f5)

Deskripsi Program :
Program ini bertujuan sebagai sistem penilaian sederhana. Program ini menggunakan struktur percabangan if-else untuk menentukan grade dari sebuah rentang nilai. Program ini akan meminta pengguna untuk memasukan nilai secara numerik, kemudian program akan menentukan grade dari rentang nilai yang dimasukan pengguna, misalnya nilai 85 akan mendapatkan grade A. Hasil akhir dari program ini berupa nilai yang telah dimasukan dan grade yang telah disesuaikan.

## <strong> <span style="color:#E5D9F2"> UNGUIDED </span> </strong>

### <span style="color:#FFCFB3"> ── Unguided 1 </span>

#### Study Case

Suatu pita (string) berisi kumpulan nama-nama bunga yang dipisahkan oleh spasi dan ‘-’, contoh pita diilustrasikan seperti berikut ini.

Pita : mawar - melati - tulip - teratal - kamboja - anggrek

Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita.

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
	bunga := []string{}
	reader := bufio.NewReader(os.Stdin)

	for perulangan := 1; ; perulangan++ {
		fmt.Printf("Bunga %d : ", perulangan)

		input, _ := reader.ReadString('\n')

		input = strings.TrimSpace(input)

		if input == "selesai" {
			break
		}

		bunga = append(bunga, input)
	}

	result := strings.Join(bunga, " - ")
	fmt.Println("Pita : " + result)
	fmt.Printf("Bunga : %d\n", len(bunga))
}
```

#### Output

![image](https://github.com/user-attachments/assets/c81df273-951b-4916-8a5c-35689581d1f8)

Deskripsi Program :
Program ini bertujuan untuk membuat pita dari nama nama bunga yang telah diinputkan pengguna. Program ini menggunakan struktur perulangan. Pengguna dapat terus memasukan nama bunga hingga mereka memasukan kata "selesai". Setelah proses input selesai, program akan menampilkan semua nama bunga yang telah diinputkan pengguna dengan tanda hubung "-" sebagai sebuah pita, serta menampilkan jumlah total bungga yang telah diinputkan.

### <span style="color:#FFCFB3"> ──  Unguided 2 </span>

#### Study Case

Setiap hari Pak Andi membawa banyak barang belanjaan dari pasar dengan mengendarai sepeda motor. Barang belanjaan tersebut dibawa dalam kantong terpal di kiri-kanan motor. Sepeda motor tidak akan oleng jika selisih berat barang di kedua kantong sisi tidak lebih dari 9 kg.

Buatlah program Pak Andi yang menerima input dua buah bilangan real positif yang menyatakan berat total masing-masing isi kantong terpal. Program akan terus meminta input bilangan tersebut hingga salah satu kantong terpal berisi 9 kg atau lebih.

Pada modifikasi program tersebut, program akan menampilkan True jika selisih kedua isi kantong lebih dari atau sama dengan 9 kg. Program berhenti memproses apabila total berat isi kedua kantong melebihi 150 kg atau salah satu kantong beratnya negatif.

#### Source Code

```go
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for i := 0; i < 4; i++ {
		fmt.Printf("Masukkan berat belanjaan di kedua kantong (pisahkan dengan spasi) : ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		parts := strings.Split(input, " ")

		if len(parts) != 2 {
			fmt.Println("Harap masukkan dua angka yang dipisahkan oleh spasi")
			continue
		}

		nilai1, err1 := strconv.Atoi(parts[0])
		nilai2, err2 := strconv.Atoi(parts[1])

		if err1 != nil || err2 != nil {
			fmt.Println("Angka yang diberikan tidak valid")
			continue
		}

		if nilai1 < 0 || nilai2 < 0 {
			fmt.Println("Berat yang diberikan tidak boleh negatif")
			continue
		}

		selisih := int(math.Abs(float64(nilai1 - nilai2)))

		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng : TRUE")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng : FALSE")
		}
	}

	fmt.Println("proses selesai")
}

```

#### Output
![image](https://github.com/user-attachments/assets/8a5dd220-5f81-4fae-b09d-420553de4030)

Deskripsi Program :
Program ini bertujuan untuk menentukan apakah sepeda motor Pak Andi akan oleng berdasarkan berat belanjaan di dua kantongnya. Program ini menggunakan struktur perulangan. Pengguna dapat memasukan berat belanjaan dia dua kantong sebanyak 4 kali, kemudia program akan menghitung selisih berat keduanya. Jika selisihnya lebih besar atau sama dengan 9, maka program akan mencetak "Sepeda motor pak Andi akan oleng : TRUE". Jika tidak, program akan mencetak "Sepeda motor pak Andi akan oleng : FALSE".

### <span style="color:#FFCFB3"> ── Unguided 3 </span>

#### Study Case

$$f(k) = \frac{(4k + 2)^2}{(4k + 1)(4k + 3)}$$

Buatlah sebuah program yang menerima input sebuah bilangan sebagai K, kemudian menghitung dan menampilkan nilai f(k) sesuai persamaan di atas.

#### Source Code

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Nilai K : ")

	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	k, err := strconv.ParseFloat(input, 64)
	if err != nil {
		fmt.Println("Nilai yang diberikan tidak valid")
		return
	}

	divK1 := ( k * 4 + 2 ) * ( k * 4 + 2 )
	divK2 := ( 4 * k + 1 ) * ( 4 * k + 3 )

	total := divK1 / divK2

	fmt.Printf("Nilai F(K) = %.10f\n", total)
}

```

#### Output
![image](https://github.com/user-attachments/assets/a42b970e-88a6-408b-b8d2-8e281ab4aeb5)

Deskripsi Program :
Program ini bertujuan untuk menghitung nilai fungsi F(k) berdasarkan rumus yang diberikan. Program ini menggunakan struktur if statement dan operasi aritmatika sederhana. Program akan meminta pengguna untuk memasukan nilai K, kemudian program akan menghitung F(k) sesuai dengan rumus yang diberikan. Hasil dari perhitunggan kemudian ditampilkan dengan 10 digit desimal.


### <span style="color:#FFCFB3"> ── Unguided 4 </span>

#### Study Case

PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BiayaPos untuk menghitung biaya pengiriman tersebut dengan ketentuan sebagai berikut:

Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram).
Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 15,- per gram saja. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.

#### Source Code

```go
package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sisa_berat int
	var biaya_tambahan int

	reader := bufio.NewReader(os.Stdin)

	fmt.Printf("Berat parsel (gram) : ")
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)

	berat_parsel, err := strconv.Atoi(input)

	if err != nil {
		fmt.Println("Berat yang dimasukkan tidak valid")
		return
	}

	sisa_berat = berat_parsel % 1000

	if sisa_berat < 500 {
		biaya_tambahan = sisa_berat * 15
	} else {
		biaya_tambahan = sisa_berat * 5
	}

	berat_kg := math.Floor(float64(berat_parsel / 1000))

	fmt.Printf("Detail berat : %.0f kg + %d gram\n", berat_kg, sisa_berat)
	fmt.Printf("Detail Biaya : %d + %d\n", berat_parsel * 10, biaya_tambahan)

	var totalBiaya float64
	if berat_parsel > 10000 {
		totalBiaya = float64(berat_parsel * 10)
	} else {
		totalBiaya = berat_kg*10000 + float64(biaya_tambahan)
	}

	fmt.Printf("Total Biaya: %.0f\n", totalBiaya)
}

```

#### Output
![image](https://github.com/user-attachments/assets/91e9d6e7-39b3-436b-bf87-9b1ffb80d98f)

Deskripsi Program :
Program ini digunakan untuk menghitung biaya pengiriman barang berdasarkan beratnya. Program ini menggunakan struktur if-else dan operasi aritmatika. Program akan meminta pengguna memasukan berat parsel, kemudian menghitung biaya dasar berdasarkan berat dalam kilogram, dan menambahkan biaya tambahan untuk sisa berat yang kurang dari 1 kilogram. Hasil akhir dari program ini berupa rincian berat, biaya dasar, biaya tambahan, dan total biaya yang harus dibayar.


### <span style="color:#FFCFB3"> ── Unguided 5 </span>

#### Study Case

Diberikan sebuah nilai akhir mata kuliah (NAM) [0..100] dan standar penilaian nilai mata kuliah (NMK) sebagai berikut:

| NAM             | NMK |
| --------------- | --- |
| NAM ≥ 80        | A   |
| 72.5 < NAM < 80 | AB  |
| 65 < NAM ≤ 72.5 | B   |
| 57.5 < NAM ≤ 65 | BC  |
| 50 < NAM ≤ 57.5 | C   |
| 40 < NAM ≤ 50   | D   |
| NAM < 40        | E   |

Program berikut menerima input sebuah bilangan riil yang menyatakan NAM. Program menghitung NMK dan menampilkannya.

Jawablah pertanyaan-pertanyaan berikut:

a. Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal?

b. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya!

c. Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; dan 49.5. Seharusnya keluaran yang diperoleh adalah 'A', 'B', dan 'D'.

#### Source Code

```go
package main

import "fmt"

func main() {
	var nam float64
	var nmk string

	fmt.Print("Nilai Akhir mata kuliah : ")
	fmt.Scanln(&nam)

	if nam >= 0 && nam <= 100 {
		if nam > 80 {
			nmk = "A"
		}else if nam > 72.5 {
			nmk = "AB"
		}else if nam > 65 {
			nmk = "B"
		}else if nam > 57.5 {
			nmk = "BC"
		}else if nam > 50 {
			nmk = "C"
		}else if nam > 40 {
			nmk = "D"
		}else{
			nmk = "E"
		}

		fmt.Println("Nilai Mata Kuliah : ", nmk)
	}else{
		fmt.Println("Nilai yang diberikan tidak valid.")
	}

}
```

#### Output
![image](https://github.com/user-attachments/assets/600b4bd7-e98b-40e3-ad21-63c7d1e42016)

Deskripsi Program :
Program ini bertujuan sebagai sistem penilaian sederhana. Program ini menggunakan struktur percabangan if-else untuk menentukan grade dari sebuah rentang nilai. Program ini akan meminta pengguna untuk memasukan nilai secara numerik, kemudian program akan menentukan grade dari rentang nilai yang dimasukan pengguna, misalnya nilai 85 akan mendapatkan grade A. Hasil akhir dari program ini berupa nilai yang telah dimasukan dan grade yang telah disesuaikan.

a. Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal?
<br>
= Ketika User memberikan input nam senilai 80.1, sistem menampilkan nilai mata kuliah berupa A, dimana nmk yang ditampilkan sudah sesuai dengan spesifikasi soal. Hal ini menandakan bahwa program dan seluruh kondisi percabangan sudah sesuai

b. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya!
<br>
= Terdapat 3 kesalahan pada baris kode awal yang diberikan. Kesalahan pertama ialah tidak adanya kontrol percabangan untuk mengatasi nilai input yang diberikan di bawah 0 ataupun di atas 100. Kesalahan kedua terletak pada seluruh kontrol percabangan yang memberikan nilai akhir, dimana didalam kontrol percabangan variabel nam di tugaskan ulang untuk mengganti nilai sebelumnya. Seharusnya variabel nmk yang ditugaskan untuk memberi nilai mata kuliah. Kesalahan ketiga terletak ketidak konsistenan penggunaan else if dan else, dimana pada baris kode terbaru, kontrol percabangan sudah konsisten menggunakan if, else if dan else.

c. Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; dan 49.5. Seharusnya keluaran yang diperoleh adalah 'A', 'B', dan 'D'.
<br>
= Pada program tersebut sudah tidak perlu ada yang diperbaiki jika menggunakan	if else if dan else yang konsisten. dimana jika memberikan input 93.5 nilai yang tampil sudah A, 70.6 nilai yang tampil sudah B, dan 49.5 nilai yang tampil sudah D.

### <span style="color:#FFCFB3"> ── Unguided 6 </span>

#### Study Case

Sebuah bilangan bulat b memiliki faktor bilangan f > 0 jika f habis membagi b. Contoh: 2 merupakan faktor dari bilangan 6 karena 6 habis dibagi 2.

Buatlah program yang menerima input sebuah bilangan bulat b dan b > 1. Program harus dapat mencari dan menampilkan semua faktor dari bilangan tersebut!

Bilangan bulat b > 0 merupakan bilangan prima p jika dan hanya jika memiliki persis dua faktor bilangan saja, yaitu 1 dan dirinya sendiri.

Lanjutkan program sebelumnya. Setelah menerima masukan sebuah bilangan bulat b > 0. Program tersebut mencari dan menampilkan semua faktor bilangan tersebut. Kemudian, program menentukan apakah b merupakan bilangan prima.

#### Source Code

```go
package main

import (
	"fmt"
)

func main() {
	var bilangan int
	fmt.Printf("Bilangan : ")
	fmt.Scanln(&bilangan)

	jumlahFaktor := 0

	fmt.Print("Faktor : ")
	for i := 1; i <= bilangan; i++ {
		if bilangan%i == 0 {
			fmt.Print(i, " ")
			jumlahFaktor++
		}
	}
	fmt.Println()


	if jumlahFaktor == 2 {
		fmt.Println("Prima : TRUE")
	} else {
		fmt.Println("Prima : FALSE")
	}
}
```

#### Output
![image](https://github.com/user-attachments/assets/a3e85717-231a-4bb9-a18e-819ff14400ed)

Deskripsi Program :
Program ini digunakan untuk menentukan apakah sebuah bilangan yang dimasukkan oleh pengguna merupakan bilangan prima atau bukan. Program akan mencari semua faktor dari bilangan tersebut dengan cara membagi bilangan tersebut dengan angka-angka dari 1 hingga bilangan itu sendiri. Jika jumlah faktor yang ditemukan hanya 2 (yaitu 1 dan bilangan itu sendiri), maka bilangan tersebut adalah bilangan prima. Program ini juga akan menampilkan semua faktor dari bilagan tersebut.

## <strong> <span style="color:#E5D9F2"> REFERENSI </span> </strong>

#### [1] Novalagung. (n.d.). Dasar Pemrograman Golang. Diakses dari https://dasarpemrogramangolang.novalagung.com/

#### [2] Golang Official Documentation. Diakses dari https://go.dev/doc/

#### [3] A Tour of Go. Diakses dari https://tour.golang.org/
