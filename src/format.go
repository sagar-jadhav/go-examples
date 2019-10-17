package main

import "fmt"

type Food struct {
	fruit, vegetable string
}

type BookInfo struct {
	isAvailable bool
	published int
	price float32
	bookTitle string
}

func main() {

	// fmt verbs for general go values
	fruits := Food{"apple", "broccoli"}

	fmt.Printf("%v\n", fruits) // prints struct in default value format
	fmt.Printf("%+v\n", fruits) // prints with struct field names included
	fmt.Printf("%T\n", fruits) // prints the type of the value

	// specific go fmt verb formats
	book := BookInfo{false, 2018, 20.95, "The Overstory"}

	fmt.Printf("This book is available: %t\n", book.isAvailable) // to print booleans
	fmt.Printf("This book was published in: %d\n", book.published) // to print integers
	fmt.Printf("This book costs $%g\n", book.price) // to print float/decimal values without precision
	fmt.Printf("The title of this book is: %s\n", book.bookTitle) // to print string values

	// Go fmt verbs for slices
	books := [4]string{
		"The Overstory",
		"Alias Grace",
		"East of Eden",
		"The Moonstone",
	}

	topBooks := books[0:2]

	fmt.Printf("The top selling books are: %q\n", topBooks) // to print slice values in a safely-escaped string

	// Go fmt verbs with precision formatting
	pi := 3.14159
	eulers := 2.71828

	fmt.Printf("%.2f | %.2f\n", pi, eulers) // print float values in specific precision format
}

