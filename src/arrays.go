package main

import "fmt"

/*
 * arrays program demonstrates use of arrays & syntax for arrays
 * in Go programming language
 */
func main() {
	// Defining Array Variation #1
	var avengers [5]string

	avengers[0] = "Iron Man"
	avengers[1] = "Captain America"
	avengers[2] = "Ant Man"
	avengers[3] = "Spider Man"
	avengers[4] = "Thor"

	for index := 0; index < len(avengers); index++ {
		fmt.Printf("Avenger at index [%d] is '%s'", index, avengers[index])
		fmt.Println()
	}

	// Defining Array Variation #2
	cricketers := [5]string{
		"Sachin Tendulkar",
		"MS Dhoni",
		"Rahul Dravid",
		"Anil Kumble",
		"Virat Kohli",
	}

	for index, cricketer := range cricketers {
		fmt.Printf("Cricketer at index [%d] is '%s'", index, cricketer)
		fmt.Println()
	}
}
