package main

import "fmt"

type rectangle struct {
	length, breadth float64
}

func (rec *rectangle) area() float64 {
	return rec.length * rec.breadth
}

func (rec *rectangle) perimeter() float64 {
	return 2 * (rec.length + rec.breadth)
}

func main() {
	rec := rectangle{5, 10}
	fmt.Println("Perimeter of rectangle = ", rec.perimeter())
	fmt.Println("Area of rectangle = ", rec.area())

}
