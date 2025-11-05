import asyncio
from util import main


def part_one(input: str) -> int:
    count = 0
    for ln in input.split("\n"):
        ln = ln.strip()
        sec_1, sec_2 = ln.split(",")
        arr1 = sec_1.split("-")
        a, b = int(arr1[0]), int(arr1[1])
        arr2 = sec_2.split("-")
        x, y = int(arr2[0]), int(arr2[1])
        if (a <= x and y <= b) or (x <= a and b <= y):
            count += 1
    return count


def part_two(input: str) -> int:
    count = 0
    for ln in input.split("\n"):
        ln = ln.strip()
        sec_1, sec_2 = ln.split(",")
        arr1 = sec_1.split("-")
        a, b = int(arr1[0]), int(arr1[1])
        arr2 = sec_2.split("-")
        x, y = int(arr2[0]), int(arr2[1])
        if (
            (a <= x and x <= b)
            or (a <= y and y <= b)
            or (x <= a and a <= y)
            or (x <= b and b <= y)
        ):
            count += 1
    return count


asyncio.run(main(2022, 4, part_one, part_two))
