import asyncio
from util import main


def part_one(input: str) -> int:
    total = 0
    for line in input.split("\n"):
        vals = [int(n) for n in line.split("x")]
        l, w, h = vals[0], vals[1], vals[2]
        total += (2 * l * w) + (2 * w * h) + (2 * h * l)
        total += min(l * w, w * h, h * l)
    return total


def part_two(input: str) -> int:
    total = 0
    for line in input.split():
        vals = [int(n) for n in line.split("x")]
        l, w, h = vals[0], vals[1], vals[2]
        total += min(2 * (l + w), 2 * (w + h), 2 * (h + l))
        total += l * w * h
    return total


asyncio.run(main(2015, 2, part_one, part_two))
