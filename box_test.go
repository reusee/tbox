package tbox

import "testing"

func TestLayout(t *testing.T) {
	ui := New()
	root := ui.NewBox(func() (Point, Point) {
		return Point{0, 0}, Point{400, 300}
	})
	b0 := ui.NewBox(func() (Point, Point) {
		return root.TopLeft().Move(1, 1), root.BottomRight().Move(-1, -1)
	})
	if (root.TopLeft() != Point{0, 0}) || (root.BottomRight() != Point{400, 300}) {
		t.Fatal("root pos")
	}
	if (b0.TopLeft() != Point{1, 1}) || (b0.BottomRight() != Point{399, 299}) {
		t.Fatal("b0 pos")
	}
	root.Repos(func() (Point, Point) {
		return Point{0, 0}, Point{800, 600}
	})
	if (root.TopLeft() != Point{0, 0}) || (root.BottomRight() != Point{800, 600}) {
		t.Fatal("root pos")
	}
	if (b0.TopLeft() != Point{1, 1}) || (b0.BottomRight() != Point{799, 599}) {
		t.Fatal("b0 pos")
	}
}

func TestLayout2(t *testing.T) {
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
	root.Repos(func() (Point, Point) {
		return Point{0, 0}, Point{800, 600}
	})
}
