package main

import "fmt"

/*
 * boolean-expressions program demonstrates different
 * boolean expressions in Go like ||,&&,!
 */
func main() {
	fmt.Println("true && false = ", (true && false))
	fmt.Println("true || false = ", (true || false))
	fmt.Println("!true = ", !true)
}
