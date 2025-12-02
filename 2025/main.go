package main

import (
	"advent-of-code/solutions"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
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
	return responseString
}

func run_day(day int) {
	input := get_input(day)
	var part_one func(input string) string
	var part_two func(input string) string
	switch day {
	case 1:
		part_one, part_two = solutions.DayOne()
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
	fmt.Printf("Time One: %dms\n\n", time_one.Milliseconds())
	fmt.Printf("Part Two: %s\n", res_two)
	fmt.Printf("Time Two: %dms\n", time_two.Milliseconds())
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
