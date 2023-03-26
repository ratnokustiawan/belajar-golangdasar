package main

import (
	"fmt"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Width + 2*r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func main() {
	var s Shape
	s = Rectangle{Width: 3, Height: 4}
	fmt.Println("Rectangle area:", s.Area())
	fmt.Println("Rectangle perimeter:", s.Perimeter())

	s = Circle{Radius: 5}
	fmt.Println("Circle area:", s.Area())
	fmt.Println("Circle perimeter:", s.Perimeter())
}
