package main

import (
	"fmt"
	"math"
)

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

func main() {
	a := Circle{3.0}
	b := Rectangle{20, 5}
	fmt.Println(a.area())
	fmt.Println(b.area())
}
