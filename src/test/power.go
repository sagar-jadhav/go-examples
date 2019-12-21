package main

import "math"

/*
 * power demonstrate how to calculate a^b
 */

func main() {
	fmt.Println(pow(2, 4))
}

func pow(base, exp float64) float64 {
	return math.Pow(base, exp)
}

