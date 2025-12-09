package solutions

import (
	"advent-of-code/solutions/day1"
	"advent-of-code/solutions/day2"
	"advent-of-code/solutions/day3"
	"advent-of-code/solutions/day4"
	"advent-of-code/solutions/day5"
	"advent-of-code/solutions/day6"
	"advent-of-code/solutions/day7"
	"advent-of-code/solutions/day8"
	"log"
)

func GetPuzzles(day int) (func(input string) string, func(input string) string) {
	switch day {
	case 1:
		return day1.Puzzles()
	case 2:
		return day2.Puzzles()
	case 3:
		return day3.Puzzles()
	case 4:
		return day4.Puzzles()
	case 5:
		return day5.Puzzles()
	case 6:
		return day6.Puzzles()
	case 7:
		return day7.Puzzles()
	case 8:
		return day8.Puzzles()
	default:
		log.Fatalf("Day %d is not implemented.\n", day)
	}
	return nil, nil
}
