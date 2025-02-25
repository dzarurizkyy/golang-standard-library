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

		/*
			- Argumen pertama berisi lokasi file biner sementara setelah kompilasi 
			- Contoh: /var/folders/6y/mm1qmxbd64v3dq75xw_63q5w0000gn/T/go-build164785341/b001/exe/main
		*/
	}

	hostname, err := os.Hostname()
	if err != nil {
		fmt.Println("Error: ", err.Error())
	} else {
		fmt.Println("Hostname:", hostname)
	}
}
