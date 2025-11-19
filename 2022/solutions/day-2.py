import asyncio
from util import main


def rock_paper_scissors(n: str, m: str) -> int:
    if m == "X":
        match n:
            case "A":
                return 1 + 3
            case "B":
                return 1
            case "C":
                return 1 + 6
            case _:
                pass
    elif m == "Y":
        match n:
            case "A":
                return 2 + 6
            case "B":
                return 2 + 3
            case "C":
                return 2
            case _:
                pass
    elif m == "Z":
        match n:
            case "A":
                return 3
            case "B":
                return 3 + 6
            case "C":
                return 3 + 3
            case _:
                pass
    return 0


def part_one(input: str) -> int:
    points = 0
    for line in input.split("\n"):
        arr = line.split()
        points += rock_paper_scissors(arr[0], arr[1])
    return points


def part_two(input: str) -> int:
    points = 0
    for line in input.split("\n"):
        arr = line.split()
        if arr[1] == "X":
            match arr[0]:
                case "A":
                    points += 3
                case "B":
                    points += 1
                case "C":
                    points += 2
                case _:
                    pass
        elif arr[1] == "Y":
            points += 3
            match arr[0]:
                case "A":
                    points += 1
                case "B":
                    points += 2
                case "C":
                    points += 3
                case _:
                    pass
        elif arr[1] == "Z":
            points += 6
            match arr[0]:
                case "A":
                    points += 2
                case "B":
                    points += 3
                case "C":
                    points += 1
                case _:
                    pass
    return points


asyncio.run(main(2022, 2, part_one, part_two))
