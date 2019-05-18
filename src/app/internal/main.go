package main

import (
	"fmt"

	"app/internal/geometry"
)

func main() {

	fmt.Println("Starting program")
	s := geometry.Square{Width: 3, Height: 5}
	geometry.Calculate(s)
}
