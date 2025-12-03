package solutions

import (
	"fmt"
	"strconv"
	"strings"
)

func parseInput3(input string) [][]int {
	var batteries [][]int
	lines := strings.SplitSeq(input, "\n")
	for line := range lines {
		var bank []int
		for _, c := range line {
			battery, _ := strconv.Atoi(fmt.Sprintf("%c", c))
			bank = append(bank, battery)
		}
		batteries = append(batteries, bank)
	}
	return batteries
}

func maximumJoltage(bank []int, n int) int {
	max_battery := 0
	var max_battery_idx int
	for i := 0; i <= len(bank)-n; i++ {
		if max_battery < bank[i] {
			max_battery = bank[i]
			max_battery_idx = i
		}
	}
	if n == 1 {
		return max_battery
	}
	joltage_str := fmt.Sprintf("%d%d", max_battery, maximumJoltage(bank[max_battery_idx+1:], n-1))
	joltage, _ := strconv.Atoi(joltage_str)
	return joltage
}

func day3PartOne(input string) string {
	batteries := parseInput3(input)
	total_joltage := 0
	for _, bank := range batteries {
		total_joltage += maximumJoltage(bank, 2)
	}
	return strconv.Itoa(total_joltage)
}

func day3PartTwo(input string) string {
	batteries := parseInput3(input)
	total_joltage := 0
	for _, bank := range batteries {
		total_joltage += maximumJoltage(bank, 12)
	}
	return strconv.Itoa(total_joltage)
}

func Day3() (func(string) string, func(string) string) {
	return day3PartOne, day3PartTwo
}
