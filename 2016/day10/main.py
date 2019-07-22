#!/usr/bin/python3

import sys
import re

file = "input"

class BalanceBot:
    def __init__(self, index, lowCommand, highCommand):
        self.number = index
        self.values = []
        self.lowCommand = lowCommand
        self.highCommand = highCommand

    def __str__(self):
        return "[{}]: {} - low -> {} - high -> {}".format(self.number, self.values, self.lowCommand, self.highCommand)

    def addValue(self, value):
        self.values.append(value)
        self.values.sort()
        return len(self.values) >= 2

class BalanceBots:

    def __str__(self):
        output = "\nLooking for {} and {}\n".format(self.low, self.high)
        output += "Bots: \n"
        for index in self.bots:
            output += "    " + self.bots[index].__str__() + "\n"
        output += "Output: {}\n".format(self.output)
        return output

    def __init__(self, low, high):
        self.bots = {}
        self.output = {}
        self.result = -1
        self.low = low
        self.high = high

    def addBot(self, index, lowCommand, highCommand):
        self.bots[index] = BalanceBot(index, lowCommand, highCommand)

    def addValue(self, index, value):
        activate = self.bots[index].addValue(value)
        if activate:
            lowValue, highValue = self.bots[index].values
            lowTarget, lowNumber = self.bots[index].lowCommand
            highTarget, highNumber = self.bots[index].highCommand
            if lowValue == self.low and highValue == self.high:
                self.result = index
            if lowTarget == "bot":
                print("[{}] Giving value {} to bot {}".format(index, lowValue, lowNumber))
                self.addValue(lowNumber, lowValue)
            else:
                self.output[lowNumber] = lowValue
            if highTarget == "bot":
                print("[{}] Giving value {} to bot {}".format(index, highValue, highNumber))
                self.addValue(highNumber, highValue)
            else:
                self.output[highNumber] = highValue

def getResult(part, file, low, high):

    firstPart = part == "A"

    with open(file, mode="rt", encoding="utf-8") as f:
        input = f.read().splitlines()

    values = []
    commands = []
    print()
    for line in input:
        matchValue = re.search(r"value (?P<value>\d+) goes to bot (?P<bot>\d+)", line)
        if matchValue:
            values.append((int(matchValue.group("bot")), int(matchValue.group("value"))))

        matchCommand = re.search(
            r"bot (?P<bot>\d+) gives low to (?P<lowTarget>\D+) (?P<lowNumber>\d+) and high to (?P<highTarget>\D+) (?P<highNumber>\d+)", line)
        if matchCommand:
            commands.append(
                (
                    int(matchCommand.group("bot")),
                    matchCommand.group("lowTarget"),
                    int(matchCommand.group("lowNumber")),
                    matchCommand.group("highTarget"),
                    int(matchCommand.group("highNumber"))
                )
            )

    if firstPart:
        result = calculateResultA(values, commands, low, high)
    else:
        result = calculateResultB(values, commands, low, high)

    return result


def calculateResultA(values, commands, low, high):

    bots = BalanceBots(low, high)

    for command in commands:
        bots.addBot(command[0], (command[1], command[2]), (command[3], command[4]))

    print(bots)

    for value in values:
        print("Bot {} starts with value {}".format(value[0], value[1]))
        bots.addValue(value[0], value[1])

    print(bots)

    return bots.result


def calculateResultB(values, commands, low, high):

    bots = BalanceBots(low, high)

    for command in commands:
        bots.addBot(command[0], (command[1], command[2]), (command[3], command[4]))

    print(bots)

    for value in values:
        print("Bot {} starts with value {}".format(value[0], value[1]))
        bots.addValue(value[0], value[1])

    print(bots)

    return bots.output[0] * bots.output[1] * bots.output[2]


if __name__ == "__main__":
    if len(sys.argv) < 2:
        part = "A"
    else:
        part = sys.argv[1]

    print(getResult(part, file, 17, 61))
