import asyncio
from util import main


def find_marker(input: str, n: int) -> int:
    count = 0
    marker = ""
    for c in input:
        count += 1
        if c in marker:
            idx = marker.index(c) + 1
            if idx == len(marker):
                marker = ""
            else:
                marker = marker[idx:]
        if len(marker) < n:
            marker += c
        if len(marker) == n:
            return count
    return -1


asyncio.run(
    main(
        2022,
        6,
        lambda input: find_marker(input, 4),
        lambda input: find_marker(input, 14),
    )
)
