import numpy as np
import time
import copy as cp

fileName = "day17_data.txt"
testData = ["2413432311323","3215453535623","3255245654254","3446585845452","4546657867536","1438598798454",
            "4457876987766","3637877979653","4654967986887","4564679986453","1224686865563","2546548887735",
            "4322674655533"]
testSol1 = 102
testSol2 = 94

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

    print()
    matrix = [[ str(int(y)).rjust(1) for y in x] for x in matrix]
    matrix = np.array(matrix)
    matrix[matrix == "1"] = "."
    [print("".join(list(x))) for x in matrix]
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

def getNewNextStraight(currLoc, nextLoc):

    i,j = currLoc
    i2,j2 = nextLoc

    if i == i2:
        return [i2, j2 + j2-j]
    elif j == j2:
        return [i2 + i2-i, j2]

def isCostBetterThanHistory(history, memo, cost0, queue, maxDepth=1000):

    i,j, nStraight0, m, n = memo
    nextStraightLoc0 = [m,n]
    currLoc0 = [i,j]

    i = 0
    while i < len(queue):
        currLoc, nextStraightLoc, cost, nStraight, path = queue[i]

        if currLoc == currLoc0 and nStraight == nStraight0 and nextStraightLoc == nextStraightLoc0:
            if cost > cost0:
                del queue[i]
                continue
            # print("better stuff")
            return True

        # if i > maxDepth:
        #     break
        i=i+1

    return False

def addToSortedQueue(queue, newElement, newCost):

    # print("\nHello")
    # print(queue)
    # print(newElement)

    if len(queue) == 0:
        # print("\t\t add to empty queue")
        queue.append(newElement)
    else:
        added = False
        for k,el in enumerate(queue):
            _cost = el[2]
            # _cost = el[0]
            if _cost > newCost:
                # print("\t\tadd to:", k)
                queue.insert(k, newElement)
                added = True
                break

        if added == False:
            # print("\t\t append to end of the queue")
            queue.append(newElement)

    return

def getNextLocations0(ii, currLoc, nStraight, maxNstraight, nextStraightLoc, N,M):

    nextLocations = []
    # We are at start
    if ii == 0:
        print("on start")
        nextLocations.append([[0,1], True])
        nextLocations.append([[1,0], True])

    ## We are somewhere not on the start
    else:

        if nStraight < maxNstraight-1:
            # print("add new straight")
            i2, j2 = nextStraightLoc
            nextLocations.append([nextStraightLoc, False])

        ## turning - from rotation matrix
        i,j = currLoc
        si,sj = nextStraightLoc
        dirY, dirX = si-i, sj-j
        # print("\t", dirX, dirY)

        i2, j2 = i - dirX, j + dirY
        nextLocations.append([[i2,j2], True])

        i2, j2 = i + dirX, j - dirY
        nextLocations.append([[i2,j2], True])

    return nextLocations

def getNextLocations(ii, currLoc, nStraight, maxNstraight, nextStraightLoc, N,M):

    nextLocations = []
    # We are at start
    if ii == 0:
        print("on start")
        nextLocations.append([[0,1], True])
        nextLocations.append([[1,0], True])

    ## We are somewhere not on the start
    else:

        if nStraight < maxNstraight-1:
            # print("add new straight")
            i2, j2 = nextStraightLoc
            if i2 >= 0 and j2 >= 0 and i2 < N and j2 < M:
                nextLocations.append([nextStraightLoc, False])

        ## turning - from rotation matrix
        i,j = currLoc
        si,sj = nextStraightLoc
        dirY, dirX = si-i, sj-j
        # print("\t", dirX, dirY)

        i2, j2 = i - dirX, j + dirY
        if i2 >= 0 and j2 >= 0 and i2 < N and j2 < M:
            nextLocations.append([[i2,j2], True])

        i2, j2 = i + dirX, j - dirY
        if i2 >= 0 and j2 >= 0 and i2 < N and j2 < M:
            nextLocations.append([[i2,j2], True])

    return nextLocations

