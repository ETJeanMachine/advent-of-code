import asyncio
from typing import override
from util import main


class Directory:
    @property
    def full_path(self) -> str:
        """The full path to the directory."""
        dir = self
        path = dir.name
        while dir.parent is not None:
            path = f"{dir.parent.name}/{path}"
            dir = dir.parent
        return path

    @property
    def file_sizes(self) -> int:
        """The size of all files directly contained within this directory."""
        file_sizes = 0
        for child in self.children.values():
            if isinstance(child, int):
                file_sizes += child
        return file_sizes

    def __init__(self, name: str, parent: Directory | None = None):
        self.name: str = name
        self.parent: Directory | None = parent
        self.children: dict[str, Directory | int] = dict()

    @override
    def __eq__(self, __o: object) -> bool:
        if isinstance(__o, Directory):
            return self.name == __o.name and self.parent == __o.parent
        return False

    @override
    def __hash__(self) -> int:
        return self.full_path.__hash__()


class FileSystem:
    def __init__(self) -> None:
        self.source: Directory = Directory("/")
        self.curr_dir: Directory = self.source

    def cd(self, loc: str):
        """Navigates the file system via the cd command.

        Args:
            loc (str): The location that cd is directing us to.
        """
        if loc == "/":
            self.curr_dir = self.source
        elif loc == ".." and self.curr_dir.parent:
            self.curr_dir = self.curr_dir.parent
        else:
            dir = self.curr_dir.children.get(loc)
            if not isinstance(dir, Directory):
                dir = Directory(loc, self.curr_dir)
                self.curr_dir.children[loc] = dir
            self.curr_dir = dir

    def ls(self, item: str):
        """Adds an item to a directory after it has been recognized by the ls command.

        Args:
            item (str): The directory or file that ls has listed.
        """
        split = item.split()
        if split[0] == "dir":
            self.curr_dir.children[split[1]] = Directory(split[1], self.curr_dir)
        else:
            self.curr_dir.children[split[1]] = int(split[0])

    def puzzle_one(self):
        visited: set[Directory] = set()
        total = 0

        def dfs(dir: Directory) -> int:
            visited.add(dir)
            size = dir.file_sizes
            nonlocal total
            for child in dir.children.values():
                if isinstance(child, Directory) and child not in visited:
                    size += dfs(child)
            total += size if size <= 100000 else 0
            return size

        _ = dfs(self.source)
        return total

    def puzzle_two(self):
        visited: set[Directory] = set()
        all_sizes: list[int] = []

        def dfs(dir: Directory) -> int:
            visited.add(dir)
            size = dir.file_sizes
            for child in dir.children.values():
                if isinstance(child, Directory) and child not in visited:
                    size += dfs(child)
            all_sizes.append(size)
            return size

        # Somehow condense this into the dfs function.
        available_space = 70000000 - dfs(self.source)
        min_size = 70000000
        for size in all_sizes:
            if size + available_space >= 30000000:
                min_size = min(size, min_size)
        return min_size


def construct_filesystem(input: str) -> FileSystem:
    """Reads from input and constructs the filesystem off of the input file.

    Returns:
        FileSystem: The virtual filesystem generated.
    """
    filesystem = FileSystem()
    for line in input.split("\n"):
        split = line.split()
        if split[0] == "$":
            if split[1] == "cd":
                filesystem.cd(split[2])
            else:
                continue
        else:
            filesystem.ls(line)
    return filesystem


asyncio.run(
    main(
        2022,
        7,
        lambda input: construct_filesystem(input).puzzle_one(),
        lambda input: construct_filesystem(input).puzzle_two(),
    )
)
