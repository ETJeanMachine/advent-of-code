package day10

import (
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type Machine struct {
	lights    []bool
	goal      []bool
	buttons   [][]int
	joltage   []int
	jolt_goal []int
}

func NewMachine(goal []bool, buttons [][]int, jolt_goal []int) *Machine {
	lights := make([]bool, len(goal))
	joltage := make([]int, len(jolt_goal))
	return &Machine{lights, goal, buttons, joltage, jolt_goal}
}

func (m Machine) pressLightButton(idx int) []bool {
	button := m.buttons[idx]
	newLights := make([]bool, len(m.lights))
	copy(newLights, m.lights)
	for _, i := range button {
		newLights[i] = !newLights[i]
	}
	return newLights
}

func (m Machine) pressJoltButton(idx int) []int {
	button := m.buttons[idx]
	newJolt := make([]int, len(m.joltage))
	copy(newJolt, m.joltage)
	for _, i := range button {
		newJolt[i] += 1
	}
	return newJolt
}

func lightString(lights []bool) string {
	lightStr := ""
	for _, lit := range lights {
		if lit {
			lightStr += "#"
		} else {
			lightStr += "."
		}
	}
	return lightStr
}

func joltString(lights []int) string {
	joltStr := ""
	for _, jolt := range lights {
		joltStr += fmt.Sprintf("%d", jolt)
	}
	return joltStr
}

func (m *Machine) ConfigureLights() int {
	seenStates := make(map[string]bool)
	seenStates[lightString(m.lights)] = true
	queue := [][]bool{m.lights}
	depthQueue := []int{0}
	goalStr := lightString(m.goal)
	for len(queue) > 0 {
		m.lights = queue[0]
		currStr := lightString(m.lights)
		queue = queue[1:]

		currDepth := depthQueue[0]
		depthQueue = depthQueue[1:]

		if currStr == goalStr {
			return currDepth
		}
		for idx := range m.buttons {
			next := m.pressLightButton(idx)
			nextStr := lightString(next)
			if _, ok := seenStates[nextStr]; !ok {
				seenStates[nextStr] = true
				queue = append(queue, next)
				depthQueue = append(depthQueue, currDepth+1)
			}
		}
	}
	return -1
}

func (m *Machine) ConfigureJoltage() int {
	seenStates := make(map[string]bool)
	seenStates[joltString(m.joltage)] = true
	queue := [][]int{m.joltage}
	depthQueue := []int{0}
	goalStr := joltString(m.jolt_goal)
	for len(queue) > 0 {
		m.joltage = queue[0]
		currStr := joltString(m.joltage)
		queue = queue[1:]

		currDepth := depthQueue[0]
		depthQueue = depthQueue[1:]

		if currStr == goalStr {
			return currDepth
		}
		for idx := range m.buttons {
			next := m.pressJoltButton(idx)
			nextStr := joltString(next)
			if _, ok := seenStates[nextStr]; !ok {
				seenStates[nextStr] = true
				queue = append(queue, next)
				depthQueue = append(depthQueue, currDepth+1)
			}
		}
	}
	return -1
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
		count += machine.ConfigureLights()
	}
	return strconv.Itoa(count)
}

func partTwo(input string) string {
	machines := parseInput(input)
	count := 0
	for _, machine := range machines {
		count += machine.ConfigureJoltage()
	}
	return strconv.Itoa(count)
}

func Puzzles() (func(string) string, func(string) string) {
	return partOne, partTwo
}
