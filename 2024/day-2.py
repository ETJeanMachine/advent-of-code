# THIS WAS ORIGINALLY DONE IN RUST AND CONVERTED TO PYTHON AFTER THE FACT

import asyncio
from util import main


def part_one(input: str) -> int:
    reports = input.split("\n")
    safe_count = len(reports)

    for report in reports:
        report = [int(n) for n in report.split()]
        dir = 0
        for i in range(0, len(report) - 1):
            step = report[i + 1] - report[i]

            if dir == 0 and step != 0:
                dir = abs(step) / step
            if step == 0 or dir != abs(step) / step or abs(step) > 3:
                safe_count -= 1
                break

    return safe_count


def part_two(input: str) -> int:
    reports = input.split("\n")
    safe_count = len(reports)
    for report in reports:
        report = [int(n) for n in report.split()]
        slopes = [report[i + 1] - report[i] for i in range(len(report) - 1)]
        sign = (sum(slopes) > 0) - (sum(slopes) < 0)

        def valid(n: int):
            return (n > 0) - (n < 0) == sign and 1 <= abs(n) <= 3

        problem_idx = next((i for i, n in enumerate(slopes) if not valid(n)), None)

        if problem_idx is not None:
            i = problem_idx

            shift_right = slopes.copy()
            right_step = shift_right.pop(i)
            if i < len(shift_right):
                shift_right[i] += right_step

            shift_left = slopes.copy()
            left_step = shift_left.pop(i)
            if i > 0:
                shift_left[i - 1] += left_step

            right_problem = any(not valid(n) for n in shift_right)
            left_problem = any(not valid(n) for n in shift_left)

            if right_problem and left_problem:
                safe_count -= 1

    return safe_count


asyncio.run(main(2024, 2, part_one, part_two))
