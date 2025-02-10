package main

import (
	"errors"
	"fmt"
)

var (
	ValidationError = errors.New("validation error")
	NotFoundError   = errors.New("not found error")
)

func GetByID(id string) error {
	if id == "" {
		return ValidationError
	}

	if id != "dzaru" {
		return NotFoundError
	}

	return nil
}

func main() {
	err := GetByID("")

	/*
		errors.ls
		Untuk mengecek jenis type error nya
	*/

	if err != nil {
		if errors.Is(err, ValidationError) {
			fmt.Println("Error: validation error")
		} else if errors.Is(err, NotFoundError) {
			fmt.Println("Error: not found error")
		} else {
			fmt.Println("Error: unknown error")
		}
	} else {
		fmt.Println("success")
	}
}
