import asyncio
from util import main


def part_one(input: str) -> int:
    code_count, mem_count = 0, 0
    for string in input.split():
        code_count += len(string)
        string = string[1:-1]
        decode_str = bytes(string, "utf-8").decode("unicode-escape")
        mem_count += len(decode_str)
    return code_count - mem_count


def part_two(input: str) -> int:
    original_count, encoded_count = 0, 0
    for string in input.split():
        original_count += len(string)
        encode_str = string.replace("\\", "\\\\").replace('"', '\\"')
        encode_str = f'"{encode_str}"'
        encoded_count += len(encode_str)
    return encoded_count - original_count


asyncio.run(main(2015, 8, part_one, part_two))
