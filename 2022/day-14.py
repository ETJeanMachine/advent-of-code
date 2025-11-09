import asyncio
import curses
import time
from enum import Enum

from util import get_input


class Material:
    class Type(Enum):
        ROCK = "#"
        SAND = "o"
        AIR = " "
        SOURCE = "+"

    def __init__(self, x: int, y: int, type=Type.AIR) -> None:
        self.x = x
        self.y = y
        self.type = type
        super().__init__()


class FallingSand:
    def __init__(self) -> None:
        self.cliff_face: list[list[Material]] = [[]]
        self.cliff_face[0].append(Material(500, 0, Material.Type.SOURCE))

    def __repr__(self) -> str:
        repr_str = ""
        for row in self.cliff_face:
            for m in row:
                repr_str += m.type.value
            repr_str += "\n"
        return repr_str

    @property
    def x_min(self) -> int:
        return self.cliff_face[0][0].x

    @property
    def x_max(self) -> int:
        return self.cliff_face[0][-1:][0].x

    @property
    def y_min(self) -> int:
        return 0

    @property
    def y_max(self) -> int:
        return self.cliff_face[-1:][0][0].y

    def run_simulation_1(self, stdscr: curses.window):
        count = 0
        over_edge = False
        height = stdscr.getmaxyx()[0] - 1
        offset = 0
        curses.curs_set(0)
        pad = curses.newpad(len(self.cliff_face) + 1, len(self.cliff_face[0]))

        for i in range(len(self.cliff_face)):
            for j in range(len(self.cliff_face[i])):
                pad.addch(i, j, ord(self.cliff_face[i][j].type.value))
        pad.refresh(0, 0, 0, 0, height, len(self.cliff_face[0]))

        while not over_edge:
            sand = Material(500, 0, type=Material.Type.SAND)
            while True:
                row, col = sand.y, sand.x - self.x_min
                if self.cliff_face[row + 1][col].type == Material.Type.AIR:
                    sand.y += 1
                elif self.cliff_face[row + 1][col - 1].type == Material.Type.AIR:
                    sand.y += 1
                    sand.x -= 1
                elif self.cliff_face[row + 1][col + 1].type == Material.Type.AIR:
                    sand.y += 1
                    sand.x += 1
                else:
                    self.cliff_face[row][col] = sand
                    pad.addch(row, col, ord(sand.type.value))
                    pad.refresh(offset, 0, 0, 0, height, len(self.cliff_face[0]))
                    count += 1
                    break
                pad.addch(row, col, ord(self.cliff_face[row][col].type.value))
                # checking if this new location is offscreen
                if sand.y < self.y_max and sand.x <= self.x_max and sand.x > self.x_min:
                    pad.addch(
                        sand.y, sand.x - self.x_min, ord(sand.type.value), curses.A_BOLD
                    )
                else:
                    over_edge = True
                    break
                # changing the cursor position, if necessary
                if sand.y > height + offset:
                    offset = sand.y - height + 1
                pad.refresh(offset, 0, 0, 0, height, len(self.cliff_face[0]))
                time.sleep(0.0005)
        pad.addstr(len(self.cliff_face), 0, f"Count: {count}", curses.A_BOLD)
        pad.refresh(
            len(self.cliff_face) - height, 0, 0, 0, height, len(self.cliff_face)
        )

    def expand_cliff(self, start, end):
        x_low, x_high = min(int(start[0]), int(end[0])), max(int(start[0]), int(end[0]))
        y_high = max(int(start[1]), int(end[1]))
        if x_high > self.x_max:
            # we need to expand the cliff face further left.
            rng = range(self.x_max, self.x_max + (x_high - self.x_max))
            for i in range(len(self.cliff_face)):
                self.cliff_face[i] += [Material(j, i) for j in rng]
        if x_low < self.x_min:
            # we need to expand the cliff face further right.
            rng = range(self.x_min - (self.x_min - x_low), self.x_min)
            for i in range(len(self.cliff_face)):
                self.cliff_face[i] = [Material(j, i) for j in rng] + self.cliff_face[i]
        if y_high > self.y_max:
            # we need to expand the cliff deeper down.
            rng = range(self.x_min, self.x_max + 1)
            for i in range(self.y_max + 1, y_high + 1):
                self.cliff_face.append([Material(j, i) for j in rng])

    def add_rock(self, start: list[str], end: list[str]):
        x_1, y_1 = int(start[0]), int(start[1])
        x_2, y_2 = int(end[0]), int(end[1])
        if x_1 != x_2:
            # drawing a rock face along the x-axis.
            x_low, x_high = min(x_1, x_2), max(x_1, x_2)
            for i in range(x_low, x_high + 1):
                self.cliff_face[y_1][i - self.x_min].type = Material.Type.ROCK
        else:
            # drawing a rock face along the y-axis.
            y_low, y_high = min(y_1, y_2), max(y_1, y_2)
            for i in range(y_low, y_high + 1):
                self.cliff_face[i][x_1 - self.x_min].type = Material.Type.ROCK


def parse_input(input: str) -> FallingSand:
    f_sand = FallingSand()
    for line in input.splitlines():
        paths = line.split("->")
        curr_path: list[str] = []
        while len(paths) > 0 or len(curr_path) > 1:
            if len(curr_path) == 2:
                start = curr_path.pop(0).split(",")
                end = curr_path[0].split(",")
                f_sand.expand_cliff(start, end)
                f_sand.add_rock(start, end)
            if len(paths) > 0:
                coords = paths.pop(0).strip()
                curr_path.append(coords)
    return f_sand


async def main():
    input = await get_input(2022, 14)
    f_sand = parse_input(input)
    stdscr = curses.initscr()
    f_sand.run_simulation_1(stdscr)


asyncio.run(main())
