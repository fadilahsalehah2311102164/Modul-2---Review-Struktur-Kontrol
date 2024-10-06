package main

import "fmt"

func main() {
	
	var pita string     
	var banyakBunga int 

	
	for {
		var bunga string
		
		fmt.Print("Bunga (ketik 'SELESAI' untuk berhenti): ")
		fmt.Scanln(&bunga) 

		
		if bunga == "SELESAI" {
			break
		}

		if pita != "" {
			pita += " – "
		}
	
		pita += bunga

		banyakBunga++
	}


	fmt.Println("Pita:", pita)                   
	fmt.Println("Banyaknya bunga:", banyakBunga) 
}