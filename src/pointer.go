package main

import (
	"fmt"
)

func main() {
	b := 255
	var a *int = &b //Saving b's address into variable 'a', so 'a' is variable of type pointer to the base type that is integer.
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	// Creating pointers using the new function

	size := new(int) // new(Type) is a built in function that allocates memory and return address of it
	fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)
	// by using pointers when handing over function parameters you can even manipulate parameters and their origin values
	manipulatePointerValue(size)
	fmt.Println("New size value is", *size) // 170
}

// manipulatePointerValue simply just doubles the given input
func manipulatePointerValue(input *int) {
	*input = 2 * (*input)
}
