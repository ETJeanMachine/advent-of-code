import asyncio
import json
from io import StringIO

from util import main


def find_nums(parsed, ignore=None) -> list[int]:
    if isinstance(parsed, int):
        return [parsed]

    nums = []
    if isinstance(parsed, list):
        for x in parsed:
            nums.extend(find_nums(x, ignore))
    elif isinstance(parsed, dict):
        values = list(parsed.values())
        if ignore in values:
            return []
        for v in values:
            nums.extend(find_nums(v, ignore))

    return nums


def part_one(input: str) -> int:
    io = StringIO(input)
    parsed = json.load(io)
    return sum(find_nums(parsed))


def part_two(input: str) -> int:
    io = StringIO(input)
    parsed = json.load(io)
    return sum(find_nums(parsed, ignore="red"))


asyncio.run(main(2015, 12, part_one, part_two))
