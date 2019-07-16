#!/usr/bin/python3

import sys
from collections import OrderedDict

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
        parts = str.split(line, "-")
        letters = sorted(list("".join(parts[:-1])))
        sectorId = int(str.split(parts[-1], "[")[0])
        checksum = str.split(parts[-1], "[")[1][:-1]

        calculatedChecksum = calculateChecksum(letters)

        if calculatedChecksum == checksum:
            result += sectorId

    return result

def calculateChecksum(letters):
    letterDict = OrderedDict()
    for letter in letters:
        letterDict[letter] = letters.count(letter)

    letterIndexes = list(letterDict.keys())
    calculatedChecksum = ""
    indexAndCount = list(enumerate(letterDict.values()))
    indexAndCount.sort(key=lambda x: x[1], reverse=True)
    for i in range(5):
        letterIndex = indexAndCount[i][0]
        letter = letterIndexes[letterIndex]
        calculatedChecksum += letterIndexes[indexAndCount[i][0]]
    return calculatedChecksum


def calculateResultB(input):

    filteredInput = []
    for line in input:
        parts = str.split(line, "-")
        letters = sorted(list("".join(parts[:-1])))
        sectorId = int(str.split(parts[-1], "[")[0])
        checksum = str.split(parts[-1], "[")[1][:-1]

        calculatedChecksum = calculateChecksum(letters)

        if calculatedChecksum == checksum:
            filteredInput.append(str.split(line, "[")[0])

    for line in filteredInput:
        if decryptRoom(line) == "northpole object storage":
            return int(str.split(line, "-")[-1])

    return 0

def decryptRoom(room):
    sectorId = int(str.split(room, "-")[-1])
    words = str.split(room, "-")[:-1]
    decrypted = [ "".join([chr((ord(y)-97+sectorId)%26+97) for y in x ]) for x in words]

    return " ".join(decrypted)

if __name__ == "__main__":
    if len(sys.argv) < 2:
        part = "A"
    else:
        part = sys.argv[1]

    print(getResult(part, file))
