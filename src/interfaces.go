package main

import (
	"fmt"
	"math"
)

// Go supports interfaces in a different way than other programming languages
// like Java does. Like a struct, an interface is created using the type keyword,
// followed by a name and the keyword interface:
type Shape interface {
	// area returns the Square's area by multiplying both sides.
	area() float64
}

// Now in order to "implement" this interface, a type must implement the interface methods defined. For example:
type Square struct {
	x1, y1, x2, y2 float64
}

// We defined method area with Square receiver type, so that Square implemented those methods. 
//As Area() method is declared in Shape interface, Square implements Shape interface.
func (s Square) area() float64 {
	return math.Abs((s.x2 - s.x1) * (s.y2 - s.y1))
}

func main() {
	//creating nil Shape reference s.
	var s Shape
	//assigning struct of type Square to Shape reference s. Since Square implements Shape, this is a valid scenario.
	s = Square{0, 0, 5, 5}
	fmt.Printf("dynamic type of s at runtime is %T\n",s) // We can clearly see that at runtime, the type of s is Square.
	// print the Square's area via the Shape interface reference
	fmt.Println(s.area()) // 25
}
