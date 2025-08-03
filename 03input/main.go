package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	var firstName string
	fmt.Print("Enter your first name: ")
	fmt.Scanln(&firstName) // It stops reading at the first space or newline. It can read multiple values if you provide multiple variables:
	fmt.Printf("Hello, %s!\n", firstName)
	fmt.Println("Please enter your name: ")

	var age int
	fmt.Scanf("%d", &age)

	reader := bufio.NewReader(os.Stdin)

	// comma ok || err ok
	input, _ := reader.ReadString('\n')
	/*
		Always Trim the Input
		Input from ReadString('\n') includes the newline character.
		Use strings.TrimSpace() to remove it.
	*/
	fmt.Printf("All the best for your Go journey, %s. I feel like you are made for Go\n", input)
}