def steppingWithOrderedQueue(data, maxNstraight=3):

    N,M = len(data), len(data[0])
    start = [0,0]
    finish = [N-1, M-1]

    ## element: (current position, nextStraightPosition, cost, nStraight, path)
    firstPosition = (start, [0,0], 0, 0, [start])
    queue = [firstPosition]

    # history = set()
    history = {}
    dist = np.ones((N,M)) * -1
    # [print("".join(list(x))) for x in [[ str(int(y)).rjust(4) for y in x] for x in dist]]

    best = None

    Nmax = 9999999
    ii=0
    while len(queue) > 0 and ii < Nmax:

        currLoc, nextStraightLoc, cost, nStraight, path = queue[0]
        del queue[0]

        # print()
        # print(ii, [currLoc, nextStraightLoc, cost, nStraight])

        i,j = currLoc
        memo = (i,j, nStraight, nextStraightLoc[0], nextStraightLoc[1])
        if memo in history:
            if history[memo] <= cost:
                ii = ii+1
                continue
            else:
                print("Been here before:", currLoc, history[memo], cost)

        ## We have to memorize NOT ONLY position but also other characteristics of the path to here
        ## like direction and number of steps in straight line
        # history.add(memo)
        history[memo] = cost

        ## The number different from -1 is important, actual number is just for debuging
        if dist[i][j] == -1 or dist[i][j] > cost:
            dist[i][j] = cost

        if i == finish[0] and j == finish[1]:
            print("finished!", ii, len(queue))
            key = list(history.keys())[50]
            print("\t", key, history[key])
            # print(np.argwhere(dist < 0), len(np.argwhere(dist < 0)))

            # print("dist f:")
            # [print("".join(list(x))) for x in [[ str(int(y)).rjust(4) for y in x] for x in dist]]

            return cost, path

        nextLocations = []
        ## We are at start
        if ii == 0:
            print("on start")
            nextLocations.append([[0,1], True])
            nextLocations.append([[1,0], True])
            if ii != 0:
                print("sad")
                break

        ## We are somewhere not on the start
        else:
            if nStraight < maxNstraight-1:
                # print("add new straight")
                nextLocations.append([nextStraightLoc, False])

            ## turning - from rotation matrix
            i,j = currLoc
            i2,j2 = nextStraightLoc
            dirY, dirX = i2-i, j2-j
            # print("\t", dirX, dirY)
            nextLocations.append([[i - dirX,j + dirY], True])
            nextLocations.append([[i + dirX,j - dirY], True])


        # print("\t next locations:", nextLocations)

        for _loc in nextLocations:
            i,j = _loc[0]
            turned = _loc[1]

            if i >= 0 and j >= 0 and i < N and j < M and dist[i][j] == -1:

                newCost = cost + int(data[i][j])
                newNexStraight = getNewNextStraight(currLoc, _loc[0])
                newNstraight = 0
                if turned == False:
                    newNstraight = nStraight + 1
                newPath = path + [_loc[0]]

                newElement = (_loc[0], newNexStraight, newCost, newNstraight, newPath)
                # print("\t new element:", newElement)

                if len(queue) == 0:
                    # print("\t\t add to empty queue")
                    queue.append(newElement)
                else:
                    added = False
                    for k,el in enumerate(queue):
                        _cost = el[2]
                        if _cost > newCost:
                            # print("\t\tadd to:", k)
                            queue.insert(k, newElement)
                            added = True
                            break

                    if added == False:
                        # print("\t\t append to end of the queue")
                        queue.append(newElement)

        # print("\t new queue:", [(x[0], x[2]) for x in queue])

        # if ii > 5000:
        #     break

        ii=ii+1

    # print("dist 2:")
    # [print("".join(list(x))) for x in [[ str(int(y)).rjust(4) for y in x] for x in dist]]

    return best

