import numpy as np
import time
import copy as cp

fileName = "day11_data.txt"
testData = ["...#......",".......#..","#.........","..........",
            "......#...",".#........",".........#","..........",
            ".......#..","#...#....."]
testSol = 374

def readFile():

    data = []

    with open(fileName, "r") as f:
        i=0
        for line in f:
            data.append(line[:-1])
            # line = line[:-1]
            # print(i, line)

            i=i+1

    # print(data[:3])
    # print(data[-3:])

    return data

def expandUniverse(data):

    horExpand = []
    verExpand = []
    M = len(data[0])
    N = len(data)
    i=0
    while i < N:

        nEmptySpaceHor = 0
        nEmptySpaceVer = 0

        j=0
        while j < M:
            if data[i][j] == ".":
                nEmptySpaceHor = nEmptySpaceHor + 1

            if data[j][i] == ".":
                nEmptySpaceVer = nEmptySpaceVer + 1
            j=j+1

        if nEmptySpaceHor == M:
            horExpand.append(i)

        if nEmptySpaceVer == N:
            verExpand.append(i)

        i=i+1

    line = ["." for x in range(M)]
    i = 0
    while i < len(horExpand):
        data.insert(horExpand[i]+i, line)
        i=i+1

    data2 = np.transpose(data)
    data2 = [list(x) for x in data2]

    column = ["." for x in range(N+len(horExpand))]
    i = 0
    while i < len(verExpand):
        data2.insert(verExpand[i]+i, column)
        # data2 = np.insert(data2, verExpand[i]+i, column)
        # print()
        # print(i, verExpand[i])
        # [print(x) for x in data2]
        i=i+1


    data = np.transpose(data2)
    data = [list(x) for x in data]

    return data

def findGalaxies(data):

    galaxies = []

    M = len(data[0])
    N = len(data)
    i=0
    while i < N:
        j=0
        while j < M:

            if data[i][j] != ".":
                galaxies.append([i,j])

            j=j+1
        i=i+1

    return galaxies

def solve(data):

    data = [list(x) for x in data]
    # [print(x) for x in data]

    data = expandUniverse(data)

    # print()
    # [print(x) for x in data]

    galaxies = findGalaxies(data)
    # print(galaxies)

    dist = []

    N = len(galaxies)
    i=0
    while i < N-1:
        # print()
        # print(galaxies[i])
        for j,g in enumerate(galaxies[i+1:]):
            d = abs(galaxies[i][0] - g[0]) + abs(galaxies[i][1] - g[1])
            # print(i,j+i+1, galaxies[i], g, d)
            dist.append(d)

        i=i+1

    solution = np.sum(dist)

    print("solution:", solution)

    return

data = readFile()
solve(testData)
solve(data)

print()
print("########### PART 2 ###############")
print()


def expandUniverse2(data):

    horExpand = []
    verExpand = []
    M = len(data[0])
    N = len(data)
    i=0
    while i < N:

        nEmptySpaceHor = 0
        nEmptySpaceVer = 0

        j=0
        while j < M:
            if data[i][j] == ".":
                nEmptySpaceHor = nEmptySpaceHor + 1

            if data[j][i] == ".":
                nEmptySpaceVer = nEmptySpaceVer + 1
            j=j+1

        if nEmptySpaceHor == M:
            horExpand.append(i)

        if nEmptySpaceVer == N:
            verExpand.append(i)

        i=i+1

    return horExpand, verExpand

def solve2(data, nExpand=2):

    data = [list(x) for x in data]
    # [print(x) for x in data]

    verExpand, horExpand = expandUniverse2(data)

    galaxies = findGalaxies(data)
    # print()
    # # print(galaxies)
    # # print(len(galaxies))
    # print(horExpand)
    # print(verExpand)

    dist = []
    solution = 0

    N = len(galaxies)
    i=0
    while i < N-1:
        # print()
        # print(galaxies[i])
        for j,g in enumerate(galaxies[i+1:]):

            eVer, eHor = 0, 0
            for x in verExpand:
                xs = np.sort([galaxies[i][0], g[0]])
                # print("\t", x, galaxies[i][1], g[1], xs)

                if xs[0] < x and xs[1] > x:
                    eVer = eVer + nExpand-1
                    # print("\tyay ver")

            for x in horExpand:
                xs = np.sort([galaxies[i][1], g[1]])
                if xs[0] < x and xs[1] > x:
                    eHor = eHor + nExpand-1
                    # print("\tyay hor")


            d = abs(galaxies[i][0] - g[0]) + abs(galaxies[i][1] - g[1])
            # print(i+1, j+i+1+1, galaxies[i], g, eHor, eVer, "|", d)

            dist.append(d+eVer+eHor)
            solution = solution + d + eVer + eHor

        # break

        i=i+1

    # print(len(dist), np.min(dist), np.max(dist))
    # solution = np.sum(dist)

    print("solution:", str(solution).rjust(20))

    return

data = readFile()
solve2(testData, 2)
solve2(testData, 10)
solve2(testData, 100)
# solve2(data, 2)
solve2(data, 1000000)