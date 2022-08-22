package cube

type Cube struct {
	L, B, H float32 // only PascalCase are exported in a struct
}

func (c *Cube) Area() float32 {
	return c.L * c.B * c.H
}

func (c *Cube) Perimeter() float32 {
	return 4 * (c.L + c.B + c.H)
}
