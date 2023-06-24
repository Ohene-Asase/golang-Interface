package main

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

func main() {
	rect := Rectangle{Width: 4, Height: 3}
	circ := Circle{Radius: 2}
	PrintShapeDetails(rect)
	PrintShapeDetails(circ)
}

func PrintShapeDetails(shape Shape) {
	fmt.Printf("Area: %.2f\n", shape.Area())
	fmt.Printf("Permiter: %.2f/n", shape.Perimeter())
}
