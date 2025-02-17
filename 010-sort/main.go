package main

import (
	"fmt"
	"sort"
)

/*
	Package Sort
	- Berisikan utilitas untuk proses pengurutan
	- Agar data bisa diurutkan, harus mengimplementasikan kontrak di interface sort.Interface
*/

type User struct {
	Name string
	Age  int
}

type UserSlice []User

func(s UserSlice) Len() int {
	return len(s)
}

func(s UserSlice) Less(i, j int) bool {
	return s[i].Age < s[j].Age
}

func(s UserSlice) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func main() {
	users := []User{
		{"Dzaru", 21},
		{"Rizky", 20},
		{"Fathan", 19},
		{"Fortuna", 18},
	}

	sort.Sort(UserSlice(users))
	fmt.Println(users)
}