---
layout: default
---

# Functions

Example demonstrates how to write a function in Go programming language along with its usage and different options

```go
func main() {
	arr := [6]int{
		1, 2,
		3, 4,
		5, 6,
	}

	slice := arr[:]

	fmt.Printf("Sum of array is : %d", calculateSum(slice))
	fmt.Println()

	fmt.Println(returnSingleValue())

	firstVal, secondVal := returnMultipleValues()
	fmt.Println(firstVal + " " + secondVal)
}

/*
 * It will return the sum of integers provided in the
 * parameters arr. Function signature is comprises of
 * function name, parameters/arguments, return types & body
 */
func calculateSum(arr []int) int {
	sum := 0
	for _, value := range arr {
		sum += value
	}
	return sum
}

/*
 * In function signature we can also provide
 * the name to the return variable
 */
func returnSingleValue() (greeting string) {
	greeting = "Hello Go!"
	return
}

/*
 * We can also return multiple value from function
 * in Go programming language
 */
func returnMultipleValues() (string, string) {
	return "Hello", "World"
}
```


<a href='https://play.golang.org/p/Vedn4G1NuSu' target='_blank'>Try It Out</a>

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/function.go)

### Output

```bash
Sum of array is : 21
Hello Go!
Hello World
```

[<< Back](./)
