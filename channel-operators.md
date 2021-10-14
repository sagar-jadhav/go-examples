---
layout: default
---

# Channel operators: Send and Receive

Example demonstrates send and receive channel  operators in Go programming language. Channel works with two principal operations, one is sending and another one is receiving.
The direction of `<-` operator  indicates whether the data is received  or send. 

Click [here](https://tour.golang.org/concurrency/2) to learn more

```go
/*
* Send operation is used to send data from one goroutine to another with the help of a channel. Values like int, float64, and bool can  be safely send through a channel because      * they are copied, and there is no risk of accidental concurrent access of the same value. Similarly, strings are also safe to transfer because they are immutable. But sending   
* pointers or reference like a slice, map, etc. through a channel are not safe because the value of pointers or reference may change by sending goroutine or by the receiving    
* goroutine at the same time and the result is unpredicted.  So, when you use pointers or references in the channel you must make sure that they can only access by the ONE goroutine * at a time. For example, *Mychannel <- element* indicates that the data(element) SEND to the channel(Mychannel) with the help of a <- operator.
*/

/*
* The receive operation is used to receive the data sent by the send operator.
* Here, *element := <-Mychannel* indicates that the element RECEIVES data from the channel(Mychannel).
* You can also write a receive statement as: *<-My channel*
*/

package main

import "fmt"

func myfunc(ch chan int) {

	fmt.Println(234 + <-ch)
}
func main() {
	fmt.Println("start Main method")
	// Creating a channel
	ch := make(chan int)

	go myfunc(ch)
	ch <- 23
	fmt.Println("End Main method")
}

```
### Output

```bash
start Main method
257
End Main method
```


<a href='https://play.golang.org/p/B-54vzrROU3' target='_blank'>Try It Out</a> | <a href='' target='_blank'>Source Code</a>

### Contributors
- <a href='https://github.com/sagar-jadhav' target='_blank'>Sagar Jadhav</a>

[<< Home Page](./) | [Previous << ](./.html) | [Next >> ](./.html)





