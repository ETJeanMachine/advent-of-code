package day4

import (
	"strconv"
	"strings"

	mapset "github.com/deckarep/golang-set/v2"
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

func (f *floor) isMoveable(row int, col int) bool {
	return f.hasRoll(row, col) && f.adjacentRollCount(row, col) < 4
}

func (f *floor) adjacentRolls(row int, col int) [][2]int {
	var adjacent_rolls [][2]int
	for curr_col := col - 1; curr_col <= col+1; curr_col++ {
		if f.hasRoll(row-1, curr_col) {
			adjacent_rolls = append(adjacent_rolls, [2]int{row - 1, curr_col})
		}
		if curr_col != col && f.hasRoll(row, curr_col) {
			adjacent_rolls = append(adjacent_rolls, [2]int{row, curr_col})
		}
		if f.hasRoll(row+1, curr_col) {
			adjacent_rolls = append(adjacent_rolls, [2]int{row + 1, curr_col})
		}
	}
	return adjacent_rolls
}

func (f *floor) removeableRolls() [][2]int {
	var moveable_rolls [][2]int
	for row := range f.height {
		for col := range f.width {
			if f.isMoveable(row, col) {
				moveable_rolls = append(moveable_rolls, [2]int{row, col})
			}
		}
	}
	return moveable_rolls
}

func (f *floor) clearRolls() int {
	visited := mapset.NewSet[[2]int]()
	removeableRolls := f.removeableRolls()
	var dfs_helper func(curr_roll [2]int) int
	dfs_helper = func(curr_roll [2]int) int {
		visited.Add(curr_roll)
		row, col := curr_roll[0], curr_roll[1]
		f.floor_map[row][col] = false
		removed := 1
		for _, adj_roll := range f.adjacentRolls(row, col) {
			adj_row, adj_col := adj_roll[0], adj_roll[1]
			if !visited.Contains(adj_roll) && f.isMoveable(adj_row, adj_col) {
				removed += dfs_helper(adj_roll)
			}
		}
		return removed
	}

	total_removed := 0
	for _, roll := range removeableRolls {
		if !visited.Contains(roll) {
			total_removed += dfs_helper(roll)
		}
	}

	return total_removed
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
	moveable_count := len(floor.removeableRolls())
	return strconv.Itoa(moveable_count)
}

func partTwo(input string) string {
	floor := parseInput(input)
	removed_count := floor.clearRolls()
	return strconv.Itoa(removed_count)
}

func Puzzles() (func(string) string, func(string) string) {
	return partOne, partTwo
}
