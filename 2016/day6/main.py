#!/usr/bin/python3

import sys
import pprint

file = "input"
pp = pprint.PrettyPrinter(4, 250)

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

    original = [list(x) for x in input]
    transposed = [{ y: x.count(y) for y in x} for x in zip(*original)]

    result = ""
    for item in transposed:
        result += max(item.items(), key=lambda x : x[1])[0]

    return result



def calculateResultB(input):

    original = [list(x) for x in input]
    transposed = [{ y: x.count(y) for y in x} for x in zip(*original)]

    result = ""
    for item in transposed:
        result += min(item.items(), key=lambda x : x[1])[0]

    return result


if __name__ == "__main__":
    if len(sys.argv) < 2:
        part = "A"
    else:
        part = sys.argv[1]

    print(getResult(part, file))
