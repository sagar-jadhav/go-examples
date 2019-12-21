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
<a href='https://play.golang.org/p/fG_C_2OI03j' target='_blank'>Try It Out</a>

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/recursion.go)

### Output

```bash
120
```

[<< Back](./)

