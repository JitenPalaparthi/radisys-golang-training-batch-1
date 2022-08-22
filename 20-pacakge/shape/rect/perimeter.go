package rect

// there are not access specifiers
// there is only a concept called exported or not exported
// any memeber/type starts with Uppercase is exported
// any member/type stats with lowercase is not exported

func Perimeter(l, b float32) float32 {
	return 2 * (l + b)
}
