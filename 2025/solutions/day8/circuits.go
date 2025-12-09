package day8

import "math"

type JunctionBox struct {
	X, Y, Z int
}

type BoxPair struct {
	box1, box2 *JunctionBox
}

func (p *BoxPair) Distance() float64 {
	x_diff := p.box1.X - p.box2.X
	y_diff := p.box1.Y - p.box2.Y
	z_diff := p.box1.Z - p.box2.Z
	return math.Sqrt(float64(x_diff*x_diff) + float64(y_diff*y_diff) + float64(z_diff*z_diff))
}

func (p *BoxPair) Connected(o *BoxPair) bool {
	return o.box1 == p.box1 || o.box1 == p.box2 || o.box2 == p.box1 || o.box2 == p.box2
}

type MinMaxHeap struct {
	slice [1000]*BoxPair
}

func NewMinMaxHeap() *MinMaxHeap {
	return &MinMaxHeap{[1000]*BoxPair{}}
}

func (h *MinMaxHeap) pushUp(i int) {}

func (h *MinMaxHeap) pushDn(i int) {
	// minChildOrGrandchild := func(m int) int {
	// 	return 0
	// }

	m := i
	// Child nodes are stored at i*2+1 in the slice; this is just
	// checking if there's children.
	for m*2+1 < len(h.slice) {
		i := m
		// "Min" values are stored at even indices, "Max" values at odd.
		if i%2 == 0 {
			m = i*2 + 1
		} else {

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
