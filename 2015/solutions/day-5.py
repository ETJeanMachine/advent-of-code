import asyncio
import re
from util import main


def part_one(input: str) -> int:
    count = 0
    for line in input.split():
        vowel_match = re.match(r"(?:.*[aeiou]){3,}", line)
        repeat_match = re.match(r".*(.)\1", line)
        anti_match = re.match(r".*(?:ab|cd|pq|xy)", line)
        if vowel_match and repeat_match and not anti_match:
            count += 1
    return count


def part_two(input: str) -> int:
    count = 0
    for line in input.split():
        pair_match = re.match(r".*(.{2}).*\1", line)
        between_match = re.match(r".*(.).\1", line)
        if pair_match and between_match:
            count += 1
    return count


asyncio.run(main(2015, 5, part_one, part_two))
