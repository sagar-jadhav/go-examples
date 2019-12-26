---
layout: default
---

# If-Elseif-Else

Example demonstrates usage of if, elseif and else clause in go. if-elseif-else program demonstrates the conditional statements in the Go programming language. In this program evenOdd function will explain if else block and signToWord program demonstrates if-elseif-else block.
Click [here](https://tour.golang.org/flowcontrol/7) to learn more

```go
package main

import "fmt"

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
```

### Output

```bash
Number 1 is Odd
Number 2 is Even
Number 3 is Odd
Number 4 is Even
Number 5 is Odd
Number 6 is Even
Number 7 is Odd
Number 8 is Even
Number 9 is Odd
Number 10 is Even
Sign = is equal to
Apple
```

<a href='https://play.golang.org/p/F2vRhPoQU-e' target='_blank'>Try It Out</a> | <a href='https://github.com/sagar-jadhav/go-examples/blob/master/src/if-elseif-else.go' target='_blank'>Source Code</a>

### Contributors
- <a href='https://github.com/sagar-jadhav' target='_blank'>Sagar Jadhav</a>

[<< Home Page](./) | [Previous << Variable Declaration](./simple-calculator.html) | [Next >> Switch Statement](./switch.html)
