package main

import "fmt"

/*
 * recursion.go demonstrate how to use recursion in go
 * programming language. Recursion is calling a function itself.
 * Recursion is key aspect in functional programming.
 */
func main() {
	fmt.Println(calculateFactorial(5))
}

func calculateFactorial(val int) int {
	if val == 0 {
		return 1
	}
	return val * calculateFactorial(val-1)
}
