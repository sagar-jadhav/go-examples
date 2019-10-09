---
layout: default
---

# If-Elseif-Else

The example demonstrates usage of if, elseif and else clause in go. if-elseif-else program demonstrates the conditional statements in the Go programming language. In this program evenOdd function will explain if else block and signToWord program demonstrates if-elseif-else block.

```
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

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/if-elseif-else.go)

# Output

```
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

[Back](./)