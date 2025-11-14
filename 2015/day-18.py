import asyncio
import copy

from util import main


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
            (x, y - 1),
            (x - 1, y),
            (x + 1, y),
            (x - 1, y + 1),
            (x, y + 1),
            (x + 1, y + 1),
        ]
        neighbours = [
            (x, y) for x, y in neighbours if x >= 0 and y >= 0 and x < 100 and y < 100
        ]
        return neighbours

    next_grid = copy.deepcopy(grid)
    for x in range(len(grid)):
        for y in range(len(grid[x])):
            n_lights = [grid[nx][ny] for nx, ny in neighbours(x, y)]
            lit_neighbours = sum(n_lights)
            if grid[x][y]:
                next_grid[x][y] = int(lit_neighbours == 3 or lit_neighbours == 2)
            else:
                next_grid[x][y] = int(lit_neighbours == 3)
    return next_grid


def part_one(input: str) -> int:
    grid = parse_input(input)
    for _ in range(100):
        grid = animate(grid)
    return count_on(grid)


def part_two(input: str) -> int:
    return 0


asyncio.run(main(2015, 18, part_one, part_two))
