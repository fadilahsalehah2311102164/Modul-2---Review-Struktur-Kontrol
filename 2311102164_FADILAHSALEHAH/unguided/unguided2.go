package main

import (
	"fmt"
	"math"
)

func main() {
	var beratKantong1, beratKantong2 float64
	const batasTotalBerat = 150.0

	for {
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&beratKantong1, &beratKantong2)

		if beratKantong1 < 0 || beratKantong2 < 0 {
			fmt.Println("Proses selesai.")
			break
		}

		totalBerat := beratKantong1 + beratKantong2
		if totalBerat > batasTotalBerat {
			fmt.Println("Proses selesai.")
			break
		}

		selisih := math.Abs(beratKantong1 - beratKantong2)

		if selisih >= 9 {
			fmt.Println("Sepeda motor Pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor Pak Andi akan oleng: false")
		}
	}
}
