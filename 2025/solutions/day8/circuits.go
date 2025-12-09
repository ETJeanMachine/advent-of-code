package day8

import "math"

type JunctionBox struct {
	X, Y, Z int
}

// The "absolute" distance between two junction boxes
// (faster to compute than the actual straight line distance)
func (jb *JunctionBox) Distance(o *JunctionBox) int {
	x_diff := math.Abs(float64(jb.X - o.X))
	y_diff := math.Abs(float64(jb.Y - o.Y))
	z_diff := math.Abs(float64(jb.Z - o.Z))
	return int(x_diff + y_diff + z_diff)
}

type BoxPair struct {
	box1, box2 *JunctionBox
	absDist    int
}

func (p *BoxPair) Connected(o *BoxPair) bool {
	return o.box1 == p.box1 || o.box1 == p.box2 || o.box2 == p.box1 || o.box2 == p.box2
}

func NewBoxPair(box1 *JunctionBox, box2 *JunctionBox) *BoxPair {
	return &BoxPair{box1, box2, box1.Distance(box2)}
}

type MinMaxHeap struct {
	slice [1000]*BoxPair
}

func NewMinMaxHeap() *MinMaxHeap {
	return &MinMaxHeap{[1000]*BoxPair{}}
}

func (h *MinMaxHeap) pushUp(i int) {}

func (h *MinMaxHeap) pushDn(i int) {
	childrenAndGrandchildren := func(m int) []int {
		children := []int{m*2 + 1, m*2 + 1}
		for _, c := range children {
			children = append(children, c*2+1, c*2+2)
		}
		idx := len(h.slice)
		for i := range children {
			if children[i] >= len(h.slice) {
				idx = i
				break
			}
		}
		return children[:idx]
	}

	// Child nodes are stored at i*2+1 in the slice; this is just
	// checking if there's children.
	m := i
	for len(children) > 0 {
		i = m
		children := childrenAndGrandchildren(i)
		// "Min" values are stored at even indices, "Max" values at odd.
		if i%2 == 0 {
			m = children[0]
			if len(children) > 1 {
				for _, child := range children[1:] {
					if h.slice[m].absDist > h.slice[child].absDist {
						m = child
					}
				}
			}
			if h.slice[m].absDist < h.slice[i].absDist {
				tmp := *h.slice[m]
				h.slice[m] = h.slice[i]
				h.slice[i] = &tmp
				if m > i*2+2 {

				}
			} else {
				break
			}
		} else {
			m = children[0]
			if len(children) > 1 {
				for _, child := range children[1:] {
					if h.slice[m].absDist < h.slice[child].absDist {
						m = child
					}
				}
			}
			if h.slice[m].absDist > h.slice[i].absDist {
				tmp := *h.slice[m]
				h.slice[m] = h.slice[i]
				h.slice[i] = &tmp
			} else {
				break
			}
		}

	}
}

func (h *MinMaxHeap) Insert(pair *BoxPair) {}

func (h *MinMaxHeap) PopMin() *BoxPair {
	return nil
}

func (h *MinMaxHeap) PopMax() *BoxPair {
	return nil
}
