import asyncio

from util import main


class Point:
    def __init__(self, x: int, y: int) -> None:
        self.x = x
        self.y = y

    def __repr__(self) -> str:
        return f"({self.x}, {self.y})"

    def __eq__(self, __o: object) -> bool:
        if isinstance(__o, Point):
            return (self.x, self.y) == (__o.x, __o.y)
        return False

    def __hash__(self) -> int:
        return hash((self.x, self.y))

    def is_adjacent(self, p: "Point") -> bool:
        """Finds out if two points are adjacent to each other.

        Args:
            p (Point): The point we're comparing this point against.

        Returns:
            bool: True if the points are adjacent, False otherwise.
        """
        return abs(self.x - p.x) < 2 and abs(self.y - p.y) < 2

    def become_adjacent(self, p: "Point"):
        # If we're already adjacent, we simply return.
        adj = self
        if not self.is_adjacent(p):
            if self.is_adjacent(adj := Point(p.x, p.y + 1)):
                pass
            elif self.is_adjacent(adj := Point(p.x, p.y - 1)):
                pass
            elif self.is_adjacent(adj := Point(p.x + 1, p.y)):
                pass
            elif self.is_adjacent(adj := Point(p.x - 1, p.y)):
                pass
            elif self.is_adjacent(adj := Point(p.x + 1, p.y + 1)):
                pass
            elif self.is_adjacent(adj := Point(p.x - 1, p.y + 1)):
                pass
            elif self.is_adjacent(adj := Point(p.x + 1, p.y - 1)):
                pass
            elif self.is_adjacent(adj := Point(p.x - 1, p.y - 1)):
                pass
        return adj


class RopeBridge:
    def __init__(self, motions: list[tuple[str, int]]) -> None:
        self.head = Point(0, 0)
        self.tail = Point(0, 0)
        self.tail_list: list[Point] = []
        self.motions = motions
        self.visited = {self.tail}
        for i in range(0, 9):
            self.tail_list.append(Point(0, 0))

    def puzzle_one(self) -> int:
        self.visited = {self.tail}
        for move in self.motions:
            for i in range(0, move[1]):
                match move[0]:
                    case "U":
                        self.head.y += 1
                    case "D":
                        self.head.y -= 1
                    case "R":
                        self.head.x += 1
                    case "L":
                        self.head.x -= 1
                self.tail = self.tail.become_adjacent(self.head)
                self.visited.add(self.tail)
        return len(self.visited)

    def puzzle_two(self) -> int:
        self.head = Point(0, 0)
        self.visited = {Point(0, 0)}
        for move in self.motions:
            for i in range(0, move[1]):
                match move[0]:
                    case "U":
                        self.head.y += 1
                    case "D":
                        self.head.y -= 1
                    case "R":
                        self.head.x += 1
                    case "L":
                        self.head.x -= 1
                self.tail_list[0] = self.tail_list[0].become_adjacent(self.head)
                for j in range(1, 9):
                    self.tail_list[j] = self.tail_list[j].become_adjacent(
                        self.tail_list[j - 1]
                    )
                self.visited.add(self.tail_list[8])
        return len(self.visited)


def motions(input: str) -> list[tuple[str, int]]:
    motions: list[tuple[str, int]] = []
    for line in input.splitlines():
        move = line.split()
        motions.append((move[0], int(move[1])))
    return motions


asyncio.run(
    main(
        2022,
        9,
        lambda input: RopeBridge(motions(input)).puzzle_one(),
        lambda input: RopeBridge(motions(input)).puzzle_two(),
    )
)
