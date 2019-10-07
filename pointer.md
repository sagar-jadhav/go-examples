---
layout: default
---

# Pointer

The example demonstrates pointer. Pointer allowing you to pass references to values and records within your program.

```
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


}
```

[Source Code](https://github.com/sagar-jadhav/go-examples/blob/master/src/pointer.go)

# Output

```
Type of a is *int
address of b is 0xc00001a0d0
Size value is 0, type is *int, address is 0xc00001a0d8
New size value is 85
```

[Back](./)