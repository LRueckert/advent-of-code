#!/usr/bin/python3

import sys

file = "input"

directions = {
    "U": (-1, 0),
    "D": (1, 0),
    "R": (0, 1),
    "L": (0, -1)
}


def getResult(part, file):

    firstPart = part == "A"

    with open(file, mode="rt", encoding="utf-8") as f:
        input = f.read().splitlines()

    if firstPart:
        result = calculateResultA(input)
    else:
        result = calculateResultB(input)

    return result


def calculateResultA(input):

    result = ""
    position = (1, 1)
    numPad = [[x + 3*i for x in range(1, 4)] for i in range(3)]

    for steps in input:
        for step in steps:
            position = tuple(map(sum, zip(position, directions[step])))
            position = (min(max(0, position[0]), 2), min(max(0, position[1]), 2))
        result += str(numPad[position[0]][position[1]])

    return result


transitions = {
    "1": {
        "U": "1",
        "D": "3",
        "R": "1",
        "L": "1"
    },
    "2": {
        "U": "2",
        "D": "6",
        "R": "3",
        "L": "2"
    },
    "3": {
        "U": "1",
        "D": "7",
        "R": "4",
        "L": "2"
    },
    "4": {
        "U": "4",
        "D": "8",
        "R": "4",
        "L": "3"
    },
    "5": {
        "U": "5",
        "D": "5",
        "R": "6",
        "L": "5"
    },
    "6": {
        "U": "2",
        "D": "A",
        "R": "7",
        "L": "5"
    },
    "7": {
        "U": "3",
        "D": "B",
        "R": "8",
        "L": "6"
    },
    "8": {
        "U": "4",
        "D": "C",
        "R": "9",
        "L": "7"
    },
    "9": {
        "U": "9",
        "D": "9",
        "R": "9",
        "L": "8"
    },
    "A": {
        "U": "6",
        "D": "A",
        "R": "B",
        "L": "A"
    },
    "B": {
        "U": "7",
        "D": "D",
        "R": "C",
        "L": "A"
    },
    "C": {
        "U": "8",
        "D": "C",
        "R": "C",
        "L": "B"
    },
    "D": {
        "U": "B",
        "D": "D",
        "R": "D",
        "L": "D"
    },
}


def calculateResultB(input):
    position = "5"
    result = ""

    for steps in input:
        for step in steps:
            position = transitions[position][step]
        result += position

    return result


if __name__ == "__main__":
    if len(sys.argv) < 2:
        part = "A"
    else:
        part = sys.argv[1]

    print(getResult(part, file))
