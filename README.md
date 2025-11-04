# Advent of Code

This is a repository containing my code solutions for every years Advent of Code, dating back to 2015. This is something I began putting together in 2025 mostly out of boredom and an attempt to challenge myself by going back at these old problems and having a stab at them. I will also update this repository with future advent of codes in addition to working on previous year's ones.

## How to run

To run this project, you'll need [`uv`](https://docs.astral.sh/uv/) installed. Once you have it installed on your system, you'll also need to clone `.env.example` into a `.env` file; and populate it with the browser session cookie when you log into Advent of Code (This, by proxy, means it will solve for *your advent of code solutions, not mine*. If you aren't looking to be spoiled - don't use this project!). To get any particular solution that has been completed so far, run:

```bash
uv run -m {YEAR}.day-{DAY}
```

Where `YEAR` is the year of the problem/folder the solution is in, and `DAY` is the date of the problem/name of the file you are running.
