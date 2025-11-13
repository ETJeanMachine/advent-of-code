import asyncio
import itertools
import re

from util import main

Graph = dict[str, dict[str, int]]


def pref_matrix(input: str):
    # the index of the prefs array each name refers to.
    names: dict[str, int] = {}
    # matrix of combined preferences; e.g., the combined value if Bob is next Alice
    # is bob's happiness + alice's happiness.
    prefs: list[list[int]] = []
    for line in input.splitlines():
        match = re.findall(r"^([A-Z][a-z]+).*(gain|lose) (\d+).*([A-Z][a-z]+)", line)[0]
        p_1, p_2, happy = (
            match[0],
            match[3],
            int(match[2]) if match[1] == "gain" else -int(match[2]),
        )
        if (i := names.get(p_1)) is None:
            i = len(names)
            names[p_1] = i
        if (j := names.get(p_2)) is None:
            j = len(names)
            names[p_2] = j
        if i >= len(prefs):
            prefs.append([])
        if j >= len(prefs[i]):
            prefs[i].append(happy)
        else:
            prefs[i][j] += happy
    # sitting next to oneself is a value of 0
    for v in names.values():
        prefs[v].insert(v, 0)
    # adding together all edges that are bi-directional
    for i in range(len(prefs)):
        for j in range(i + 1, len(prefs[i])):
            combined = prefs[i][j] + prefs[j][i]
            prefs[i][j] = combined
            prefs[j][i] = combined
    return prefs, names


# i am a lazy pos and am gonna brute force this one.
def max_happiness(prefs: list[list[int]]) -> int:
    max_h = 0
    # all possible orders of seating
    orders = itertools.permutations(range(len(prefs)), len(prefs))
    for comb in orders:
        happiness = prefs[comb[-1]][comb[0]]
        for x in range(len(comb) - 1):
            i, j = comb[x], comb[x + 1]
            happiness += prefs[i][j]
        max_h = max(happiness, max_h)
    return max_h


def part_one(input: str) -> int:
    prefs, _ = pref_matrix(input)
    max_h = max_happiness(prefs)
    return max_h


def part_two(input: str) -> int:
    prefs, _ = pref_matrix(input)
    # adding myself as person "0"
    prefs.insert(0, [0 for _ in range(len(prefs) + 1)])
    for i in range(1, len(prefs)):
        prefs[i].insert(0, 0)
    max_h = max_happiness(prefs)
    return max_h


asyncio.run(main(2015, 13, part_one, part_two))
