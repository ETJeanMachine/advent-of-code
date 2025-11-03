import os
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
            return await response.text()
