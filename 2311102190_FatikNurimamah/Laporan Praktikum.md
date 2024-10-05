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
Struktur kontrol dari Go berkaitan dengan bahasa C namun berbeda dalam hal-hal tertentu[1]. Tidak ada pengulangan do atau while, hanya for; switch yang lebih fleksibel; if dan switch bisa menggunakan perintah inisialisasi seperti halnya pada for; perintah break dan continue memiliki label identifikasi yang opsional; dan ada beberapa kontrol struktur baru termasuk switch pada tipe dan komunikasi multiplexer, select[1]. Sintaksnya juga sedikit berbeda: tidak ada tanda kurung dan bagian badan dari kontrol harus selalu dibatasi oleh kurung kurawal[1].

**1. `if`**
   
Dalam Go, bentuk if yang paling sederhana adalah sebagai berikut:

```go
if x > 0 {
	return y
}
```

Kewajiban penggunaan kurung kurawal mendorong penulisan perintah `if` menjadi lebih terstruktur dengan beberapa baris. Gaya penulisan ini sangat efektif, terutama jika badan kondisi mengandung perintah kontrol seperti `return` atau `break`[1].

Karena `if` dan `switch` dapat melakukan inisialisasi, sering kali kita melihat penggunaan kedua struktur ini untuk mendeklarasikan variabel lokal[1].

```go
if err := file.Chmod(0664); err != nil {
	log.Print(err)
	return err
}
```

Dalam pustaka Go, dapat dijumpai bahwa jika perintah `if` tidak melanjutkan ke perintah berikutnya—misalnya, ketika badan kondisi diakhiri dengan `break`, `continue`, `goto`, atau `return`—maka kondisi `else` yang tidak diperlukan akan dihilangkan[1].

```go
f, err := os.Open(name)
if err != nil {
	return err
}
codeUsing(f)
```

Ini adalah contoh situasi umum di mana kode perlu mempertahankan urutan penanganan kesalahan. Kode akan lebih mudah dibaca jika kontrol yang berhasil terus mengalir ke bawah, sehingga menghilangkan kasus kesalahan saat mereka muncul. Karena kasus kesalahan cenderung diakhiri dengan perintah `return`, maka kode tidak memerlukan perintah `else`[1].

```go
f, err := os.Open(name)
if err != nil {
	return err
}
d, err := f.Stat()
if err != nil {
	f.Close()
	return err
}
codeUsing(f, d)
```

**2. Deklarasi dan Penemparan Ulang**

Contoh terakhir dari bagian sebelumnya menunjukkan detail tentang cara kerja deklarasi singkat `:=`. Deklarasi yang menggunakan `os.Open` terdengar seperti ini:

```go
f, err := os.Open(name)
```

Perintah tersebut mendeklarasikan dua variabel, yaitu `f` dan `err`. Beberapa baris kemudian, terdapat pembacaan panggilan ke `f.Stat`[1].

```go
d, err := f.Stat()
```

Pernyataan tersebut tampak seperti mendeklarasikan variabel `d` dan `err`. Perhatikan bahwa `err` muncul di kedua pernyataan[1]. Duplikasi seperti ini diperbolehkan: `err` dideklarasikan dalam pernyataan pertama, tetapi dapat digunakan kembali dalam pernyataan kedua[1]. Ini berarti panggilan ke `f.Stat` menggunakan variabel `err` yang sama yang dideklarasikan sebelumnya, dan diberi nilai baru[1].

Ketika mendeklarasikan sebuah variabel `v` dengan `:=`, variabel tersebut dapat digunakan kembali meskipun telah dideklarasikan sebelumnya, asalkan:

1. Deklarasi berada dalam cakupan yang sama dengan deklarasi sebelumnya dari `v` (jika `v` dideklarasikan di luar cakupan, deklarasi ini akan menciptakan variabel baru)[1].
2. Nilai yang diinisialisasi dapat disimpan dalam `v`[1].
3. Setidaknya ada satu variabel baru lainnya dalam deklarasi tersebut[1].

Sifat yang tidak biasa ini bersifat pragmatis, memudahkan kita untuk menggunakan nilai tunggal `err`, misalnya, dalam beberapa pernyataan `if-else`. Anda akan sering menemukan penggunaan seperti ini[1].

