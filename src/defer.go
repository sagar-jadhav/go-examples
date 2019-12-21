package main

import "fmt"

func main() {
	//A defer statement delays the execution of a function until the surrounding function returns.
	defer fmt.Println("world") // will not execute until control return from main function.
	fmt.Print("Hello ")
}
