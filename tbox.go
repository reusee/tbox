package tbox

import "sync"

type UI struct {
	lock  *sync.Mutex
	boxes []*Box
	wg    *sync.WaitGroup
	Root  *RootBox
}

func New() *UI {
	lock := new(sync.Mutex)
	return &UI{
		lock: new(sync.Mutex),
		wg:   new(sync.WaitGroup),
		Root: &RootBox{
			baseBox: &baseBox{
				lock: lock,
				cond: sync.NewCond(lock),
			},
		},
	}
}

func (u *UI) Relayout() {
	u.lock.Lock()
	defer u.lock.Unlock()
	for _, box := range u.boxes {
		box.x0 = -1
		box.y0 = -1
		box.x1 = -1
		box.y1 = -1
	}
	u.wg.Add(len(u.boxes))
	for _, box := range u.boxes {
		box := box
		go func() {
			topLeft, bottomRight := box.pos()
			box.cond.L.Lock()
			box.x0 = topLeft.X
			box.y0 = topLeft.Y
			box.x1 = bottomRight.X
			box.y1 = bottomRight.Y
			box.cond.L.Unlock()
			box.cond.Broadcast()
			u.wg.Done()
		}()
	}
	u.wg.Wait()
}
