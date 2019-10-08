---
layout: default
---

# Closure

closure.go program demonstrate how to declare and use the function inside function in Go programming language. Closures are one of the key aspects of functional programming.

```
/*
 * closure.go program demonstrate how to declare
 * and use the function inside function in Go
 * programming language. Closures are one of the
 * key aspects of functional programming
 */
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

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/closure.go)

### Ouput

```
Hello-World!
2
Greetings from Go!
```

[Back](./)
