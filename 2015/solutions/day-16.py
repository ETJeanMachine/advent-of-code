import asyncio
import re

from util import main

my_sue = {
    "children": 3,
    "cats": 7,
    "samoyeds": 2,
    "pomeranians": 3,
    "akitas": 0,
    "vizslas": 0,
    "goldfish": 5,
    "trees": 3,
    "cars": 2,
    "perfumes": 1,
}


def parse_input(input: str) -> list[dict[str, int]]:
    sues: list[dict[str, int]] = []
    for line in input.splitlines():
        matches: list[tuple[str, str]] = re.findall(r"([a-z]+): (\d+)", line)
        sue_data = {k: int(v) for k, v in matches}
        sues.append(sue_data)
    return sues


def part_one(input: str) -> int:
    sues = parse_input(input)
    for i in range(len(sues)):
        sue = sues[i]
        similar = True
        for k in sue.keys():
            similar &= sue[k] == my_sue[k]
        if similar:
            return i + 1
    return -1


def part_two(input: str) -> int:
    sues = parse_input(input)
    for i in range(len(sues)):
        sue = sues[i]
        similar = True
        for k in sue.keys():
            match k:
                case "cats" | "trees":
                    similar &= sue[k] > my_sue[k]
                case "pomeranians" | "goldfish":
                    similar &= sue[k] < my_sue[k]
                case _:
                    similar &= sue[k] == my_sue[k]
        if similar:
            return i + 1
    return -1


asyncio.run(main(2015, 16, part_one, part_two))
