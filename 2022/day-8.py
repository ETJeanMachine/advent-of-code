import asyncio

from util import main


class TreeMap:
    def __init__(self, map: list[list[int]]) -> None:
        """Generates a tree map of which we can run the puzzles off of.

        Args:
            map (List[List[int]]): A 2D array of integers representing the height
            values of the trees.
        """
        self.map = map
        self.num_rows = len(map)
        self.num_cols = len(map[0])
        self.visible = set()

    def __is_edge(self, row, col) -> bool:
        """Checks if a given coordinate is an edge row.

        Args:
            row (int): The row of the tree.
            col (int): The column of the tree.

        Returns:
            bool: Whether the row or column is along the edge of the map.
        """
        return (
            row == 0 or col == 0 or row == self.num_rows - 1 or col == self.num_cols - 1
        )

    def __scenic_score(self, row, col) -> int:
        """Calculates the scenic score from a given coordinate on the map, calculated
        by multiplying together the viewing distance in all four cardinal directions.

        Args:
            row (int): The row of the tree.
            col (int): The column of the tree.

        Returns:
            int: The scenic score of the tree.
        """
        # Passing immediately if it's an edge.
        if self.__is_edge(row, col):
            return 0
        height, i = self.map[row][col], 1
        check_u = check_d = check_r = check_l = True
        u_score = d_score = r_score = l_score = 0
        while check_u or check_d or check_l or check_r:
            if check_u:
                check_u = (self.map[row][col + i] < height) and not self.__is_edge(
                    row, col + i
                )
                u_score += 1
            if check_d:
                check_d = (self.map[row][col - i] < height) and not self.__is_edge(
                    row, col - i
                )
                d_score += 1
            if check_r:
                check_r = (self.map[row + i][col] < height) and not self.__is_edge(
                    row + i, col
                )
                r_score += 1
            if check_l:
                check_l = (self.map[row - i][col] < height) and not self.__is_edge(
                    row - i, col
                )
                l_score += 1
            i += 1
        return u_score * d_score * r_score * l_score

    def __look_from_dir(self, dir):
        """Observes the map via a cardinal direction as a person would. Adds the
        visible trees to an internal set of visible trees, if they have not already
        been observed.

        Args:
            dir (str): A character representing the direction we are observing.
        """
        end_i = self.num_cols if dir == "N" or "S" else self.num_rows
        end_j = self.num_rows if dir == "N" or "S" else self.num_cols
        for i in range(0, end_i):
            seen = set()
            largest_seen: int = 0
            for j in range(0, end_j):
                row, col = i, j
                if dir == "E" or dir == "S":
                    col = (end_j - 1) - col
                if dir == "N" or dir == "S":
                    row, col = col, row
                if self.__is_edge(row, col):
                    seen.add((row, col))
                    largest_seen = self.map[row][col]
                elif self.map[row][col] > largest_seen:
                    seen.add((row, col))
                    largest_seen = self.map[row][col]
                if largest_seen == 9:
                    break
            self.visible = self.visible.union(seen)

    def puzzle_one(self) -> int:
        """Finds the number of trees that are visible by looking from the outside
        of the map.

        Returns:
            int: The number of visible trees.
        """
        self.__look_from_dir("N")
        self.__look_from_dir("S")
        self.__look_from_dir("E")
        self.__look_from_dir("W")
        return len(self.visible)

    def puzzle_two(self) -> int:
        """Finds the maximum viewing distance within the tree map, and calculates
        this via running through every tree in the map and calculating it's value.

        Returns:
            int: The maximum viewing distance in the tree map.
        """
        max_score = 0
        for row in range(0, self.num_rows):
            for col in range(0, self.num_cols):
                max_score = max(max_score, self.__scenic_score(row, col))
        # Technically the below runs twice as fast but it's not necessarily
        # provable it'll always be right, so it's commented out.
        # for coords in self.visible:
        #     max_score = max(max_score, self.scenic_score(coords[0], coords[1]))
        return max_score


def create_map(input: str) -> list[list[int]]:
    map = []
    for line in input.splitlines():
        row = []
        for h in line.strip():
            row.append(int(h))
        map.append(row)
    return map


asyncio.run(
    main(
        2022,
        8,
        lambda input: TreeMap(create_map(input)).puzzle_one(),
        lambda input: TreeMap(create_map(input)).puzzle_two(),
    )
)
