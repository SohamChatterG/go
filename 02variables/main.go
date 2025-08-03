package main

import "fmt"

const LogInToken string = "iqrciinirnini" // Public global variable
// in go by making the letter of the variable name capitalized, we can make it public
// in the same package, we can access it without the package name
// in other packages, we can access it with the package name
// in go, we can use const for constant variables. Constants are always compile-time evaluated.
// You cannot use := with constants.

var x int = 10

func main() {
	var username string = "Soham"
	fmt.Println("username:", username)
	fmt.Printf("Variable is of type: %T \n", username)

	var isLoggedIn bool = true
	fmt.Printf("Variable is of type: %T \n", isLoggedIn)

	var intVal int = 254
	fmt.Println(intVal)
	fmt.Printf("Variable is of type : %T \n", intVal)

	//  Multiple Variable Declarations
	var a, b, c int = 1, 2, 3
	fmt.Println("a:", a, "b:", b, "c:", c)

	// default values and some aliases
	var anotherVariable int
	fmt.Println(anotherVariable) // no garbage value default value for int is always 0

	var anotherString string
	fmt.Println(anotherString) // no garbage value default value for string is always ""

	// implicit typing

	var website = "soham-portfolio.com"
	fmt.Println(website)

	// Short Variable Declaration

	/*	:= is only used inside functions.
		// It is shorthand for declaring and initializing a variable.
		// It infers the type. */
	viewerInLasr24Hours := 3000 // valorous operator
	fmt.Println(viewerInLasr24Hours)

	fmt.Println("LogINToken ", LogInToken)

	// Variable Shadowing
	x := 20        // shadows the global x
	fmt.Println(x) // prints 20
}
