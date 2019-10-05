package main

import (  
    "fmt"
)

func main() {  
    b := 255
    var a *int = &b                       //Saving b's address into variable a
    fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
	

   // Creating pointers using the new function

   size := new(int)  // we can imagine that "new" keyword is a pointer
   fmt.Printf("Size value is %d, type is %T, address is %v\n", *size, size, size)
   *size = 85 
   fmt.Println("New size value is", *size)

   // Pass a pointer to a function so we can manipulate the value inside the function
   var p = 0
   fmt.Printf("Original value  for the variable p: %d\n", p)
   manipulatePointer(&p)
   fmt.Printf("The value of the variable p has been updated inside the function: %d\n", p)
}

func manipulatePointer(p *int) {
   *p++;
}
