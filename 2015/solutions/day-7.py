import asyncio
import re
from dataclasses import dataclass

from util import main


@dataclass
class Gate:
    op: str
    inputs: list[str | int]


@dataclass
class Wire:
    value: int = -1
    input: None | Gate | str = None


def process_wires(input: str) -> dict[str, Wire]:
    wires: dict[str, Wire] = {}
    for line in input.split("\n"):
        split = line.split("-> ")
        input, output = split[0], split[1]

        # Regex fun stuff
        input_match: list[str] = re.findall(r"([a-z]+|\d+)", input)
        inputs = [int(n) if n.isdigit() else n for n in input_match]
        op_match: list[str] = re.findall(r"[A-Z]+", input)

        if len(op_match) > 0:
            # This input is a gate
            op = op_match[0]
            gate = Gate(op, inputs)
            wires[output] = Wire(input=gate)
        elif isinstance(inputs[0], int):
            # This input is a fixed value
            wires[output] = Wire(value=inputs[0])
        else:
            # This input is another wire
            wires[output] = Wire(input=inputs[0])
    return wires


def solve_circuit(wires: dict[str, Wire]):
    def gate_helper(input: str | int) -> int:
        if isinstance(input, str):
            return solve_helper(input)
        else:
            return input

    def solve_helper(key: str) -> int:
        wire = wires[key]
        if wire.input is None or wire.value != -1:
            return wire.value
        if isinstance(wire.input, Gate):
            inputs, op = wire.input.inputs, wire.input.op
            match op:
                case "NOT":
                    wires[key].value = ~gate_helper(inputs[0])
                case "AND":
                    wires[key].value = gate_helper(inputs[0]) & gate_helper(inputs[1])
                case "OR":
                    wires[key].value = gate_helper(inputs[0]) | gate_helper(inputs[1])
                case "LSHIFT":
                    wires[key].value = gate_helper(inputs[0]) << gate_helper(inputs[1])
                case "RSHIFT":
                    wires[key].value = gate_helper(inputs[0]) >> gate_helper(inputs[1])
                case _:
                    print("Oopsies D:")
            return wires[key].value
        else:
            wires[key].value = solve_helper(wire.input)
            return wires[key].value

    for key in wires.keys():
        _ = solve_helper(key)


def part_one(input: str) -> int:
    wires = process_wires(input)
    solve_circuit(wires)
    if a := wires["a"]:
        return a.value
    return -1


def part_two(input: str) -> int:
    wires = process_wires(input)
    solve_circuit(wires)
    if a := wires["a"]:
        wires["b"].value = a.value
        for key in wires.keys():
            if wires[key].input is not None:
                wires[key].value = -1
    solve_circuit(wires)
    if a := wires["a"]:
        return a.value
    return -1


asyncio.run(main(2015, 7, part_one, part_two))
