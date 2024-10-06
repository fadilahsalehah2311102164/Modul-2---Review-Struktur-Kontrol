<h2 align="center">LAPORAN PAKTIKUM ALGORITMA DAN PEMROGRAMAN 2</h2>
<br></br>
<h2 align="center">MODUL 2</h2>
<h2 align="center">REVIEW STRUKTUR KONTROL</h2>

<p align="center"><img src=https://github.com/user-attachments/assets/37e9c953-078b-4ef4-97e7-a5d25344e50b alt="Logo" width="300"/><p>

<p align="center">Disusun Oleh : </p>
<p align="center">Irshad Benaya Fardeca / 2311102199</p>
<p align="center">IF-11-05</p>
<br></br>
<p align="center">Dosen Pengampu : </p>
<p align="center">Arif Amrulloh, S.Kom., M.Kom.</p>
<br></br>
<h3 align="center">PROGRAM STUDI S1 TEKNIK INFORMATIKA</h3>
<h3 align="center">FAKULTAS INFORMATIKA</h3>
<h3 align="center">TELKOM UNIVERSITY PURWOKERTO</h3>
<h3 align="center">2024</p>

#### I. DASAR TEORI


<br></br>


#### II. GUIDED
##### 1. soal
##### Source code
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
##### Screenshoot Output
![Screenshot 2024-10-06 183039](https://github.com/user-attachments/assets/adf2fd2b-6f04-46c1-a48e-06b446da0462)
##### Deskripsi Program


##### 2. soal
Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturutan adalah 'merah', 'kuning', 'hijau', dan 'ungu' selama 5 kali percobaan berulang. Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna warna sesuai sesuai d dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk Telkom University ngan informasi yang d urutan warna lainnya.
##### Source code
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
	success := true

	for i := 1; i <= 5; i++ {
		fmt.Printf("Percobaan %d: ", i)

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		colors := strings.Split(input, " ")

		for j := 0; j < 4; j++ {
			if colors[j] != correctOrder[j] {
				success = false
				break
			}
		}

		if !success {
			break
		}
	}

	if success {
		fmt.Println("BERHASIL: true")
	} else {
		fmt.Println("BERHASIL: false")
	}

}
```
##### Screenshoot Output
![Screenshot 2024-10-06 183644](https://github.com/user-attachments/assets/79d6b75b-faf1-4e60-b2f7-c28ea504574c)
##### Deskripsi Program


##### 3. soal
Diberikan sebuah nilai akhir mata kullah (NAM) [0..100] dan standar penilaian nilai mata kuliah (NMK)Program berikut menerima input sebuah bilangan riil yang menyatakan NAM. Program menghitung NMK dan menampilkannya.
##### Source code
```go
package main

import "fmt"

func main() {
	var nam float32
	var nmk string

	fmt.Println("Masukkan nilai: ")
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

	fmt.Printf("Nilai Indeks untuk nilai %.2f adalah %s\n", nam, nmk)
}
```
##### Screenshoot Output
![Screenshot 2024-10-06 191041](https://github.com/user-attachments/assets/6b592e5c-3028-41a3-8502-6b9416c1ce6c)
##### Deskripsi Program


##### 3. soal
##### Source code
```go
package main

import "fmt"

func main() {

	var a, b, c, d, e int
	var hasil int
	fmt.Scanln(&a, &b, &c, &d, &e)

	hasil = a + b + c + d + e
	fmt.Println("Hasil Penjumlahan ", a, b, c, d, e, "adalah = ", hasil)
}

```
##### Screenshoot Output
![Screenshot 2024-10-06 190844](https://github.com/user-attachments/assets/b1c7f4a4-0a40-43d4-8a0f-3e8d310bc7c1)
##### Deskripsi Program


<br></br>
#### III. UNGUIDED

##### 2B. soal
Suatu pita (string) berisi kumpulan nama-nama bunga yang dipisahkan oleh spasi dan '-', contoh pita diilustrasikan seperti berikut ini. Pita: mawar - melati - tulip-teratal - kamboja-anggrek
Buatlah sebuah program yang menerima input sebuah bilangan bulat positif (dan tidak nol) N, kemudian program akan meminta input berupa nama bunga secara berulang sebanyak N kali dan nama tersebut disimpan ke dalam pita. (Petunjuk: gunakan operasi penggabungan string dengan operator "+"). Tampilkan isi pita setelah proses input selesai. Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergarls bawah adalah input/read):

Modifikasi program sebelumnya, proses input akan berhenti apabila user mengetikkan 'SELESAI'. Kemudian tampilkan isi pita beserta banyaknya bunga yang ada di dalam pita Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergarls bawah adalah input/read):
##### Source code
```go
package main

import "fmt"

func main() {
	var jum int
	var bunga, pita string

	for i := 0; i < 10; i++ {
		fmt.Printf("Bunga %d: ", i+1)
		fmt.Scanln(&bunga)
		if bunga == "selesai" {
			break
		}
		jum++
		pita += bunga + "-"
	}

	fmt.Println("Pita :", pita)
	fmt.Println("Bunga :", jum)
}

