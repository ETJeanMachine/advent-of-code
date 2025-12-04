package solutions

import (
	"advent-of-code/solutions/day1"
	"advent-of-code/solutions/day2"
	"advent-of-code/solutions/day3"
	"advent-of-code/solutions/day4"
	"log"
)

func GetPuzzles(day int) (func(input string) string, func(input string) string) {
	var part_one func(input string) string
	var part_two func(input string) string
	switch day {
	case 1:
		part_one, part_two = day1.Puzzles()
	case 2:
		part_one, part_two = day2.Puzzles()
	case 3:
		part_one, part_two = day3.Puzzles()
	case 4:
		part_one, part_two = day4.Puzzles()
	default:
		log.Fatalf("Day %d is not implemented.\n", day)
	}
	return part_one, part_two
}
