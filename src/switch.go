package main

import "fmt"

/*
 * switch.go program will demonstrate use of switch statements
 * using the sign to word program
 */
func main() {
	var sign string = "="
	switch sign {
	case "=":
		fmt.Println("Equal To")
	case "<":
		fmt.Println("Less Than")
	case ">":
		fmt.Println("Greater Than")
	default:
		fmt.Println("Unknown")
	}
}
