package day8

import (
	"fmt"
	"regexp"
	"strconv"
)

func parseInput(input string) []*JunctionBox {
	re := regexp.MustCompile(`(\d+),(\d+),(\d+)`)
	matches := re.FindAllStringSubmatch(input, -1)
	boxes := make([]*JunctionBox, len(matches))
	for idx, match := range matches {
		x, _ := strconv.Atoi(match[1])
		y, _ := strconv.Atoi(match[2])
		z, _ := strconv.Atoi(match[3])
		boxes[idx] = &JunctionBox{x, y, z}
	}
	return boxes
}

func partOne(input string) string {
	boxes := parseInput(input)
	heap := NewMinMaxHeap(1000)
	for idx, box1 := range boxes[:len(boxes)-1] {
		for _, box2 := range boxes[idx+1:] {
			pair := NewBoxPair(box1, box2)
			heap.Insert(pair)
		}
	}
	min := heap.PopMin()
	for min != nil {
		fmt.Printf("Dist: %d, Len: %d\n", min.sqDist, heap.Size())
		min = heap.PopMin()
	}
	return strconv.Itoa(heap.Size())
}

func partTwo(input string) string {
	return "0"
}

func Puzzles() (func(string) string, func(string) string) {
	return partOne, partTwo
}
