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

func (s *Square) area() float64 {
	return math.Abs((s.x2 - s.x1) * (s.y2 - s.y1))
}

func measure(s Shape) float64 {
	return s.area()
}

func main() {
	// create new 5x5 Square
	s := Square{0, 0, 5, 5}
	// print the Square's area via the Shape interface
	fmt.Println(measure(&s)) // 25
}