```
##### Screenshoot Output
![Screenshot 2024-10-06 201002](https://github.com/user-attachments/assets/df8266cd-f57c-4b78-9847-65c107f4716b)
##### Deskripsi Program


##### 3B. soal
Setiap hari Pak Andi membawa banyak barang belanjaan dari pasar dengan mengendarai sepeda motor. Barang belanjaan tersebut dibawa dalam kantong terpal di kiri-kanan motor. Sepeda motor tidak akan oleng jika selisih berat barang di kedua kantong sisi tidak lebih dari 9 kg. Mawar Tulip Bunga: 0 14 15 16 17 Buatlah program Pak Andi yang menerima input dua buah bilangan real positif yang menyatakan berat total masing-masing isi kantong terpal. Program akan terus meminta input bilangan tersebut hingga salah satu kantong terpal berisi 9 kg atau lebih. Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergarls bawah adalah input/read):

Pada modifikasi program tersebut, program akan menampilkan true jika selisih kedua isi kantong lebih dari atau sama dengan 9 kg. Program berhenti memproses apabila total berat isi kedua kantong melebihi 150 kg atau salah satu kantong beratnya negatif. Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergarls bawah adalah input/read):
##### Source code
##### Screenshoot Output
##### Deskripsi Program


##### 4B. soal
Buatlah sebuah program yang menerima input sebuah bilangan sebagai K, kemudian menghitung dan menampilkan nilai f(K) sesuai persamaan di atas. Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergarls bawah adalah input/read):
##### Source code
```go
package main

import "fmt"

func main() {
	var k float64

	fmt.Print("Nilai k = ")
	fmt.Scanln(&k)

	fk := ((4*k + 2) * (4*k + 2)) / ((4*k + 1) * (4*k + 3))

	fmt.Printf("Nilai f(k) = %.10f\n", fk)
}

```
##### Screenshoot Output
![Screenshot 2024-10-06 193055](https://github.com/user-attachments/assets/6fe9c47e-fe00-437e-93b6-37ca95e8c38f)
##### Deskripsi Program


##### 1C. soal
PT POS membutuhkan aplikasi perhitungan biaya kirim berdasarkan berat parsel. Maka, buatlah program BlayaPos untuk menghitung blaya pengiriman tersebut dengan ketentuan sebagai berikut! Dari berat parsel (dalam gram), harus dihitung total berat dalam kg dan sisanya (dalam gram). Biaya jasa pengiriman adalah Rp. 10.000,- per kg. Jika sisa berat tidak kurang dari 500 gram, maka tambahan biaya kirim hanya Rp. 5,- per gram saja. Tetapi jika kurang dari 500 gram, maka tambahan biaya akan dibebankan sebesar Rp. 15,- per gram. Sisa berat (yang kurang dari 1kg) digratiskan biayanya apabila total berat ternyata lebih dari 10kg. Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergaris bawah adalah input/read):
##### Source code
```go
package main

import "fmt"

func main() {
	var berat int
	var kg, sisaKg int
	var biaya, sisaBiaya, total int

	fmt.Print("Berat parsel (gram): ")
	fmt.Scanln(&berat)

	kg = berat / 1000
	sisaKg = berat % 1000

	biaya = kg * 10000

	if sisaKg >= 500 {
		sisaBiaya = sisaKg * 5
	} else {
		sisaBiaya = sisaKg * 15
	}

	if kg > 10 {
		sisaBiaya = 0
	}

	total = biaya + sisaBiaya

	fmt.Printf("\nDetail berat: %d kg + %d gr\n", kg, sisaKg)
	fmt.Printf("Detail biaya: Rp. %d + Rp. %d\n", biaya, sisaBiaya)
	fmt.Printf("Total biaya: Rp. %d\n", total)
}
```
##### Screenshoot Output
![Screenshot 2024-10-06 200522](https://github.com/user-attachments/assets/44683eaf-d96d-44a6-9a19-735d3d26c433)
##### Deskripsi Program


##### 3C. soal
Sebuah bilangan bulat b memiliki faktor bilangan f > O jika f habis membagi b. Contoh: 2 merupakan faktor dari bilangan 6 karena 6 habis dibagi 2. Buatlah program yang menerima input sebuah bilangan bulat b dan b > 1. Program harus dapat mencari dan menampilkan semua faktor dari bilangan tersebut! Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergarls bawah adalah input/read):

Bilangan bulat b> O merupakan bilangan prima p jika dan hanya jika memiliki persis dua faktor bilangan saja, yaitu 1 dan dirinya sendiri. Lanjutkan program sebelumnya. Setelah menerima masukan sebuah bilangan bulat b > 0. Program tersebut mencari dan menampilkan semua faktor bilangan tersebut. Kemudian, program menentukan apakah b merupakan bilangan prima. Perhatikan contoh sesi interaksi program seperti di bawah ini (teks bergarls bawah adalah input/read):
##### Source code
```go
package main

import "fmt"

func main() {

	var bil int
	var jum int

	fmt.Print("Bilangan: ")
	fmt.Scan(&bil)

	if bil <= 1 {
		fmt.Println("Bilangan harus lebih besar dari 1")
		return
	}

	fmt.Printf("Faktor : ")
	for i := 1; i <= bil; i++ {
		if bil%i == 0 {
			fmt.Print(" ", i)
			jum++
		}
	}

	fmt.Println("")
	if jum == 2 {
		fmt.Println("Prima: True")
	} else {
		fmt.Println("Prima: False")
	}

}
```
##### Screenshoot Output
![Screenshot 2024-10-06 202229](https://github.com/user-attachments/assets/ed5d2695-3ba9-4e01-975f-ef26b788ecf7)
##### Deskripsi Program
