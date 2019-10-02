//Using panic example to demonstrate working of "recover"
package main

import (
	"fmt"
)

func recoverForm() {
	// recover() is a builtin function which is used to regain control of a panicking goroutine.

	// recover() is useful only when called inside deferred functions (here, recoverForm()).
	// Executing a call to recover inside a deferred function stops the panicking sequence
	// by restoring normal execution and retrieves the error value passed to the call of panic.
	recoverObj := recover()

	if recoverObj != nil {
		fmt.Println("Recovered from: ", recoverObj) //prints the error message thrown by panic()
	}
}

func processForm(name *string, age *int8, zipcode *int32) {
	defer recoverForm()

	if name == nil {
		panic("runtime error: name cannot be nil")
	}
	if name == nil {
		panic("runtime error: age cannot be nil")
	}
	if name == nil {
		panic("runtime error: zipcode cannot be nil")
	}
	//code under this line will not execute if panic()

	fmt.Printf("Name: %s \n", *name)
	fmt.Printf("Age: %d \n", *age)
	fmt.Printf("Zipcode: %d \n", *zipcode)
}

func main() {
	var age int8
	var zipcode int32

	age = 23
	zipcode = 60613

	processForm(nil, &age, &zipcode) //panic in first if{}

	//code below this line will not execute if panic()
	fmt.Println("processForm() executed, ending in main()")
}
