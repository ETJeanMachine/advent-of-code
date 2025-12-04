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

This is a monorepo containing my code solutions for every years [Advent of Code](https://adventofcode.com), dating back to 2015. This is something I began putting together in 2025 mostly out of boredom and an attempt to challenge myself by going back at these old problems and having a stab at them. I will also update this repository with future advent of codes in addition to working on previous year's ones.

## Running Solutions
In general, all subdirectories/years have a `.env.example` file that needs to be copied into a `.env` file with a session token in order to retrieve solutions from Advent of Code - this repo does not directly store input files on it. You can retrieve this session cookie from any modern web browser's dev tools, under the `Application` tab and `Cookies` section for chromium-based browsers, for instance.

The session token is tied to *your login* of Advent of Code, which means this repository will attempt to solve for your Advent of Code problems. If you don't want to be spoiled - don't run this repo! Not all problems will necessarily work; as some rely on additional input that may not be uniform between users playing Advent of Code, and these have been hardcoded into the solutions.

- 2015 (Python):
  - [Instructions](2015/README.md)
  - [Solutions](2015/solutions)
- 2022 (Python):
  - [Instructions](2022/README.md)
  - [Solutions](2022/solutions)
- 2024 (Rust):
  - [Instructions](2024/README.md)
  - [Solutions](2024/src/solutions)
- 2025 (Go):
  - [Instructions](2025/README.md)
  - [Solutions](2025/solutions)

## Progress
- 2015: 42 Stars
- 2022: 27 Stars
  - 2022's day's one through fourteen part one were done privately in 2022; and then imported and refactored to fit to this repository.
- 2024: 27 Stars
- 2025: 8 Stars
