import asyncio
from util import main


Instruction = tuple[str, tuple[int, int], tuple[int, int]]


def parse_input(input: str) -> list[Instruction]:
    res: list[Instruction] = []
    for line in input.split("\n"):
        ins = None
        split = line.split()
        if split[0] == "turn":
            ins = split[1]
            split.remove(split[1])
        else:
            ins = split[0]
        start_vals, end_vals = split[1].split(","), split[3].split(",")
        start = (int(start_vals[0]), int(start_vals[1]))
        end = (int(end_vals[0]), int(end_vals[1]))
        res.append((ins, start, end))
    return res


def part_one(input: str) -> int:
    lights = [[0 for _ in range(1000)] for _ in range(1000)]
    instructions = parse_input(input)
    for ins in instructions:
        for x in range(ins[1][0], ins[2][0] + 1):
            for y in range(ins[1][1], ins[2][1] + 1):
                if ins[0] == "on":
                    lights[x][y] = 1
                elif ins[0] == "off":
                    lights[x][y] = 0
                elif ins[0] == "toggle":
                    lights[x][y] = not lights[x][y]
    count = 0
    for light_row in lights:
        count += light_row.count(1)
    return count


def part_two(input: str) -> int:
    return 0


asyncio.run(main(2015, 6, part_one, part_two))
