---
layout: default
---

# Comparison Operators

Example demonstrates different comparison operators in Go. Comparison operators are used to evaluate values to boolean value of either true or false.
Click [here](https://golang.org/ref/spec#Arithmetic_operators) to learn more

```go
func main() {
	x := 3
	y := 5

	fmt.Println("x == y =", x == y)
	fmt.Println("x != y =", x != y)
	fmt.Println("x < y  =", x < y)
	fmt.Println("x > y  =", x > y)
	fmt.Println("x <= y =", x <= y)
	fmt.Println("x >= y =", x >= y)
}
```

### Output

```bash
x == y = false
x != y = true
x < y  = true
x > y  = false
x <= y = true
x >= y = false
```

<a href='https://play.golang.org/p/4DDLxqsNZmg' target='_blank'>Try It Out</a> | <a href='https://github.com/sagar-jadhav/go-examples/blob/master/src/comparison-operators.go' target='_blank'>Source Code</a>

### Contributors
- <a href='https://github.com/robihdy' target='_blank'>Robi Hidayat</a>

[<< Home Page](./) | [Previous << Boolean Expressions](./boolean-expressions.html) | [Next >> Math functions](./math-functions.html)
