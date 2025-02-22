package main

import (
	"fmt"
	"path/filepath"
)

/*
	Package Path
	Digunakan untuk memanipulasi path di URL
*/

/*
	Package Filepath
	Digunakan untuk memanipulasi path di File System
*/

func main() {
	fmt.Println("Directory :", filepath.Dir("hello/world.go"))
	fmt.Println("File      :", filepath.Base("hello/world.go"))
	fmt.Println("Extension :", filepath.Ext("hello/world.go"))
	fmt.Println("Join      :", filepath.Join("hello", "world", "main.go"))
}
