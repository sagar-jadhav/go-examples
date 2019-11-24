---
layout: default
---

# Comparison-operators

 program demonstrates different comparison operators in Go

```
func main() {
	x := 3
	y := 5

	// Comparison operators are used to evaluate values to boolean value of
	// either true or false.
	fmt.Println("x == y =", x == y)
	fmt.Println("x != y =", x != y)
	fmt.Println("x < y  =", x < y)
	fmt.Println("x > y  =", x > y)
	fmt.Println("x <= y =", x <= y)
	fmt.Println("x >= y =", x >= y)
}
```
<a href='https://play.golang.org/p/lbr8W98IwfW' target='_blank'>Try It Out</a>

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/comparison-operators.go)

### Output

```
x == y = false
x != y = true
x < y  = true
x > y  = false
x <= y = true
x >= y = false
```

[Back](./)
