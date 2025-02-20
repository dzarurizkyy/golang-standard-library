package main

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
)

/*
	Package Reflect
	Digunakan untuk melihat struktur kode program pada saat aplikasi sedang berjalan
*/

type Person struct {
	Name    string `required:"true"`
	Email   string `required:"true"`
	Address string `required:"true"`
	Age     int    `required:"true" max:"80"`
}

func ReadField(value any) {
	valueType := reflect.TypeOf(value)
	fmt.Printf("Struct (%s): \n", valueType.Name())

	for i := 0; i < valueType.NumField(); i++ {
		valueField := valueType.Field(i)
		if valueField.Tag.Get("max") == "" {
			fmt.Printf(" - Field %d: %s (%s)(required:%s) \n", i+1, valueField.Name, valueField.Type, valueField.Tag.Get("required"))
		} else {
			fmt.Printf(" - Field %d: %s (%s)(required:%s)(max:%s) \n", i+1, valueField.Name, valueField.Type, valueField.Tag.Get("required"), valueField.Tag.Get("max"))
		}
	}
}

func IsValid(value any) (isValid []error) {
	result := true
	valueType := reflect.TypeOf(value)

	for i := 0; i < valueType.NumField(); i++ {
		valueField := valueType.Field(i)
		data := reflect.ValueOf(value).Field(i).Interface()

		if valueField.Tag.Get("required") == "true" {
			result = data != ""
			if result == false {
				isValid = append(isValid, errors.New(valueField.Name + " field cannot be empty"))
			}
		}

		if valueField.Tag.Get("max") != "" {
			maxValue, _ := strconv.Atoi(valueField.Tag.Get("max"))
			if dataValue, ok := data.(int); ok {
				result = dataValue <= maxValue
				if result == false {
					isValid = append(isValid, errors.New(valueField.Name + " field cannot more than " + strconv.Itoa(maxValue)))
				}
			}
		}
	}

	if len(isValid) == 0 {
		isValid = append(isValid, nil)
	}

	return isValid
}

func main() {
	person := Person{
		Name:    "Dzaru Rizky Fathan Fortuna",
		Email:   "dzaurizkybusiness@gmail.com",
		Address: "Surabaya, Indonesia",
		Age:     20,
	}

	ReadField(person)

	fmt.Println("-------------------------------")

	for _, isValid := range IsValid(person) {
		if isValid == nil {
			fmt.Println(" - Success")
		} else {
			fmt.Println(" - Error:", isValid)
		}
	}
}
