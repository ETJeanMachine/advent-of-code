package day5

import (
	"fmt"
	"slices"

	mapset "github.com/deckarep/golang-set/v2"
)

// Struct that holds overlapping intervals sorted by start and by end.
type overlaps struct {
	byStart [][2]int
	byEnd   [][2]int
}

func newOverlaps(center_intervals [][2]int) *overlaps {
	if len(center_intervals) == 0 {
		return nil
	}
	byStart := make([][2]int, len(center_intervals))
	copy(byStart, center_intervals)
	slices.SortFunc(byStart, func(a, b [2]int) int {
		return a[0] - b[0]
	})
	byEnd := make([][2]int, len(center_intervals))
	copy(byEnd, center_intervals)
	slices.SortFunc(byEnd, func(a, b [2]int) int {
		return b[1] - a[1]
	})
	return &overlaps{byStart, byEnd}
}

type Node struct {
	intervals *overlaps // all intervals that overlap the center point of the node
	center    int       // the center point
	left      *Node     // the left node
	right     *Node     // the right node
}

func (n *Node) iterByStart() [][2]int {
	if n.intervals == nil {
		return [][2]int{}
	}
	return n.intervals.byStart
}

func (n *Node) iterByEnd() [][2]int {
	if n.intervals == nil {
		return [][2]int{}
	}
	return n.intervals.byEnd
}

// finds all intervals that a value falls within
func (n *Node) overlaps(value int, overlaps [][2]int) [][2]int {
	if n == nil {
		return overlaps
	} else if value < n.center {
		for _, interval := range n.iterByStart() {
			if interval[0] <= value {
				overlaps = append(overlaps, interval)
			} else {
				break
			}
		}
		return n.left.overlaps(value, overlaps)
	} else if value > n.center {
		for _, interval := range n.iterByEnd() {
			if interval[1] >= value {
				overlaps = append(overlaps, interval)
			} else {
				break
			}
		}
		return n.right.overlaps(value, overlaps)
	}
	return append(overlaps, n.iterByStart()...)
}

func (n *Node) NodeString() string {
	return fmt.Sprintf("%.2f", float32(n.center)/1.0e13)
}

func (n *Node) PrettyPrint(prefix string, isLeft bool) {
	if n == nil {
		return
	}

	fmt.Print(prefix)
	if isLeft {
		fmt.Print("├── ")
	} else {
		fmt.Print("└── ")
	}
	fmt.Println(n.NodeString())

	extension := prefix
	if isLeft {
		extension += "│   "
	} else {
		extension += "    "
	}

	if n.left != nil || n.right != nil {
		if n.left != nil {
			n.left.PrettyPrint(extension, true)
		} else {
			fmt.Print(extension + "├── ")
			fmt.Println("nil")
		}

		if n.right != nil {
			n.right.PrettyPrint(extension, false)
		} else {
			fmt.Print(extension + "└── ")
			fmt.Println("nil")
		}
	}
}

func (it *IntervalTree) Print() {
	if it.root == nil {
		fmt.Println("Empty tree")
		return
	}
	fmt.Printf("%s\n", it.root.NodeString())
	if it.root.left != nil || it.root.right != nil {
		if it.root.left != nil {
			it.root.left.PrettyPrint("", true)
		} else {
			fmt.Println("├── nil")
		}
		if it.root.right != nil {
			it.root.right.PrettyPrint("", false)
		} else {
			fmt.Println("└── nil")
		}
	}
}

type IntervalTree struct {
	root *Node
}

// public function for finding all intervals a value falls within
func (it *IntervalTree) Intervals(value int) [][2]int {
	return it.root.overlaps(value, [][2]int{})
}

func (it *IntervalTree) IntervalSpan() int {
	uniqueIntervals := mapset.NewSet[[2]int]()
	var findUnique func(n *Node)
	findUnique = func(n *Node) {

	}
	findUnique(it.root)
	spanTotal := 0
	for interval := range uniqueIntervals.Iter() {
		spanTotal += (interval[1] - interval[0]) + 1
	}
	return spanTotal
}

func NewTree(all_intervals [][2]int) *IntervalTree {
	var addIntervals func(intervals [][2]int) *Node
	addIntervals = func(intervals [][2]int) *Node {
		if len(intervals) == 0 {
			return nil
		}
		var min, max int
		for idx := range intervals {
			if idx == 0 {
				min, max = intervals[idx][0], intervals[idx][1]
			}
			if intervals[idx][0] < min {
				min = intervals[idx][0]
			}
			if max < intervals[idx][1] {
				max = intervals[idx][1]
			}
		}
		center := (min + max) / 2
		left_intervals, right_intervals := [][2]int{}, [][2]int{}
		center_intervals := [][2]int{}
		for idx := range intervals {
			start, end := intervals[idx][0], intervals[idx][1]
			if start <= center && center <= end {
				center_intervals = append(center_intervals, intervals[idx])
			} else if end < center {
				left_intervals = append(left_intervals, intervals[idx])
			} else if center < start {
				right_intervals = append(right_intervals, intervals[idx])
			}
		}

		return &Node{
			intervals: newOverlaps(center_intervals),
			center:    center,
			left:      addIntervals(left_intervals),
			right:     addIntervals(right_intervals),
		}
	}
	root := addIntervals(all_intervals)
	return &IntervalTree{root}
}
