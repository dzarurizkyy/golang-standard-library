package main

import (
	"fmt"
	"strconv"
)

/*
	Package Strconv
	Mengubah tipe data nilai dengan tipe data yang berbeda
*/

func main() {
	/*
		strconv.parseBool(string)
		Mengubah string ke bool
	*/
	if result, err := strconv.ParseBool("true"); err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(result)
	}

	/*
		strconv.Atoi(string)
		Mengubah string ke int
	*/
	if result, err := strconv.Atoi("1000"); err != nil {
		fmt.Println("Error:", err.Error())
	} else {
		fmt.Println(result)
	}

	/*
		strconv.Itoa(int)
		mengubah int ke string
	*/
	result := strconv.Itoa(1000)
	fmt.Println(result)

	/* 
		strconv.FormatInt(int64, base)
		Mengubah int ke string dengan basis tertentu
	*/
	result = strconv.FormatInt(999, 2)
	fmt.Println(result)
}
