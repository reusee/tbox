package tbox

type Point struct {
	X, Y int
}

func (p Point) Move(x, y int) Point {
	x = p.X + x
	if x < 0 {
		x = 0
	}
	y = p.Y + y
	if y < 0 {
		y = 0
	}
	return Point{x, y}
}
