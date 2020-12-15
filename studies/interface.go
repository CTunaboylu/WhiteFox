package main

import (
	"fmt"
	"math"
)

// Every type that has an area() method returning a float64 is of type shape
type shape interface {
	area() float64
	// if there would be another function here, then circle and rectangle
	// would not satisfy the condition thus they do not implement the shape
	//interface
}
type circle struct {
	radius float64
}
type rectangle struct {
	width  float64
	height float64
}

func (r rectangle) area() float64 {
	return r.width * r.height
}
func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}
func main() {
	c1 := circle{4.5}
	r1 := rectangle{5, 7}
	shapes := []shape{c1, r1}

	for _, shape := range shapes {
		fmt.Printf("%f\n", shape.area())
	}

}
