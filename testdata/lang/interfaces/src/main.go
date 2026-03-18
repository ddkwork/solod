package main

type Shape interface {
	Area() int
	Perim(n int) int
}

type Rect struct {
	width, height int
}

func (r *Rect) Area() int {
	return r.width * r.height
}

func (r *Rect) Perim(n int) int {
	return n * (2*r.width + 2*r.height)
}

func calcShape(s Shape) int {
	return s.Perim(2) + s.Area()
}

func shapeIsRect(s Shape) bool {
	_, ok := s.(*Rect)
	return ok
}

func shapeAsRect(s Shape) *Rect {
	_, ok := s.(*Rect)
	if !ok {
		return nil
	}
	r := s.(*Rect)
	return r
}

func rectAsShape(r *Rect) Shape {
	return r
}

func shapeCheckAssign(s Shape) bool {
	var ok bool
	_, ok = s.(*Rect)
	return ok
}

func nilShape() Shape {
	return nil
}

func main() {
	r := Rect{width: 10, height: 5}
	{
		// Shape interface is implemented by *Rect pointer.
		s := Shape(&r)
		var s2 Shape = &r // also works
		_ = s2

		calcShape(s)
		calcShape(Shape(&r)) // also works
		calcShape(&r)        // also works

		_ = shapeIsRect(s)
		_ = shapeCheckAssign(s)
		rval := shapeAsRect(s)
		_ = rval
	}
	{
		// Wrap Rect value into Shape via function.
		s := rectAsShape(&r)
		_ = s
	}
	{
		// Nil interface.
		var s1 Shape
		if s1 != nil {
			panic("want nil interface")
		}
		var s2 Shape = nil
		if s2 != nil {
			panic("want nil interface")
		}
		s3 := nilShape()
		if s3 != nil {
			panic("want nil interface")
		}
		isRect := shapeIsRect(nil)
		if isRect {
			panic("want isRect == false")
		}
		var r Rect
		var s4 Shape = &r
		if s4 == nil {
			panic("want non-nil interface")
		}
	}
}
