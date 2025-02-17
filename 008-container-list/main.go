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
	data := list.New()

	data.PushBack("Dzaru")
	data.PushBack("Rizky")
	data.PushBack("Fathan")
	data.PushBack("Fortuna")

	head := data.Front()
	fmt.Println(head)

	for e := data.Front(); e != nil; e = e.Next() {
		fmt.Println(e.Value)
	}
}
