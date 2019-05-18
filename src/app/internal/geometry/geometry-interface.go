package geometry

import "fmt"

type Geometry interface {
	area() float64
	perimeter() float64
}

func Calculate(g Geometry) {
	fmt.Println("Area: ", g.area())
	fmt.Println("Perimter: ", g.perimeter())
}
