import numpy as np
import time
import copy as cp

fileName = "day9_data.txt"

testData = ["0 3 6 9 12 15","1 3 6 10 15 21","10 13 16 21 30 45"]
testSolution = 114
testSolution2 = 2

def readFile():

    data = []

    with open(fileName, "r") as f:
        i=0
        for line in f:
            data.append(line[:-1])
            # line = line[:-1]
            # print(i, line)

            i=i+1

    return data

def getDiffTree(history):

    tree = [history]
    level = cp.deepcopy(history)

    N = 100
    i = 0
    while i < N:

        level = []
        M = len(tree[-1])
        j = 1
        while j < M:

            level.append(tree[-1][j]-tree[-1][j-1])

            j=j+1

        tree.append(level)

        if all(x == 0 for x in level):
            break

        i=i+1

    return tree

def solve(data):

    solution = 0

    for line in data:

        hist = [int(x) for x in line.split()]
        tree = getDiffTree(hist)

        # print()
        # [print(x) for x in tree]

        s = 0
        for level in tree[::-1]:
            s = s+level[-1]
            # print(s)

        solution = solution + s


    # print(sols)
    print("solution:", solution)

    return

data = readFile()
solve(testData)
solve(data)

print()
print("########### PART 2 ###############")
print()

def solve2(data):

    solution = 0

    for line in data:

        hist = [int(x) for x in line.split()]
        tree = getDiffTree(hist)

        # print()
        # [print(x) for x in tree]

        sa = []
        s = 0
        for level in tree[::-1]:
            s = level[0] - s
            sa.append(s)

        # print("+", sa)
        solution = solution + s


    # print(sols)
    print("solution:", solution)

    return

data = readFile()
solve2(testData)
solve2(data)