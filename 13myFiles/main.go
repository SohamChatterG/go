package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to files in Gooooooooo!")

	content := "GooooooooooooooooLang!"

	file, err := os.Create("./Gooooooooooooooo.txt")

	if err != nil {
		panic(err)
	}

	len, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("the len is", len)
	readFile("./Gooooooooooooooo.txt")
	file.Close()

}

func readFile(filename string) {
	databyte, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("text data\n", string(databyte))
}
