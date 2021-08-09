package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

type circle struct {
	radius float64
}

type rectangle struct {
	len     float64
	breadth float64
}

func getArea(s shape) float64 {
	return s.area()
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (r rectangle) area() float64 {
	return r.len * r.breadth
}

func main() {
	c := circle{7}
	r := rectangle{5, 10}

	// fmt.Println(c.getArea())
	// fmt.Println(r.getArea())

	shapes := []shape{c, r}

	for _, shape := range shapes {
		// fmt.Println(shape.getArea())
		fmt.Println(getArea(shape))
	}
}
