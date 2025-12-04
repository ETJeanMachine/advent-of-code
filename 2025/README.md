# Advent of Code 2025

To run this project, have the latest version of go installed. You will also need to create a `.env` file based on `.env.example` with the session cookie taken from your browser instance of Advent of Code. Once you do both of those things, run:
```bash
go run . <DAY>
```

Where day is the day of problem solutions you want to solve.

You can also run a benchmark on the solutions, which will run them 1000x and provide the median & minimum time it took to run, with:
```bash
go run . --bench <DAY>
```
