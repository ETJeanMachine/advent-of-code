import asyncio
from util import get_input


def part_one(input: str) -> int:
    visited: set[tuple[int, int]] = set()
    x, y = 0, 0
    visited.add((x, y))
    for c in input:
        if c == "^":
            y += 1
        elif c == "v":
            y -= 1
        elif c == ">":
            x += 1
        elif c == "<":
            x -= 1
        visited.add((x, y))
    return len(visited)


def part_two(input: str) -> int:
    visited: set[tuple[int, int]] = set()
    x1, y1, x2, y2 = 0, 0, 0, 0
    robo = False
    visited.add((x1, y1))
    for c in input:
        x, y = 0, 0
        if c == "^":
            y += 1
        elif c == "v":
            y -= 1
        elif c == ">":
            x += 1
        elif c == "<":
            x -= 1
        if not robo:
            x1 += x
            y1 += y
            visited.add((x1, y1))
        else:
            x2 += x
            y2 += y
            visited.add((x2, y2))
        robo = not robo
    return len(visited)


async def main():
    input = await get_input(2015, 3)
    print(f"Part One: {part_one(input)}")
    print(f"Part Two: {part_two(input)}")


asyncio.run(main())
