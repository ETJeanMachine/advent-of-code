```
     ⋆
    >O<          ___       __                 __           ____   ______          __
   >>@o<        /   | ____/ /   _____  ____  / /_   ____  / __/  / ____/___  ____/ /__
  >>o><<<      / /| |/ __  / | / / _ \/ __ \/ __/  / __ \/ /_   / /   / __ \/ __  / _ \
 >>@<<o><<    / ___ / /_/ /| |/ /  __/ / / / /_   / /_/ / __/  / /___/ /_/ / /_/ /  __/
>>><<O>@<<<  /_/  |_\__,_/ |___/\___/_/ /_/\__/   \____/_/     \____/\____/\__,_/\___/
   |___|
```
---

This is a repository containing my code solutions for every years [Advent of Code](https://adventofcode.com), dating back to 2015. This is something I began putting together in 2025 mostly out of boredom and an attempt to challenge myself by going back at these old problems and having a stab at them. I will also update this repository with future advent of codes in addition to working on previous year's ones.

## How to run

To run this project, you'll need [`uv`](https://docs.astral.sh/uv/) installed. Once you have it installed on your system, you'll also need to clone `.env.example` into a `.env` file; and populate it with the browser session cookie when you log into Advent of Code (This, by proxy, means it will solve for *your advent of code solutions, not mine*. If you aren't looking to be spoiled - don't use this project!). To get any particular solution that has been completed so far, run:

```bash
uv run -m {YEAR}.day-{DAY}
```

Where `YEAR` is the year of the problem/folder the solution is in, and `DAY` is the date of the problem/name of the file you are running.

## Progress
- 2015: 36 Stars
- 2022: 27 Stars
  - 2022's day's one through fourteen part one were done privately in 2022; and then imported and refactored to fit to this repository.
- 2024: 4 Stars
  - 2024's first two day's were originally done in Rust, and can be accessed [at this repository](https://github.com/ETJeanMachine/Advent-of-Code-2024).
