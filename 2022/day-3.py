import asyncio
from util import main


def priority(c: str):
    if c.isupper():
        return (ord(c) - 65) + 27
    else:
        return (ord(c) - 97) + 1


def part_one(input: str) -> int:
    priority_sum = 0
    for line in input.split("\n"):
        line = line.strip()
        rucksack_1 = line[0 : len(line) // 2]
        rucksack_2 = line[len(line) // 2 : len(line)]
        for c in rucksack_1:
            if c in rucksack_2:
                priority_sum += priority(c)
                break
    return priority_sum


def part_two(input: str) -> int:
    priority_sum = 0
    curr_group: list[str] = []
    for line in input.split("\n"):
        rucksack = line.strip()
        if len(curr_group) < 3:
            # Making sure the shortest rucksack is at the beginning of the group.
            if len(curr_group) > 0 and len(curr_group[0]) > len(rucksack):
                curr_group.insert(0, rucksack)
            else:
                curr_group.append(rucksack)
        if len(curr_group) == 3:
            for c in curr_group[0]:
                if (c in curr_group[1]) and (c in curr_group[2]):
                    priority_sum += priority(c)
                    break
            curr_group = []
    return priority_sum


asyncio.run(main(2022, 3, part_one, part_two))
