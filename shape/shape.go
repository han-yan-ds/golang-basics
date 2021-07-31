package shape

type Shape interface {
	getArea() float32
}

type Triangle struct {
	Height float32
	Base   float32
}

type Rectangle struct {
	Height float32
	Base   float32
}

func (t Triangle) GetArea() float32 {
	return (t.Height * t.Base) / 2
}

func (r Rectangle) GetArea() float32 {
	return r.Height * r.Base
}
