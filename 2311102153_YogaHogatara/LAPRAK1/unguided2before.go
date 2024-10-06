package main

import (
	"fmt"
	
)

func main() {
	for { 
		var kantong1_153, kantong2_153 float64 
		fmt.Print("Masukan berat belanjaan di kedua kantong: ") 
		fmt.Scan(&kantong1_153, &kantong2_153) 
		if kantong1_153 >= 9 || kantong2_153 >= 9 { 
		fmt.Println("Proses selesai.") 
		break 
		} 
	}	 
}
