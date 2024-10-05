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
