package tbox

import "sync"

type baseBox struct {
	x0, y0, x1, y1 int
	lock           *sync.Mutex
	cond           *sync.Cond
}

type IBox interface {
	X0() int
	Y0() int
	X1() int
	Y1() int
	TopLeft() Point
	TopMiddle() Point
	TopRight() Point
	MiddleLeft() Point
	Middle() Point
	MiddleRight() Point
	BottomLeft() Point
	BottomMiddle() Point
	BottomRight() Point
}

type Box struct {
	*baseBox
	ui  *UI
	pos func() (Point, Point)
}

type RootBox struct {
	*baseBox
}

func (u *UI) NewBox(pos func() (Point, Point)) *Box {
	lock := new(sync.Mutex)
	topLeft, bottomRight := pos()
	box := &Box{
		baseBox: &baseBox{
			lock: lock,
			cond: sync.NewCond(lock),
			x0:   topLeft.X,
			y0:   topLeft.Y,
			x1:   bottomRight.X,
			y1:   bottomRight.Y,
		},
		ui:  u,
		pos: pos,
	}
	u.boxes = append(u.boxes, box)
	return box
}

func (b *baseBox) withCondLock(fn func()) {
	b.cond.L.Lock()
	fn()
	b.cond.L.Unlock()
}

func (b *baseBox) X0() (ret int) {
	b.withCondLock(func() {
		for b.x0 == -1 {
			b.cond.Wait()
		}
		ret = b.x0
	})
	return
}

func (b *baseBox) Y0() (ret int) {
	b.withCondLock(func() {
		for b.y0 == -1 { //NOCOVER
			b.cond.Wait()
		}
		ret = b.y0
	})
	return
}

func (b *baseBox) X1() (ret int) {
	b.withCondLock(func() {
		for b.x1 == -1 { //NOCOVER
			b.cond.Wait()
		}
		ret = b.x1
	})
	return
}

func (b *baseBox) Y1() (ret int) {
	b.withCondLock(func() {
		for b.y1 == -1 { //NOCOVER
			b.cond.Wait()
		}
		ret = b.y1
	})
	return
}

func (b *baseBox) TopLeft() Point {
	return Point{b.X0(), b.Y0()}
}

func (b *baseBox) TopMiddle() Point {
	return Point{(b.X0() + b.X1()) / 2, b.Y0()}
}

func (b *baseBox) TopRight() Point {
	return Point{b.X1(), b.Y0()}
}

func (b *baseBox) MiddleLeft() Point {
	return Point{b.X0(), (b.Y0() + b.Y1()) / 2}
}

func (b *baseBox) Middle() Point {
	return Point{(b.X0() + b.X1()) / 2, (b.Y0() + b.Y1()) / 2}
}

func (b *baseBox) MiddleRight() Point {
	return Point{b.X1(), (b.Y0() + b.Y1()) / 2}
}

func (b *baseBox) BottomLeft() Point {
	return Point{b.X0(), b.Y1()}
}

func (b *baseBox) BottomMiddle() Point {
	return Point{(b.X0() + b.X1()) / 2, b.Y1()}
}

func (b *baseBox) BottomRight() Point {
	return Point{b.X1(), b.Y1()}
}

func (b *Box) Repos(pos func() (Point, Point)) {
	b.ui.withLock(func() {
		b.pos = pos
		topLeft, bottomRight := pos()
		b.withCondLock(func() {
			b.x0 = topLeft.X
			b.y0 = topLeft.Y
			b.x1 = bottomRight.X
			b.y1 = bottomRight.Y
		})
	})
	b.ui.Relayout()
}
