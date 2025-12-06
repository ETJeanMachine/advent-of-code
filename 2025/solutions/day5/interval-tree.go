package day5

import "fmt"

type IntervalNode struct {
	left     *IntervalNode
	right    *IntervalNode
	interval [2]int
	center   int
}

func (n *IntervalNode) NodeString() string {
	return fmt.Sprintf("%d", n.center)
}

func (n *IntervalNode) PrettyPrint(prefix string, isLeft bool) {
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
	root *IntervalNode
}

func (it *IntervalTree) InRange(value int) bool {
	var inNodeRange func(node *IntervalNode) bool
	inNodeRange = func(node *IntervalNode) bool {
		if node == nil {
			return false
		}
		start, end := node.interval[0], node.interval[1]
		if start <= value && value <= end {
			return true
		} else if value < start {
			return inNodeRange(node.left)
		}
		return inNodeRange(node.right)
	}
	return inNodeRange(it.root)
}

func (it *IntervalTree) IntervalSpan() int {
	var spanTotal func(node *IntervalNode, parents []*IntervalNode) int
	spanTotal = func(node *IntervalNode, parents []*IntervalNode) int {
		if node == nil {
			return 0
		}
		start, end := node.interval[0], node.interval[1]
		span := (end - start) + 1
		for _, parentNode := range parents {
			if parentNode != nil {
				p_start, p_end := parentNode.interval[0], parentNode.interval[1]
				if (p_start <= start && start <= p_end) && (p_start <= end && end <= p_end) {
					span = 0
					break
				} else if p_start <= start && start <= p_end {
					span -= (p_end - start) + 1
					break
				} else if p_start <= end && end <= p_end {
					span -= (end - p_start) + 1
					break
				}
			}
		}
		parents = append(parents, node)
		return span +
			spanTotal(node.left, parents) +
			spanTotal(node.right, parents)
	}
	return spanTotal(it.root, []*IntervalNode{})
}

func NewTree(all_intervals [][2]int) *IntervalTree {
	var addIntervals func(intervals [][2]int) *IntervalNode
	addIntervals = func(intervals [][2]int) *IntervalNode {
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
		start, end := center, center
		for idx := range intervals {
			other_start, other_end := intervals[idx][0], intervals[idx][1]
			if other_start <= center && center <= other_end {
				if other_start < start {
					start = other_start
				}
				if end < other_end {
					end = other_end
				}
			} else if other_end < center {
				left_intervals = append(left_intervals, intervals[idx])
			} else if center < other_start {
				right_intervals = append(right_intervals, intervals[idx])
			}
		}
		return &IntervalNode{
			left:     addIntervals(left_intervals),
			right:    addIntervals(right_intervals),
			interval: [2]int{start, end},
			center:   center,
		}
	}
	root := addIntervals(all_intervals)
	return &IntervalTree{root}
}
