/* read_from_console.go will demonstrate how to take
input from the console and continue doing the same
till we enter the key string(here "hi") to terminate the program */
package main

import (
	"bufio" /* Package bufio implements buffered I/O. */
	"fmt"
	"os"
	"strings"
)

/* We’ll use Go’s while loop equivalent of a for loop without any parameters to
ensure our program continues on forever. In this example every time text is entered
and then enter is pressed, we assign text to equal everything up to and including
the \n special character. If we want to do a comparison on the string that has
just been entered then we can use the strings.Replace method in order to remove
this trailing \n character with nothing and then do the comparison. */

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Simple Shell")
	fmt.Println("---------------------")
	for {
		fmt.Print("-> ")
		text, _ := reader.ReadString('\n')
		text = strings.Replace(text, "\r\n", "", -1) // use this for windows
		// text = strings.Replace(text, "\n", "", -1) use this for Unix
		if strings.Compare("hi", text) == 0 {
			fmt.Println("hello, Yourself")
			break
		}
	}
}

/* Output :
Simple Shell
---------------------
-> input 1
-> input 2
-> input 3
-> next step will be terminating step
-> hi
hello, Yourself */
