#!/usr/bin/python3

import sys

file = "input"

def getResult(part, file):

    firstPart = part == "A"

    with open("input", mode="rt", encoding="utf-8") as f:
        input = f.readlines()

    if firstPart:
        result = calculateResultA(input)
    else:
        result = calculateResultB(input)

    return result

def calculateResultA(input):
    return 0


def calculateResultB(input):
    return 0

if __name__ == "__main__":
    if len(sys.argv) < 2:
        part = "A"
    else:
        part = sys.argv[1]

    print(getResult(part, file))
