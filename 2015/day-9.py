import asyncio
import math
from dataclasses import dataclass

from heapdict import heapdict

from util import main

Graph = dict[str, dict[str, int]]


def held_karp(graph: Graph) -> tuple[list[str], int | float]:
    n, N = len(graph), 1 << len(graph)
    costs = [[math.inf] * n for _ in range(N)]
    prevs = [[""] * n for _ in range(N)]
    nodes = []
    nodes.extend(graph.keys())
    costs[1][0] = 0

    for mask in range(1, N):
        if not (mask & 1):
            continue
        for j in range(1, n):
            if not (mask & (1 << j)):
                continue
            prev_mask = mask ^ (1 << j)
            for k in range(n):
                if prev_mask & (1 << k):
                    weight = graph[nodes[k]][nodes[j]]
                    cost = costs[prev_mask][k] + weight
                    if cost < costs[mask][j]:
                        costs[mask][j] = cost
                        prevs[mask][j] = nodes[k]
                k += 1

    full_mask = (1 << n) - 1
    min_cost = math.inf
    last = None
    for j in range(1, n):
        cost = costs[full_mask][j] + graph[nodes[j]][nodes[0]]
        if cost < min_cost:
            min_cost = cost
            last = j

    path: list[str] = []
    mask = full_mask
    curr = last
    while True:
        path.append(nodes[curr])
        prev = prevs[mask][curr]
        mask ^= 1 << curr
        if prev in nodes:
            curr = nodes.index(prev)
        else:
            break

    path.reverse()

    return path, min_cost


def construct_graph(input: str) -> Graph:
    graph: Graph = {}
    for line in input.split("\n"):
        locs, weight = line.split(" = ")
        weight = int(weight)
        loc_a, loc_b = locs.split(" to ")
        if edges := graph.get(loc_a):
            edges[loc_b] = weight
        else:
            graph[loc_a] = {loc_b: weight}
        if edges := graph.get(loc_b):
            edges[loc_a] = weight
        else:
            graph[loc_b] = {loc_a: weight}
    return graph


def part_one(input: str) -> int | float:
    graph = construct_graph(input)
    path, cost = held_karp(graph)
    print(path)
    return cost


def part_two(input: str) -> int:
    return 0


asyncio.run(main(2015, 9, part_one, part_two))
