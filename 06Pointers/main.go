package main

import "fmt"

func main() {
	fmt.Println("Hello from Go, Soham!")
	var two int = 2
	var ptr *int = &two

	fmt.Println("Value of pointer is:", ptr)    // memory address
	fmt.Println("Value it points to is:", *ptr) // actual value: 2

	*ptr = *ptr * 3

	fmt.Print("new value is:", two) // 4

}
