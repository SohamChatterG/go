// package main: Every executable Go program starts with this package.
package main

import (
	"fmt"
	"sort"
)

// The main function is the entry point of the program.
func main() {
	fmt.Println("--- Go Slices ---")

	// 1. Declaration and Initialization
	// Slices are dynamic and do not have a fixed size.
	// A slice literal `[]string` does not specify a length.
	var fruitList = []string{"apple", "tomato", "peach"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)
	fmt.Printf("Initial slice: %v, Length: %d, Capacity: %d\n", fruitList, len(fruitList), cap(fruitList))

	// 2. The `append` Function
	// `append` adds elements to a slice. If the underlying array is full,
	// a new, larger array is created and the elements are copied over.
	fruitList = append(fruitList, "guava", "banana")
	fmt.Printf("After append: %v, Length: %d, Capacity: %d\n", fruitList, len(fruitList), cap(fruitList))

	// 3. Slicing with `[low:high]`
	// This creates a new slice that is a "view" into a portion of the original.
	// It includes `low` but excludes `high`. The new slice shares the same underlying array.
	subSlice := fruitList[1:4]
	fmt.Printf("Sub-slice from index 1 to 4: %v\n", subSlice)
	fmt.Printf("The sub-slice has length: %d and capacity: %d\n", len(subSlice), cap(subSlice))

	// 4. Using `make` to create a slice
	// `make` allows you to pre-allocate memory with a specific length and capacity.
	// This is a performance optimization to avoid reallocations when appending.
	// `make([]int, 5, 10)` creates a slice of length 5 and capacity 10.
	highScores := make([]int, 5, 10)
	highScores[0] = 234
	highScores[1] = 515
	highScores[2] = 121
	highScores[3] = 789
	highScores[4] = 148
	fmt.Printf("\nHigh scores slice: %v, Length: %d, Capacity: %d\n", highScores, len(highScores), cap(highScores))

	// Appending new values. The capacity is sufficient, so no new allocation occurs.
	highScores = append(highScores, 343)
	fmt.Printf("After appending one value: %v, Length: %d, Capacity: %d\n", highScores, len(highScores), cap(highScores))

	// 5. Sorting Slices
	// The `sort` package provides functions for sorting slices of basic types.
	sort.Ints(highScores)
	fmt.Printf("Sorted high scores: %v\n", highScores)

	// 6. Removing an Element
	// A common idiom for removing an element at a given index.
	// It appends the slice up to the index with the slice from the index+1 to the end.
	courses := []string{"reactjs", "javascript", "swift", "python"}
	fmt.Println("\nCourses before removal:", courses)

	var indexToRemove = 2 // We want to remove "swift"
	courses = append(courses[:indexToRemove], courses[indexToRemove+1:]...)
	fmt.Println("Courses after removal:", courses)

	// 7. Slices are Reference Types (to the underlying array)
	// When you pass a slice to a function, you are passing a copy of the slice header,
	// which contains a pointer to the underlying array.
	// Changes made inside the function will be reflected in the original slice.
	fmt.Println("\nOriginal slice before function call:", courses)
	modifySlice(courses)
	fmt.Println("Original slice after function call: ", courses) // The original slice is now modified.
}

// modifySlice receives a copy of the slice header.
func modifySlice(sl []string) {
	// The `sl` variable points to the same underlying array as the original.
	sl[0] = "golang"
	fmt.Println("Inside modifySlice, the local copy is:", sl)
}