Perlu diingat bahwa dalam Go, cakupan parameter fungsi dan nilai kembalian sama dengan tubuh fungsi, meskipun tampaknya secara leksikal berada di luar kurung kurawal yang menutup tubuh fungsi[1].

**3. `For`**

Pengulangan `for` di Go mirip—tetapi tidak identik—dengan yang ada di C[1]. Ia menggabungkan fungsi `for` dan `while`, dan tidak ada bentuk `for-while`[1]. Terdapat tiga bentuk pengulangan `for`, namun hanya satu yang menggunakan titik koma[1].

```go
// Seperti "for" pada C
for inisialisasi; kondisi; selanjutnya { }

// Seperti "while" pada C
for kondisi { }

// Seperti "for(;;)" pada C
for { }
```

Deklarasi singkat membuatnya mudah mendeklarasikan variabel index di dalam pengulangan[1].

```go
sum := 0
for i := 0; i < 10; i++ {
	sum += i
}
```

Ketika melakukan pengulangan pada array, slice, string, atau map, atau saat membaca dari sebuah channel, klausa `range` dapat digunakan dalam pengulangan[1].

```go
for key, value := range oldMap {
	newMap[key] = value
}
```

Jika hanya memerlukan item pertama dalam `range` (yaitu, kunci dari map atau indeks dari array/slice/string), variabel kembalian kedua dapat dihilangkan[1].

```go
for key := range m {
	if key.expired() {
		delete(m, key)
	}
}

```

Jika hanya memerlukan item kedua (nilainya), gunakan pengidentifikasi kosong, yaitu garis bawah, untuk mengabaikan yang pertama[1].

```go
sum := 0
for _, value := range array {
	sum += value
}
```

Go tidak memiliki operator koma, dan perintah `dan` serta `--` bukanlah ekspresi. Oleh karena itu, jika ingin menggunakan beberapa variabel dalam satu pernyataan `for`, kita perlu menggunakan penempatan paralel, karena cara ini tidak mengizinkan penggunaan `dan` atau `--`[1].

```go
// Reverse a
for i, j := 0, len(a)-1; i < j; i, j = i+1, j-1 {
	a[i], a[j] = a[j], a[i]
}
```

**4. `Switch`**

Kontrol `switch` di Go lebih generik dibandingkan dengan C. Ekspresi `switch` di Go tidak harus berupa konstanta atau integer; bagian kondisi `case` dievaluasi dari atas ke bawah hingga menemukan kondisi yang cocok. Jika ekspresi `switch` tidak ada, maka akan memeriksa kondisi `case` yang bernilai true[1]. Oleh karena itu, sangat memungkinkan—dan menjadi idiomatis—untuk menulis kondisi `if-else-if-else` menggunakan `switch`[1].

```go
func unhex(c byte) byte {
	switch {
	case '0' <= c && c <= '9':
		return c - '0'
	case 'a' <= c && c <= 'f':
		return c - 'a' + 10
	case 'A' <= c && c <= 'F':
		return c - 'A' + 10
	}
	return 0
}
```

Perintah `break` di Go dapat digunakan untuk mengakhiri blok `switch`[1]. Terkadang, perlu juga untuk keluar dari pengulangan, tetapi bukan dari `switch`, dan dalam Go, hal ini dapat dilakukan dengan memberikan label pada pengulangan dan "keluar" dari label tersebut[1]. Contoh berikut menunjukkan penggunaan kedua jenis `break` tersebut.

```go
Loop:
	for n := 0; n < len(src); n += size {
		switch {
		case src[n] < sizeOne:
			if validateOnly {
				break // keluar dari switch, tapi tetap dalam pengulangan `for`
			}
			size = 1
			update(src[n])

		case src[n] < sizeTwo:
			if n+1 >= len(src) {
				err = errShortInput
				break Loop // keluar dari `switch` dan pengulangan `for`
			}
			if validateOnly {
				break // keluar dari switch, tapi tetap dalam pengulangan `for`
			}
			size = 2
			update(src[n] + src[n+1]<&lt;shift)
		}
	}
```

Untuk mengakhiri bagian ini, berikut adalah fungsi yang membandingkan dua slice byte dengan menggunakan dua perintah switch:

