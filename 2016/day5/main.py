#!/usr/bin/python3

import sys
from hashlib import md5


input = "ffykfhsq"

def getResult(part, input):

    firstPart = part == "A"

    if firstPart:
        result = calculateResultA(input)
    else:
        result = calculateResultB(input)

    return result

def calculateResultA(input):
    password = ""
    i = 0

    while len(password) < 8:
        hash = hashIndex(i, input)
        if hash[:5] == "00000":
            password += hash[5]
        i += 1
        if i%1e6 == 0:
            print("{} mil iterations".format(i/1e6))

    return password


def hashIndex(index, input):
    return md5((input+str(index)).encode("utf-8")).hexdigest()


def calculateResultB(input):
    password = [0 for x in range(8)]
    found = 0
    i = 0

    while found < 8:
        hash = hashIndex(i, input)
        if hash[:5] == "00000":
            position = int(hash[5], 16)
            if position < len(password) and password[position] == 0:
                print("Found: {}".format(hash[:7]))
                password[position] = hash[6]
                found += 1
        i += 1
        if i%1e6 == 0:
            print("{} mil iterations".format(i/1e6))

    return "".join(password)

if __name__ == "__main__":
    if len(sys.argv) < 2:
        part = "A"
    else:
        part = sys.argv[1]

    print(getResult(part, input))
