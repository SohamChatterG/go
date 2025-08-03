package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Hello from Go, Soham!")
	fmt.Println("Please enter some rating between 1 to 10")

	scanner := bufio.NewReader(os.Stdin)
	rating, _ := scanner.ReadString('\n')
	rating = strings.TrimSpace(rating)
	fmt.Printf("type of rating is : %T\n", rating)

	numRating, err := strconv.ParseFloat(rating, 64)
	if err != nil {
		fmt.Println("Error while converting rating to float64", err)
		return
	} else {
		fmt.Println("Your rating is + 1 is : ", numRating+1)
	}
	// reader := bufio.NewReader(os.Stdin)

	// 	// comma ok || err ok
	// 	input, _ := reader.ReadString('\n')

	// strings
	text := "   Hello, Soham!\n"
	trimmed := strings.TrimSpace(text)
	fmt.Println("Trimmed:", trimmed)

	if strings.Contains(trimmed, "Soham") {
		fmt.Println("Yes, your name is in the string!")
	}

	words := strings.Split(trimmed, ",")
	fmt.Println("Split result:", words)

	// strconv
	i, err := strconv.Atoi("42")
	if err == nil {
		fmt.Println("Integer:", i)
	}

	// int to string
	s := strconv.Itoa(123)
	fmt.Println("String:", s)

	// string to float64
	f, _ := strconv.ParseFloat("3.14", 64)
	fmt.Println("Float:", f)

}

/*
| Function                                     | Description                     | Example            |
| -------------------------------------------- | ------------------------------- | ------------------ |
| `strconv.Atoi(s)`                            | Converts string to `int`        | `"123"` → `123`    |
| `strconv.Itoa(i)`                            | Converts `int` to string        | `123` → `"123"`    |
| `strconv.ParseFloat(s, bitSize)`             | Converts string to `float32/64` | `"3.14"` → `3.14`  |
| `strconv.FormatFloat(f, fmt, prec, bitSize)` | `float` → `string`              | `3.14` → `"3.140"` |
| `strconv.ParseBool(s)`                       | `"true"/"false"` → `bool`       | `"true"` → `true`  |

*/
