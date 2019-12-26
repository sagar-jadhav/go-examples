---
layout: default
---

# Format Verbs

Example demonstrates the use of fmt verbs in Go. Click [here](https://godoc.org/golang.org/x/exp/errors/fmt) to learn more

```go
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
	
	// prints struct in default value format
	fmt.Printf("%v\n", fruits)
	
	// prints with struct field names included
	fmt.Printf("%+v\n", fruits)

	// prints the type of the value	
	fmt.Printf("%T\n", fruits)

	// specific go fmt verb formats
	book := BookInfo{false, 2018, 20.95, "The Overstory"}

	// to print booleans
	fmt.Printf("This book is available: %t\n", book.isAvailable)
	
	// to print integers
	fmt.Printf("This book was published in: %d\n", book.published)

	// to print float/decimal values without precision	
	fmt.Printf("This book costs $%g\n", book.price)
	
	// to print string values
	fmt.Printf("The title of this book is: %s\n", book.bookTitle)

	// Go fmt verbs for slices
	books := [4]string{
		"The Overstory",
		"Alias Grace",
		"East of Eden",
		"The Moonstone",
	}

	topBooks := books[0:2]

	// to print slice values in a safely-escaped string	
	fmt.Printf("The top selling books are: %q\n", topBooks)

	// Go fmt verbs with precision formatting
	pi := 3.14159
	eulers := 2.71828

	// print float values in specific precision format
	fmt.Printf("%.2f | %.2f\n", pi, eulers)
}
```

### Output

```bash
{apple broccoli}
{fruit:apple vegetable:broccoli}
main.Food
This book is available: false
This book was published in: 2018
This book costs $20.95
The title of this book is: The Overstory
The top selling books are: ["The Overstory" "Alias Grace"]
3.14 | 2.72
```

<a href='https://play.golang.org/p/_CDBhxBHGnV' target='_blank'>Try It Out</a> | <a href='https://github.com/sagar-jadhav/go-examples/blob/master/src/format-verbs.go' target='_blank'>Source Code</a>

[<< Home Page](./) | [Previous << Math functions](./math-functions.html) | [Next >> String length](./string-length.html)
