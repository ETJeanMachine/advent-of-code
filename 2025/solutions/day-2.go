package solutions

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func parseInput2(input string) [][]int {
	var parsed [][]int
	lines := strings.SplitSeq(input, ",")
	for line := range lines {
		split := strings.Split(line, "-")
		start, _ := strconv.Atoi(split[0])
		end, _ := strconv.Atoi(split[1])
		parsed = append(parsed, []int{start, end})
	}
	return parsed
}

func day2part2(input string) string {
	return ""
}
func doubleEndedInRange(start int, end int) []int {
	half_to_invalid := func(n int) int {
		invalid, _ := strconv.Atoi(fmt.Sprintf("%d%d", n, n))
		return invalid
	}

	var invalids []int
	var curr_half int
	start_str := strconv.Itoa(start)
	if len(start_str)%2 == 1 {
		curr_half = int(math.Pow10(len(start_str) / 2))
	} else {
		curr_half, _ = strconv.Atoi(start_str[0 : len(start_str)/2])
	}
	curr_invalid := half_to_invalid(curr_half)
	for curr_invalid <= end {
		if curr_invalid >= start {
			invalids = append(invalids, curr_invalid)
		}
		curr_half += 1
		curr_invalid = half_to_invalid(curr_half)
	}
	return invalids
}

func day2part1(input string) string {
	parsed := parseInput2(input)
	invalid_sum := 0
	for _, id_range := range parsed {
		start, end := id_range[0], id_range[1]
		invalid_ids := doubleEndedInRange(start, end)
		for _, id := range invalid_ids {
			invalid_sum += id
		}
	}
	return strconv.Itoa(invalid_sum)
}

func Day2() (func(string) string, func(string) string) {
	return day2part1, day2part2
}
