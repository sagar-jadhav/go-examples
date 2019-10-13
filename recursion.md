---
layout: default
---

# Recursion

The example demonstrates how to use recursion in the Go programming language.
Recursion is a function calling itself. Recursion is a key aspect in functional programming.

```
package main

import "fmt"

func main() {
	fmt.Println(calculateFactorial(5))
}

func calculateFactorial(val int) int {
	if val == 0 {
		return 1
	}
	return val * calculateFactorial(val-1)
}
```

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/recursion.go)

### Output

```
120
```

[Back](./)

