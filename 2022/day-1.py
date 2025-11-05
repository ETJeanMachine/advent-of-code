import asyncio
from util import main


def part_one(input: str):
    largest = 0
    for line in input.split("\n\n"):
        calories = sum([int(n) for n in line.split()])
        largest = max(calories, largest)
    return largest


def part_two(input: str) -> int:
    calories = 0
    largest_3 = [0, 0, 0]
    for line in input.split("\n\n"):
        calories = sum([int(n) for n in line.split()])
        if calories > largest_3[0]:
            largest_3[2] = largest_3[1]
            largest_3[1] = largest_3[0]
            largest_3[0] = calories
        elif calories > largest_3[1]:
            largest_3[2] = largest_3[1]
            largest_3[1] = calories
        elif calories > largest_3[2]:
            largest_3[2] = calories
        calories = 0
    return sum(largest_3)


asyncio.run(main(2022, 1, part_one, part_two))
