import asyncio
import re

from util import main

Graph = dict[str, dict[str, int]]


def construct_graph(input: str) -> Graph:
    graph: Graph = {}
    for line in input.split("\n"):
        groups = re.findall(r"^([A-Z][a-z]+).*(gain|lose) (\d+).*([A-Z][a-z]+)", line)
        groups = groups[0]
        happiness = -int(groups[2]) if groups[1] == "lose" else int(groups[2])
        if edges := graph.get(groups[0]):
            edges[groups[3]] = happiness
        else:
            graph[groups[0]] = {groups[3]: happiness}
    return graph


def part_one(input: str) -> int:
    graph = construct_graph(input)
    return 0


def part_two(input: str) -> int:
    return 0


asyncio.run(main(2015, 13, part_one, part_two))
