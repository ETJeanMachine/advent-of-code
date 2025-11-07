import os
import time
from collections.abc import Callable

import aiohttp
from dotenv import load_dotenv

_ = load_dotenv()


async def get_input(year: int, day: int) -> str:
    proxies: dict[str, str] = {}
    cookies: dict[str, str] = {}

    if http_proxy := os.getenv("HTTP_PROXY") or os.getenv("http_proxy"):
        proxies["http"] = http_proxy

    if https_proxy := os.getenv("HTTPS_PROXY") or os.getenv("https_proxy"):
        proxies["https"] = https_proxy

    if session_cookie := os.getenv("SESSION_COOKIE"):
        cookies["session"] = session_cookie

    connector = aiohttp.TCPConnector() if not proxies else None
    async with aiohttp.ClientSession(connector=connector) as session:
        async with session.get(
            f"https://adventofcode.com/{year}/day/{day}/input",
            proxy=proxies.get("https") if proxies else None,
            cookies=cookies if cookies else None,
        ) as response:
            input = await response.text()
            input = input.rstrip()
            return input


PartFn = Callable[[str], int | float | str]


async def main(year: int, day: int, part_one: PartFn, part_two: PartFn):
    input = await get_input(year, day)
    print(f"Solutions for {year} Day {day}:\n")
    time_one = time.perf_counter_ns()
    print(f"Part One: {part_one(input)}")
    time_one = time.perf_counter_ns() - time_one
    print(f"Time One: {time_one / 1e6:.2f}ms\n")
    time_two = time.perf_counter_ns()
    print(f"Part Two: {part_two(input)}")
    time_two = time.perf_counter_ns() - time_two
    print(f"Time Two: {time_two / 1e6:.2f}ms")
