import asyncio
import heapq
import math
from dataclasses import dataclass

from util import main


@dataclass
class Edge:
    neighbour: str
    weight: int


Graph = dict[str, list[Edge]]


def dijkstra(graph: Graph):
    cost: dict[tuple[str, set[str]], int | float] = {}
    queue: list[tuple[int | float, tuple[str, set[str]]]] = []

    for node in graph.keys():
        cost[(node, set())] = math.inf
        heapq.heappush(queue, (math.inf, (node, set())))

    while len(queue) > 0:
        _, (curr_node, subset) = heapq.heappop(queue)
        visited = subset.copy()
        visited.add(curr_node)

        for edge in graph[curr_node]:
            alt = cost[(curr_node, visited)] + edge.weight
            if alt < cost[(edge.neighbour, subset)]:
                cost[(edge.neighbour, visited)] = alt
                heapq.heappush(queue, (alt, (edge.neighbour, visited)))

    return cost


def construct_graph(input: str) -> Graph:
    graph: Graph = {}
    for line in input.split("\n"):
        locs, weight = line.split(" = ")
        weight = int(weight)
        loc_a, loc_b = locs.split(" to ")
        if edges := graph.get(loc_a):
            edges.append(Edge(loc_b, weight))
        else:
            graph[loc_a] = [Edge(loc_b, weight)]
        if edges := graph.get(loc_b):
            edges.append(Edge(loc_a, weight))
        else:
            graph[loc_b] = [Edge(loc_a, weight)]
    return graph


def part_one(input: str) -> int:
    graph = construct_graph(input)
    source = ""
    for node in graph.keys():
        source = node
        break
    cost = dijkstra(graph)
    print(cost)
    return 0


def part_two(input: str) -> int:
    return 0


asyncio.run(main(2015, 9, part_one, part_two))
