package main

import (
	"flag"
	"fmt"
)

/*
	Package Flag
	Berisikan fungsionalitas untuk memparsing command line argument
*/

func main() {
	/*
		flag.type(name, value, usage)
		- name  : nama flag yang akan digunakan dalam command line
		- value : nilai default untuk flag tersebut
		- usage : deskripsi penggunaan untuk flag tersebut yang akan ditampilkan dalam help
	*/

	username := flag.String("username", "root", "database username")
	password := flag.String("password", "root", "database password")
	host 	 := flag.String("host", "localhost", "database host")
	port     := flag.Int("port", 0, "database port")
	help     := flag.Bool("help", false, "display help")

	flag.Parse()

	if *help {
		// Menampilkan bantuan kepada pengguna tentang cara menggunakan flag dalam program
		flag.PrintDefaults()
	} else {
		fmt.Println("Username :", *username)
		fmt.Println("Password :", *password)
		fmt.Println("Host     :", *host)
		fmt.Println("Port     :", *port)
	}
}
