import numpy as np
import time
import copy as cp
import matplotlib.pyplot as plt

fileName = "day21_data.txt"
testData1 = ["...........",".....###.#.",".###.##..#.","..#.#...#..","....#.#....",".##..S####.",
             ".##..#...#.",".......##..",".##.#.####.",".##..##.##.","..........."]
testSol1 = 16

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
    # print()
    # print(data[-3:])

    return data

def getStartPosition(data):

    pos = None

    M = len(data[0])
    N = len(data)
    i=0
    while i < N:
        j=0
        while j < M:

            if data[i][j] == "S":
                return [i,j]

            j=j+1
        i=i+1

    return pos

def getPossibleSteps(data, position, N,M):

    steps = []
    i,j = position

    # print(position, N,M)

    if i > 0 and data[i-1][j] != "#":
        steps.append([i-1, j])
    if i < N and data[i+1][j] != "#":
        steps.append([i+1, j])

    if j > 0 and data[i][j-1] != "#":
        steps.append([i, j-1])
    if j < M and data[i][j+1] != "#":
        steps.append([i, j+1])

    return steps

def removeDuplicateSteps(steps):

    # print("b", steps)
    _steps = []

    for step in steps:
        if len(_steps) == 0:
            _steps.append(step)
        else:

            add = True
            for s in _steps:
                if step[0] == s[0] and step[1] == s[1]:
                    add = False
                    break
            if add == True:
                _steps.append(step)

    # print("a", _steps)
    return _steps

def solve(data, maxSteps=6):

    solution = 0

    # print(data)
    [print("".join(list(x))) for x in [[ y for y in x] for x in data]]

    start = getStartPosition(data)
    # print("start", start)

    M = len(data[0])
    N = len(data)
    positions = [start]

    i=0
    while i < maxSteps:
        _pos = []
        for pos in positions:
            steps = getPossibleSteps(data, pos, N, M)
            # print("possible", pos, steps)
            _pos = _pos + steps

        _pos = removeDuplicateSteps(_pos)
        # print(_pos)

        positions = cp.deepcopy(_pos)

        # if i > 1:
        #     break
        i=i+1


    solution = len(positions)
    print("solution:", solution)

    return

data = readFile()
solve(testData1)
# solve(data, maxSteps=64)

print()
print("########### PART 2 ###############")
print()

def getPossibleSteps2(data, position, N,M):

    steps = []
    i,j = position

    i2 = i%N
    j2 = j%M

    # print("yolo", position, N,M, "|", i2,j2)

    if data[i2-1][j2] != "#":
        if i2-1 == -1:
            steps.append([i-1, j])
        else:
            steps.append([i-1, j])

    if i2 < N-1 and data[i2+1][j2] != "#":
        steps.append([i+1, j])
    elif i2 == N-1 and data[0][j2] != "#":
        steps.append([i+1, j])

    if data[i2][j2-1] != "#":
        if j2-1 == -1:
            steps.append([i, j-1])
        else:
            steps.append([i, j-1])

    if j2 < M-1 and data[i2][j2+1] != "#":
        steps.append([i, j+1])
    elif j2 == M-1 and data[i2][0] != "#":
        steps.append([i, 0])

    return steps

def isStepsInHistory(history, steps):

    for i,_steps in enumerate(history):

        if len(steps) == len(_steps) and all(step in _steps for step in steps):
                return True, i

    return False, None

def solve2(data, maxSteps=26501365):

    solution = 0

    # print(data)
    data = [list(x) for x in data]
    [print(" ".join(list(x))) for x in [[ y for y in x] for x in data]]

    start = getStartPosition(data)

    M = len(data[0])
    N = len(data)
    positions = [start]
    history = [[start]]
    sols = [1]

    # print(getPossibleSteps2(data, [N,5], N,M))
    # return

    i=0
    while i < maxSteps:

        # print()
        print(i)
        _pos = []
        for pos in positions:
            steps = getPossibleSteps2(data, pos, N, M)
            # print("possible", pos, steps)
            _pos = _pos + steps

        _pos = removeDuplicateSteps(_pos)

        positions = cp.deepcopy(_pos)

        history.append(cp.deepcopy(_pos))
        sols.append(len(positions))
        # print("history", isHistory, ind)

        if i > 50:
            break

        i=i+1

    print(sols)
    xs = [6, 10, 50]
    [print(x, sols[x]) for x in xs]

    x = np.arange(len(sols))
    plt.figure()
    plt.plot(x, sols, "x-")
    plt.show()

    print("solution:", solution)

    return


data = readFile()
solve2(testData1)