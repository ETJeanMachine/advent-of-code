package main

import (
	"advent-of-code/solutions/day1"
	"advent-of-code/solutions/day2"
	"advent-of-code/solutions/day3"
	"advent-of-code/solutions/day4"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/joho/godotenv"
)

func get_input(day int) string {
	// Getting environment variables.
	env_vars, err := godotenv.Read()
	if err != nil {
		log.Fatal("Fatal error in reading .env file!")
	}

	// Setting up our get request and headers.
	url := fmt.Sprintf("https://adventofcode.com/2025/day/%d/input", day)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		log.Fatal("Fatal error in generating HTTP request.")
	}
	session_cookie := fmt.Sprintf("session=%s", env_vars["SESSION_COOKIE"])
	req.Header.Add("Cookie", session_cookie)

	// Sending the request.
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Fatalf("Fatal error in fetching input for day %d!", day)
	}

	responseData, err := io.ReadAll(resp.Body)
	responseString := string(responseData)
	return strings.TrimSpace(responseString)
}

func run_day(day int) {
	input := get_input(day)
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
	fmt.Printf("Advent of Code 2025 Day %d\n", day)
	now := time.Now()
	res_one := part_one(input)
	time_one := time.Since(now)
	now = time.Now()
	res_two := part_two(input)
	time_two := time.Since(now)
	fmt.Printf("Part One: %s\n", res_one)
	fmt.Printf("Time One: %.2fms\n\n", time_one.Seconds()*1000)
	fmt.Printf("Part Two: %s\n", res_two)
	fmt.Printf("Time Two: %.2fms\n", time_two.Seconds()*1000)
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Missing necessary command line arguments!")
	}
	cliArgs := os.Args[1:]
	day, err := strconv.Atoi(cliArgs[0])
	if err != nil {
		log.Fatalf("Day \"%s\" is not an integer!", cliArgs[0])
	}
	run_day(day)
}
