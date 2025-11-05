import asyncio
from util import main


def load_crates(input: str):
    crates: dict[int, list[str]] = {}
    crate_input = input.split("\n\n")[0]
    for line in crate_input.split("\n"):
        if len(line) > 0 and line[1] == "1":
            continue
        num = 1
        for i in range(1, len(line), 4):
            if line[i] != " ":
                if crates.get(num) is None:
                    crates[num] = [line[i]]
                else:
                    crates[num].insert(0, line[i])
            num += 1
    return crates


def part_one(input: str) -> str:
    crates = load_crates(input)
    instructions = input.split("\n\n")[1]
    for line in instructions.split("\n"):
        cmds = line.split()
        _move, _from, _to = int(cmds[1]), int(cmds[3]), int(cmds[5])
        for i in range(_move):
            popped = crates[_from].pop()
            crates[_to].append(popped)
    tops = ""
    for i in range(1, len(crates) + 1):
        tops += crates[i][len(crates[i]) - 1]
    return tops


def part_two(input: str) -> str:
    crates = load_crates(input)
    instructions = input.split("\n\n")[1]
    for line in instructions.split("\n"):
        cmds = line.split()
        _move, _from, _to = int(cmds[1]), int(cmds[3]), int(cmds[5])
        picked_up = crates[_from][-_move:]
        crates[_from] = crates[_from][0 : len(crates[_from]) - _move]
        crates[_to] += picked_up
    tops = ""
    for i in range(1, len(crates) + 1):
        tops += crates[i][len(crates[i]) - 1]
    return tops


asyncio.run(main(2022, 5, part_one, part_two))
