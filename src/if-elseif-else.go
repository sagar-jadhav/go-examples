package main

import "fmt"

/*
 * if-elseif-else program demonstrates the conditional statements in the
 * Go programming language. In this program evenOdd function will explain
 * if else block and signToWord program demonstrates if-elseif-else block
 */
func main() {
	// It will demonstrate if-else statements
	evenOdd()
	// It will demonstrate if-elseif-else statements
	signToWord()

	//Special case when accessing a key from Map
	//which is not present in Map
	alphabets := make(map[string]string)
	alphabets["A"] = "Apple"
	if value, isSuccessful := alphabets["A"]; isSuccessful {
		fmt.Println(value)
	}

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
