package main

import (
	"fmt"

	"app/internal/geometry"
)

func main() {

	fmt.Println("Hello World!")

	s := geometry.Square{Width: 3, Height: 5}
	geometry.Calculate(s)
}
