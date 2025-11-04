# Advent of Code

This is a repository containing my code solutions for every years Advent of Code, dating back to 2015. This is something I began putting together in 2025 mostly out of boredom and an attempt to challenge myself by going back at these old problems and having a stab at them. I will also update this repository with future advent of codes in addition to working on previous year's ones.

## How to run

To run this project, you'll need [`uv`](https://docs.astral.sh/uv/) installed. Once you have it installed on your system, you'll also need to clone `.env.example` into a `.env` file; and populate it with the browser session cookie when you log into Advent of Code. To get any particular solution that has been completed so far, run:

```bash
uv run -m YEAR.day
```

Where `YEAR` is the name of the folder and `day` is the name of the day's file (all lowercase).
