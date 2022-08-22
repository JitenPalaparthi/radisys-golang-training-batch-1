package rect

import "errors"

type Rect struct {
	L, B *float32 // using pointers here
	A, P float32
}

func New(l, b *float32) (*Rect, error) {
	if l == nil || b == nil {
		return nil, errors.New("nil length or bredth")
	}
	if *l == 0 || *b == 0 {
		return nil, errors.New("invalid length and bredth")
	}

	return &Rect{L: l, B: b}, nil
}

func (r *Rect) Area() float32 {
	r.A = (*r.L * *r.B)
	return r.A
}

func (r *Rect) Perimeter() float32 {
	r.P = 2 * (*r.L + *r.B)
	return r.P
}
