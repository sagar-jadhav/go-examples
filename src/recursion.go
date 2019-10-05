package main

import "fmt"

/*
 * recursion.go demonstrates how to use recursion in the Go
 * programming language. Recursion is a function calling itself.
 * Recursion is a key aspect in functional programming.
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
