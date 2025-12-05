package main

import (
	"advent-of-code/solutions"
	"cmp"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"slices"
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

func format_duration(duration time.Duration) string {
	if duration.Milliseconds() == 0 {
		return fmt.Sprintf("%.2fμs", duration.Seconds()*1e6)
	} else if duration.Milliseconds() < 1000 {
		return fmt.Sprintf("%.2fms", duration.Seconds()*1000)
	}
	return fmt.Sprintf("%.2fs", duration.Seconds())
}

func run_day(day int) {
	input := get_input(day)
	part_one, part_two := solutions.GetPuzzles(day)

	fmt.Printf("Advent of Code 2025 Day %d\n", day)

	now := time.Now()
	res_one := part_one(input)
	time_one := time.Since(now)
	fmt.Printf("Part One: %s\n", res_one)
	fmt.Printf("Time One: %s\n\n", format_duration(time_one))

	now = time.Now()
	res_two := part_two(input)
	time_two := time.Since(now)
	fmt.Printf("Part Two: %s\n", res_two)
	fmt.Printf("Time Two: %s\n", format_duration(time_two))
}

func benchmark_day(day int) {
	bench_puzzle := func(puzzle func(input string) string, input string) (time.Duration, time.Duration, time.Duration) {
		var benchmarks []time.Duration
		// 1000 samples
		for range 1000 {
			start := time.Now()
			puzzle(input)
			benchmarks = append(benchmarks, time.Since(start))
		}
		slices.SortFunc(benchmarks, func(a, b time.Duration) int {
			return cmp.Compare(a, b)
		})
		median := (benchmarks[499] + benchmarks[500]) / 2
		p25 := benchmarks[250]
		p75 := benchmarks[750]
		return median, p25, p75
	}

	fmt.Printf("Advent of Code 2025 Day %d Benchmarks (1000x)\n", day)
	input := get_input(day)
	part_one, part_two := solutions.GetPuzzles(day)
	median, p25, p75 := bench_puzzle(part_one, input)
	fmt.Printf("Part One P25 Time: %s\n", format_duration(p25))
	fmt.Printf("Part One P75 Time: %s\n", format_duration(p75))
	fmt.Printf("Part One Median Time: %s\n\n", format_duration(median))
	median, p25, p75 = bench_puzzle(part_two, input)
	fmt.Printf("Part Two P25 Time: %s\n", format_duration(p25))
	fmt.Printf("Part Two P75 Time: %s\n", format_duration(p75))
	fmt.Printf("Part Two Median Time: %s\n", format_duration(median))
}

func main() {
	if len(os.Args) <= 1 {
		log.Fatal("Missing necessary command line arguments!")
	}
	cliArgs := os.Args[1:]
	if cliArgs[0] == "--bench" {
		day, err := strconv.Atoi(cliArgs[1])
		if err != nil {
			log.Fatalf("Day \"%s\" is not an integer!", cliArgs[1])
		}
		benchmark_day(day)
	} else {
		day, err := strconv.Atoi(cliArgs[0])
		if err != nil {
			log.Fatalf("Day \"%s\" is not an integer!", cliArgs[0])
		}
		run_day(day)
	}
}
