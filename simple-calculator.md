---
layout: default
---

# Simple Calculator

The example demonstrates different approaches to assigning values to variables, and performing mathematical operations on those variables.

Here we explicitly initialize `variableX` and `variableY` using the `var` declaration, and also explicitly define the type `int`.
```
	// type #1 variable declaration
	var variableX int = 24
	var variableY int = 12
```

This snippet demonstrates the short assignment `:=` which replaces the `var` declaration. This also allows inferred typing based on the value on the right-hand side. This cannot be used outside of a function.

```
	// type #2 variable declaration
	variableA := 24
	variableB := 12
```

Here we explicitly initialize `variableF` and `variableG` but type is inferred based on the value on the right-hand side.

```
	// type #3 variable declaration
	var variableF = 24
	var variableG = 12
```
This snippet demonstrates the initializing of multiple variables that may vary in type. 

```
	// type #4 variable declaration
	var (
		variableJ = 24
		variableK = 12
	)
```

Here we output calculations using these variables.

```
	fmt.Println("Result of Addition of 24,12 : ", variableX+variableY)
	fmt.Println("Result of Multiplication of 24,12 : ", variableA*variableB)
	fmt.Println("Result of Division of 24,12 : ", variableF/variableG)
	fmt.Println("Result of Subtraction of 24,12 : ", variableJ-variableK)
```

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/simple-calculator.go)

### Output

```
Result of Addition of 24,12 :  36
Result of Multiplication of 24,12 :  288
Result of Division of 24,12 :  2
Result of Subtraction of 24,12 :  12
```

[Back](./)
