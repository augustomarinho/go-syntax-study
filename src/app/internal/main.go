package main

import (
	"fmt"

	"app/internal/files"
	"app/internal/geometry"
)

func main() {

	fmt.Println("Starting program")
	s := geometry.Square{Width: 3, Height: 5}
	geometry.Calculate(s)

	banner := new(files.Banner)
	banner.NewBanner("/home/augustomarinho/dev/workspace_go/go-syntax-study/src/app/internal/banner")

	fmt.Println("BANNER")
	fmt.Println(banner.GetContent())
}
