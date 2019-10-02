package main

import (
	"fmt"
)

func processForm(name *string, age *int8, zipcode *int32) {
	//A panic typically means something went unexpectedly
	//wrong. Mostly used to fail fast on errors that shouldn't
	//occur during normal operation, or that aren't prepared
	//to be handled gracefully

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
	var name string
	var age int8
	var zipcode int32
	name = "Jon Smith"
	age = 23
	zipcode = 60613

	processForm(&name, &age, &zipcode)

	fmt.Println("processForm() executed")

	processForm(nil, &age, &zipcode)  //panic in first if{}
	processForm(&name, nil, &zipcode) //panic in second if{}
	processForm(&name, &age, nil)     //panic in third if{}

	//code below this line will not execute if panic()
	fmt.Println("processForm() executed, ending in main()")
}
