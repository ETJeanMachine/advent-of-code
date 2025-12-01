package solutions

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
)

type safe struct {
	dial int
}

func (s *safe) rotate_dial(dist int) int {
	clicks := 0
	if dist > 99 {
		clicks += dist / 100
	} else if dist < -99 {
		clicks -= dist / 100
	}
	dist %= 100
	prev := s.dial
	s.dial += dist
	if prev != 0 && (s.dial > 99 || s.dial < 0) {
		clicks += 1
	} else if prev != 0 && s.dial == 0 {
		clicks += 1
	}
	s.dial = ((s.dial % 100) + 100) % 100
	return clicks
}

func part_one(instructions []int) int {
	safe := safe{50}
	password := 0
	for _, dist := range instructions {
		safe.rotate_dial(dist)
		if safe.dial == 0 {
			password += 1
		}
	}
	return password
}

func part_two(instructions []int) int {
	safe := safe{50}
	password := 0
	for _, dist := range instructions {
		password += safe.rotate_dial(dist)
	}
	return password
}

func parse_input(input string) []int {
	var parsed []int
	r := regexp.MustCompile(`(L|R)(\d+)`)
	matches := r.FindAllStringSubmatch(input, -1)
	for _, match := range matches {
		switch match[1] {
		case "L":
			dist, _ := strconv.Atoi(match[2])
			parsed = append(parsed, -dist)
		case "R":
			dist, _ := strconv.Atoi(match[2])
			parsed = append(parsed, dist)
		default:
			log.Fatal("Invalid input!")
		}
	}
	return parsed
}

func Day1(input string) {
	// test := "L68\nL30\nR48\nL5\nR60\nL55\nL1\nL99\nR14\nL82"
	instructions := parse_input(input)
	one_sol := part_one(instructions)
	two_sol := part_two(instructions)
	fmt.Printf("Part One: %d\n", one_sol)
	fmt.Printf("Part Two: %d\n", two_sol)
}
