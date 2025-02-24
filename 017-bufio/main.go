package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

/*
	Package Bufio
	Digunakan untuk membuat data IO seperti Reader dan Writer
*/

func main() {
	// Reader
	input := strings.NewReader("This is long string\nwith new line\n")
	reader := bufio.NewReader(input)

	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		fmt.Println(string(line))
	}

	// Writer
	writer := bufio.NewWriter(os.Stdout)
	writer.WriteString("Hello world\n")
	writer.WriteString("Have a good study\n")
	writer.Flush()
}
