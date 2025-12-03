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

func maximumJoltage(bank []int) int {
	battery_one := 0
	var battery_one_idx int
	for i := 0; i < len(bank)-1; i++ {
		if battery_one < bank[i] {
			battery_one = bank[i]
			battery_one_idx = i
		}
	}
	battery_two := 0
	for i := battery_one_idx + 1; i < len(bank); i++ {
		if battery_two < bank[i] {
			battery_two = bank[i]
		}
	}
	return (battery_one * 10) + battery_two
}

func day3PartOne(input string) string {
	// input = "987654321111111\n811111111111119\n234234234234278\n818181911112111"
	batteries := parseInput3(input)
	total_joltage := 0
	for _, bank := range batteries {
		total_joltage += maximumJoltage(bank)
	}
	return strconv.Itoa(total_joltage)
}

func day3PartTwo(input string) string {
	return ""
}

func Day3() (func(string) string, func(string) string) {
	return day3PartOne, day3PartTwo
}
