package day5

import (
	"errors"
	"fmt"
)

type Node struct {
	overlapping [][2]int // all intervals that overlap the center point of the node
	center      int      // the center point
	left        *Node    // the left node
	right       *Node    // the right node
}

func (n *Node) inInterval(value int) bool {
	if n == nil {
		return false
	}
	inOverlapping := false
	for idx := range n.overlapping {
		start, end := n.overlapping[idx][0], n.overlapping[idx][1]
		if start <= value && value <= end {
			inOverlapping = true
			break
		}
	}
	if inOverlapping {
		return true
	} else if value < n.center {
		return n.left.inInterval(value)
	}
	return n.right.inInterval(value)
}

func (n *Node) fullInterval() (int, int, error) {
	if len(n.overlapping) == 0 {
		return 0, 0, errors.New("Empty overlapping set.")
	}
	min, max := n.overlapping[0][0], n.overlapping[0][1]
	for idx := 1; idx < len(n.overlapping); idx++ {
		if n.overlapping[idx][0] < min {
			min = n.overlapping[idx][0]
		}
		if n.overlapping[idx][1] > max {
			max = n.overlapping[idx][1]
		}
	}
	return min, max, nil
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

func (it *IntervalTree) InInterval(value int) bool {
	return it.root.inInterval(value)
}

func (it *IntervalTree) IntervalSpan() int {
	var spanTotal func(n *Node) int
	spanTotal = func(n *Node) int {
		if n == nil {
			return 0
		}
		start, end, err := n.fullInterval()
		var span int
		if err != nil {
			span = 0
		} else {
			span = (end - start) + 1
		}
		if n.left != nil {
			_, l_end, err := n.left.fullInterval()
			if err == nil && l_end >= start {
				span -= (l_end - start) + 1
			}
		}
		if n.right != nil {
			r_start, _, err := n.right.fullInterval()
			if err == nil && r_start <= end {
				span -= (end - r_start) + 1
			}
		}
		return span + spanTotal(n.left) + spanTotal(n.right)
	}
	return spanTotal(it.root)
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
		overlapping := [][2]int{}
		for idx := range intervals {
			start, end := intervals[idx][0], intervals[idx][1]
			if start <= center && center <= end {
				overlapping = append(overlapping, intervals[idx])
			} else if end < center {
				left_intervals = append(left_intervals, intervals[idx])
			} else if center < start {
				right_intervals = append(right_intervals, intervals[idx])
			}
		}

		return &Node{
			overlapping: overlapping,
			center:      center,
			left:        addIntervals(left_intervals),
			right:       addIntervals(right_intervals),
		}
	}
	root := addIntervals(all_intervals)
	return &IntervalTree{root}
}
