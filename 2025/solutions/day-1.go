package solutions

import (
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

func parseInput(input string) []int {
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

func dayOnePartOne(input string) string {
	parsed := parseInput(input)
	safe := safe{50}
	password := 0
	for _, dist := range parsed {
		safe.rotate_dial(dist)
		if safe.dial == 0 {
			password += 1
		}
	}
	return strconv.Itoa(password)
}

func dayOnePartTwo(input string) string {
	parsed := parseInput(input)
	safe := safe{50}
	password := 0
	for _, dist := range parsed {
		password += safe.rotate_dial(dist)
	}
	return strconv.Itoa(password)
}

func Day1() (func(string) string, func(string) string) {
	return dayOnePartOne, dayOnePartTwo
}
