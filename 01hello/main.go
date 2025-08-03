package main

import "fmt"

// | Function        | Description                                           | Adds newline? | Formatting? |
// | --------------- | ----------------------------------------------------- | ------------- | ----------- |
// | `fmt.Print()`   | Prints values as-is                                   | ❌ No          | ❌ No        |
// | `fmt.Println()` | Prints values with space between, ends with newline   | ✅ Yes         | ❌ No        |
// | `fmt.Printf()`  | Formatted print (like `printf` in C)                  | ❌ No          | ✅ Yes       |
// | `fmt.Sprint()`  | Returns formatted string (no printing)                | ❌ No          | ❌ No        |
// | `fmt.Sprintf()` | Returns formatted string (like `printf`, no printing) | ❌ No          | ✅ Yes       |
// sprint and sprintf are generally used for concatenating strings or creating formatted strings without printing them directly
func main() {
	fmt.Println("Hello from Go, Soham!")

	name := "Soham"
	age := 22

	fmt.Print("Hello")           // → Hello
	fmt.Println("Soham", age)    // → Soham 21\n
	fmt.Printf("Age: %d\n", age) // → Age: 21\n

	msg := fmt.Sprint("My age is ", age)
	fmt.Println(msg) // Output: My age is 21

	msg2 := fmt.Sprintf("My name is %s and I am %d years old", name, age)
	fmt.Println(msg2)
	// Output: My name is Soham and I am 21 years old

}
