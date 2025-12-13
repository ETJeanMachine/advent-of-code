package day10

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Machine struct {
	lights  []bool
	goal    []bool
	buttons [][]int
	joltage []int
}

func NewMachine(goal []bool, buttons [][]int, joltage []int) *Machine {
	lights := make([]bool, len(goal))
	return &Machine{lights, goal, buttons, joltage}
}

func (m *Machine) configure() int {
	return 0
}

// Helper func for parsing a comma-seperated string of digits.
func parseDigitCommaString(str string) []int {
	digits := strings.SplitSeq(str, ",")
	res := []int{}
	for d := range digits {
		digit, _ := strconv.Atoi(d)
		res = append(res, digit)
	}
	return res
}

func parseInput(input string) []*Machine {
	lines := strings.SplitSeq(input, "\n")
	machines := []*Machine{}
	for line := range lines {
		// Parsing the "goal" position of the lights.
		goalRe := regexp.MustCompile(`\[((?:\.|#)+)\]`)
		goalMatch := goalRe.FindStringSubmatch(line)[1]
		goal := []bool{}
		for _, c := range goalMatch {
			switch c {
			case '.':
				goal = append(goal, false)
			case '#':
				goal = append(goal, true)
			default:
				log.Fatalf("Error parsing string!")
			}
		}
		// Parsing the buttons
		buttonRe := regexp.MustCompile(`\(((?:\d,?)+)\)`)
		buttonMatches := buttonRe.FindAllStringSubmatch(line, -1)
		buttons := [][]int{}
		for _, buttonStr := range buttonMatches {
			buttons = append(buttons, parseDigitCommaString(buttonStr[1]))
		}
		// Parsing the joltage
		joltRe := regexp.MustCompile(`\{((?:\d+,?)+)\}`)
		joltageStr := joltRe.FindStringSubmatch(line)[1]
		joltage := parseDigitCommaString(joltageStr)
		machines = append(machines, NewMachine(goal, buttons, joltage))
	}
	return machines
}

func partOne(input string) string {
	machines := parseInput(input)
	count := 0
	for _, machine := range machines {
		count += machine.configure()
	}
	return strconv.Itoa(count)
}

func partTwo(input string) string {
	return "0"
}

func Puzzles() (func(string) string, func(string) string) {
	return partOne, partTwo
}
