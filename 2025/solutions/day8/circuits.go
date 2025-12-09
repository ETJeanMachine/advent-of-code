package day8

import (
	"math/bits"
)

type JunctionBox struct {
	X, Y, Z int
}

// The "square" distance between two junction boxes - faster to compute than the actual
// straight line distance, which has a costly square root call.
func (jb *JunctionBox) squareDistance(o *JunctionBox) int {
	x_diff := (jb.X - o.X) * (jb.X - o.X)
	y_diff := (jb.Y - o.Y) * (jb.Y - o.Y)
	z_diff := (jb.Z - o.Z) * (jb.Z - o.Z)
	return int(x_diff + y_diff + z_diff)
}

type BoxPair struct {
	box1, box2 *JunctionBox
	sqDist     int
}

// Checks if two boxpairs are connected together (e.g., both contain a shared Junctionbox)
func (p *BoxPair) Connected(o *BoxPair) bool {
	return o.box1 == p.box1 || o.box1 == p.box2 || o.box2 == p.box1 || o.box2 == p.box2
}

// Constructs a new BoxPair from JunctionBoxes
func NewBoxPair(box1 *JunctionBox, box2 *JunctionBox) *BoxPair {
	return &BoxPair{box1, box2, box1.squareDistance(box2)}
}

// A circuit is a bunch of junction boxes connected together.
type Circuit struct {
	boxes map[*JunctionBox]bool
}

// Creates a new circuit containing an original pair of JB's.
func NewCircuit(pair *BoxPair) *Circuit {
	boxes := make(map[*JunctionBox]bool)
	boxes[pair.box1] = true
	boxes[pair.box2] = true
	return &Circuit{boxes}
}

// Inserts a box pair into the circuit (only if it can be).
func (c *Circuit) Insert(pair *BoxPair) bool {
	if _, ok := c.boxes[pair.box1]; ok {
		c.boxes[pair.box2] = true
		return true
	}
	if _, ok := c.boxes[pair.box2]; ok {
		c.boxes[pair.box1] = true
		return true
	}
	return false
}

// Extends a circuit with another circuit.
func (c *Circuit) Extend(o *Circuit) {
	for jb := range o.boxes {
		c.boxes[jb] = true
	}
}

// Returns the number of elements in a circuit.
func (c Circuit) Size() int { return len(c.boxes) }

// A min-max heap struct for our junction boxes; capped at a size of 1000
// and removing the largest elements as we attempt to expand it beyond 1000
// elements.
type MinMaxHeap struct {
	slice   []*BoxPair
	maxSize int
}

// Creates a new MinMaxHeap struct.
func NewMinMaxHeap(maxSize int) *MinMaxHeap {
	return &MinMaxHeap{[]*BoxPair{}, maxSize}
}

// Swaps the elements at positions i and m.
func (h *MinMaxHeap) swap(i int, m int) {
	h.slice[i], h.slice[m] = h.slice[m], h.slice[i]
}

// Returns if the level we're on is a min-level or not.
func (h MinMaxHeap) isMin(i int) bool {
	level := bits.Len(uint(i+1)) - 1
	return level%2 == 0
}

// Returns the children of index i.
func (h MinMaxHeap) children(i int) []int {
	children := []int{i*2 + 1, i*2 + 2}
	idx := 2
	for i, child := range children {
		if child >= h.Size() {
			idx = i
			break
		}
	}
	return children[:idx]
}

// Returns the grandchildren of index i.
func (h MinMaxHeap) grandchildren(i int) []int {
	grandchildren := []int{}
	children := h.children(i)
	for _, child := range children {
		grandchildren = append(grandchildren, h.children(child)...)
	}
	return grandchildren
}

// Returns the parent index of an index; -1 if it is the root.
func (h MinMaxHeap) parent(i int) int {
	if i == 0 {
		return -1
	}
	return (i - 1) / 2
}

// Returns the grandparent of an index; -1 if it has none.
func (h MinMaxHeap) grandparent(i int) int {
	parent := h.parent(i)
	if parent == -1 || parent == 0 {
		return -1
	}
	return h.parent(parent)
}

// Given an array of indicies; this returns the index in our min-max heap that
// has the smallest value. If the array is empty; it returns -1.
func (h MinMaxHeap) min(indicies []int) int {
	if len(indicies) == 0 {
		return -1
	}
	minIndex := indicies[0]
	if len(indicies) > 1 {
		for _, i := range indicies[1:] {
			if h.slice[i].sqDist < h.slice[minIndex].sqDist {
				minIndex = i
			}
		}
	}
	return minIndex
}

// Given an array of indicies; this returns the index in our min-max heap that
// has the smallest value. If the array is empty; it returns -1.
func (h MinMaxHeap) max(indicies []int) int {
	if len(indicies) == 0 {
		return -1
	}
	maxIndex := indicies[0]
	if len(indicies) > 1 {
		for _, i := range indicies[1:] {
			if h.slice[i].sqDist > h.slice[maxIndex].sqDist {
				maxIndex = i
			}
		}
	}
	return maxIndex
}

