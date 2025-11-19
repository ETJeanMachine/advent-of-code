# Advent of Code 2015

To run this project, you'll need [`uv`](https://docs.astral.sh/uv/) installed. Once you have it installed on your system, you'll also need to clone `.env.example` into a `.env` file; and populate it with the browser session cookie when you log into Advent of Code (This, by proxy, means it will solve for *your advent of code solutions, not mine*. If you aren't looking to be spoiled - don't use this project!). To get any particular solution that has been completed so far, run:

```bash
uv run -m solutions.day-{DAY}
```

Where `YEAR` is the year of the problem/folder the solution is in, and `DAY` is the date of the problem/name of the file you are running.
