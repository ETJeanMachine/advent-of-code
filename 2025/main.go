package main

import (
	"advent-of-code/solutions"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"

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
	switch day {
	case 1:
		solutions.Day1(input)
	default:
		log.Fatalf("Day %d is not implemented.\n", day)
	}
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
