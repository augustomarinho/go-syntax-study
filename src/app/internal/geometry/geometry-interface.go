package geometry

import "fmt"

type Geometry interface {
	area() float64
	perimeter() float64
}

type Square struct {
	Width, Height float64
}

type Circle struct {
	Radius float64
}

func (s Square) area() float64 {
	return s.Height * s.Width
}

func (s Square) perimeter() float64 {
	return 2*s.Height + 2*s.Width
}

func Calculate(g Geometry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())
}
