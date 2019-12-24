---
layout: default
---

# Pointer

Example demonstrates pointer. Pointer allowing you to pass references to values and records within your program.
Click [here](https://tour.golang.org/moretypes/1) to learn more

```go
func main() {
	b := 255
	//Saving b's address into variable 'a', 
	//so 'a' is variable of type pointer to 
	//the base type that is integer.
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	//Creating pointers using the new function

	//new(Type) is a built in function that 
	//allocates memory and return address of it
	size := new(int)
	fmt.Printf("Size value is %d, type is %T, 
	address is %v\n", *size, size, size)
	*size = 85
	fmt.Println("New size value is", *size)
	// by using pointers when handing over function 
	//parameters you can even manipulate parameters and their origin values
	manipulatePointerValue(size)
	// 170
	fmt.Println("New size value is", *size) 
}

// manipulatePointerValue simply just doubles the given input
func manipulatePointerValue(input *int) {
	*input = 2 * (*input)
}
```

<a href='https://play.golang.org/p/RlzXQRU2WJl' target='_blank'>Try It Out</a>

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/pointer.go)

# Output

```bash
Type of a is *int
address of b is 0x40e020
Size value is 0, type is *int, address is 0x40e024
New size value is 85
New size value is 170
```

[<< Back](./)
