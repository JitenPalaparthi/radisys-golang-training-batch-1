package cube

func Area(l, b, h float32) float32 {
	return l * b * h
}

func Perimeter(l, b, h float32) float32 {
	return 4 * (l + b + h)
}
