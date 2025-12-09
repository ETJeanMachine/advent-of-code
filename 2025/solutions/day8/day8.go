package day8

import (
	"regexp"
	"slices"
	"strconv"
)

func updateCircuits(circuits []*Circuit, pair *BoxPair) []*Circuit {
	mergeWith := []int{}
	for i := range circuits {
		if circuits[i].Insert(pair) {
			mergeWith = append(mergeWith, i)
		}
	}
	merged := NewCircuit(pair)
	shrunk := []*Circuit{}
	prev := 0
	for _, idx := range mergeWith {
		merged.Extend(circuits[idx])
		shrunk = append(shrunk, circuits[prev:idx]...)
		prev = idx + 1
	}
	shrunk = append(shrunk, circuits[prev:]...)
	shrunk = append(shrunk, merged)
	return shrunk
}

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
	circuits := []*Circuit{}
	min := heap.PopMin()
	for min != nil {
		circuits = updateCircuits(circuits, min)
		min = heap.PopMin()
	}
	slices.SortFunc(circuits, func(a, b *Circuit) int {
		return b.Size() - a.Size()
	})
	res := 1
	numCircuits := slices.Min([]int{3, len(circuits)})
	for _, circuit := range circuits[:numCircuits] {
		res *= circuit.Size()
	}
	return strconv.Itoa(res)
}

func partTwo(input string) string {
	// input = "162,817,812\n57,618,57\n906,360,560\n592,479,940\n352,342,300\n466,668,158\n542,29,236\n431,825,988\n739,650,466\n52,470,668\n216,146,977\n819,987,18\n117,168,530\n805,96,715\n346,949,466\n970,615,88\n941,993,340\n862,61,35\n984,92,344\n425,690,689"
	boxes := parseInput(input)
	heap := NewMinMaxHeap(-1)
	for idx, box1 := range boxes[:len(boxes)-1] {
		for _, box2 := range boxes[idx+1:] {
			pair := NewBoxPair(box1, box2)
			heap.Insert(pair)
		}
	}
	circuits := []*Circuit{}
	var min *BoxPair
	for !(len(circuits) == 1 && circuits[0].Size() == len(boxes)) {
		min = heap.PopMin()
		circuits = updateCircuits(circuits, min)
	}
	res := min.box1.X * min.box2.X
	return strconv.Itoa(res)
}

func Puzzles() (func(string) string, func(string) string) {
	return partOne, partTwo
}
