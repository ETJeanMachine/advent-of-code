import asyncio
import heapq
import math
import re
from collections import defaultdict

from heapdict import heapdict

from util import main


def parse_input(input: str) -> tuple[dict[str, str], str]:
    r_strs, medicine = input.split("\n\n")
    replacements: dict[str, str] = {}
    for line in r_strs.splitlines():
        value, key = line.split(" => ")
        replacements[key] = value
    return replacements, medicine


def part_one(input: str) -> int:
    replacements, medicine = parse_input(input)
    molecules: set[str] = set()
    for k in replacements.keys():
        for m in re.finditer(rf"({replacements[k]})", medicine):
            start, end = m.span()
            repl = f"{medicine[:start]}{k}{medicine[end:]}"
            molecules.add(repl)
    return len(molecules)


def a_star(replacements: dict[str, str], root: str, goal="e"):
    def neighbours(molecule: str):
        neighbours: set[str] = set()
        for k, v in replacements.items():
            for m in re.finditer(rf"({k})", molecule):
                start, end = m.span()
                repl = f"{molecule[:start]}{v}{molecule[end:]}"
                neighbours.add(repl)
        return neighbours

    # pretty shoddy heuristic that sometimes runs for farrrrr too long.
    # should be improved.
    def heuristic(molecule: str):
        m_res = re.findall(r"[A-Z][a-z]?", molecule)
        return len(m_res)

    discovered = heapdict()
    discovered[root] = heuristic(root)

    came_from: dict[str, str] = {}

    g_score: defaultdict[str, int | float] = defaultdict(lambda: math.inf)
    g_score[root] = 0

    while len(discovered) > 0:
        curr, _ = discovered.popitem()
        if curr == goal:
            return g_score[curr]
        for n in neighbours(curr):
            ten_g_score = g_score[curr] + 1
            if ten_g_score < g_score[n]:
                came_from[n] = curr
                g_score[n] = ten_g_score
                discovered[n] = ten_g_score + heuristic(n)
        pass

    return None


def part_two(input: str) -> int | float:
    replacements, medicine = parse_input(input)
    if depth := a_star(replacements, medicine):
        return depth
    return 0


asyncio.run(main(2015, 19, part_one, part_two))
