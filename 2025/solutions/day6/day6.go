package day6

import (
	"errors"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type MathProblem struct {
	numbers  []int
	operator rune
}

func (p *MathProblem) Solve() (int, error) {
	if p == nil {
		return -1, errors.New("Undefined problem.")
	}
	var solution int
	switch p.operator {
	case '+':
		solution = 0
	case '*':
		solution = 1
	default:
		return -1, errors.New("Undefined operator.")
	}
	for _, value := range p.numbers {
		switch p.operator {
		case '+':
			solution += value
		case '*':
			solution *= value
		}
	}
	return solution, nil
}

func solveProblems(mathProblems []MathProblem) int {
	total := 0
	for _, problem := range mathProblems {
		solution, error := problem.Solve()
		if error == nil {
			total += solution
		} else {
			fmt.Printf("%v %c", problem.numbers, problem.operator)
			log.Fatalf("Error when solving: %s\n", error.Error())
		}
	}
	return total
}

func partOne(input string) string {
	lines := strings.Split(input, "\n")
	mathProblems := []MathProblem{}
	for _, line := range lines {
		values := strings.Fields(line)
		for idx, value := range values {
			if idx >= len(mathProblems) {
				mathProblems = append(mathProblems, MathProblem{[]int{}, '?'})
			}
			num, err := strconv.Atoi(value)
			if err == nil {
				mathProblems[idx].numbers = append(mathProblems[idx].numbers, num)
			} else {
				bytes := []byte(value)
				mathProblems[idx].operator = rune(bytes[0])
			}
		}
	}
	solution := solveProblems(mathProblems)
	return strconv.Itoa(solution)
}

func rotateNumbers(numStrs []string, length int) []int {
	nums := []int{}
	for i := range length {
		rotatedStr := ""
		for _, str := range numStrs {
			if rune(str[i]) != ' ' {
				rotatedStr += string(str[i])
			}
		}
		rotated, error := strconv.Atoi(rotatedStr)
		if error != nil {
			log.Fatal("Integer conversion error when rotating.")
		}
		nums = append(nums, rotated)
	}
	return nums
}

func partTwo(input string) string {
	lines := strings.Split(input, "\n")
	// re-adding whitespace i trimmed when fetching input :p
	lines[len(lines)-1] += strings.Repeat(" ", len(lines[0])-len(lines[len(lines)-1]))
	re := regexp.MustCompile(`((\+|\*) *)`)
	operators := re.FindAll([]byte(lines[len(lines)-1]), 1000)
	lens := []int{}
	mathProblems := []MathProblem{}
	for idx, bytes := range operators {
		if idx < len(operators)-1 {
			lens = append(lens, len(bytes)-1)
		} else {
			// hardcoding this bc i dont wanna fix my shit where i trim whitespace in
			// main.go :p
			lens = append(lens, len(bytes))
		}
		mathProblems = append(mathProblems, MathProblem{[]int{}, rune(bytes[0])})
	}
	numLines := lines[:len(lines)-1]
	start := 0
	for idx := range mathProblems {
		numStrs := []string{}
		for _, line := range numLines {
			end := start + lens[idx]
			numStrs = append(numStrs, line[start:end])
		}
		start += lens[idx] + 1
		mathProblems[idx].numbers = rotateNumbers(numStrs, lens[idx])
	}
	solution := solveProblems(mathProblems)
	return strconv.Itoa(solution)
}

func Puzzles() (func(string) string, func(string) string) {
	return partOne, partTwo
}
