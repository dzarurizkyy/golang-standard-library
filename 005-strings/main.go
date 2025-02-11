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
	fmt.Println(strings.ToUpper("Dzaru Rizky Fathan Fortuna"))
	fmt.Println(strings.Trim("   Dzaru Rizky Fathan Fortuna   ", " "))
	fmt.Println(strings.ToLower("Dzaru Rizky Fathan Fortuna"))
	fmt.Println(strings.ReplaceAll("Dzaru Rizky Fathan Fortuna", "Fathan Fortuna", "F. F"))
	fmt.Println(strings.Split("Dzaru Rizky Fathan Fortuna", " "))
	fmt.Println(strings.Contains("Dzaru Rizky Fathan Fortuna", "Dzaru"))
}
