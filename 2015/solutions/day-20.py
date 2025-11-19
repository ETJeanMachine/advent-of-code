import asyncio
from functools import reduce

from util import main


def factors(n):
    return set(
        reduce(
            list.__add__, ([i, n // i] for i in range(1, int(n**0.5) + 1) if n % i == 0)
        )
    )


def part_one(input: str) -> int:
    def total_presents(n: int):
        return 10 * sum(factors(n))

    presents = int(input)
    curr = 1  # current house we're at
    while total_presents(curr) < presents:
        curr += 1
    return curr


def part_two(input: str) -> int:
    def total_presents(n: int):
        red = set([x for x in factors(n) if n // x <= 50])
        return 11 * sum(red)

    presents = int(input)
    curr = 1  # current house we're at
    while total_presents(curr) < presents:
        curr += 1
    return curr


asyncio.run(main(2015, 20, part_one, part_two))
