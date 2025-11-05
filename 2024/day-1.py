# THIS WAS ORIGINALLY DONE IN RUST AND CONVERTED TO PYTHON AFTER THE FACT

import asyncio
from util import main


def part_one(input: str) -> int:
    l_list: list[int] = []
    r_list: list[int] = []
    append_left = True
    for val in input.split():
        val = int(val)
        if append_left:
            l_list.append(val)
        else:
            r_list.append(val)
        append_left = not append_left
    l_list.sort()
    r_list.sort()
    dist = 0
    for i in range(0, len(l_list)):
        dist += abs(l_list[i] - r_list[i])
    return dist


def part_two(input: str) -> int:
    l_list: list[int] = []
    r_list: list[int] = []
    append_left = True
    for val in input.split():
        val = int(val)
        if append_left:
            l_list.append(val)
        else:
            r_list.append(val)
        append_left = not append_left
    similarity = 0
    for val in l_list:
        similarity += val * r_list.count(val)
    return similarity


asyncio.run(main(2024, 1, part_one, part_two))
