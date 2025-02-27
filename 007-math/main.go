package main

import (
	"fmt"
	"math"
)

/*
	Package Math
	Berisikan constant dan fungsi matematika
*/

func main() {
	// Membulatkan float64 ke atas
	fmt.Println("Ceil  :", math.Ceil(1.40))
	// Membulatkan float64 ke bawah
	fmt.Println("Floor :", math.Floor(1.60))
	// Membulatkan float64 ke atas atau ke bawah, sesuai dengan yang paling dekat
	fmt.Println("Round :", math.Round(1.60))
	// Mengembalikan nilai float64 yang paling besar
	fmt.Println("Max   :", math.Max(10, 11))
	// Mengembalikan nilai float64 yang paling kecil
	fmt.Println("Min   :", math.Min(10, 11))
}