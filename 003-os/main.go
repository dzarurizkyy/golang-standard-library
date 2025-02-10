package main

import (
	"fmt"
	"os"
)

/*
	Package OS
	Berisikan fungsionalitas untuk mengakses fitur sistem operasi secara independen (bisa digunakan di semua sistem operasi)
*/

func main() {
	args := os.Args
	for _, arg := range args {
		fmt.Println("Args:", arg)
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		fmt.Println("Hostname: ", hostname)
	}
}
