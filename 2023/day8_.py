import numpy as np
import copy as cp
import time

fileName = "day8__data.txt"

testData0 = ["RL","AAA = (BBB, CCC)","BBB = (DDD, EEE)","CCC = (ZZZ, GGG)","DDD = (DDD, DDD)","EEE = (EEE, EEE)",
            "GGG = (GGG, GGG)","ZZZ = (ZZZ, ZZZ)]"]
testData1 = ["LLR","AAA = (BBB, BBB)","BBB = (AAA, ZZZ)","ZZZ = (ZZZ, ZZZ)"]
testSolution0 = 2
testSolution1 = 6

testData2 = ["LR","11A = (11B, XXX)","11B = (XXX, 11Z)","11Z = (11B, XXX)","22A = (22B, XXX)","22B = (22C, 22C)",
             "22C = (22Z, 22Z)","22Z = (22B, 22B)","XXX = (XXX, XXX)"]
testSolution2 = 6

def readFile():

    data = []

    with open(fileName, "r") as f:

        for line in f:
            # print(line[:-1])
            data.append(line[:-1])

    # print(data[:5])
    # print(data[-5:])

    return data

def saveToDictionary(data):

    dataOut = {}

    for i,line in enumerate(data):

        if i == 0:

            dirs = []
            for d in line:
                if d == "L":
                    dirs.append(0)
                elif d == "R":
                    dirs.append(1)
                else:
                    print("ERROR in left or right")
            dataOut["instructions"] = dirs


        elif line != "":
            _line = line.split(" = ")
            dataOut[_line[0]] = _line[1][1:-1].split(", ")

    # print(dataOut)
    return dataOut


def solve(data):

    dataDict = saveToDictionary(data)
    directions = dataDict["instructions"]

    solution = 0
    maxIter = 1000000
    loc = "AAA"

    N = len(directions)
    i = 0
    while loc != "ZZZ" and solution < maxIter:

        loc = dataDict[loc][directions[i]]

        i=i+1
        solution = solution + 1

        if i == N:
            i = 0

    # print()
    # print(np.array(orderedData))

    print()
    print("Solution:", solution)

    return


data = readFile()
solve(testData0)
solve(testData1)
solve(data[:])

print()
print("//////////////////////// Part II ////////////////")
print()

def checkPeriodInInstructions(data):


    for line in enumerate(data):
        directions = line[1]
        break

    print(directions)

    # directions = dataDict["instructions"]

    i = 1
    while i < len(directions) / 2 +2:

        subD = directions[:i]
        # print("sad" ,subD)

        if subD in directions[i:]:
            print(i, subD)

        i=i+1

    return

def isThisNodeCycle(directions, graph, curLoc, currDirInd):

    i = len(graph) - 1
    while i >= 0:

        if curLoc == graph[i][0] and currDirInd == graph[i][1]:
            return True, i

        i=i-1

    return False, -1

def findMinimumEnd(graphs, zInds, cycleStarts):

    startingCycles = 1258304364256//3
    # startingCycles = 10000000

    cycleSizes = [len(graphs[i]) - cycleStarts[i] for i,x in enumerate(graphs)]
    steps = [zInds[i][-1] + startingCycles*cycleSizes[i] for i,x in enumerate(zInds)]

    N = len(graphs)
    i=0
    while i < N:
        print("\n'"+str(i)+"'")
        # print(graphs[i])
        print(zInds[i])
        print(cycleStarts[i], cycleSizes[i])
        i=i+1

    ## minimize: c1 + z1*cs2*k1 = c2 + z2*cs2*k2
    # print()
    # print("start:", steps)
    maxSteps = 10000000000
    # M = len(cycleStarts)
    i = 0
    while i < maxSteps:

        if i % 5000000 == 0:
            print("progress", i,"/",maxSteps, "("+str(round(100.*i/maxSteps, 2))+" %)")
            print("\t", steps, len(str(np.min(steps))), np.max([x - steps[0] for x in steps]))

        # print()
        j = np.argmin(steps)
        # print(j,steps[j])
        steps[j] = steps[j] + cycleSizes[j]

        # print(steps)

        if all(x == steps[0] for x in steps):
            return steps[0]

        i=i+1

    # Last known steps
    # progress 999000000 / 1000000000 (99.9 %)
    # 	 [2516608743959, 2516608735558, 2516608739352, 2516608736913, 2516608736100, 2516608728512]

    print("last", steps)

    return -1

