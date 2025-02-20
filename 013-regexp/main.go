package main

import (
	"fmt"
	"regexp"
)

/*
	Package Regexp
	Utilitas untuk melakukan pencarian regular expression
*/

func main() {
	// regexp.MustCompile(string): membuat regexp
	regex := regexp.MustCompile(`d[a-z]*u`)

	// regexp.MatchString(string)bool : mengecek apakah regexp match dengan string
	fmt.Println("Match:", regex.MatchString("dzaru"))
	fmt.Println("Match:", regex.MatchString("dhalia"))
	fmt.Println("Match:", regex.MatchString("fortuna"))

	// regexp.FindAllString(string, max): mencari string yang match dengan maximum jumlah hasil
	fmt.Println("Match:", regex.FindAllString("dzaru dhalia fortuna danu darmu daiju daru", 10))
}
