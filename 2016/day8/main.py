#!/usr/bin/python3

import pprint
import sys
import re

file = "input"
width = 50
tall = 6
pp = pprint.PrettyPrinter(4, 300).pprint


class Screen:

    def __init__(self, width, tall):
        self.width = width
        self.tall = tall
        self.data = [[0 for y in range(width)] for x in range(tall)]

    def print(self):
        printable = ["".join(["#" if cell==1 else "." for cell in row]) for row in self.data]
        pp(printable)

    def on(self, width, height):
        for y in range(height):
            for x in range(width):
                self.data[y][x] = 1

    def rotateRow(self, row, value):
        self.data[row] = self.data[row][-value:] + self.data[row][:-value]

    def rotateColumn(self, col, value):
        transposed = [list(x) for x in zip(*self.data)]
        transposed[col] = transposed[col][-value:] + transposed[col][:-value]
        self.data = [list(x) for x in zip(*transposed)]

    def value(self):
        return sum([sum(x) for x in self.data])


def getResult(part, file, width, tall):

    firstPart = part == "A"

    with open(file, mode="rt", encoding="utf-8") as f:
        input = f.read().splitlines()

    return calculateResultA(input, width, tall)


def calculateResultA(input, width, tall):
    print()
    screen = Screen(width, tall)

    for command in input:
        onMatch = re.search(r"rect (?P<width>\d+)x(?P<height>\d+)", command)
        if onMatch:
            screen.on(int(onMatch.group("width")), int(onMatch.group("height")))

        rowMatch = re.search(r"rotate row y=(?P<row>\d+) by (?P<value>\d+)", command)
        if rowMatch:
            screen.rotateRow(int(rowMatch.group("row")), int(rowMatch.group("value")))

        colMatch = re.search(r"rotate column x=(?P<col>\d+) by (?P<value>\d+)", command)
        if colMatch:
            screen.rotateColumn(int(colMatch.group("col")), int(colMatch.group("value")))

    screen.print()
    return screen.value()


def calculateResultB(input, width, tall):
    return 0


if __name__ == "__main__":
    if len(sys.argv) < 2:
        part = "A"
    else:
        part = sys.argv[1]

    print(getResult(part, file, width, tall))
