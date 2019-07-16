#!/usr/bin/python3

import sys

file = "input"


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

    result = 0

    for line in input:
        numbers = list(map(int, line.split()))
        numbers.sort()
        if numbers[0] + numbers[1] > numbers[2]:
            result += 1

    return result


def calculateResultB(input):
    result = 0

    it = iter(input)
    for one in it:
        two = next(it)
        three = next(it)
        numbersOne = list(map(int, one.split()))
        numbersTwo = list(map(int, two.split()))
        numbersThree = list(map(int, three.split()))
        numbers = [numbersOne, numbersTwo, numbersThree]
        numbers = [[numbers[inner][outer] for inner in range(3) ] for outer in range(3)]
        for col in numbers:
            col.sort()
            if col[0] + col[1] > col[2]:
                result += 1

    return result


if __name__ == "__main__":
    if len(sys.argv) < 2:
        part = "A"
    else:
        part = sys.argv[1]

    print(getResult(part, file))
