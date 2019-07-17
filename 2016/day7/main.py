#!/usr/bin/python3

import sys
import re

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
        if supportsTLS(line):
            result += 1

    return result


def calculateResultB(input):
    result = 0

    for line in input:
        if supportsSSL(line):
            result += 1

    return result


def supportsTLS(line):
    parts = re.split(r'\[|\]', line)
    result = False
    for key, segment in enumerate(parts):
        if containsABBA(segment):
            if key % 2 == 1:
                return False
            else:
                result = True
    return result


def containsABBA(string):
    for i in range(len(string)-3):
        if string[i] != string[i+1] and string[i] == string[i+3] and string[i+1] == string[i+2]:
            return True
    return False


def supportsSSL(line):
    parts = re.split(r'\[|\]', line)
    supernet = parts[::2]
    hypernet = parts[1::2]
    abas = []
    for segment in supernet:
        abas += findABAs(segment)

    if len(abas) == 0:
        return False

    for segment in hypernet:
        for aba in abas:
            if matchesABA(segment, aba):
                return True

    return False


def findABAs(string):
    result = []
    for i in range(len(string)-2):
        if string[i] != string[i+1] and string[i] == string[i+2]:
            result.append(string[i:i+3])
    return result


def matchesABA(segment, aba):
    bab = aba[1]+aba[0]+aba[1]
    if str.find(segment, bab) == -1:
        return False
    return True


if __name__ == "__main__":
    if len(sys.argv) < 2:
        part = "A"
    else:
        part = sys.argv[1]

    print(getResult(part, file))
