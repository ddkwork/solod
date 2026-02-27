package main

type Shape interface {
	Area() int
	Perim(n int) int
}

type Rect struct {
	width, height int
}

func (r Rect) Area() int {
	return r.width * r.height
}

func (r Rect) Perim(n int) int {
	return n * (2*r.width + 2*r.height)
}

func calc(s Shape) int {
	return s.Perim(2) + s.Area()
}

func isRect(s Shape) bool {
	_, ok := s.(Rect)
	return ok
}

func asRect(s Shape) int {
	_, ok := s.(Rect)
	if !ok {
		return 0
	}
	r := s.(Rect)
	return r.Area()
}

func main() {
	r := Rect{width: 10, height: 5}

	s := Shape(r)
	calc(s)
	calc(Shape(r)) // also works
	calc(r)        // also works

	_ = isRect(s)

	a := asRect(s)
	_ = a
}
