package main

import (
	"fmt"
	"strconv"
)

/*
	Package Strconv
	Mengubah tipe data nilai dengan tipe data yang berbeda

	- parseBool(string)        : mengubah string ke bool
	- parseFloat(string)       : mengubah string ke float
	- parseInt(string)         : mengubah string ke int64
	- formatBool(bool)         : mengubah bool ke string
	- formatFloat(float, ...)  : mengubah float64 ke string
	- formatInt(int, ...)      : mengubah int64 ke string
*/

func main() {
	if result, err := strconv.ParseBool("true"); err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(result)
	}

	if result, err := strconv.Atoi("1000"); err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(result)
	}

	result := strconv.Itoa(1000);
	fmt.Println(result)

	result = strconv.FormatInt(999, 2)
	fmt.Println(result)
}
