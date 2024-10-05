### III UNGUIDED ###

### 1. Siswa kelas IPA di salah satu sekolah menengah atas di Indonesia sedang mengadakan praktikum kimia. Di setiap percobaan akan menggunakan 4 tabung reaksi, yang mana susunan warna cairan di setiap tabung akan menentukan hasil percobaan. Siswa diminta untuk mencatat hasil percobaan tersebut. Percobaan dikatakan berhasil apabila susunan warna zat cair pada gelas 1 hingga gelas 4 secara berturut-turut adalah ‘merah’, ‘kuning’, ‘hijau’, dan ‘ungu’ selama 5 kali percobaan berulang.<br/> Buatlah sebuah program yang menerima input berupa warna dari ke 4 gelas reaksi sebanyak 5 kali percobaan. Kemudian program akan menampilkan true apabila urutan warna sesuai dengan informasi yang diberikan pada paragraf sebelumnya, dan false untuk urutan warna lainnya.<br/>

``` 
package main
import "fmt"

func main() {
	var (
		satu, dua, string // Deklarasi tiga variabel string
		tempt string // Variabel sementara untuk membantu pertukaran nilai
	)
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&satu)   // Menerima input string pertama dan menyimpannya di variabel 'satu'
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&dua)  // Menerima input string kedua dan menyimpannya di variabel 'dua'
	fmt.Print("Masukan input string: ")
	fmt.Scanln(&tiga) // Menerima input string ketiga dan menyimpannya di variabel 'tiga'
	fmt.Println("Output awal = " + satu + " " + dua + " " + tiga) // Menampilkan output awal sebelum pertukaran
	// Proses pertukaran nilai antar variabel
	tempt = satu
	satu = dua
	dua = tiga
	tiga = temp
	fmt.Println ("Output akhir = " + satu + " " + dua + " " + tiga) // Menampilkan output setelah pertukaran
}
```
### Output: 

Kode di atas adalah
