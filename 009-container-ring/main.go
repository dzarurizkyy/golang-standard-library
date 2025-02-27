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
	// Membuat sebuah circular list baru
	data := ring.New(5)

	for i := range data.Len() {
		// Mengatur nilai untuk setiap elemen dalam circular list
		data.Value = "Value " + strconv.Itoa(i+1)
		// Berpindah ke elemen berikutnya dalam circular list
		data = data.Next()
	}

	data.Do(func(value any) {
		fmt.Println(value)
	})
}
