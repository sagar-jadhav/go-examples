package main

import (
	"fmt"
	//OS module provides APIs for os level commands
	"os"
)

func main() {
	//os.Args will fetch the executable name and the command line arguments
	//after the executable name as an array of arguments

	//get the arguments from the shell
	allArgs := os.Args

	fmt.Println(allArgs)
	fmt.Println("Program name: ", allArgs[0])
	fmt.Println("Arguments: ", allArgs[1:]) //slicing from index 1 onwards (starting from 2nd element)
}
