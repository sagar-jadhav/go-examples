package main

import "fmt"

/*
 * for-loop program demonstrates for loop syntax &
 * usage by printing number from 1 to 50 & 50 to 100
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

	// Variation #3 Iterating over array
	var numbers [5]int

	numbers[0] = 0
	numbers[1] = 1
	numbers[2] = 2
	numbers[3] = 3
	numbers[4] = 4

	for index, number := range numbers {
		fmt.Printf("Number at index [%d] is %d", index, number)
		fmt.Println()
	}

	// In case if we don't require index then we can use '_' as
	// placeholder for it

	for _, number := range numbers {
		fmt.Println(number)
	}

	//Special case iterating over Map
	//Special case when accessing a key from Map
	//which is not present in Map
	alphabets := make(map[string]string)
	alphabets["A"] = "Apple"
	alphabets["B"] = "Ball"

	for key, value := range alphabets {
		fmt.Println(key, value)
	}
}