def steppingWithOrderedQueue2(data, maxNstraight=3):

    N,M = len(data), len(data[0])
    start = [0,0]
    finish = [N-1, M-1]

    ## element: (current position, nextStraightPosition, cost, nStraight, path)
    firstPosition = (start, [0,0], 0, 0, [start])
    queue = [firstPosition]

    # history = set()
    history = {}
    dist = np.ones((N,M)) * -1
    # [print("".join(list(x))) for x in [[ str(int(y)).rjust(4) for y in x] for x in dist]]

    Nmax = 9999999
    k=0
    while len(queue) > 0 and k < Nmax:

        currLoc, nextStraightLoc, cost, nStraight, path = queue[0]
        del queue[0]

        # print()
        # print(ii, [currLoc, nextStraightLoc, cost, nStraight])

        i,j = currLoc
        memo = (i,j, nStraight, nextStraightLoc[0], nextStraightLoc[1])

        if memo in history and history[memo] <= cost:
            k = k+1
            continue

        ## We have to memorize NOT ONLY position but also other characteristics of the path to here
        ## like direction and number of steps in straight line
        # history.add(memo)
        history[memo] = cost

        ## The number different from -1 is important, actual number is just for debuging
        if dist[i][j] == -1 or dist[i][j] > cost:
            dist[i][j] = cost

        if i == finish[0] and j == finish[1]:
            print("finished!", k, len(queue))
            key = list(history.keys())[50]
            print("\t", key, history[key])
            # print(np.argwhere(dist < 0), len(np.argwhere(dist < 0)))

            # print("dist f:")
            # [print("".join(list(x))) for x in [[ str(int(y)).rjust(4) for y in x] for x in dist]]

            return cost, path

        nextLocations = getNextLocations(k, currLoc, nStraight, maxNstraight, nextStraightLoc, N,M)
        # print("\nhello", currLoc)
        # print("next locations:", nextLocations)
        # print("2", getCandidates(i, j, path, N, M))
        # if k > 1:
        #     raise Exception("trolling")

        for _loc in nextLocations:
            i2,j2 = _loc[0]
            turned = _loc[1]
            newCost = cost + int(data[i2][j2])
            newNexStraight = getNewNextStraight(currLoc, _loc[0])
            newNstraight = 0
            if turned == False:
                newNstraight = nStraight + 1

            nextMemo = (i2,j2, newNstraight, newNexStraight[0], newNexStraight[1])
            if dist[i2][j2] == -1 or not (nextMemo in history and history[nextMemo] <= newCost):

                newPath = path + [_loc[0]]
                newElement = (_loc[0], newNexStraight, newCost, newNstraight, newPath)
                # print("\t new element:", newElement)

                addToSortedQueue(queue, newElement, newCost)

        # print("\t new queue:", [(x[0], x[2]) for x in queue])

        # if ii > 5000:
        #     break

        k=k+1

    return None,None

def solve(data):

    solution = 0

    # [print("  "+"   ".join(list(x))) for x in data]
    # [print(""+"".join(list(x))) for x in data]

    M=len(data[0])
    N=len(data)

    allPossible = np.zeros((N,M))
    allPossible[0,0] = 1

    # solution, path = steppingWithOrderedQueue(data)
    solution, path = steppingWithOrderedQueue2(data)
    # print("path\n", path)

    # [print("  "+"   ".join(list(x))) for x in data]
    # [print(""+"".join(list(x))) for x in data]

    print("solution:", solution)

    testSolutionPath = [[0,0], [0,1], [0,2], [1,2], [1,3], [1,4], [1,5], [0,5], [0,6], [0,7],
                [0,8], [1,8], [2,8], [2,9], [2,10], [3,10], [4,10], [4,11], [5,11],
                [6,11], [7,11], [7,12], [8,12], [9,12], [10,12], [10,11], [11,11],
                [12,11], [12,12]]

    # printPath(data, testSolutionPath, N, M)
    # printPath(data, path, N, M)

    return

data = readFile()
solve(testData)

# t1 = time.time()
# solve(data)
# print("Elapsed:", round(time.time() - t1, 2), "s")

## 948 is to high

print()
print("########### PART 2 ###############")
print()

def getNextLocations2(ii, currLoc, nStraight, nextStraightLoc, minNstraight, maxNstraight, N,M):

    nextLocations = []
    # We are at start
    if ii == 0:
        print("on start")
        nextLocations.append([[0,1], True])
        nextLocations.append([[1,0], True])

    ## We are somewhere not on the start
    else:

        if nStraight < maxNstraight-1:
            # print("add new straight")
            i2, j2 = nextStraightLoc
            if i2 >= 0 and j2 >= 0 and i2 < N and j2 < M:
                nextLocations.append([nextStraightLoc, False])

        ## turning - from rotation matrix
        if nStraight >= minNstraight -1:
            i,j = currLoc
            si,sj = nextStraightLoc
            dirY, dirX = si-i, sj-j
            # print("\t", dirX, dirY)

            i2, j2 = i - dirX, j + dirY
            if i2 >= 0 and j2 >= 0 and i2 < N and j2 < M:
                nextLocations.append([[i2,j2], True])

            i2, j2 = i + dirX, j - dirY
            if i2 >= 0 and j2 >= 0 and i2 < N and j2 < M:
                nextLocations.append([[i2,j2], True])

    return nextLocations

