package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {
	fmt.Println("Welcome to functions in GoLang")
	greeter2()

	result := adder(3, 5)
	fmt.Println("Resut is : ", result)

	proRes := proAdder(2, 4, 5, 6, 6)
	fmt.Println("preRes: ", proRes)

	soham := User{"soham", "soham@54gmail.com", true, 54}
	soham.GetStatus()
	soham.chanegeMail() // Note : this does not change the mail of the actual soham variable // something like pass by variable
	fmt.Println("Printing mail from the main function :", soham.Email)
}
func adder(a int, b int) int {
	return a + b
}
func greeter2() {
	fmt.Println("Namastey from the Gooooooo")
}

func proAdder(values ...int) int {
	var sum int
	for _, ele := range values {
		sum += ele
	}
	return sum
}

// we pass a type in the function for the function to be a method
func (u User) GetStatus() {
	fmt.Println("Is User Active : ", u.Status)
}

func (u User) chanegeMail() { // passed by variable
	u.Email = "soham299@gmail.com"
	fmt.Println("new email: ", u.Email)
}
