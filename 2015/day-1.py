import asyncio
from util import main


def part_one(input: str) -> int:
    floor = 0
    for c in input:
        if c == "(":
            floor += 1
        elif c == ")":
            floor -= 1
    return floor


def part_two(input: str) -> int:
    floor = 0
    i = 1
    for c in input:
        if c == "(":
            floor += 1
        elif c == ")":
            floor -= 1
        if floor == -1:
            return i
        else:
            i += 1
    return -1


asyncio.run(main(2015, 1, part_one, part_two))
