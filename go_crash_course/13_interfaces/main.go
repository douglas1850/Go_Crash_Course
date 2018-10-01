package main

import (
	"fmt"
	"math"
)

//interfaces are data types that represent a set of method signatures for structs

//define Interface
type Shape interface {
	area() float64
}

//now different types of area methods can be defined, and will act in a certain way
//depending on the type of struct passed to them

type Circle struct {
	x, y, radius float64
}

type Rectangle struct {
	width, height float64
}

func (c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2) //p * radius squares
}

func (r Rectangle) area() float64 {
	return r.width * r.height
}

func getArea(s Shape) float64 {
	return s.area()
}

func main() {
	circle := Circle{x: 0, y: 0, radius: 5}

	rectangle := Rectangle{width: 10, height: 20}

	fmt.Printf("Circle Area: %f\n", getArea(circle))
	fmt.Printf("Rectangle Area: %f\n", getArea(rectangle))

}
