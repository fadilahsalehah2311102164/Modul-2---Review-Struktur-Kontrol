package main

import (
	"fmt"
	"math"
)

func main() {
	for {
		
		var kantong1_153,kantong2_153 float64
		fmt.Print("Masukkan berat belanjaan di kedua kantong: ")
		fmt.Scan(&kantong1_153,&kantong2_153)
		
		if kantong1_153 < 0 || kantong2_153 < 0 {
			fmt.Println("Proses selesai.")
			break
		}
		
		totalBerat := kantong1_153 + kantong2_153
		
		if totalBerat > 150 {
			fmt.Println("Proses selesai.")
			break
		}
		
		selisih2 := math.Abs(kantong1_153 - kantong2_153)
		
		olengco := selisih2 >= 9
		fmt.Printf("Sepeda motor Pak Andi akan oleng: %t\n", olengco)
	}
}
