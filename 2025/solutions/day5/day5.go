package day5

import (
	"strconv"
	"strings"
)

func parseValues(str string) []int {
	var values []int
	split := strings.SplitSeq(str, "\n")
	for line := range split {
		value, _ := strconv.Atoi(line)
		values = append(values, value)
	}
	return values
}

func parseIntervals(str string) [][2]int {
	var intervals [][2]int
	split := strings.SplitSeq(str, "\n")
	for line := range split {
		line_split := strings.Split(line, "-")
		start, _ := strconv.Atoi(line_split[0])
		end, _ := strconv.Atoi(line_split[1])
		intervals = append(intervals, [2]int{start, end})
	}
	return intervals
}

func partOne(input string) string {
	// input = "3-5\n10-14\n16-20\n12-18\n\n\n1\n5\n8\n11\n17\n32"
	split := strings.Split(input, "\n\n")
	intervals := parseIntervals(split[0])
	ingredients := parseValues(split[1])
	tree := NewTree(intervals)
	freshCount := 0
	for _, value := range ingredients {
		overlaps := tree.InIntervals(value)
		if len(overlaps) > 0 {
			freshCount += 1
		}
	}
	return strconv.Itoa(freshCount)
}

func partTwo(input string) string {
	// input = "3-5\n10-14\n16-20\n12-18"
	interval_str := strings.Split(input, "\n\n")[0]
	intervals := parseIntervals(interval_str)
	tree := NewTree(intervals)
	// tree.Print()
	return strconv.Itoa(tree.IntervalSpan())
}

func Puzzles() (func(string) string, func(string) string) {
	return partOne, partTwo
}
