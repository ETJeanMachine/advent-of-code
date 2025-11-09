import asyncio

from util import main


class CPU:
    def __init__(self):
        self.X = 1
        self.sum = 0
        self.display = ""
        self.__cycle = 0

    def cycle(self):
        self.__cycle += 1
        pos = (self.__cycle - 1) % 40
        if pos == 0:
            self.display += "\n"
        if pos in range(self.X - 1, self.X + 2):
            self.display += "#"
        else:
            self.display += " "
        if (self.__cycle - 20) % 40 == 0:
            self.sum += self.__cycle * self.X

    def addx(self, V: int):
        self.cycle()
        self.cycle()
        self.X += V


def init_cpu(input: str) -> CPU:
    cpu = CPU()
    for line in input.splitlines():
        arr = line.split()
        if arr[0] == "noop":
            cpu.cycle()
        else:
            cpu.addx(int(arr[1]))
    return cpu


asyncio.run(
    main(
        2022,
        10,
        lambda input: init_cpu(input).sum,
        lambda input: init_cpu(input).display,
    )
)
