package main

import "fmt"

/*
 * simple-calculator program demonstrates different mathematical
 * operations in Go along with different way of defining a
 * variables in Go
 */
func main() {
	// type #1 variable declaration
	var variableX int = 24
	var variableY int = 12

	// type #2 variable declaration
	variableA := 24
	variableB := 12

	// type #3 variable declaration
	var variableF = 24
	var variableG = 12

	// type #4 variable declaration
	var (
		variableJ = 24
		variableK = 12
	)

	fmt.Println("Result of Addition of 24,12 : ", variableX+variableY)
	fmt.Println("Result of Multiplication of 24,12 : ", variableA*variableB)
	fmt.Println("Result of Division of 24,12 : ", variableF/variableG)
	fmt.Println("Result of Subtraction of 24,12 : ", variableJ-variableK)
}
