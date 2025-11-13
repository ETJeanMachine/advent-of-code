import asyncio
import re
from dataclasses import dataclass

from util import async_main


@dataclass
class Reindeer:
    speed: int
    endurance: int
    rest: int
    fly_time: int = 0
    rest_time: int = 0
    resting: bool = False
    dist_travelled: int = 0
    score: int = 0


async def fly_reindeer(reindeer: Reindeer, time: int):
    dist_travelled = 0
    resting = False
    fly_time, rest_time = 0, 0
    for _ in range(time):
        if not resting:
            rest_time = 0
            dist_travelled += reindeer.speed
            fly_time += 1
            resting = fly_time == reindeer.endurance
        else:
            fly_time = 0
            rest_time += 1
            resting = rest_time != reindeer.rest
    return dist_travelled


def parse_input(input: str) -> dict[str, Reindeer]:
    reindeer = {}
    for line in input.splitlines():
        groups = re.findall(
            r"([A-Z][a-z]+).+(?:fly (\d+) km\/s) (?:for (\d+) sec).+(?:for (\d+) sec).+$",
            line,
        )[0]
        reindeer[groups[0]] = Reindeer(int(groups[1]), int(groups[2]), int(groups[3]))
    return reindeer


async def part_one(input: str) -> int:
    all_reindeer = parse_input(input)
    time = 2503
    async with asyncio.TaskGroup() as tg:
        tasks: dict[str, asyncio.Task[int]] = {}
        for k, v in all_reindeer.items():
            tasks[k] = tg.create_task(fly_reindeer(v, time))
    max_dist = 0
    for v in tasks.values():
        max_dist = max(v.result(), max_dist)
    return max_dist


async def part_two(input: str) -> int:
    all_reindeer = parse_input(input)
    time = 2503
    max_score = 0
    for _ in range(time):
        for r in all_reindeer.values():
            if not r.resting:
                r.rest_time = 0
                r.dist_travelled += r.speed
                r.fly_time += 1
                r.resting = r.fly_time == r.endurance
            else:
                r.fly_time = 0
                r.rest_time += 1
                r.resting = r.rest_time != r.rest
        leading: list[str] = []
        for k, v in all_reindeer.items():
            if len(leading) == 0:
                leading.append(k)
            elif v.dist_travelled > all_reindeer[leading[0]].dist_travelled:
                leading.clear()
                leading.append(k)
            elif v.dist_travelled == all_reindeer[leading[0]].dist_travelled:
                leading.append(k)
        for r in leading:
            all_reindeer[r].score += 1
    for r in all_reindeer.values():
        max_score = max(r.score, max_score)
    return max_score


asyncio.run(main=async_main(2015, 14, part_one, part_two))
