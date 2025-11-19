import asyncio
import math

from heapdict import heapdict

from util import main


class HeightMap:
    def __init__(self) -> None:
        self.heights: list[list[int]] = []
        self.lowest: set[tuple[int, int]] = set()
        self.start: tuple[int, int] | None = None
        self.end: tuple[int, int] | None = None
        self.rows = 0
        self.cols = 0

    def add_tile(self, c: str, row: int, col: int):
        """Adds a tile to the HeightMap.

        Args:
            c (str): The tile as a character (a-z or S or E).
            row (int): Row of the tile.
            col (int): Column of the tile.
        """
        if self.rows == len(self.heights):
            self.heights.append([])
        height = ord(c) - 97
        if c == "S":
            self.start = (row, col)
            height = 0
        elif c == "E":
            self.end = (row, col)
            height = 25
        if height == 0:
            self.lowest.add((row, col))
        self.heights[row].append(height)

    def neighbours(self, coords: tuple[int, int]) -> list[tuple[int, int]]:
        """Finds the neighbours to a coordinate

        Args:
            coords (Tuple[int, int]): The coords we're finding the neighbours of.

        Returns:
            List[Tuple[int, int]]: A list of neighbouring coords (at most 4).
        """
        row, col = coords[0], coords[1]
        neighbouring = []
        if row - 1 >= 0:
            neighbouring.append((row - 1, col))
        if row + 1 < self.rows:
            neighbouring.append((row + 1, col))
        if col - 1 >= 0:
            neighbouring.append((row, col - 1))
        if col + 1 < self.cols:
            neighbouring.append((row, col + 1))
        return neighbouring

    def distance(self, p1: tuple[int, int], p2: tuple[int, int]) -> float:
        """The distance between two points is 1 if it's less than one space up,
        and infinity otherwise.

        Args:
            p1 (Tuple[int, int]): The first point for comparison.
            p2 (Tuple[int, int]): The second point for comparison.

        Returns:
            float: 1 if it's able to be traveled to, infinity otherwise.
        """
        h1 = self.heights[p1[0]][p1[1]]
        h2 = self.heights[p2[0]][p2[1]]
        if h2 - h1 <= 1:
            return 1
        return math.inf

    def dijkstra(self, *, source: tuple[int, int] | None = None):
        """Performs dijkstra's algorithm.

        Args:
            source (Tuple[int, int], optional): The source tile. Defaults to the tile labeled 'S'.

        Returns:
            Dict[Tuple[int, int], float]: A dictionary of the distance from all observed values to
            the end tile.
            Dict[Tuple[int, int], Tuple[int, int]]: A traceback dictionary of tiles to the start tile,
            if you start from the end.
        """
        if source is None:
            source = self.start
        heap = heapdict()
        dist: dict[tuple[int, int] | None, int | float] = {source: 0}
        prev: dict[tuple[int, int], tuple[int, int]] = {}
        for row in range(self.rows):
            for col in range(self.cols):
                curr = (row, col)
                if curr != source:
                    dist[curr] = math.inf
                heap[curr] = dist[curr]
        while len(heap) != 0:
            curr, val = heap.popitem()
            if curr == self.end or val == math.inf:
                return dist, prev
            for n in self.neighbours(curr):
                alt = dist[curr] + self.distance(curr, n)
                if n in heap and alt < heap[n]:
                    dist[n] = alt
                    prev[n] = curr
                    heap[n] = alt
        return dist, prev

    def lowest_possible(self):
        """Finds the lowest possible distance to the end, given you start at any point with
        a height of 0.

        Returns:
            Dict[Tuple[int, int], float]: A dictionary of the distance from all observed values to
            the end tile.
            Dict[Tuple[int, int], Tuple[int, int]]: A traceback dictionary of tiles to the start tile,
            if you start from the end.
        """
        dist, prev = dict(), dict()
        source = None
        dist[self.end] = math.inf
        # generating the viable checking values (will have a step one up immediately in its neighbours)
        viable = set()
        for low in self.lowest:
            for n in self.neighbours(low):
                h1 = self.heights[low[0]][low[1]]
                h2 = self.heights[n[0]][n[1]]
                if h2 - h1 == 1:
                    viable.add(low)
        # checking every possible viable start point.
        for v in viable:
            n_dist, n_prev = self.dijkstra(source=v)
            if n_dist[self.end] < dist[self.end]:
                dist, prev = n_dist, n_prev
                source = v
        return dist, prev, source

    def traceback(
        self,
        prev_dict: dict[tuple[int, int], tuple[int, int]],
        *,
        source: tuple[int, int] | None = None,
    ):
        """Generates a string that represents the path taken.

        Args:
            prev_dict (Dict[Tuple[int, int], Tuple[int, int]]): The dictionary generated by dijkstra.
            source (Tuple[int, int], optional): The source tile. Defaults to the tile labeled 'S'.

        Returns:
            str: A string showing the path taken.
        """
        if source is None:
            source = self.start
        trace_map = [[]]
        curr = (0, 0)
        if source is not None and self.end is not None:
            trace_map = [["." for j in range(self.cols)] for i in range(self.rows)]
            trace_map[source[0]][source[1]] = "S"
            trace_map[self.end[0]][self.end[1]] = "E"
            curr = prev_dict[self.end]
        while curr != source:
            prev = prev_dict[curr]
            p_row, p_col = prev[0], prev[1]
            c_row, c_col = curr[0], curr[1]
            c: str
            if p_row < c_row:
                c = "↓"
            elif p_row > c_row:
                c = "↑"
            elif p_col < c_col:
                c = "→"
            else:
                c = "←"
            trace_map[c_row][c_col] = c
            curr = prev
        _out = ""
        for row in trace_map:
            for c in row:
                _out += c
            _out += "\n"
        return _out

    def __repr__(self) -> str:
        repr_str = ""
        for row in range(self.rows):
            for col in range(self.cols):
                if (row, col) == self.start:
                    repr_str += "S"
                elif (row, col) == self.end:
                    repr_str += "E"
                else:
                    repr_str += chr(self.heights[row][col] + 97)
            repr_str += "\n"
        return repr_str


def parse_input(input: str) -> HeightMap:
    h_map = HeightMap()
    for line in input.splitlines():
        line = line.strip()
        h_map.cols = len(line)
        for col in range(h_map.cols):
            h_map.add_tile(line[col], h_map.rows, col)
        h_map.rows += 1
    return h_map


def part_one(input: str) -> int | float:
    h_map = parse_input(input)
    dist, _ = h_map.dijkstra()
    return dist[h_map.end]


def part_two(input: str) -> int | float:
    h_map = parse_input(input)
    dist, _, _ = h_map.lowest_possible()
    return dist[h_map.end]


asyncio.run(main(2022, 12, part_one, part_two))
