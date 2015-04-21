package tbox

import "testing"

func TestLayout(t *testing.T) {
	ui := New()
	ui.Resize(400, 300)
	b0 := ui.NewBox(func() (Point, Point) {
		return ui.Root.TopLeft().Move(1, 1), ui.Root.BottomRight().Move(-1, -1)
	})
	if (ui.Root.TopLeft() != Point{0, 0}) || (ui.Root.BottomRight() != Point{400, 300}) {
		t.Fatal("root pos")
	}
	if (b0.TopLeft() != Point{1, 1}) || (b0.BottomRight() != Point{399, 299}) {
		t.Fatal("b0 pos")
	}
	ui.Resize(800, 600)
	if (ui.Root.TopLeft() != Point{0, 0}) || (ui.Root.BottomRight() != Point{800, 600}) {
		t.Fatal("root pos")
	}
	if (b0.TopLeft() != Point{1, 1}) || (b0.BottomRight() != Point{799, 599}) {
		t.Fatal("b0 pos")
	}
}

func TestLayout2(t *testing.T) {
	ui := New()
	ui.Resize(400, 300)
	var last IBox = ui.Root
	n := 512
	boxes := []*Box{}
	for i := 0; i < n; i++ {
		var follow IBox = last
		b := ui.NewBox(func() (Point, Point) {
			return follow.TopLeft().Move(1, 1), follow.BottomRight().Move(-1, -1)
		})
		boxes = append(boxes, b)
		last = b
	}
	ui.Resize(800, 600)
	last = ui.Root
	for i := 0; i < n; i++ {
		box := boxes[i]
		if box.TopLeft() != last.TopLeft().Move(1, 1) {
			t.Fatalf("top left of %d", i)
		}
		if box.BottomRight() != last.BottomRight().Move(-1, -1) {
			t.Fatalf("bottom right of %d", i)
		}
		last = box
	}
}

func TestBoxPoints(t *testing.T) {
	ui := New()
	box := ui.NewBox(func() (Point, Point) {
		return Point{1, 2}, Point{3, 4}
	})
	if box.TopLeft() != (Point{1, 2}) {
		t.Fatal("top left")
	}
	if box.TopMiddle() != (Point{2, 2}) {
		t.Fatal("top middle")
	}
	if box.TopRight() != (Point{3, 2}) {
		t.Fatal("top right")
	}
	if box.MiddleLeft() != (Point{1, 3}) {
		t.Fatal("middle left")
	}
	if box.Middle() != (Point{2, 3}) {
		t.Fatal("middle")
	}
	if box.MiddleRight() != (Point{3, 3}) {
		t.Fatal("middle right")
	}
	if box.BottomLeft() != (Point{1, 4}) {
		t.Fatal("bottom left")
	}
	if box.BottomMiddle() != (Point{2, 4}) {
		t.Fatal("bottom middle")
	}
	if box.BottomRight() != (Point{3, 4}) {
		t.Fatal("bottom right")
	}
}
