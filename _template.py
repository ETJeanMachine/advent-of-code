import asyncio
from util import get_input


def part_one(input: str) -> int:
    return 0


def part_two(input: str) -> int:
    return 1


async def main():
    input = await get_input(2015, 3)
    print(f"Part One: {part_one(input)}")
    print(f"Part Two: {part_two(input)}")


asyncio.run(main())
