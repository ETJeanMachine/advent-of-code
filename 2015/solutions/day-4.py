import asyncio
import hashlib
from util import main


def part_one(input: str) -> int:
    i = 1
    while True:
        key = f"{input}{i}"
        hash = hashlib.md5(key.encode())
        hash_hex = hash.hexdigest()
        if hash_hex.startswith("00000"):
            break
        else:
            i += 1
    return i


def part_two(input: str) -> int:
    i = 1
    while True:
        key = f"{input}{i}"
        hash = hashlib.md5(key.encode())
        hash_hex = hash.hexdigest()
        if hash_hex.startswith("000000"):
            break
        else:
            i += 1
    return i


asyncio.run(main(2015, 4, part_one, part_two))