def findMinimumEnd2(graphs, zInds, cycleStarts):

    cycleSizes = [len(graphs[i]) - cycleStarts[i] for i,x in enumerate(graphs)]
    steps = [zInds[i][-1] for i,x in enumerate(zInds)]

    matrix = []
    vecN = []

    N = len(graphs)
    i=0
    while i < N:
        print("\n'"+str(i)+"'")
        print(graphs[i])
        print(zInds[i])
        print(cycleStarts[i], cycleSizes[i])

        matrix.append([1, -cycleSizes[i]])
        vecN.append(zInds[i][-1])

        i=i+1


    matrix = np.array(matrix)
    vecN = np.array(vecN)

    print()
    print(matrix)
    print(vecN)
    print(np.linalg.solve(matrix, vecN))

    return -1

def solve2(data):

    dataDict = saveToDictionary(data)
    directions = dataDict["instructions"]

    solution = 0
    maxSteps = 10000000000
    # maxSteps = 10

    locations = [x for x in dataDict.keys() if x != "instructions"]
    locs = np.sort([x for x in locations if x[2] == "A"])
    # end = np.sort([x for x in locations if x[2] == "Z"])
    # locs = ["AAA"]
    # end = ["ZZZ"]

    graphs = [[(x,directions[0])] for x in locs]
    zInds = [[] for x in locs]
    foundCycle = [-1 for x in locs]

    # print(len(locs), len(end))
    print("start", locs)

    N = len(directions)
    i = 0
    n = 0
    while n < maxSteps:

        # if n % 10000 == 0:
        #     print("progress", n,"/",maxSteps, "("+str(round(100.*n/maxSteps, 2))+" %)")
        #     print("\t", locs, "|", end)
        #     print("\t",foundCycle)

        #     j=0
        #     while j < M:
        #         print("\t",i, zInds[i])
        #         # break
        #         j=j+1

        _locs = []
        for j,loc in enumerate(locs):

            if foundCycle[j] == -1:
                # print()
                # print(j)

                ## Step forward
                newLoc = dataDict[loc][directions[i]]

                ## Making graph
                isCycle = isThisNodeCycle(directions, graphs[j], newLoc, (i+1)%N)
                if isCycle[0] == True:
                    # print("Found cycle", j, newLoc, (i+1)%N, n+1)
                    foundCycle[j] = isCycle[1]
                    _locs.append(None)
                    continue
                # else:
                #     print("NOT cycle", j, newLoc, i, n+1)

                _locs.append(newLoc)
                graphs[j].append((newLoc, (i+1)%N))
                if newLoc[2] == "Z":
                    # print("found Z:", [newLoc, j, n+1, (i+1)%N])
                    zInds[j].append(n+1)

            else:
                # print("alredy found cycle for", j)
                _locs.append(None)

        # _locs = np.sort(_locs)
        locs = cp.deepcopy(_locs)
        # print("locs",locs)

        i=i+1
        n = n +1

        # if all(loc[2] == "Z" for loc in locs):
        #     break

        if -1 not in foundCycle:
            break

        if i == N:
            i = 0

    print("found all cycles")

    print("Z location:", [x[-1] for x in zInds])
    print("cycle size:", [len(graphs[i]) - foundCycle[i] for i,x in enumerate(graphs)])

    # solution = findMinimumEnd(graphs, zInds, foundCycle)
    # solution = findMinimumEnd2(graphs, zInds, foundCycle)

    import math
    out = math.lcm(*[x[-1] for x in zInds])

    # print()
    # [print(x) for x in graphs]

    print()
    print("Solution:", out)

    return

solve2(testData2)

# [('11A', 0), ('11B', 1), ('11Z', 0)]###, ('11B', 1), ('11Z', 0), ('11B', 1), ('11Z', 0)]
# [2]
# 1

# [('22A', 0), ('22B', 1), ('22C', 0), ('22Z', 1), ('22B', 0), ('22C', 1), ('22Z', 0)]###, ('22B', 1)]
# [3, 6]
# 1

# checkPeriodInInstructions(data)

t1 = time.time()
solve2(data)
print("elapsed", round((time.time()-t1)/60, 2), "min")
