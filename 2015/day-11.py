import asyncio
import re

from util import main


def base_10(num: str) -> int:
    table = {}
    base = ord("a")
    for i in range(26):
        table[chr(base + i)] = i
    res = 0
    for d in num:
        res = res * 26 + table[d]
    return res


def base_26(num: int) -> str:
    table = []
    base = ord("a")
    for i in range(26):
        table.append(chr(base + i))
    res = ""
    dec = num
    while dec > 0:
        rem = dec % 26
        res = f"{table[rem]}{res}"
        dec //= 26
    return res


def is_valid(password: str) -> bool:
    straight = False
    for i in range(len(password) - 2):
        c1, c2, c3 = ord(password[i]), ord(password[i + 1]), ord(password[i + 2])
        straight = c1 == c2 - 1 == c3 - 2
        if straight:
            break
    banned_letters = re.match(r"i|o|l", password) is None
    repeat_match = re.findall(r"(.)\1.*(.)\2", password)
    repeat = False
    for match in repeat_match:
        repeat = match[0] != match[1]
        if repeat:
            break
    return straight and banned_letters and repeat


def part_one(input: str) -> str:
    next_pass = input
    while True:
        pass_int = base_10(next_pass)
        next_pass = base_26(pass_int + 1)
        if is_valid(next_pass):
            return next_pass


def part_two(input: str) -> str:
    next_pass = part_one(input)
    while True:
        pass_int = base_10(next_pass)
        next_pass = base_26(pass_int + 1)
        if is_valid(next_pass):
            return next_pass


asyncio.run(main(2015, 11, part_one, part_two))
