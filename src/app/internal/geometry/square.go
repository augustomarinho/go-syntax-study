package geometry

type Square struct {
	Width, Height float64
}

func (s Square) area() float64 {
	return s.Height * s.Width
}

func (s Square) perimeter() float64 {
	return 2*s.Height + 2*s.Width
}
