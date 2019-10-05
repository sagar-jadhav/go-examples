---
layout: default
---

# Variadic functions

The example demonstrates variadic functions. Variadic functions are functions which accept a variable number of parameters.

```
package main

import "fmt"

/*
 * variadic.go program demonstrates the use of
 * variadic parameters in the Go programming language
 */
func main() {
	//passing zero parameter
	printNumbers()

	//passing two parameters
	printNumbers(5, 6)

	//passing slice
	slice := []int{1, 2}
	printNumbers(slice...)
}

/*
 * This function can take zero or more parameters
 */
func printNumbers(numbers ...int) {
	for index, number := range numbers {
		fmt.Printf("Number at index %d is %d", index, number)
		fmt.Println()
	}
}
```

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/variadic.go)

### Output

```
Number at index 0 is 5
Number at index 1 is 6
Number at index 0 is 1
Number at index 1 is 2
```

[Back](./)
