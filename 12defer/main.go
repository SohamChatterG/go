package main

import "fmt"

func main() {
	// remember two things about defer
	// whenever we put defer before a statement it is placed just befre the last curly braces of the function
	// they work in last in first out fashion like a stack they are not run immediately they are put into stack
	// and whenever the function is about to end this stack is executed
	defer fmt.Println("world")
	defer fmt.Println("one")
	defer fmt.Println("two")
	fmt.Println("hello")
	myDefer()
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
