package day5

import (
	"strconv"
	"strings"
)

type freshChecker struct {
	ranges [][2]int
}

func (self *freshChecker) addRange(start int, end int) {
	self.ranges = append(self.ranges, [2]int{start, end})
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

func parseInput(input string) (freshChecker, []int) {
	var freshChecker freshChecker
	var ids []int
	split := strings.Split(input, "\n\n")
	range_strs, ing_strs := strings.Split(split[0], "\n"), strings.Split(split[1], "\n")
	for _, range_str := range range_strs {
		fresh_ids := strings.Split(range_str, "-")
		start, _ := strconv.Atoi(fresh_ids[0])
		end, _ := strconv.Atoi(fresh_ids[1])
		freshChecker.addRange(start, end)
	}
	for _, ing_str := range ing_strs {
		ingredient, _ := strconv.Atoi(ing_str)
		ids = append(ids, ingredient)
	}
	return freshChecker, ids
}

func partOne(input string) string {
	freshChecker, ids := parseInput(input)
	freshCount := 0
	for _, ingredient := range ids {
		if freshChecker.isFresh(ingredient) {
			freshCount += 1
		}
	}
	return strconv.Itoa(freshCount)
}

func partTwo(input string) string {
	return "0"
}

func Puzzles() (func(string) string, func(string) string) {
	return partOne, partTwo
}
