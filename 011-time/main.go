package main

import (
	"fmt"
	"time"
)

func main() {
	/*
		Package Time
		Berisikan fungsionalitas untuk management waktu
	*/
	{
		// time.Now() -> Untuk mendapatkan waktu saat ini
		now := time.Now()
		fmt.Println(now.Local())

		// time.Date(...) -> Untuk membuat waktu
		date := time.Date(2003, time.June, 27, 0, 0, 0, 0, time.Local)
		fmt.Println(date.Local())
		fmt.Println(date.UTC())

		// time.Parse(layout, string) -> Untuk memparsing waktu dari string
		formatter := "2006-01-02 15:04:05"
		value := "2020-04-15 15:00:00"

		valueTime, err := time.Parse(formatter, value)
		if err != nil {
			fmt.Println("Error", err.Error())
		} else {
			fmt.Println(valueTime)
		}
	}

	/* 
		Duration 
		- Saat mengggunakan tipe data waktu, kadang membutuhkan data berupa durasi
		- Package time memiliki type Duration, yang sebenarnya adalah alias untuk int64
		- Namun terdapat banyak method yang bisa digunakan untuk memanipulasi Duration
	*/
	{
		duration1 := 100 * time.Second
		duration2 := 10 * time.Millisecond
		duration3 := duration1 - duration2

		fmt.Printf("Duration: %d\n", duration3)
	}
}
