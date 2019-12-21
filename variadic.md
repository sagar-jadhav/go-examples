---
layout: default
---

# Variadic functions

Example demonstrates variadic functions. Variadic functions are functions which accept a variable number of parameters.

```go
package main

import "fmt"

func main() {
	//passing zero parameter
	printNumbers()

	//passing two parameters
	printNumbers(5, 6)

	//passing slice
	slice := []int{1, 2}
	printNumbers(slice...)
}

func printNumbers(numbers ...int) {
	for index, number := range numbers {
		fmt.Printf("Number at index %d is %d", index, number)
		fmt.Println()
	}
}
```
<a href='https://play.golang.org/p/jO0EqOanaQF' target='_blank'>Try It Out</a>

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/variadic.go)

### Output

```bash
Number at index 0 is 5
Number at index 1 is 6
Number at index 0 is 1
Number at index 1 is 2
```

[<< Back](./)
