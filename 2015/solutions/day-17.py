import asyncio
import re
from itertools import combinations

from util import main


def parse_input(input: str) -> list[int]:
    matches: list[str] = re.findall(r"(\d+)", input)
    return [int(m) for m in matches]


def part_one(input: str) -> int:
    containers = parse_input(input)
    total = 0
    for r in range(1, len(containers) + 1):
        for c in combinations(containers, r):
            if sum(c) == 150:
                total += 1
    return total


def part_two(input: str) -> int:
    containers = parse_input(input)
    total = 0
    for r in range(1, len(containers) + 1):
        for c in combinations(containers, r):
            if sum(c) == 150:
                total += 1
        if total != 0:
            return total
    return total


asyncio.run(main(2015, 17, part_one, part_two))
