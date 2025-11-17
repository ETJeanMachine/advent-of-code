import asyncio
import re

import heapdict

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


# ill be totally honest. this is a modified bfs, but i have zero clue
# why this works.
# this runs in an infinite loop sometimes. other times it works!
# i dont get it right now. but ill figure it out.
def trace(replacements: dict[str, str], root: str, goal="e") -> int:
    def molecule_len(molecule: str) -> int:
        return len(re.findall(r"[A-Z][a-z]?", molecule))

    queue = heapdict.heapdict()
    root_len = molecule_len(root)
    queue[(root, 0)] = root_len
    curr_depth = -1
    visited: set[str] = set()
    visited.add(root)
    while len(queue) > 0:
        (curr, depth), c_len = queue.popitem()
        if depth > curr_depth:
            curr_depth = depth
            print(f"Depth: {curr_depth}, Visited: {len(visited)}")
        neighbours: set[str] = set()
        for k in replacements.keys():
            for m in re.finditer(rf"({k})", curr):
                start, end = m.span()
                repl = f"{curr[:start]}{replacements[k]}{curr[end:]}"
                neighbours.add(repl)
        for n in neighbours:
            if n == goal:
                return depth + 1
            n_len = molecule_len(n)
            if n not in visited:
                visited.add(n)
                queue[(n, depth + 1)] = (depth + 1) * (n_len - root_len)
    return 0


def part_two(input: str) -> int:
    replacements, medicine = parse_input(input)
    return trace(replacements, medicine)


asyncio.run(main(2015, 19, part_one, part_two))
