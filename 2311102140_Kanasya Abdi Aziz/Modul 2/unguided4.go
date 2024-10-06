package main

import (
	"fmt"
)

func main() {
	// Contoh data parsel
	parcels := []map[string]interface{}{
		{"berat": 8500, "detailBerat": []int{8, 500}},
		{"berat": 9250, "detailBerat": []int{9, 250}},
		{"berat": 11750, "detailBerat": []int{11, 750}},
	}

	// Hitung biaya untuk setiap parsel
	for i, parcel := range parcels {
		totalBiaya := hitungBiayaParcel(parcel["berat"].(int))
		parcels[i]["totalBiaya"] = totalBiaya
		fmt.Printf("Contoh #%d\n", i+1)
		fmt.Printf("Berat parsel (gram): %d\n", parcel["berat"])
		fmt.Printf("Detail berat: %d kg + %d gr\n", parcel["detailBerat"].([]int)[0], parcel["detailBerat"].([]int)[1])
		fmt.Printf("Total biaya: Rp. %d\n", totalBiaya)
		fmt.Println("-----")
	}
}

// Fungsi untuk menghitung biaya parsel berdasarkan berat
func hitungBiayaParcel(berat int) int {
	biayaPerKg := 10000
	biayaTambahanPerGram := 5
	beratKg := berat / 1000
	sisaGram := berat % 1000

	biaya := (beratKg * biayaPerKg) + (sisaGram * biayaTambahanPerGram)
	return biaya
}
