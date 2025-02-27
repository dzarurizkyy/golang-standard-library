package main

import (
	"fmt"
	"strings"
)

/*
	Package Strings
	Berisikan function-function untuk memanipulasi tipe data String
*/

func main() {
	// Membuat semua karakter string menjadi upper case
	fmt.Println(strings.ToUpper("Dzaru Rizky Fathan Fortuna"))

	// Memotong cutset di awal dan dan akhir string
	fmt.Println(strings.Trim("   Dzaru Rizky Fathan Fortuna   ", " "))

	// Membuat semua karakter string menjadi lower case
	fmt.Println(strings.ToLower("Dzaru Rizky Fathan Fortuna"))

	// Mengubah semua string dari from ke to
	fmt.Println(strings.ReplaceAll("Dzaru Rizky Fathan Fortuna", "Fathan Fortuna", "F. F"))

	// Memotong string berdasarkan separator
	fmt.Println(strings.Split("Dzaru Rizky Fathan Fortuna", " "))

	// Mengecek apakah string mengandung string lain
	fmt.Println(strings.Contains("Dzaru Rizky Fathan Fortuna", "Dzaru"))
}
