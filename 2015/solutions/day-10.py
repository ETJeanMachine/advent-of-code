import asyncio

from util import main


def look_say(input: str) -> str:
    digits: list[list[str]] = []
    prev_d = 0
    for d in input:
        if d != prev_d:
            prev_d = d
            digits.append([d])
        else:
            digits[-1].append(d)
    result = ""
    for d in digits:
        result += f"{len(d)}{d[0]}"
    return result


def part_one(input: str) -> int:
    result = input
    for _ in range(40):
        result = look_say(result)
    return len(result)


def part_two(input: str) -> int:
    result = input
    for _ in range(50):
        result = look_say(result)
    return len(result)


asyncio.run(main(2015, 10, part_one, part_two))
