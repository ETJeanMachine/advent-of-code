package solutions

import (
	"fmt"
	"log"
	"math"
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
)

func parseInput2(input string) [][]int {
	var parsed [][]int
	lines := strings.SplitSeq(input, ",")
	for line := range lines {
		split := strings.Split(line, "-")
		start, err_1 := strconv.Atoi(split[0])
		end, err_2 := strconv.Atoi(split[1])
		if err_1 != nil || err_2 != nil {
			log.Fatal("Error in parsing input!")
		}
		parsed = append(parsed, []int{start, end})
	}
	return parsed
}

func invalidInRange(start int, end int, n int) []int {
	half_to_invalid := func(chunk int) int {
		invalid_str := strconv.Itoa(chunk)
		for range n - 1 {
			invalid_str = fmt.Sprintf("%s%d", invalid_str, chunk)
		}
		invalid, err := strconv.Atoi(invalid_str)
		if err != nil {
			log.Fatal("Fatal error!")
		}
		return invalid
	}

	var invalids []int
	var curr_chunk int
	start_str := strconv.Itoa(start)
	if len(start_str)%n != 0 {
		curr_chunk = int(math.Pow10(len(start_str) / n))
	} else {
		var err error
		curr_chunk, err = strconv.Atoi(start_str[0 : len(start_str)/n])
		if err != nil {
			log.Fatal("Fatal error!")
		}
	}
	curr_invalid := half_to_invalid(curr_chunk)
	for curr_invalid <= end {
		if curr_invalid >= start {
			invalids = append(invalids, curr_invalid)
		}
		curr_chunk += 1
		curr_invalid = half_to_invalid(curr_chunk)
	}
	return invalids
}

func day2part2(input string) string {
	parsed := parseInput2(input)
	invalid_sum := 0
	for _, id_range := range parsed {
		start, end := id_range[0], id_range[1]
		invalid_ids := mapset.NewSet[int]()
		for n := 2; n <= len(strconv.Itoa(end)); n++ {
			invalid_ids.Append(invalidInRange(start, end, n)...)
		}
		for _, id := range invalid_ids.ToSlice() {
			invalid_sum += id
		}
	}
	return strconv.Itoa(invalid_sum)
}

func day2part1(input string) string {
	parsed := parseInput2(input)
	invalid_sum := 0
	for _, id_range := range parsed {
		start, end := id_range[0], id_range[1]
		invalid_ids := invalidInRange(start, end, 2)
		for _, id := range invalid_ids {
			invalid_sum += id
		}
	}
	return strconv.Itoa(invalid_sum)
}

func Day2() (func(string) string, func(string) string) {
	return day2part1, day2part2
}
