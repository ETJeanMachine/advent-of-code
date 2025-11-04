import asyncio
from util import get_input


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


async def main():
    input = await get_input(2015, 1)
    print(f"Part One: {part_one(input)}")
    print(f"Part Two: {part_two(input)}")


asyncio.run(main())
