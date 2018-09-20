package main

import "fmt"

/*
 * for-loop program demonstrates for loop syntax &
 * usage by priniting number from 1 to 50 & 50 to 100
 */
func main() {
	// Variation #1
	var indexA int = 1
	for indexA <= 50 {
		fmt.Println(indexA)
		indexA++
	}

	// Variation #2
	for indexB := 51; indexB <= 100; indexB++ {
		fmt.Println(indexB)
	}
}
