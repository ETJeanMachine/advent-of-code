import asyncio
import math
from functools import reduce

from util import main


def factors(n):
    return set(
        reduce(
            list.__add__, ([i, n // i] for i in range(1, int(n**0.5) + 1) if n % i == 0)
        )
    )


def part_one(input: str) -> int:
    presents = int(input)
    n = 1  # current house we're at
    while (t := sum(factors(n))) * 10 < presents:
        n += 1
    return n


def part_two(input: str) -> int:
    return 0


asyncio.run(main(2015, 20, part_one, part_two))
