package main

import (
	"fmt"
	"time"
)

func startJob() {
	//Defer is used to ensure that a function call is performed later
	//in a program’s execution,
	//usually for purposes of cleanup.

	//Cleanup() will be executed at the end of the enclosing function (startJob),
	//after startJob() has finished.
	defer cleanUp()

	startSomething()
	doSomethingElse()
	somethingAfterThis()
	finallyDoThis()
}

func startSomething() {
	fmt.Println("Started doing something ... ")
	time.Sleep(2 * time.Second)
	fmt.Println("Done something!")
}

func doSomethingElse() {
	fmt.Println("Started doing something else... ")
	time.Sleep(2 * time.Second)
	fmt.Println("Done something else!")
}

func somethingAfterThis() {
	fmt.Println("Started something after this... ")
	time.Sleep(2 * time.Second)
	fmt.Println("Done something after this!")
}

func finallyDoThis() {
	fmt.Println("Started something finally... ")
	time.Sleep(2 * time.Second)
	fmt.Println("Finally done something!")
}

func cleanUp() {
	fmt.Println("Cleaning up this mess... ")
	time.Sleep(4 * time.Second)
	fmt.Println("Squeaky clean now!")
}

func main() {
	fmt.Println("Starting job .... ")
	startJob()
	fmt.Println("Job Complete!")
}
