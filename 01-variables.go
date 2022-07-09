package main

// This lets us print a few things to the console
import "fmt"

func main() {
	fmt.Println("Hello my friends!")

	// Declaring and assigning the variable
	var favoriteBook = "Harry Potter"
	fmt.Println(favoriteBook)

	// Reassigning
	favoriteBook = "Power of Habits"
	fmt.Println(favoriteBook)

	// Go is a strongly typed language
	// This is not allowed:		
	// favoriteBook = 12

	// Go is able to infer the type without us
	// having to specify it. However, we can still
	// explicitly defining the type
	var otherFavoriteBook string = "Bad Blood"
	fmt.Println(otherFavoriteBook)

	// You have to specify the type if you're
	// not specifying a value to it
	// If you don't assign it to anything 
	// it is given the default value
	var thirdFavoriteBook string
	var myAge int
	var amICool bool

	thirdFavoriteBook = "Diary of Wimpy Kid"

	fmt.Println(thirdFavoriteBook, myAge, amICool)

	// Compound creation
	// Creating multiple variables
	// var favoriteNumber, favoriteChocolate, favoriteTeam = 213, "KitKat", "Knicks" 

	// Block creation
	var (
		favoriteNumber = 27
		favoriteChocolate = "KitKat"
		favoriteTeam = "Knicks"
	)

	fmt.Println(favoriteNumber, favoriteChocolate, favoriteTeam)

	// Shortcut for declaring/assigning a variable:
	favoriteAnimal := "Tiger"
	fmt.Println(favoriteAnimal)

	pet1, pet2, pet3 := "cat", "dog", "rat"
	fmt.Println(pet1, pet2, pet3)

	// Constants
	const myName = "Cesare"
	// Not possible:
	// myName = "Sarah"
}