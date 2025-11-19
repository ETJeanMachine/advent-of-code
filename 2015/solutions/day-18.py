import asyncio
import copy

from util import main


def print_grid(grid: list[list[int]]):
    for row in grid:
        lights = "".join(["#" if bool(x) else "." for x in row])
        print(lights)


def parse_input(input: str) -> list[list[int]]:
    grid: list[list[int]] = []
    for line in input.splitlines():
        row = []
        for c in line:
            if c == "#":
                row.append(1)
            else:
                row.append(0)
        grid.append(row)
    return grid


def count_on(grid: list[list[int]]) -> int:
    count = 0
    for row in grid:
        for lit in row:
            if bool(lit):
                count += 1
    return count


def animate(grid: list[list[int]]):
    def neighbours(x: int, y: int) -> list[tuple[int, int]]:
        neighbours = [
            (x - 1, y - 1),
            (x, y - 1),
            (x + 1, y - 1),
            (x - 1, y),
            (x + 1, y),
            (x - 1, y + 1),
            (x, y + 1),
            (x + 1, y + 1),
        ]
        neighbours = [
            (x, y)
            for x, y in neighbours
            if x >= 0 and y >= 0 and x < len(grid) and y < len(grid)
        ]
        return neighbours

    next_grid = copy.deepcopy(grid)
    for y in range(len(grid)):
        for x in range(len(grid[y])):
            n_lights = [grid[ny][nx] for nx, ny in neighbours(x, y)]
            lit_count = sum(n_lights)
            if grid[y][x]:
                next_grid[y][x] = int(lit_count == 3 or lit_count == 2)
            else:
                next_grid[y][x] = int(lit_count == 3)
    return next_grid


def part_one(input: str) -> int:
    grid = parse_input(input)
    for _ in range(100):
        grid = animate(grid)
    return count_on(grid)


def part_two(input: str) -> int:
    grid = parse_input(input)
    for _ in range(100):
        grid = animate(grid)
        # turn our corners back on
        grid[0][0] = 1
        grid[0][-1] = 1
        grid[-1][0] = 1
        grid[-1][-1] = 1
    return count_on(grid)


asyncio.run(main(2015, 18, part_one, part_two))
