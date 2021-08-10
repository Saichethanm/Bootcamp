package main

import (
	"fmt"
	"math"
)

type geometry interface {
	area() float64
	perimeter() float64
}

type Square struct {
	a float64
}

type Circle struct {
	radius float64
}

func (s Square) area() float64 {
	// func (s *Square) area() float64 {
	return s.a * s.a
}

func (s Square) perimeter() float64 {
	// func (s *Square) perimeter() float64 {
	return s.a * 4
}

func (c Circle) area() float64 {
	// func (c *Circle) area() float64 {
	return c.radius * c.radius * math.Pi
}

func (c Circle) perimeter() float64 {
	// func (c *Circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func measure(g geometry) {
	fmt.Println(g)
	fmt.Println("Perimeter :", g.perimeter())
	fmt.Println("Area :", g.area())
}

func main() {
	s := Square{5}
	c := Circle{10}
	measure(&s)
	measure(&c)
}
