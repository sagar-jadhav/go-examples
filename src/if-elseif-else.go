package main

import "fmt"

/*
 * if-elseif-else program will demonstrates the conditional statements
 * in Go programming language. In this program even-odd program will explain
 * if else block and sign to word program  demonstrate if-elseif-else block
 */
func main() {
	// It will demonstrate if-else statements
	evenOdd()
	// It will demonstrate if-elseif-else statements
	signToWord()
}

func evenOdd() {
	for n := 1; n <= 10; n++ {
		if n%2 == 0 {
			fmt.Printf("Number %d is Even", n)
		} else {
			fmt.Printf("Number %d is Odd", n)
		}
		fmt.Println()
	}
}

func signToWord() {
	var sign string = "="
	if sign == "=" {
		fmt.Printf("Sign %s is equal to", sign)
	} else if sign == "<" {
		fmt.Printf("Sign %s is less than", sign)
	} else {
		fmt.Printf("Not a valid sign")
	}
	fmt.Println()
}
