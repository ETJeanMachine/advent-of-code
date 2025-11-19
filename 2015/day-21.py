import asyncio
import math
import re

from util import main

weapons = {
    "Dagger": (8, 4),
    "Shortsword": (10, 5),
    "Warhammer": (25, 6),
    "Longsword": (40, 7),
    "Greataxe": (74, 8),
}
armor = {
    "None": (0, 0),
    "Leather": (13, 1),
    "Chainmail": (31, 2),
    "Splintmail": (53, 3),
    "Bandedmail": (75, 4),
    "Platemail": (102, 5),
}
rings = {
    "None": (0, 0, 0),
    "Damage +1": (25, 1, 0),
    "Damage +2": (50, 2, 0),
    "Damage +3": (100, 3, 0),
    "Defense +1": (20, 0, 1),
    "Defense +2": (40, 0, 2),
    "Defense +3": (80, 0, 3),
}


def fight(player: list[int], boss: list[int]):
    player_turn: bool = True

    while player[0] > 0:
        attacker, defender = (player, boss) if player_turn else (boss, player)
        damage = attacker[1] - defender[2]
        defender[0] -= damage if damage > 1 else 1

        if boss[0] <= 0:
            return True

        player_turn = not player_turn

    return False


def part_one(input: str):
    min_gold = math.inf
    loadout: list[str] = []
    for wk, wv in weapons.items():
        for ak, av in armor.items():
            for r1k, r1v in rings.items():
                for r2k, r2v in rings.items():
                    if r1k == r2k and r1k != "None":
                        continue
                    gold = wv[0] + av[0] + r1v[0] + r2v[0]
                    if gold >= min_gold:
                        continue
                    p_damage = wv[1] + r1v[1] + r2v[1]
                    p_armor = av[1] + r1v[2] + r2v[2]
                    player = [100, p_damage, p_armor]
                    boss = [int(x) for x in re.findall(r"\d+", input)]
                    if fight(player, boss):
                        loadout = [wk, ak, r1k, r2k]
                        min_gold = gold
    print(loadout)
    return min_gold


def part_two(input: str) -> int:
    max_gold = 0
    loadout: list[str] = []
    for wk, wv in weapons.items():
        for ak, av in armor.items():
            for r1k, r1v in rings.items():
                for r2k, r2v in rings.items():
                    if r1k == r2k and r1k != "None":
                        continue
                    gold = wv[0] + av[0] + r1v[0] + r2v[0]
                    if gold <= max_gold:
                        continue
                    p_damage = wv[1] + r1v[1] + r2v[1]
                    p_armor = av[1] + r1v[2] + r2v[2]
                    player = [100, p_damage, p_armor]
                    boss = [int(x) for x in re.findall(r"\d+", input)]
                    if not fight(player, boss):
                        loadout = [wk, ak, r1k, r2k]
                        max_gold = gold
    print(loadout)
    return max_gold


asyncio.run(main(2015, 21, part_one, part_two))
