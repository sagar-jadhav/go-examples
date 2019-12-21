---
layout: default
---

# Pointer

Example demonstrates pointer. Pointer allowing you to pass references to values and records within your program.

```go
package main

import (  
    "fmt"
)

func main() {  
    b := 255
	//Saving b's address into variable a
    var a *int = &b                       
    fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)
	

   // Creating pointers using the new function
   // we can imagine that "new" keyword is a pointer
   size := new(int)  
   fmt.Printf("Size value is %d, type is %T, 
    address is %v\n", *size, size, size)
   *size = 85 
   fmt.Println("New size value is", *size)


}
```
https://play.golang.org/p/RlzXQRU2WJl

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/pointer.go)

# Output

```bash
Type of a is *int
address of b is 0xc00001a0d0
Size value is 0, type is *int, address is 0xc00001a0d8
New size value is 85
```

[<< Back](./)
