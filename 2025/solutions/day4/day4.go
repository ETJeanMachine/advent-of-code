package day4

import (
	"strconv"
	"strings"
)

type floor struct {
	height    int
	width     int
	floor_map [][]bool
}

func (f *floor) hasRoll(row int, col int) bool {
	if (row >= 0 && row < f.height) && (col >= 0 && col < f.width) {
		return f.floor_map[row][col]
	}
	return false
}

func (f *floor) adjacentRollCount(row int, col int) int {
	roll_count := 0
	for curr_col := col - 1; curr_col <= col+1; curr_col++ {
		if f.hasRoll(row-1, curr_col) {
			roll_count += 1
		}
		if curr_col != col && f.hasRoll(row, curr_col) {
			roll_count += 1
		}
		if f.hasRoll(row+1, curr_col) {
			roll_count += 1
		}
	}
	return roll_count
}

func parseInput(input string) floor {
	lines := strings.Split(input, "\n")
	var floor floor
	floor.height = len(lines)
	floor.width = len(lines[0])
	for _, line := range lines {
		var row []bool
		for _, c := range line {
			row = append(row, c == '@')
		}
		floor.floor_map = append(floor.floor_map, row)
	}
	return floor
}

func partOne(input string) string {
	floor := parseInput(input)
	moveable_count := 0
	for row := range floor.height {
		for col := range floor.width {
			if floor.hasRoll(row, col) && floor.adjacentRollCount(row, col) < 4 {
				moveable_count += 1
			}
		}
	}
	return strconv.Itoa(moveable_count)
}

func partTwo(input string) string {
	floor := parseInput(input)
	moveable_count := 1
	removed_count := 0
	for moveable_count > 0 {
		moveable_count = 0
		for row := range floor.height {
			for col := range floor.width {
				if floor.hasRoll(row, col) && floor.adjacentRollCount(row, col) < 4 {
					floor.floor_map[row][col] = false
					removed_count += 1
					moveable_count += 1
				}
			}
		}
	}
	return strconv.Itoa(removed_count)
}

func Puzzles() (func(string) string, func(string) string) {
	return partOne, partTwo
}
