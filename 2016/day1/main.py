#!/usr/bin/python3

import sys
import pprint
from enum import Enum

file = "input"


class Direction(Enum):
    North = 1
    East = 2
    South = 3
    West = 4


def getResult(part, file):

    firstPart = part == "A"

    with open("input", mode="rt", encoding="utf-8") as f:
        file = f.read().strip()
        input = str.split(file, ", ")

    if firstPart:
        result = calculateResultA(input)
    else:
        result = calculateResultB(input)

    return result


def calculateResultA(input):
    facing = Direction.North
    position = (0, 0)

    for step in input:
        (length, direction, facing) = move(position, facing, step)
        movement = tuple(length * i for i in direction)
        position = tuple(map(sum,zip(position,movement)))

    return abs(position[0]) + abs(position[1])


def calculateResultB(input):
    facing = Direction.North
    pp = pprint.PrettyPrinter(4)
    position = (0, 0)
    grid = [x[:] for x in [[False] * 1000] * 1000]
    grid[0][0] = True
    print()

    for step in input:
        (length, direction, facing) = move(position, facing, step)
        for i in range(1, length+1):
            position = tuple(map(sum,zip(position,direction)))
            if grid[position[0]][position[1]] == True:
                return abs(position[0]) + abs(position[1])
            else: grid[position[0]][position[1]] = True


def move(position, facing, step):
    direction = step[0]
    length = int(step[1:])

    if direction == "R" and facing == Direction.North:
        movement = (1, 0)
        facing = Direction.East
    elif direction == "R" and facing == Direction.East:
        movement = (0, -1)
        facing = Direction.South
    elif direction == "R" and facing == Direction.South:
        movement = (-1, 0)
        facing = Direction.West
    elif direction == "R" and facing == Direction.West:
        movement = (0, 1)
        facing = Direction.North
    elif direction == "L" and facing == Direction.North:
        movement = (-1, 0)
        facing = Direction.West
    elif direction == "L" and facing == Direction.East:
        movement = (0, 1)
        facing = Direction.North
    elif direction == "L" and facing == Direction.South:
        movement = (1, 0)
        facing = Direction.East
    elif direction == "L" and facing == Direction.West:
        movement = (0, -1)
        facing = Direction.South

    return (length, movement, facing)


if __name__ == "__main__":
    if len(sys.argv) < 2:
        part = "A"
    else:
        part = sys.argv[1]

    print(getResult(part, file))