// Min-Max heap "push down" or "heapify" function. Implemented iteratively
// to save space complexity.
func (h *MinMaxHeap) pushDn(i int) {
	m := i
	children := h.children(i)
	for len(children) > 0 {
		i = m
		if h.isMin(i) {
			m = h.min(children)
			smallestGrandchild := h.min(h.grandchildren(i))
			isGrandchild := smallestGrandchild != -1 && h.slice[smallestGrandchild].sqDist < h.slice[m].sqDist
			if isGrandchild {
				m = smallestGrandchild
			}
			if h.slice[m].sqDist < h.slice[i].sqDist {
				h.swap(m, i)
				if isGrandchild && h.slice[m].sqDist > h.slice[h.parent(m)].sqDist {
					h.swap(m, h.parent(m))
				}
			} else {
				break
			}
		} else {
			m = h.max(children)
			largestGrandchild := h.max(h.grandchildren(i))
			isGrandchild := largestGrandchild != -1 && h.slice[largestGrandchild].sqDist > h.slice[m].sqDist
			if isGrandchild {
				m = largestGrandchild
			}
			if h.slice[m].sqDist > h.slice[i].sqDist {
				h.swap(m, i)
				if isGrandchild && h.slice[m].sqDist < h.slice[h.parent(m)].sqDist {
					h.swap(m, h.parent(m))
				}
			} else {
				break
			}
		}
		children = h.children(m)
	}
}

// Internal "push up" or "bubble up" function for the Min-Max heap.
func (h *MinMaxHeap) pushUp(i int) {
	pushUpMin := func(i int) {
		grandparent := h.grandparent(i)
		for grandparent != -1 && h.slice[i].sqDist < h.slice[grandparent].sqDist {
			h.swap(i, grandparent)
			i = grandparent
			grandparent = h.grandparent(i)
		}
	}
	pushUpMax := func(i int) {
		grandparent := h.grandparent(i)
		for grandparent != -1 && h.slice[i].sqDist > h.slice[grandparent].sqDist {
			h.swap(i, grandparent)
			i = grandparent
			grandparent = h.grandparent(i)
		}
	}
	if i != 0 {
		parent := h.parent(i)
		if h.isMin(i) {
			if h.slice[i].sqDist > h.slice[parent].sqDist {
				h.swap(i, parent)
				pushUpMax(parent)
			} else {
				pushUpMin(i)
			}
		} else {
			if h.slice[i].sqDist < h.slice[parent].sqDist {
				h.swap(i, parent)
				pushUpMin(parent)
			} else {
				pushUpMax(i)
			}
		}
	}
}

// Returns the size of the heap.
func (h MinMaxHeap) Size() int { return len(h.slice) }

// Min-Max heap insertion function; but modified to allow for two addtional properties:
//  1. Removes & destroys the largest element when the cardinality of the heap exceeds
//     the maxSize value of the heap.
//  2. Does not insert elements that exceed the size of the largest element when the heap
//     is full; instead return a boolean informing the user if it was inserted or not.
func (h *MinMaxHeap) Insert(pair *BoxPair) bool {
	if h.Size() == h.maxSize {
		max := h.Max()
		if max == nil || pair.sqDist >= max.sqDist {
			return false
		}
		h.PopMax()
	}
	h.slice = append(h.slice, pair)
	h.pushUp(h.Size() - 1)
	return true
}

// Returns the minimum BoxPair on the heap without removing it.
func (h MinMaxHeap) Min() *BoxPair {
	if h.Size() == 0 {
		return nil
	}
	return h.slice[0]
}

// Returns the maximum Boxpair on the heap without removing it.
func (h MinMaxHeap) Max() *BoxPair {
	if h.Size() == 0 {
		return nil
	} else if h.Size() == 1 {
		return h.slice[0]
	}
	children := h.children(0)
	max := children[0]
	if len(children) > 1 && h.slice[children[1]].sqDist > h.slice[max].sqDist {
		max = children[1]
	}
	return h.slice[max]
}

// Removes and returns the minimum value from the
// heap.
func (h *MinMaxHeap) PopMin() *BoxPair {
	if h.Size() == 0 {
		return nil
	} else if h.Size() == 1 {
		min := h.slice[0]
		h.slice = h.slice[:0]
		return min
	}
	min := h.Min()
	h.swap(0, h.Size()-1)
	h.slice = h.slice[:h.Size()-1]
	h.pushDn(0)
	return min
}

// Removes and returns the maximum value from the heap.
func (h *MinMaxHeap) PopMax() *BoxPair {
	if h.Size() == 0 {
		return nil
	} else if h.Size() == 1 {
		min := h.slice[0]
		h.slice = h.slice[:0]
		return min
	}
	children := h.children(0)
	idx := children[0]
	if len(children) > 1 && h.slice[children[1]].sqDist > h.slice[idx].sqDist {
		idx = children[1]
	}
	max := h.slice[idx]
	h.swap(idx, h.Size()-1)
	h.slice = h.slice[:h.Size()-1]
	h.pushDn(idx)
	return max
}
