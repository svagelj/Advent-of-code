import numpy as np
import time
import copy as cp

fileName = "day14_data.txt"
testData = ["O....#....","O.OO#....#",".....##...","OO.#O....O",".O.....O#.",
            "O.#..O.#.#","..O..#O..O",".......O..","#....###..","#OO..#...."]

testSol1 = 136
testSol2 = 64

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

def moveStones(data):

    M = len(data[0])
    N = len(data)
    i=1
    while i < N:
        # print()
        # print(data[i])

        j=0
        while j < M:

            if data[i][j] == "O":
                k = 1
                n = 0
                while k < N:
                    # print(k)
                    if data[i-k][j] == ".":
                        n=n+1
                    if i-k == 0:
                        # print("top")
                        break
                    if data[i-k][j] != ".":
                        # print("bellow top")
                        break
                    k=k+1

                # print("move", i,j, "to", i-n,j, k,n)
                data[i][j] = "."
                data[i-n][j] = "O"

            j=j+1

        # return data
        # if i > 2:
        #     return data
        i=i+1

    return data

def solve(data):

    solution = 0

    data = [list(x) for x in data]
    # [print(" ".join(x)) for x in data]

    data = moveStones(data)

    M = len(data[0])
    N = len(data)
    i=0
    while i < N:
        # print()
        # print(data[i])

        j=0
        while j < M:

            if data[i][j] == "O":
                solution = solution + N-i

            j=j+1
        i=i+1

    # print()
    # [print(" ".join(x)) for x in data]

    print()
    print("solution:", solution)

    return

data = readFile()
solve(testData)
solve(data)

print()
print("########### PART 2 ###############")
print()

def slideFor(data, i,j, N,M):

    k = 1
    n = 0
    while k < N:
        # print(k)
        if data[i-k][j] == ".":
            n=n+1
        if i-k == 0:
            # print("top")
            break
        if data[i-k][j] != ".":
            # print("bellow top")
            break
        k=k+1

    return n

def initializeData(data):

    movingPosition = []
    slideFor = []

    M = len(data[0])
    N = len(data)
    i=0
    while i < N:

        slideFor.append([])

        j=0
        while j < M:

            if data[i][j] == "O":
                movingPosition.append([i,j])

            if data[i][j] != "#":
                ## to north

                ## to south


                slideFor.append([0,0,0,0])

            return 1,1

            j=j+1
        i=i+1

    return movingPosition, slideFor

def moveStones2(data):

    """
    O O O O . # . O . .
    O O . . # . . . . #
    O O . . O # # . . O
    O . . # . O O . . .
    . . . . . . . . # .
    . . # . . . . # . #
    . . O . . # . O . O
    . . O . . . . . . .
    # . . . . # # # . .
    # . . . . # . . . .

    -------------------

    . . . . . # . . . .
    . . . . # . . . O #
    . . . O O # # . . .
    . O O # . . . . . .
    . . . . . O O O # .
    . O # . . . O # . #
    . . . . O # . . . .
    . . . . . . O O O O
    # . . . O # # # . .
    # . . O O # . . . .
    """

    M = len(data[0])
    N = len(data)
    i=1
    while i < N:
        # print()
        # print(data[i])

        j=0
        while j < M:

            if data[i][j] == "O":
                k = 1
                n = 0
                while k < N:
                    # print(k)
                    if data[i-k][j] == ".":
                        n=n+1
                    if i-k == 0:
                        # print("top")
                        break
                    if data[i-k][j] != ".":
                        # print("bellow top")
                        break
                    k=k+1

                # print("move", i,j, "to", i-n,j, k,n)
                data[i][j] = "."
                data[i-n][j] = "O"

            j=j+1

        # return data
        # if i > 2:
        #     return data
        i=i+1

    return data

def checkIfSameGrid(data1, data2):

    M = len(data1[0])
    N = len(data1)
    i=0
    while i < N:
        j=0
        while j < M:
            # print(i,j, N,M)
            if data1[i][j] != data2[i][j]:
                return False
            j=j+1
        i=i+1

    return True

def getCyclingPeriod(data, nMax):

    states = [cp.deepcopy(data)]

    k=1
    while k < nMax:

        data = moveStones2(data)

        data = [list(x) for x in np.rot90(np.array(data), -1)]
        data = moveStones2(data)

        data = [list(x) for x in np.rot90(np.array(data), -1)]
        data = moveStones2(data)

        data = [list(x) for x in np.rot90(np.array(data), -1)]
        data = moveStones2(data)

        data = [list(x) for x in np.rot90(np.array(data), -1)]

        for i,state in enumerate(states):
            same = checkIfSameGrid(data, state)
            # print(k, i, same)
            if same == True:
                return i, k, states

        states.append(cp.deepcopy(data))

        k=k+1

    # for i, state in enumerate(states):
    #     print()
    #     print(i)
    #     [print(" ".join(x)) for x in state]

    return None, None, None

def solve2(data, nCycle=1):

    solution = 0

    data = [list(x) for x in data]
    # [print(" ".join(x)) for x in data]

    # movingPositions, slideFor = initializeData(data)

    start, end, states = getCyclingPeriod(data, max(nCycle, 100))
    print(start, end, "|", nCycle % end)

    if nCycle <= start:
        ind = nCycle
    else:
        ind = start + (nCycle-start) % (end - start)
    print("wanted ind =", ind)

    data = states[ind]

    M = len(data[0])
    N = len(data)
    i=0
    while i < N:
        # print()
        # print(data[i])

        j=0
        while j < M:

            if data[i][j] == "O":
                solution = solution + N-i

            j=j+1
        i=i+1

    # print()
    # [print(" ".join(x)) for x in data]

    # print()
    print("solution:", solution)

    return

data = readFile()

t1 = time.time()
# solve2(testData, nCycle=11)
solve2(testData, 1000000000)
# for x in range(25):
#     solve2(testData, nCycle=x)
solve2(data, 1000000000)
print("elapsed:", round((time.time()-t1)/60., 2), "min")