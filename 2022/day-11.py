import asyncio
import re

from util import main


class Monkey:
    def __init__(
        self,
        id: int,
        items: list[int],
        op: str,
        op_val: str,
        div_by: int,
        test_t: int,
        test_f: int,
    ) -> None:
        self.id = id
        self.items = items
        self.inspecting = -1
        self.inspect_count = 0
        self.div_by = div_by
        self.__op = op
        self.__op_val = op_val
        self.__test_t = test_t
        self.__test_f = test_f

    def inspect(self):
        self.inspecting = self.items.pop(0)
        self.inspect_count += 1

    def operation(self):
        val: int
        if self.__op_val == "old":
            val = self.inspecting
        else:
            val = int(self.__op_val)
        if self.__op == "*":
            self.inspecting *= val
        elif self.__op == "+":
            self.inspecting += val

    def bored(self):
        self.inspecting //= 3

    def test(self) -> int:
        if self.inspecting % self.div_by == 0:
            throw_to = self.__test_t
        else:
            throw_to = self.__test_f
        return throw_to

    @property
    def inspected_str(self) -> str:
        return f"Monkey {self.id} inspected items {self.inspect_count} times."


def puzzle(monkeys: list[Monkey], rounds: int, get_bored: bool = True) -> int:
    # the largest value we can get worried by.
    max_worry = 1
    for m in monkeys:
        max_worry *= m.div_by
    # running the rounds of the monkey business.
    for i in range(rounds):
        for m in monkeys:
            num_items, j = len(m.items), 0
            while j < num_items:
                m.inspect()
                m.inspecting %= max_worry
                m.operation()
                if get_bored:
                    m.bored()
                throw_to = m.test()
                monkeys[throw_to].items.append(m.inspecting)
                j += 1
    largest, largest_2 = 0, 0
    for m in monkeys:
        if m.inspect_count > largest:
            largest_2 = largest
            largest = m.inspect_count
        elif m.inspect_count > largest_2:
            largest_2 = m.inspect_count
        # print(m.inspected_str)
    return largest * largest_2


def process_input(input: str) -> list[Monkey]:
    monkeys: list[Monkey] = []
    matches: list[str] = re.findall(r"Monkey \d+:\n((?:.+\n?)+)", input)
    for match in matches:
        lines = match.splitlines()
        items = [int(n) for n in re.findall(r"(\d+)", lines[0])]
        op_arr = lines[1].split()
        op, op_val = op_arr[4], op_arr[5]
        test = lines[2] + lines[3] + lines[4]
        test_match = re.findall(r"(\d+)", test)
        div_by, test_t, test_f = (
            int(test_match[0]),
            int(test_match[1]),
            int(test_match[2]),
        )
        monkey = Monkey(len(monkeys), items, op, op_val, div_by, test_t, test_f)
        monkeys.append(monkey)
    return monkeys


asyncio.run(
    main(
        2022,
        11,
        lambda input: puzzle(process_input(input), 20),
        lambda input: puzzle(process_input(input), 10000, False),
    )
)
