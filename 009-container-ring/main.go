package main

import (
	"container/ring"
	"fmt"
	"strconv"
)

/*
	Package Container/Ring
	Implementasi struktur data circular list, dimana diakhir elemen akan kembali ke elemen awal (HEAD)
*/

func main() {
	data := ring.New(5)

	for i := 0; i < data.Len(); i++ {
		data.Value = "Value " + strconv.Itoa(i+1)
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
