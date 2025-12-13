package day9

import (
	"regexp"
	"strconv"
)

type Point struct{ x, y int }

type Rectangle struct {
	topLeft, botRight Point
}

func NewRectangle(p1 Point, p2 Point) Rectangle {
	minX, maxX := p1.x, p2.x
	if p2.x < p1.x {
		minX, maxX = p2.x, p1.x
	}
	minY, maxY := p1.y, p2.y
	if p2.y < p1.y {
		minY, maxY = p2.y, p1.y
	}
	topLeft := Point{minX, minY}
	botRight := Point{maxX, maxY}
	return Rectangle{topLeft, botRight}
}

// Returns the points for the rectangle, ordered top-bottom, left-right.
func (r Rectangle) Points() [4]Point {
	botLeft := Point{r.topLeft.x, r.botRight.y}
	topRight := Point{r.botRight.x, r.topLeft.y}
	return [4]Point{r.topLeft, topRight, botLeft, r.botRight}
}

func (r Rectangle) ContainsPoint(p Point) bool {
	fitsWidth := r.topLeft.x <= p.x && p.x <= r.botRight.x
	fitsHeight := r.topLeft.y <= p.y && p.y <= r.botRight.y
	return fitsHeight && fitsWidth
}

func (r Rectangle) Area() int {
	width := (r.botRight.x - r.topLeft.x) + 1
	height := (r.botRight.y - r.topLeft.y) + 1
	return width * height
}

type Shape struct{ rectangles []Rectangle }

func NewShape(points []Point) Shape {
	type vertline struct{ topY, botY int }

	return Shape{}
}

func (s Shape) ContainsRectangle(rect Rectangle) bool {
	points := rect.Points()
	for _, r := range s.rectangles {
		containsAll := true
		for _, p := range points {
			containsAll = containsAll && r.ContainsPoint(p)
		}
		if containsAll {
			return true
		}
	}
	return false
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
	for idx, p1 := range points[:len(points)-1] {
		for _, p2 := range points[idx+1:] {
			rect := NewRectangle(p1, p2)
			area := rect.Area()
			if area > max_area {
				max_area = area
			}
		}
	}
	return strconv.Itoa(max_area)
}

func partTwo(input string) string {
	input = "7,1\n11,1\n11,7\n9,7\n9,5\n2,5\n2,3\n7,3"
	points := parseInput(input)
	shape := NewShape(points)
	var max_area = 0
	for idx, p := range points[:len(points)-1] {
		for _, o := range points[idx+1:] {
			rect := NewRectangle(p, o)
			area := rect.Area()
			if area > max_area && shape.ContainsRectangle(rect) {
				max_area = area
			}
		}
	}
	return strconv.Itoa(max_area)
}

func Puzzles() (func(string) string, func(string) string) {
	return partOne, partTwo
}
