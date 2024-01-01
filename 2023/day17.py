import numpy as np
import time
import copy as cp

fileName = "day17_data.txt"
testData = ["2413432311323","3215453535623","3255245654254","3446585845452","4546657867536","1438598798454",
            "4457876987766","3637877979653","4654967986887","4564679986453","1224686865563","2546548887735",
            "4322674655533"]
testSol1 = 102

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

def printPath(data, path, N,M):

    matrix = np.zeros((N,M))

    s = 0
    for i, pos in enumerate(path):
        matrix[pos[0], pos[1]] = matrix[pos[0], pos[1]] + 1
        if i != 0:
            s = s + int(data[pos[0]][pos[1]])
    s = s + int(data[-1][-1])

    print()
    matrix = [[ str(int(y)).rjust(1) for y in x] for x in matrix]
    [print(" ".join(list(x))) for x in matrix]
    print("sum =", s)

    return

def getPossbileSteps(data, history, i,j, N,M):

    steps = []
    maxSameHeight = 4

    ind = maxSameHeight - 1
    # arr = history[i][j][-ind:]

    if i > 0:
        if history[i][j] == 0 or len(history[i][j]) == 0:
            steps.append([i-1,j])
        elif i-1 != history[i][j][-1][0] or j != history[i][j][-1][1]:
            if len(history[i][j]) < ind or not all(x[1] == j for x in history[i][j][:][-ind:]):
                steps.append([i-1,j])
                # print("yay 1", history[i][j], history[i][j][:][-ind:])
    if i < N-1:
        if history[i][j] == 0 or len(history[i][j]) == 0:
            steps.append([i+1,j])
        elif i+1 != history[i][j][-1][0] or j != history[i][j][-1][1]:
            if len(history[i][j]) < ind or not all(x[1] == j for x in history[i][j][:][-ind:]):
                steps.append([i+1,j])
                # print("yay 2", history[i][j], history[i][j][:][-ind:])
    if j > 0:
        if history[i][j] == 0 or len(history[i][j]) == 0:
            steps.append([i,j-1])
        elif i != history[i][j][-1][0] or j-1 != history[i][j][-1][1]:
            if len(history[i][j]) < ind or not all(x[0] == i for x in history[i][j][:][-ind:]):
                steps.append([i,j-1])
                # print("yay 3", history[i][j], history[i][j][:][-ind:])
    if j < N-1:
        if history[i][j] == 0 or len(history[i][j]) == 0:
            steps.append([i,j+1])
        elif i != history[i][j][-1][0] or j+1 != history[i][j][-1][1]:
            if len(history[i][j]) < ind or not all(x[0] == i for x in history[i][j][:][-ind:]):
                steps.append([i,j+1])
                # print("yay 4", history[i][j], history[i][j][:][-ind:])

    return steps

def isCurrentInAllPossible(allPossible, current):

    for pos in allPossible:
        if current[0] == pos[0] and current[1] == pos[1]:
            return True

    return False

def doOnePass(data, dist, history, allPossible, N,M):

    i=0
    while i < N:
        j=0
        while j < M:

            # print()
            # print(i,j)

            if allPossible[i,j] == 1:

                if i == 0 and j == 0:
                    dist[i,j] = 0
                    history[i][j] = []

                possibleSteps = getPossbileSteps(data, history, i, j, N,M)
                # print("\t history:", history[i][j])
                # print("\tpossible:",possibleSteps)

                for step in possibleSteps:
                    v0, v1 = step[0], step[1]
                    allPossible[v0,v1] = 1

                    # print("\t",step, i,j)
                    # print("\t",dist[i][j], data[v0][v1])

                    d = dist[i][j] + int(data[v0][v1])
                    if dist[v0][v1] == -1 or d < dist[v0][v1]:
                        # print("\thistory1 b:", v0,v1, history[v0][v1], dist[v0][v1])
                        dist[v0, v1] = d
                        history[v0][v1] = cp.deepcopy(history[i][j])
                        history[v0][v1].append([i,j])
                        # print("\thistory1 a:", v0,v1, history[v0][v1], dist[v0][v1])

                    # elif d < dist[v0][v1]:
                    #     print("\thistory2 b:", v0,v1, history[v0][v1], dist[v0][v1])
                    #     dist[v0, v1] = d
                    #     history[v0][v1] = cp.deepcopy(history[i][j])
                    #     history[v0][v1].append([i,j])
                    #     print("\thistory2 a:", v0,v1, history[v0][v1], dist[v0][v1])
            else:
                print("\tThis was not possible")
                ## TODO check other suboptimal route

            # break
            # if j>1:
            #     break

            j=j+1

        # if i>1:
        #     break
        i=i+1

    return data, dist, history, allPossible

def solve(data, maxSteps=100):

    solution = 0

    # [print("  "+"   ".join(list(x))) for x in data]
    # [print(""+"".join(list(x))) for x in data]

    M=len(data[0])
    N=len(data)

    dist = np.ones((N,M)) * -1
    history = [list(x) for x in np.zeros((N,M))]
    allPossible = np.zeros((N,M))
    allPossible[0,0] = 1

    maxIter = 200
    i=0
    while i < maxIter:
        # print(i)
        data, _dist, history, allPossible = doOnePass(data, cp.deepcopy(dist), history, allPossible, N, M)

        if np.sum(dist - _dist) == 0:
            print("nothing changed", i)
            break
        dist = cp.deepcopy(_dist)

        break
        # print(30*"#")

        i=i+1

    # print()
    # print(history[3][0])

    print()
    dist = [[ str(int(y)).rjust(4) for y in x] for x in dist]
    [print("".join(list(x))) for x in dist]
    print()
    # [print("  "+"   ".join(list(x))) for x in data]
    [print(""+" ".join(list(x))) for x in data]


    print()
    print("solution:", dist[-1][-1])

    testPath = [[0,0], [0,1], [0,2], [1,2], [1,3], [1,4], [1,5], [0,5], [0,6], [0,7],
                [0,8], [1,8], [2,8], [2,9], [2,10], [3,10], [4,10], [4,11], [5,11],
                [6,11], [7,11], [7,12], [8,12], [9,12], [10,12], [10,11], [11,11],
                [12,11]]

    printPath(data, history[-1][-1], N,M)
    printPath(data, testPath, N, M)

    return

data = readFile()
solve(testData)
# solve(data)

print()
print("########### PART 2 ###############")
print()
