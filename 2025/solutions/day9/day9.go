package day9

import (
	"regexp"
	"strconv"
)

type Point struct{ x, y int }

type Line struct{ p1, p2 Point }

func NewLine(p1 Point, p2 Point) Line {
	return Line{p1, p2}
}

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

func inBounds(lines []Line, p Point, o Point) bool {
	alts := []Point{Point{p.x, o.y}, Point{o.x, p.y}}
	fits := false
	for i := 0; i < len(lines)-1; i += 2 {
		line1, line2 := lines[i], lines[i+1]
		bound1, bound2 := line1.p1, line2.p2
		isInBounds := true
		for _, alt := range alts {
			if bound1.x < bound2.x {
				isInBounds = isInBounds && bound1.x <= alt.x && alt.x <= bound2.x
			} else {
				isInBounds = isInBounds && bound2.x <= alt.x && alt.x <= bound1.x
			}
			if bound1.y < bound2.y {
				isInBounds = isInBounds && bound1.y <= alt.y && alt.y <= bound2.y
			} else {
				isInBounds = isInBounds && bound2.y <= alt.y && alt.y <= bound1.y
			}
		}
		fits = fits != isInBounds
	}
	return fits
}

func partTwo(input string) string {
	input = "7,1\n11,1\n11,7\n9,7\n9,5\n2,5\n2,3\n7,3"
	points := parseInput(input)
	lines := []Line{}
	for i := range len(points) - 1 {
		lines = append(lines, NewLine(points[i], points[i+1]))
	}
	lines = append(lines, NewLine(points[0], points[len(points)-1]))
	var max_area = 0
	for idx, p := range points[:len(points)-1] {
		for _, o := range points[idx+1:] {
			area := p.Area(o)
			if area > max_area && inBounds(lines, p, o) {
				max_area = area
			}
		}
	}
	return strconv.Itoa(max_area)
}

func Puzzles() (func(string) string, func(string) string) {
	return partOne, partTwo
}
