#!/usr/bin/python3

import sys
import re

file = "input"


def getResult(part, file):

    firstPart = part == "A"

    with open(file, mode="rt", encoding="utf-8") as f:
        input = f.read().splitlines()

    if firstPart:
        result = calculateResultA(input[0])
    else:
        result = calculateResultB(input[0])

    return result


def calculateResultA(compressed):
    decompressed = ""
    while True:
        match = re.search(r"\((?P<length>\d+)x(?P<times>\d+)\)", compressed)
        if match:
            matched = match.group()
            length = int(match.group("length"))
            times = int(match.group("times"))
            # print("Matched: {}, length: {}, times: {}".format(matched, length, times))
            effected = match.end() + length
            # print("{} - {}".format(decompressed, compressed), end=" ")
            decompressed += compressed[:match.start()] + (compressed[match.end():effected] * times)
            compressed = compressed[effected:]
            # print("-> {} - {}".format(decompressed, compressed))
        else:
            decompressed += compressed
            break

    return len(decompressed)


def calculateResultB(compressed):

    return getSequenceLength(compressed, 1)


def getSequenceLength(compressed, mult):
    # print()
    # print("Compressed {} - Multiplier: {}".format(compressed, mult))
    match = re.search(r"\((?P<length>\d+)x(?P<times>\d+)\)", compressed)
    if match:
        matched = match.group()
        length = int(match.group("length"))
        times = int(match.group("times"))
        # print("Matched: {}, length: {}, times: {}".format(matched, length, times))
        effected = match.end() + length
        decompressed = compressed[:match.start()]
        affected = compressed[match.end():effected]
        remaining = compressed[effected:]
        # print("Decompressed: {} - Affected: {} - Remaining: {}".format(decompressed, affected, remaining))
        lenghtAffected = getSequenceLength(affected, times)
        lengthRemaining = getSequenceLength(remaining, 1)
        # print("Returning: {} * ({} + {} + {})".format(mult, len(decompressed), lenghtAffected, lengthRemaining))
        return mult * (len(decompressed) + lenghtAffected + lengthRemaining)
    else:
        # print("Returning guard: {} * {}".format(mult, len(compressed)))
        return mult * len(compressed)


if __name__ == "__main__":
    if len(sys.argv) < 2:
        part = "A"
    else:
        part = sys.argv[1]

    print(getResult(part, file))
