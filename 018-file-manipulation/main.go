package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

/*
	os.OpenFile(name, flag, permission)
	- name berisikan nama file, bisa absolute atau relative / local
	- flag merupakan penanda file, apakah untuk membaca, menulis, dan lain-lain
	- permission merupakan permission yang diperlukan ketika membyat file
*/

/*
	Permission dapat disimulasikan di https://chmod-calculator.com/
*/

// Membuat file baru
func createNewFile(name, mesaage string) error {
	file, err := os.OpenFile(name, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	file.WriteString(mesaage)
	return nil
}

// Membaca file
func readFile(name string) error {
	file, err := os.OpenFile(name, os.O_RDONLY, 0666)
	if err != nil {
		return err
	}

	defer file.Close()
	reader := bufio.NewReader(file)
	var message string

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}
		message += string(line) + "\n"
	}

	fmt.Println(message)
	return nil
}

// Membaca dan menambahkan ke file
func addToFile(name, message string) error {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_APPEND, 0666)
	if err != nil {
		return nil
	}

	defer file.Close()
	file.WriteString(message)
	return nil
}

func main() {
	createNewFile("sample.log", "this is sample log\nthis is sample log\n")
	readFile("sample.log")
	addToFile("sample.log", "this is end line")
}