def steppingWithOrderedQueue20(data, minNstraight=4, maxNstraight=10):

    N,M = len(data), len(data[0])
    start = [0,0]
    finish = [N-1, M-1]

    ## element: (current position, nextStraightPosition, cost, nStraight, path)
    firstPosition = (start, [0,0], 0, 0, [start])
    queue = [firstPosition]

    # history = set()
    history = {}
    dist = np.ones((N,M)) * -1
    # [print("".join(list(x))) for x in [[ str(int(y)).rjust(4) for y in x] for x in dist]]

    Nmax = 9999999
    k=0
    while len(queue) > 0 and k < Nmax:

        currLoc, nextStraightLoc, cost, nStraight, path = queue[0]
        del queue[0]

        # print()
        # print(ii, [currLoc, nextStraightLoc, cost, nStraight])

        i,j = currLoc
        memo = (i,j, nStraight, nextStraightLoc[0], nextStraightLoc[1])
        if memo in history:
            if history[memo] <= cost:
                # print("trol", k)
                # break
                k = k+1
                continue

            else:
                print("Been here before:", currLoc, history[memo], cost)

        ## We have to memorize NOT ONLY position but also other characteristics of the path to here
        ## like direction and number of steps in straight line
        # history.add(memo)
        history[memo] = cost

        ## The number different from -1 is important, actual number is just for debuging
        if dist[i][j] == -1 or dist[i][j] > cost:
            dist[i][j] = cost

        if i == finish[0] and j == finish[1]:
            print("finished!", k, len(queue))
            key = list(history.keys())[50]
            print("\t", key, history[key])
            # print(np.argwhere(dist < 0), len(np.argwhere(dist < 0)))

            # print("dist f:")
            # [print("".join(list(x))) for x in [[ str(int(y)).rjust(4) for y in x] for x in dist]]

            return cost, path

        nextLocations = getNextLocations2(k, currLoc, nStraight, nextStraightLoc, minNstraight, maxNstraight, N,M)
        # print("\nhello", currLoc)
        # print("next locations:", nextLocations)
        # print("2", getCandidates(i, j, path, N, M))
        # if k > 1:
        #     raise Exception("trolling")

        for _loc in nextLocations:
            i2,j2 = _loc[0]
            turned = _loc[1]
            newCost = cost + int(data[i2][j2])
            newNexStraight = getNewNextStraight(currLoc, _loc[0])
            newNstraight = 0
            if turned == False:
                newNstraight = nStraight + 1

            nextMemo = (i2,j2, newNstraight, newNexStraight[0], newNexStraight[1])

            # if i >= 0 and j >= 0 and i < N and j < M and dist[i][j] == -1:
            if dist[i2][j2] == -1 or not (nextMemo in history and history[nextMemo] <= newCost):

                newPath = path + [_loc[0]]
                newElement = (_loc[0], newNexStraight, newCost, newNstraight, newPath)
                # print("\t new element:", newElement)

                addToSortedQueue(queue, newElement, newCost)

        # print("\t new queue:", [(x[0], x[2]) for x in queue])

        # if ii > 5000:
        #     break

        k=k+1

    return None,None

def solve2(data):

    solution = 0

    # [print("  "+"   ".join(list(x))) for x in data]
    # [print(""+"".join(list(x))) for x in data]

    M=len(data[0])
    N=len(data)

    # allPossible = np.zeros((N,M))
    # allPossible[0,0] = 1

    solution, path = steppingWithOrderedQueue20(data)
    # print("path\n", path)

    # [print("  "+"   ".join(list(x))) for x in data]
    # [print(""+"".join(list(x))) for x in data]

    print("solution:", solution)

    testSolutionPath = [[0,0], [0,1], [0,2], [0,3], [0,4], [0,5], [0,6], [0,7],
                        [0,8], [1,8], [2,8], [3,8], [4,8], [4,9], [4,10], [4,11],
                        [4,12], [5,12], [6,12], [7,12], [8,12], [9,12], [10,12],
                        [11,12], [12,12]]

    # printPath(data, testSolutionPath, N, M)
    # printPath(data, path, N, M)

    return

data = readFile()
solve2(testData)

# t1 = time.time()
# solve2(data)
# print("Elapsed:", round(time.time() - t1, 2), "s")