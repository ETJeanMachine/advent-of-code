package day7

import (
	"strconv"
	"strings"
)

func parseInput(input string) *TachyonManifold {
	lines := strings.Split(input, "\n")
	splitters := [][]bool{}
	var init TachyonBeam
	for row, line := range lines {
		splitters = append(splitters, []bool{})
		for col, r := range line {
			switch r {
			case 'S':
				init = TachyonBeam{row, col}
				splitters[row] = append(splitters[row], false)
			case '^':
				splitters[row] = append(splitters[row], true)
			default:
				splitters[row] = append(splitters[row], false)
			}
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
