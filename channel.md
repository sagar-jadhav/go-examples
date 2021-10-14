---
layout: default
---

# Channel 

Example demonstrates creating channel in Go programming language. Channel is a medium through which a goroutine communicates with another goroutine.
By default, channel is bidirectional, it can only  transfer data of the same type.
Click [here](https://tour.golang.org/concurrency/2) to learn more

```go
/*
 * A channel is created using chan keyword.
 * You can also create a channel using make() function using a shorthand declaration. 
 */
package main

import "fmt"

func main() {

	// Creating a channel
	// Using var keyword
	var mychannel chan int
	fmt.Println("Value of the channel: ", mychannel)
	fmt.Printf("Type of the channel: %T ", mychannel)

	// Creating a channel using make() function
	mychannel1 := make(chan int)
	fmt.Println("\nValue of the channel1: ", mychannel1)
	fmt.Printf("Type of the channel1: %T ", mychannel1)
}

```
### Output

```bash
Value of the channel:  
Type of the channel: chan int 
Value of the channel1:  0x432080
Type of the channel1: chan int 
```

<a href='https://play.golang.org/p/HjmlW1VVynF' target='_blank'>Try It Out</a> | <a href='' target='_blank'>Source Code</a>

### Contributors
- <a href='https://github.com/sagar-jadhav' target='_blank'>Sagar Jadhav</a>

[<< Home Page](./) | [Previous << ](./.html) | [Next >> ](./.html)
