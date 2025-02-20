package main

import (
	"fmt"
	"slices"
)

/*
	Package Slices
	Digunakan untuk memanipulasi data di slice
*/

func main() {
	names := []string{"John", "Paul", "George", "Ringo"}
	values := []int{100, 95, 80, 90}

	fmt.Println("Min:", slices.Min(values))
	fmt.Println("Max:", slices.Max(values))
	fmt.Println("------------------------------")

	fmt.Println("Min:", slices.Min(names))
	fmt.Println("Max:", slices.Max(names))
	fmt.Println("------------------------------")

	fmt.Println("Index:", slices.Index(names, "George"))
	fmt.Println("Index:", slices.Index(names, "Eko"))
	fmt.Println("------------------------------")
	
	fmt.Println("Contain:", slices.Contains(names, "Paul"))
}
