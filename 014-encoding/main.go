package main

import (
	"encoding/base64"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
	Package Encoding
	Berisikan berbagai macam algortima untuk encoding, contoh yang populer adalah base64, csv, dan json
*/

func main() {
	// Base64
	value := "Dzaru Rizky Fathan Fortuna"

	encoded := base64.StdEncoding.EncodeToString([]byte(value))
	fmt.Println("Encoded:", encoded)

	decoded, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Decoded:", string(decoded))
	}

	fmt.Println("------------------------------------")

	// CSV Reader
	csvString := "Dzaru\n" + "Rizky\n" + "Fathan\n" + "Fortuna\n"
	reader := csv.NewReader(strings.NewReader(csvString))

	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		fmt.Println(record)
	}

	fmt.Println("------------------------------------")

	// CSV Writer
	writer := csv.NewWriter(os.Stdout)
	
	_ = writer.Write([]string{"Dzaru", "Rizky", "Fathan", "Fortuna"})
	_ = writer.Write([]string{"Kirigaya", "Kazuto"})

	writer.Flush()
}
