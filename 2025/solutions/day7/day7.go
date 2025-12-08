package day7

import (
	"strconv"
	"strings"
)

func parseInput(input string) *TachyonManifold {
	lines := strings.Split(input, "\n")
	height, width := len(lines), len(lines[0])
	splitters := make([][]bool, height)
	var init TachyonBeam
	for row, line := range lines {
		splitters[row] = make([]bool, width)
		if row == 0 {
			col := strings.IndexRune(line, 'S')
			init = TachyonBeam{row, col}
			continue
		} else if row%2 == 1 {
			continue
		}
		col := strings.IndexRune(line, '^')
		for {
			splitters[row][col] = true
			to_next := strings.IndexRune(line[col+1:], '^')
			if to_next == -1 {
				break
			}
			col += to_next + 1
		}
	}
	return NewManifold(init, splitters)
}

func partOne(input string) string {
	manifold := parseInput(input)
	return strconv.Itoa(manifold.CountSplits())
}

func partTwo(input string) string {
	manifold := parseInput(input)
	return strconv.Itoa(manifold.CountTimelines())
}

func Puzzles() (func(string) string, func(string) string) {
	return partOne, partTwo
}
