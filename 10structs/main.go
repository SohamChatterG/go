package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"
)

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

type Person struct {
	Name  string
	Email string
}

type Human struct {
	Person // Embedded struct (composition)
	Status bool
	Age    int
}

// Note : Go doesn't have the concept of inheritance in classes
func main() {
	soham := User{"Soham", "soham@gmail.com", true, 55}
	fmt.Printf("%v\n", soham)

	// Struct Embedding (Composition in Go)

	u := Human{
		Person: Person{ // User embeds Person, which means User gets access to Name and Email directly, without inheritance.
			Name:  "Soham",
			Email: "soham@gmail.com",
		},
		Status: true,
		Age:    55,
	}

	fmt.Println("Name:", u.Name) // Access fields directly
	fmt.Println("Email:", u.Email)
	fmt.Println("Status:", u.Status)
	fmt.Println("Age:", u.Age)
	userJson, err := json.MarshalIndent(u, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}
	fmt.Println("\nUser JSON:\n", string(userJson))

	// ----------------------------

	fmt.Println("If Else in GoLang")

	logInCount := 25
	if logInCount < 10 {
		fmt.Println("Regular User")
	} else if logInCount == 10 {
		fmt.Println("10 -> Leo!!! Go Leo!!")
	} else {
		fmt.Println("Whta is it!")
	}

	fmt.Println("Switch Case in Go Lang")

	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Dice Number :", diceNumber)
	switch diceNumber {
	case 1:
		{
			fmt.Println("dice Value is 1")
		}
	case 6:

		fmt.Println("Yes Let's go")

	default:
		fmt.Println("Damn what was that!")
	}

	fmt.Println("Loops in Gooooooooooo!")

	days := []string{"Sunday", "Monday", "TuesDay", "WednesDay", "ThursDay"}

	fmt.Println("days : ", days)

	for d := 0; d < len(days); d++ {
		fmt.Println(days[d])
	}
	for idx, day := range days {
		fmt.Println(idx, " ", day)
	}

	// while loop
	i := 0
	for i < 4 {
		if i == 4 { // we name a labes which is messi here thean we use the keyword goto for jumping there whenever required
			goto messi
		}
		if i == 2 {
			i++
			continue // break
		}
		fmt.Println("i : ", i)
		i++
	}

	// go to statement
messi:
	fmt.Println("Messi is a good player")

}
