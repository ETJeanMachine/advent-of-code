package day5

import (
	"strconv"
	"strings"
)

type freshChecker struct {
	ranges [][2]int
}

func (self *freshChecker) addRange(start int, end int) {
	start_idx := len(self.ranges)
	for idx := len(self.ranges) - 1; idx >= 0; idx-- {
		fresh_start := self.ranges[idx][0]
		if start <= fresh_start {
			start_idx = idx
			break
		}
	}
	end_idx := start_idx
	for idx := start_idx; idx < len(self.ranges); idx++ {
		fresh_end := self.ranges[idx][1]
		if fresh_end <= end {
			end_idx = idx
			break
		}
	}
	new_range := [2]int{start, end}
	front_copy := make([][2]int, len(self.ranges[:start_idx]))
	copy(front_copy, self.ranges[:start_idx])
	front_copy = append(front_copy, new_range)
	new_idx := len(front_copy) - 1
	self.ranges = append(front_copy, self.ranges[end_idx:]...)
	// fusing fronts/ends of equal adjacent ranges
	if new_idx != 0 && self.ranges[new_idx][0] <= self.ranges[new_idx-1][1] {
		fused_range := [2]int{self.ranges[new_idx-1][0], self.ranges[new_idx][1]}

		front_copy := make([][2]int, len(self.ranges[:new_idx-1]))
		copy(front_copy, self.ranges[:new_idx-1])
		front_copy = append(front_copy, fused_range)
		self.ranges = append(front_copy, self.ranges[new_idx+1:]...)
		new_idx = len(front_copy) - 1
	}
	if new_idx != len(self.ranges)-1 && self.ranges[new_idx][1] >= self.ranges[new_idx+1][0] {
		fused_range := [2]int{self.ranges[new_idx][0], self.ranges[new_idx+1][1]}

		front_copy := make([][2]int, len(self.ranges[:new_idx]))
		copy(front_copy, self.ranges[:new_idx])
		front_copy = append(front_copy, fused_range)
		self.ranges = append(front_copy, self.ranges[new_idx+2:]...)
	}
}

func (self *freshChecker) isFresh(ingredient int) bool {
	for _, fresh_range := range self.ranges {
		start, end := fresh_range[0], fresh_range[1]
		if start <= ingredient && ingredient <= end {
			return true
		}
	}
	return false
}

func (self *freshChecker) totalFresh() int {
	total := 0
	for _, fresh_range := range self.ranges {
		start, end := fresh_range[0], fresh_range[1]
		total += (end - start) + 1
	}
	return total
}

func fromString(str string) freshChecker {
	var freshChecker freshChecker
	range_strs := strings.SplitSeq(str, "\n")
	for range_str := range range_strs {
		fresh_ids := strings.Split(range_str, "-")
		start, _ := strconv.Atoi(fresh_ids[0])
		end, _ := strconv.Atoi(fresh_ids[1])
		freshChecker.addRange(start, end)
	}
	return freshChecker
}

func partOne(input string) string {
	split := strings.Split(input, "\n\n")
	freshChecker := fromString(split[0])
	var ids []int
	ing_strs := strings.SplitSeq(split[1], "\n")
	for ing_str := range ing_strs {
		ingredient, _ := strconv.Atoi(ing_str)
		ids = append(ids, ingredient)
	}
	freshCount := 0
	for _, ingredient := range ids {
		if freshChecker.isFresh(ingredient) {
			freshCount += 1
		}
	}
	return strconv.Itoa(freshCount)
}

func partTwo(input string) string {
	freshChecker := fromString(input)
	return strconv.Itoa(freshChecker.totalFresh())
}

func Puzzles() (func(string) string, func(string) string) {
	return partOne, partTwo
}
