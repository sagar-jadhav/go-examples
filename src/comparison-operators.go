package main

import "fmt"

// Comparison-operators program demonstrates different
// comparison operators in Go, like:
// == (equal to)
// != (not equal to)
//  < (less than)
//  > (greater than)
// <= (less than or equal to)
// >= (greater than or equal to)
func main() {
	x := 3
	y := 5

	// Comparison operators are used to evaluate values to boolean value of
	// either true or false.
	fmt.Println("x == y =", x == y)
	fmt.Println("x != y =", x != y)
	fmt.Println("x < y  =", x < y)
	fmt.Println("x > y  =", x > y)
	fmt.Println("x <= y =", x <= y)
	fmt.Println("x >= y =", x >= y)
}
