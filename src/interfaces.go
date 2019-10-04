package main

import (
	"fmt"
)

//Go supports interfaces in a different way than other programming languages
//like Java do. Like a struct an interface is created using the type keyword,
//followed by a name and the keyword interface:
type Shape interface {
	area() float64
}

//Now in order to "implement" this interface, a type must implement the interface methods defined. For example:

type Square struct {
	x1, y1, x2, y2 float64
}

func (s *Square) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y2) // calculate distance between 2 points
	return l * l
}
func main() {

	s := Square{0, 0, 5, 5}
	fmt.Println(s.area()) // 25

}
