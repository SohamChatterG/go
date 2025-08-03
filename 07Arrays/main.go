// package main: Every executable Go program starts with this package.
package main

import "fmt"

// The main function is the entry point of the program.
func main() {
	fmt.Println("--- Go Arrays ---")

	// 1. Declaration and Initialization
	// Arrays have a fixed size that is part of their type.
	// `var scores [5]int` declares an array of 5 integers, initialized to 0.
	var scores [5]int
	fmt.Printf("Initial array: %v\n", scores)

	// Accessing and assigning values using the zero-based index.
	scores[0] = 100
	scores[1] = 95
	scores[2] = 80
	scores[3] = 70
	scores[4] = 90
	fmt.Printf("After assignment: %v\n", scores)

	// You can also initialize an array with a literal.
	// The compiler infers the length from the number of elements.
	fruits := [...]string{"apple", "orange", "banana"}
	fmt.Printf("Fruits array: %v, Length: %d\n", fruits, len(fruits))

	// 2. Fixed Size Limitation
	// You cannot append to an array because its size is fixed.
	// The following line would cause a compile-time error:
	// fruits = append(fruits, "grape")

	// 3. Arrays are Value Types
	// When an array is assigned to another variable, a full copy is made.
	// Modifying the new array does not affect the original.
	originalArray := [3]int{10, 20, 30}
	copiedArray := originalArray // A new copy of the array is created here.

	fmt.Printf("\nOriginal array before modification: %v\n", originalArray)
	fmt.Printf("Copied array before modification:   %v\n", copiedArray)

	copiedArray[0] = 99 // Modify the copied array.

	fmt.Printf("Original array after modification:  %v\n", originalArray) // Original is unchanged.
	fmt.Printf("Copied array after modification:    %v\n", copiedArray)   // Copied is changed.

	// 4. Arrays in Functions
	// When you pass an array to a function, a copy of the entire array is passed.
	fmt.Println("\nOriginal array before function call:", originalArray)
	modifyArray(originalArray)
	fmt.Println("Original array after function call: ", originalArray) // Original remains unchanged.
}

// modifyArray receives a copy of the array.
func modifyArray(arr [3]int) {
	// Any changes made here are only to the local copy, not the original array.
	arr[0] = 1000
	fmt.Println("Inside modifyArray, the local copy is:", arr)
}
