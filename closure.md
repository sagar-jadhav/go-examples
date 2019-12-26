---
layout: default
---

# Closure

Example demonstrates how to declare and use the function inside function in Go programming language. Closures are one of the key aspects of functional programming.
Click [here](https://tour.golang.org/moretypes/25) to learn more

```go
func main() {
	fmt.Println(outerFuncOne())
	fmt.Println(outerFuncTwo())
	returnInnerFunction := returnFunction()
	fmt.Println(returnInnerFunction())
}

func outerFuncOne() string {
	message := "Hello"

	innerFuncOne := func(message string) (greetingMessage string) {
		greetingMessage = message + "-" + "World!"
		return greetingMessage
	}

	return innerFuncOne(message)
}

/*
 * Inner functions have access to local variables
 * define inside the outer function
 */
func outerFuncTwo() int {
	x := 1

	innerFuncTwo := func() {
		x += 1
	}

	innerFuncTwo()
	return x
}

/*
 * Function can also return another function similar
 * to variables.
 */
func returnFunction() func() string {
	message := "Greetings from "
	return func() string {
		message += "Go!"
		return message
	}
}

```

### Ouput

```bash
Hello-World!
2
Greetings from Go!
```

<a href='https://play.golang.org/p/8Y_XC1cgziU' target='_blank'>Try It Out</a> | <a href='https://github.com/sagar-jadhav/go-examples/blob/master/src/closure.go' target='_blank'>Source Code</a>

### Contributors
- <a href='https://github.com/sagar-jadhav' target='_blank'>Sagar Jadhav</a>

[<< Home Page](./) | [Previous << Variadic Functions](./variadic.html) | [Next >> Recursion](./recursion.html)
