package main

import "fmt"

type Shape interface {
	perimeter() float64
}

type Rectangle struct {
	length, width float64
}

func (r Rectangle) perimeter() float64 {
	return (2 * r.length) + (2 * r.width)
}

type Square struct {
	side float64
}

func (s Square) perimeter() float64 {
	return s.side * 4
}

func getPerimeter(s Shape) float64 {
	return s.perimeter()
}

func main() {
	rectangle := Rectangle{width: 5, length: 4}
	square := Square{side: 4}

	fmt.Println("Rectangle Perimeter", getPerimeter(rectangle))
	fmt.Println("Square Perimeter", getPerimeter(square))
}
