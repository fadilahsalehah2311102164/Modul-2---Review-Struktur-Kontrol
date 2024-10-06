package main

import (
	"bufio"   
	"fmt"     
	"os"      
	"strconv" 
	"strings" 

func main() {
	reader := bufio.NewReader(os.Stdin) 

	for {
		fmt.Print("Masukan berat belanjaan di kedua kantong: ")
		input, _ := reader.ReadString('\n') 
		input = strings.TrimSpace(input)    

		
		berat := strings.Split(input, " ")

		if len(berat) != 2 {
			fmt.Println("Input tidak valid. Masukkan dua angka.") 
			continue                                              

		
		berat1, _ := strconv.ParseFloat(berat[0], 64)
		berat2, _ := strconv.ParseFloat(berat[1], 64)

		if berat1 < 0 || berat2 < 0 {
			fmt.Println("Sepeda motor pak Andi akan oleng: false") 
			break                                                 
		}

		
		if berat1+berat2 > 150 {
			fmt.Println("Sepeda motor pak Andi akan oleng: false") 
			break                                                
		}

		selisih := berat1 - berat2
		if selisih < 0 {
			selisih = -selisih 
		}

		if selisih >= 9 {
			fmt.Println("Sepeda motor pak Andi akan oleng: true")
		} else {
			fmt.Println("Sepeda motor pak Andi akan oleng: false")
		}
	}

	
	fmt.Println("Proses selesai.")
}
}