package main

import (
	"container/list"
	"fmt"
)

/*
	Package Container/List
	Implementasi struktur data double linked list
*/

func main() {
	// Membuat sebuah linked list baru
	data := list.New()

	// Menambahkan elemen ke bagian belakang linked list
	data.PushBack("Dzaru")
	data.PushBack("Rizky")
	data.PushBack("Fathan")
	data.PushBack("Fortuna")

	// Mengembalikan elemen pertama dalam linked list
	head := data.Front()
	fmt.Println(head)

	for e := data.Front(); e != nil; e = e.Next() {
		// e.Next() digunakan untuk mengakses elemen berikutnya dalam linked list
		fmt.Println(e.Value)
	}
}
