package main

import (
	"fmt"
	"math"
)

//interface implements similarity in a given function e.g
type shape interface {
	area() float64
}

//struct for circles
type Circle struct {
	radius float64
}

//methods  with receiver
func (c Circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

//struct for rectangular
type Rectangle struct {
	width, length float64
}

//method and func

func (r Rectangle) area() float64 {
	return r.length * r.width
}

//getall total area
func totalArea(shapes ...shape) float64 {
	var area float64
	for _, s := range shapes {
		area += s.area()
	}
	return area
}
func main() {
	a := Circle{3.0}
	b := Rectangle{20, 5}

	shapes := []shape{a, b}

	for _, shape := range shapes {
		//fmt.Println(shape.area())
	}

	fmt.Println(totalArea(&a, &b))
	//fmt.Println(a.area())
	//fmt.Println(b.area())
}
