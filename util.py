import asyncio
import inspect
import os
import time
from collections.abc import Awaitable, Callable

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
AsyncPartFn = Callable[[str], Awaitable[int | float | str]]


async def main(
    year: int, day: int, part_one: PartFn | AsyncPartFn, part_two: PartFn | AsyncPartFn
):
    input = await get_input(year, day)
    print(f"Solutions for {year} Day {day}:\n")

    async def run_part(n: int) -> tuple[int | float | str, int]:
        fn = part_one if n == 1 else part_two
        time_start = time.perf_counter_ns()
        res: int | float | str = (
            await fn(input) if inspect.iscoroutinefunction(fn) else fn(input)
        )  # pyright:ignore
        time_elapsed = time.perf_counter_ns() - time_start
        return res, time_elapsed

    def format_time(t: int):
        if t >= 6 * 10**9:
            m = t // (60 * 10**9)
            return f"{m}m:{(t % (60 * 10**9)) / 1e9:.2f}s"
        elif t >= 10**9:
            return f"{t / 1e9:.2f}s"
        else:
            return f"{t / 1e6:.2f}ms"

    async with asyncio.TaskGroup() as tg:
        task_one = tg.create_task(run_part(1))
        task_two = tg.create_task(run_part(2))

    res_one, time_one = task_one.result()
    res_two, time_two = task_two.result()

    print(f"Part One: {res_one}")
    print(f"Time One: {format_time(time_one)}\n")
    print(f"Part Two: {res_two}")
    print(f"Time Two: {format_time(time_two)}")
