package tbox

import "testing"

func BenchmarkLayout512boxes(b *testing.B) {
	ui := New()
	root := ui.NewBox(func() (Point, Point) {
		return Point{0, 0}, Point{800, 600}
	})
	last := root
	n := 512
	for i := 0; i < n; i++ {
		follow := last
		b := ui.NewBox(func() (Point, Point) {
			return follow.TopLeft().Move(1, 1), follow.BottomRight().Move(-1, -1)
		})
		last = b
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		root.Repos(func() (Point, Point) {
			return Point{0, 0}, Point{400, 300}
		})
	}
}