```go
// Compare mengembalikan sebuah integer hasil pembandingan dari dua slice byte
// secara leksikografi.
// Hasilnya adalah 0 jika a == b, -1 jika a < b, dan +1 jika a > b .
func Compare(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}
	switch {
	case len(a) > len(b):
		return 1
	case len(a) < len(b):
		return -1
	}
	return 0
}
```

Perintah `switch` juga dapat digunakan untuk menentukan tipe dinamis dari sebuah variabel interface[1]. Tipe `switch` ini menggunakan sintaks dengan kata kunci `type` dalam tanda kurung[1]. Jika perintah `switch` mendeklarasikan sebuah variabel dalam ekspresinya, variabel tersebut akan memiliki tipe yang sesuai di setiap klausa[1]. Dalam konteks ini, menjadi idiomatis untuk menggunakan nama yang sama, sehingga efeknya adalah mendeklarasikan variabel baru dengan nama yang sama tetapi dengan tipe yang berbeda di setiap `case`[1].

```go
var t interface{}
t = functionOfSomeType()
switch t := t.(type) {
default:
	fmt.Printf("unexpected type %T\n", t)     // %T mencetak tipe dari `t`
case bool:
	fmt.Printf("boolean %t\n", t)             // t bertipe bool
case int:
	fmt.Printf("integer %d\n", t)             // t bertipe int
case *bool:
	fmt.Printf("pointer to boolean %t\n", *t) // t bertipe *bool
case *int:
	fmt.Printf("pointer to integer %d\n", *t) // t bertipe *int
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
Program ini digunakan untuk meminta pengguna meng inputkan nama mereka. Setelah input diberikan, program akan menampilkan nama yang diinputkan di layar.

### Algoritma Program
1. Inisialisasi variabel `nama` untuk menyimpan data yang diinput oleh pengguna.
2. Tampilkan pesan yang meminta pengguna untuk memasukkan nama.
3. Gunakan `fmt.Scanln` untuk membaca nama yang dimasukkan oleh pengguna.
4. Tampilkan nama yang telah dimasukkan menggunakan `fmt.Println`.

### Cara Kerja Program
1. Meminta Nama: Program meminta pengguna untuk mengetikkan nama melalui konsol.
2. Menampilkan Nama: Setelah nama diinputkan, program akan menampilkan nama tersebut.


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
Program ini dibuat untuk menguji kemampuan pengguna dalam mengingat urutan warna tertentu. Pengguna diminta untuk memasukkan urutan warna dalam lima percobaan, dan program akan memeriksa apakah urutan yang dimasukkan sesuai dengan urutan yang benar. Hasil akhirnya akan menunjukkan apakah pengguna berhasil atau tidak.

### Algoritma Program
1. Tetapkan urutan warna yang benar dalam array `correctOrder`.
2. Inisialisasi `reader` untuk membaca input dari pengguna.
3. Atur variabel `success` menjadi `true` untuk menandakan bahwa semua percobaan berhasil.
4. Jalankan loop sebanyak lima kali untuk mendapatkan input dari pengguna:
   - Tampilkan nomor percobaan.
   - Baca input pengguna dan hapus spasi di awal dan akhir.
   - Pisahkan input berdasarkan spasi menjadi array `colors`.
   - Periksa apakah setiap warna dalam colors sesuai dengan `correctOrder`.
   - Jika ada warna yang tidak sesuai, ubah `success` menjadi `false` dan keluar dari loop.
5. Setelah loop selesai, tampilkan hasil berdasarkan nilai `success`.

### Cara Kerja Program
1. Urutan Warna yang Benar: Program menetapkan urutan warna yang benar, yaitu "merah", "kuning", "hijau", dan "ungu".
2. Membaca Input: Program menggunakan `bufio.NewReader` untuk mengambil input dari pengguna dalam lima percobaan.
3. Memeriksa Input: Setiap input yang diberikan oleh pengguna akan dipisahkan menjadi kata-kata berdasarkan spasi. Program kemudian memeriksa apakah setiap warna dalam input sesuai dengan urutan yang benar.
4. Menampilkan Hasil: Jika semua percobaan sesuai dengan urutan yang ditetapkan, program akan menampilkan "BERHASIL: true". Jika terdapat satu percobaan yang salah, program akan menampilkan "BERHASIL: false".

### 3. Penjumlahan 5 Angka dari Input Pengguna

### Source Code :
```go
package main

import "fmt"

