---
layout: default
---

# Recursion

Example demonstrates how to use recursion in the Go programming language.
Recursion is a function calling itself. Recursion is a key aspect in functional programming.

```go
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

### Output

```bash
120
```

<a href='https://play.golang.org/p/fG_C_2OI03j' target='_blank'>Try It Out</a> | <a href='https://github.com/sagar-jadhav/go-examples/blob/master/src/recursion.go' target='_blank'>Source Code</a>

### Contributors
- <a href='https://github.com/sagar-jadhav' target='_blank'>Sagar Jadhav</a>
- <a href='https://github.com/ilmanzo' target='_blank'>Andrea Manzini</a>

[<< Home Page](./) | [Previous << Closure](./closure.html) | [Next >> Boolean Expressions](./boolean-expressions.html)

