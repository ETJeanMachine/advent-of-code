import asyncio
import math

from util import main

Graph = dict[str, dict[str, int]]


def held_karp(graph: Graph, source: str) -> tuple[list[str], int | float]:
    n, N = len(graph), 1 << len(graph)

    nodes: list[str] = []
    nodes.extend(graph.keys())
    # arranging our dict to have the source at the front:
    source_idx = nodes.index(source)
    nodes = nodes[source_idx:] + nodes[:source_idx]
    dists: list[list[int]] = [[0] * n for _ in range(n)]
    for i in range(n):
        for j in range(n):
            if i != j:
                dists[i][j] = graph[nodes[i]][nodes[j]]

    costs: list[list[float | int]] = [[math.inf] * n for _ in range(N)]
    prevs: list[list[None | int]] = [[None] * n for _ in range(N)]
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
                    cost = costs[prev_mask][k] + dists[k][j]
                    if cost < costs[mask][j]:
                        costs[mask][j] = cost
                        prevs[mask][j] = k

    full_mask = (1 << n) - 1
    min_cost = math.inf
    last = None
    for j in range(1, n):
        cost = costs[full_mask][j]
        if cost < min_cost:
            min_cost = cost
            last = j

    idx_path: list[int] = []
    mask = full_mask
    curr = last
    while curr is not None:
        idx_path.append(curr)
        prev = prevs[mask][curr]
        mask ^= 1 << curr
        curr = prev
    idx_path.reverse()

    path: list[str] = []
    for i in idx_path:
        path.append(nodes[i])

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


def part_one(input: str) -> str:
    graph = construct_graph(input)
    min_cost = math.inf
    min_path = []
    for source in graph:
        path, cost = held_karp(graph, source)
        if cost < min_cost:
            min_cost, min_path = cost, path
    path_str = ""
    for i in min_path:
        path_str += f"{i} -> "
    path_str = path_str.removesuffix(" -> ")
    results = f"Path: {path_str}\n          Cost: {min_cost}"
    return results


def part_two(input: str) -> int:
    return 0


asyncio.run(main(2015, 9, part_one, part_two))