func main() {
	var a, b, c, d, e int
	var hasil int
	fmt.Scanln(&a, &b, &c, &d, &e)

	hasil = a + b + c + d + e
	fmt.Println("Hasil Penjumlahan", a, b, c, d, e, "adalah =", hasil)
}
```
### Output:
![Screenshot 2024-10-05 131743](https://github.com/user-attachments/assets/d6ebbbd9-5959-45eb-90f0-b6dc491f144f)

### Full code Screenshot:
![Screenshot 2024-10-05 131800](https://github.com/user-attachments/assets/f5fc9204-d27d-4274-bef8-537eb155f4ac)

### Deskripsi Program : 
Program ini dibuat untuk meminta pengguna memasukkan lima angka bulat. Setelah input diterima, program akan menghitung total dari kelima angka tersebut dan menampilkan hasilnya di layar.

### Algoritma Program
1. Deklarasi Variabel:
   - Buat lima variabel a, b, c, d, dan e untuk menampung angka bulat yang akan dijumlahkan.
   - Siapkan variabel hasil untuk menyimpan hasil penjumlahan.
2. Mengambil Input:
   - Gunakan `fmt.Scanln` untuk menerima input dari pengguna. Program akan menunggu hingga pengguna memasukkan lima angka bulat yang dipisahkan oleh spasi.
3. Proses Penjumlahan:
   - Hitung jumlah dari kelima angka tersebut dengan menjumlahkan a, b, c, d, dan e, lalu simpan hasilnya dalam variabel hasil.
4. Menampilkan Hasil:
   - Cetak hasil penjumlahan dengan format yang sesuai di layar.

### Cara Kerja Program
1. Program mulai dijalankan dan menunggu input dari pengguna.
2. Pengguna diminta untuk menginputkan lima angka bulat
3. Program akan menjumlahkan angka-angka yang telah diinputkan 
4. Setelah perhitungan selesai, program akan menampilkan hasilnya 

### 4.Diberikan sebuah nilai akhir mata kuliah (NAM) [0..100] dan standar penilaian nilai mata kuliah (NMK) sebagai berikut:

![Screenshot 2024-10-05 185203](https://github.com/user-attachments/assets/f9c9959e-bcde-4b09-b3a5-5feab7302db3)
![Screenshot 2024-10-05 185442](https://github.com/user-attachments/assets/e7e4becb-ef34-4315-a1e3-9ec4fa7b5057)

**Jawablah pertanyaan-pertanyaan berikut:**

a. Jika nam diberikan adalah 80.1, apa keluaran dari program tersebut? Apakah eksekusi program tersebut sesuai spesifikasi soal?

b. Apa saja kesalahan dari program tersebut? Mengapa demikian? Jelaskan alur program seharusnya!

C. Perbaiki program tersebut! Ujilah dengan masukan: 93.5; 70.6; dan 49.5. Seharusnya keluaran yang diperoleh adalah 'A', 'B', dan 'D'.

**Jawaban:**


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
3. Di dalam loop:
   - Tambahkan nama bunga yang dimasukkan ke dalam variabel pita.
   - Jika pengguna mengetik "SELESAI", akhiri loop dan kurangi penghitung jumlah bunga karena input ini tidak dianggap sebagai bunga.
5. Setelah keluar dari loop, tampilkan daftar lengkap bunga yang dimasukkan serta jumlah total bunga yang dihitung.

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
Program ini membantu Pak Andi untuk memeriksa keseimbangan beban di sepeda motornya berdasarkan berat dua kantong yang dia bawa. Program akan terus menerima input berat kantong kiri dan kanan hingga salah satu kantong memiliki berat negatif atau total berat kedua kantong melebihi 150 kg. Program juga akan mengevaluasi apakah sepeda motor akan oleng berdasarkan perbedaan berat antara kedua kantong.

### Algoritma Program
1. Inisialisasi variabel untuk menyimpan berat kantong kiri dan kanan.
2. Looping: Program berjalan terus selama kondisi penghentian belum terpenuhi.
   - Minta input berat dari pengguna.
   - Jika berat salah satu kantong negatif, program berhenti.
   - Jika total berat dari kedua kantong lebih dari 150 kg, program berhenti.
   - Hitung selisih berat antara kedua kantong.
   - Jika selisih berat >= 9 kg, tampilkan bahwa sepeda motor oleng; jika tidak, tampilkan bahwa sepeda motor seimbang.
3. Program berhenti jika salah satu kondisi penghentian tercapai.

### Cara Kerja Program
1. Input Berat Kantong: Program meminta pengguna memasukkan berat untuk kantong kiri dan kanan.
2. Penghentian Program: Program berhenti jika salah satu berat kantong negatif atau jika total berat kedua kantong lebih dari 150 kg.
3. Perhitungan Selisih Berat: Program menghitung perbedaan berat antara kantong kiri dan kanan. Jika selisihnya 9 kg atau lebih, program akan menyatakan bahwa sepeda motor akan oleng.
4. Looping: Program akan terus meminta input sampai salah satu kondisi untuk menghentikan program terpenuhi.

### 3. Diberikan sebuah persamaan sebagai berikut ini.
![Screenshot 2024-10-05 114327](https://github.com/user-attachments/assets/22a7ac2a-36ca-4eaa-a185-708f80bc7bc7)

**Buatlah sebuah program yang menerima input sebuah bilangan sebagai K, kemudian menghitung dan menampilkan nilai f(k) sesuai persamaan di atas.Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah Input/read):**
![Screenshot 2024-10-05 134155](https://github.com/user-attachments/assets/837edc1f-f20e-4222-add9-133b56311548)


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
Program ini digunakan untuk menghitung nilai dari fungsi `f(k)`  berdasarkan nilai `( k )` yang dimasukkan oleh pengguna. Fungsi `f(k)` dihitung menggunakan rumus matematika tertentu, dan hasilnya ditampilkan dengan presisi hingga 10 angka di belakang koma.

### Algoritma Program
1. Program meminta pengguna memasukkan nilai `k`.
2. Fungsi `f(k)` melakukan langkah berikut:
   - Hitung pembilang  `(4k + 2)^2`.
   - Hitung penyebut  `(4k + 1)(4k + 3)`.
   - Kembalikan hasil pembagian pembilang dengan penyebut.
3. Tampilkan hasil perhitungan fungsi `f(k)` dengan format presisi 10 angka desimal.
### Cara Kerja Program
1. Meminta Input `k` : Program akan meminta pengguna untuk memasukkan nilai `k`, yang bisa berupa bilangan desimal.
2. Perhitungan Fungsi `f(k)` : Program menghitung nilai fungsi `f(k)` dengan rumus
   
    ![Screenshot 2024-10-05 114327](https://github.com/user-attachments/assets/ccf8bfc7-3591-4abe-98b2-3627d5208190)


3. Menampilkan Hasil : Setelah perhitungan selesai, program akan menampilkan hasil fungsi dengan presisi 10 angka desimal.

### 4. PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BlayaPos untuk menghitung blaya pengiriman tersebut dengan ketentuan sebagai berikut!
**Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg.**

**Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):**
![Screenshot 2024-10-05 134805](https://github.com/user-attachments/assets/73515b8c-05ac-4ca4-859d-8dab50bf3497)

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
Program ini menghitung biaya pengiriman parsel berdasarkan berat yang dimasukkan oleh pengguna dalam satuan gram. Biaya pengiriman dihitung dengan tarif tertentu per kilogram, serta biaya tambahan untuk berat yang kurang dari satu kilogram. Program juga menampilkan detail berat parsel dan rincian biaya pengiriman.

### Algoritma Program
1. Program meminta pengguna memasukkan berat parsel dalam satuan gram.
2. Berat diubah menjadi kilogram dan sisa gram.
3. Biaya dasar dihitung dari berat kilogram dengan tarif Rp 10.000,- per kilogram.
4. Biaya tambahan dihitung untuk sisa gram:
   - Jika total berat lebih dari 10 kg, tidak ada biaya tambahan.
   - Jika sisa gram ≤ 500, dikenakan Rp 5,- per gram.
   - Jika sisa gram > 500, dikenakan Rp 15,- per gram.
5. Biaya total adalah hasil penjumlahan biaya dasar dan biaya tambahan.
6. Tampilkan rincian berat, biaya, dan total biaya pengiriman.

### Cara Kerja Program
1. Memasukkan Berat: Pengguna memasukkan berat parsel dalam satuan gram.
2. Menghitung Biaya: Program mengubah berat dari gram ke kilogram, kemudian menghitung biaya pengiriman dasar sebesar Rp 10.000,- per kilogram. Jika ada sisa berat di bawah 1 kg, program akan menghitung biaya tambahan berdasarkan ketentuan:
   - Jika total berat lebih dari 10 kg, sisa berat tidak dikenakan biaya.
   - Jika sisa berat 500 gram atau kurang, dikenakan Rp 5,- per gram.
   - Jika sisa berat lebih dari 500 gram, dikenakan Rp 15,- per gram.
3. Hasil Output: Program menampilkan rincian berat dalam kilogram dan gram, rincian biaya dasar dan biaya tambahan, serta total biaya pengiriman.

### 5. Buatlah program yang menerima input sebuah bilangan bulat b dan b > 1 Program harus dapat mencari dan menampilkan semua faktor dari bilangan tersebut!
**Lanjutkan program sebelumnya. Setelah menerima masukan sebuah bilangan bulat b > 0. Program tersebut mencari dan menampilkan semua faktor bilangan tersebut. Kemudian, program menentukan apakah b merupakan bilangan prima.**
**Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):**

![Screenshot 2024-10-05 135730](https://github.com/user-attachments/assets/ba0f6794-f7a6-4dc6-9ef7-d793c5ea4d37)


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
Program ini berfungsi untuk menemukan semua faktor dari bilangan yang dimasukkan oleh pengguna dan memeriksa apakah bilangan tersebut merupakan bilangan prima. Program akan menampilkan daftar faktor serta informasi apakah bilangan tersebut adalah prima.

### Algoritma Program
1. Meminta pengguna untuk memasukkan bilangan bulat.
2. Fungsi `cariFaktor(b)`:
     - Menginisialisasi daftar kosong untuk menyimpan faktor.
     - Menggunakan loop dari 1 hingga b untuk memeriksa setiap angka:
        - Jika b dapat dibagi oleh i (tanpa sisa), tambahkan i ke dalam daftar faktor.
     - Mengembalikan daftar faktor yang ditemukan.
3. Fungsi `cekBilPrima(b)`:
    - Jika b kurang dari atau sama dengan 1, kembalikan false (bukan bilangan prima).
    - Menggunakan loop dari 2 hingga akar kuadrat b untuk memeriksa apakah ada angka yang dapat membagi b:
       - Jika ditemukan pembagi, kembalikan false (bukan prima).
    - Jika tidak ada pembagi yang ditemukan, kembalikan true (bilangan prima).
4. Menampilkan daftar faktor dan status primalitas dari bilangan tersebut.

### Cara Kerja Program
1. Input Bilangan: Pengguna diminta untuk memasukkan sebuah bilangan bulat.
2. Mencari Faktor: Program menghitung semua faktor dari bilangan tersebut dengan memeriksa setiap angka dari 1 hingga bilangan itu sendiri dan menambahkan angka yang dapat membagi bilangan tanpa sisa ke dalam daftar faktor.
3. Memeriksa Bilangan Prima: Program mengecek apakah bilangan yang dimasukkan adalah bilangan prima dengan mencari tahu apakah ada angka lain yang dapat membagi bilangan tersebut, selain dari 1 dan bilangan itu sendiri.
4. Output Hasil: Program menampilkan daftar faktor dan status apakah bilangan tersebut merupakan bilangan prima atau tidak.

## Kesimpulan 
1. Struktur Kontrol: Bahasa Go memiliki kontrol alur yang serupa dengan C, namun tidak memiliki perulangan `do` atau `while`, hanya menggunakan `for`, dan `switch` lebih fleksibel dalam mengevaluasi kondisi.
2. Inisialisasi dan Penanganan Kesalahan: Dalam Go, baik `if` maupun `switch` dapat digunakan untuk menginisialisasi variabel lokal. Selain itu, keharusan menggunakan kurung kurawal dalam `if` membantu dalam menangani kesalahan dengan menghilangkan kebutuhan untuk kondisi `else` yang tidak diperlukan.
3. Pengulangan dan Tipe Dinamis:  Pengulangan `for` memiliki tiga bentuk dan memanfaatkan klausa `range` untuk memudahkan iterasi pada struktur data, sementara `switch` dapat digunakan untuk menangani tipe dinamis dari variabel interface.

## Daftar Pustaka
[1] Effective Go. (n.d.). Retrieved from Golang Indonesia: https://golang-id.org/doc/effective_go.html 










 
