import asyncio
import re
from ast import literal_eval
from typing import List, Tuple

from util import main


class Packets:
    def __init__(self) -> None:
        self.pairs: List[Tuple[list, list]] = []
        self.sorted: List = [[[2]], [[6]]]

    def add_pair(self, left: List, right: List):
        self.pairs.append((left, right))
        self.sorted.append(left)
        self.sorted.append(right)

    def compare(self, left: int | list, right: int | list) -> bool | None:
        if isinstance(left, int) and isinstance(right, int):
            if left < right:
                return True
            elif left > right:
                return False
            else:
                return None
        elif isinstance(left, list) and isinstance(right, list):
            i = 0
            while len(left) != i and len(right) != i:
                l_val, r_val = left[i], right[i]
                _comp = self.compare(l_val, r_val)
                if _comp is not None:
                    return _comp
                i += 1
            if len(left) < len(right):
                return True
            elif len(left) > len(right):
                return False
            else:
                return None
        elif isinstance(left, int):
            return self.compare([left], right)
        else:
            return self.compare(left, [right])

    def compare_all(self):
        _sum = 0
        for i in range(len(self.pairs)):
            pair = self.pairs[i]
            if self.compare(pair[0], pair[1]):
                _sum += i + 1
        return _sum

    def sort(self):
        def merge_sort(arr):
            if len(arr) > 1:
                mid = len(arr) // 2
                _l = arr[:mid]
                _r = arr[mid:]
                merge_sort(_l)
                merge_sort(_r)
                i = j = k = 0
                while i < len(_l) and j < len(_r):
                    if self.compare(_l[i], _r[j]):
                        arr[k] = _l[i]
                        i += 1
                    else:
                        arr[k] = _r[j]
                        j += 1
                    k += 1
                while i < len(_l):
                    arr[k] = _l[i]
                    i += 1
                    k += 1
                while j < len(_r):
                    arr[k] = _r[j]
                    j += 1
                    k += 1

        merge_sort(self.sorted)
        return (self.sorted.index([[2]]) + 1) * (self.sorted.index([[6]]) + 1)


def parse_input(input: str) -> Packets:
    packets = Packets()
    matches: list[tuple[str, str]] = re.findall(r"(\[.*\])\n(\[.*\])\n?", input)
    for match in matches:
        packets.add_pair(literal_eval(match[0]), literal_eval(match[1]))
    return packets


asyncio.run(
    main(
        2022,
        13,
        lambda input: parse_input(input).compare_all(),
        lambda input: parse_input(input).sort(),
    )
)
