import asyncio
import re
from dataclasses import dataclass

from util import main


@dataclass
class Ingredient:
    capacity: int
    durability: int
    flavor: int
    texture: int
    calories: int


def parse_input(input: str) -> dict[str, Ingredient]:
    ingredients: dict[str, Ingredient] = {}
    for line in input.splitlines():
        if match := re.match(
            r"([A-Za-z]+):.+(-?\d+),.+(-?\d+),.+(-?\d+),.+(-?\d+).+.+(-?\d+)$", line
        ):
            groups = match.groups()
            ingredient = Ingredient(
                int(groups[1]),
                int(groups[2]),
                int(groups[3]),
                int(groups[4]),
                int(groups[5]),
            )
            ingredients[groups[0]] = ingredient
    return ingredients


def calculate_score(recipe: list[tuple[int, Ingredient]]) -> int:
    capacity, durability, flavor, texture = 0, 0, 0, 0
    for tsps, ing in recipe:
        capacity += tsps * ing.capacity
        durability += tsps * ing.durability
        flavor += tsps * ing.flavor
        texture += tsps * ing.texture
    return capacity * durability * flavor * texture


def n_adds_to_x(n: int, x: int) -> list[list[int]]:
    if n == 1:
        return [[x]]
    perms = []
    # we start at 1 because 0 will always multiply to a score
    # of 0
    for i in range(1, x + 1):
        sub_perms = n_adds_to_x(n - 1, x - i)
        for p in sub_perms:
            appending = [i]
            appending.extend(p)
            perms.append(appending)
    return perms


def part_one(input: str) -> int:
    ingredients = parse_input(input)
    values = ingredients.values()
    permutations = n_adds_to_x(4, 100)
    max_score = 0
    for p in permutations:
        recipe = list(zip(p, values))
        score = calculate_score(recipe)
        max_score = max(score, max_score)
        if max_score == score:
            print(f"new max score: {score}")
            print(f"recipe: {p}")
    return max_score


def part_two(input: str) -> int:
    return 0


asyncio.run(main(2015, 15, part_one, part_two))
