// package main: Every executable Go program must start with this package.
package main

import "fmt"

// The main function is the entry point of the program.
func main() {
	fmt.Println("--- Go Maps ---")

	// 1. Declaration and Initialization
	// A map is a key-value pair data structure.
	// You must use `make` or a map literal to initialize a map.
	// `make(map[KeyType]ValueType)` creates an empty map.
	// The key type must be comparable (e.g., string, int, but not slice or map).
	// `make(map[string]string)`
	languages := make(map[string]string)

	// 2. Adding Elements
	// Use `map[key] = value` syntax to add or update an element.
	languages["Js"] = "JavaScript"
	languages["Rb"] = "Ruby"
	languages["Py"] = "Python"
	languages["Go"] = "GoLang"
	fmt.Println("Initial map:", languages)

	// 3. Accessing Elements
	// Access a value using its key. If the key doesn't exist, the zero value
	// for the value type is returned (e.g., "" for string, 0 for int).
	fmt.Println("Py stands for:", languages["Py"])
	fmt.Println("A non-existent key returns:", languages["Unknown"])

	// The "comma ok" idiom is a safer way to check if a key exists.
	// The second return value is a boolean indicating existence.
	value, ok := languages["Js"]
	fmt.Printf("Is 'Js' in the map? Value: '%s', Exists: %v\n", value, ok)

	value, ok = languages["Unknown"]
	fmt.Printf("Is 'Unknown' in the map? Value: '%s', Exists: %v\n", value, ok)

	// 4. Deleting Elements
	// The `delete` function removes a key-value pair.
	// It does nothing if the key does not exist, so no check is needed.
	delete(languages, "Rb")
	fmt.Println("\nAfter deleting 'Rb':", languages)

	// 5. Iterating Over a Map
	// Use a `for...range` loop to iterate. The order of iteration is NOT guaranteed.
	fmt.Println("\nIterating over the map:")
	for key, value := range languages {
		fmt.Printf("Key: %s, Value: %s\n", key, value)
	}

	// You can also choose to only iterate over keys or values.
	fmt.Println("\nIterating over keys only:")
	for key := range languages {
		fmt.Println("Key:", key)
	}

	// 6. Maps are Reference Types
	// Similar to slices, maps are passed by reference.
	// A function that modifies a map will affect the original map.
	fmt.Println("\nOriginal map before function call:", languages)
	addLanguage(languages, "Java", "Java")
	fmt.Println("Original map after function call:", languages) // The original map is now modified.

	// 7. The `nil` Map
	// The zero value of a map is `nil`. A `nil` map cannot be used to add elements.
	var nilMap map[string]int
	fmt.Printf("\nThe nil map has value: %v, length: %d\n", nilMap, len(nilMap))
	// The following line would cause a runtime panic:
	// nilMap["key"] = 1
}

// addLanguage demonstrates that maps are reference types.
func addLanguage(m map[string]string, key string, value string) {
	// The function receives a copy of the map header, which points to the
	// same underlying data. Any change made here modifies the original map.
	m[key] = value
	fmt.Println("Inside addLanguage, the map is now:", m)
}
