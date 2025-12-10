package day9

import (
	"regexp"
	"strconv"
)

type Point struct{ x, y int }

func (p Point) Area(o Point) int {
	width := p.x - o.x
	if width < 0 {
		width = -width
	}
	height := p.y - o.y
	if height < 0 {
		height = -height
	}
	area := (width + 1) * (height + 1)
	return area
}

func parseInput(input string) []Point {
	re := regexp.MustCompile(`(\d+),(\d+)`)
	matches := re.FindAllStringSubmatch(input, -1)
	points := make([]Point, len(matches))
	for idx, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		points[idx] = Point{x, y}
	}
	return points
}

func partOne(input string) string {
	points := parseInput(input)
	var max_area = 0
	for idx, p := range points[:len(points)-1] {
		for _, o := range points[idx+1:] {
			if p.Area(o) > max_area {
				max_area = p.Area(o)
			}
		}
	}
	return strconv.Itoa(max_area)
}

func partTwo(input string) string {
	return "0"
}

func Puzzles() (func(string) string, func(string) string) {
	return partOne, partTwo
}
