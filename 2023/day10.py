import numpy as np
import time
import copy as cp

fileName = "day10_data.txt"

corners = {'F':[1,1], '7':[1,-1], 'J':[-1, -1], 'L':[-1,1]}

testData11 = ["-L|F7","7S-7|","L|7||","-L-J|","L|-JF"]
testData12 = ["7-F7-",".FJ|7","SJLL7","|F--J","LJ.LJ"]
testSolution11 = 4
testSolution12 = 8

testData21 = ["...........",".S-------7.",".|F-----7|.",".||OOOOO||.",".||OOOOO||.",
              ".|L-7OF-J|.",".|II|O|II|.",".L--JOL--J.",".....O....."]
testData22 = ["..........",".S------7.",".|F----7|.",".||OOOO||.",".||OOOO||.",
              ".|L-7F-J|.",".|II||II|.",".L--JL--J.",".........."]
testData23 = ["OF----7F7F7F7F-7OOOO","O|F--7||||||||FJOOOO","O||OFJ||||||||L7OOOO",
              "FJL7L7LJLJ||LJIL-7OO","L--JOL7IIILJS7F-7L7O","OOOOF-JIIF7FJ|L7L7L7",
              "OOOOL7IF7||L7|IL7L7|","OOOOO|FJLJ|FJ|F7|OLJ","OOOOFJL-7O||O||||OOO",
              "OOOOL---JOLJOLJLJOOO"]
testData24 = ["FF7FSF7F7F7F7F7F---7","L|LJ||||||||||||F--J","FL-7LJLJ||||||LJL-77",
              "F--JF--7||LJLJIF7FJ-","L---JF-JLJIIIIFJLJJ7","|F|F-JF---7IIIL7L|7|",
              "|FFJF7L7F-JF7IIL---7","7-L-JL7||F7|L7F-7F7|","L.L7LFJ|||||FJL7||LJ",
              "L7JLJL-JLJLJL--JLJ.L"]
testSol21, testSol22, testSol23, testSol24 = 4, 4, 8, 10

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

def getStartingData(data):

    i=0
    while i < len(data):
        j = 0
        while j < len(data[i]):
            if data[i][j] == "S":
                directions = []

                if data[i-1][j] in ['|', 'F', '7']:
                    directions.append([-1,0])
                if data[i+1][j] in ['|', 'J', 'L']:
                    directions.append([1,0])
                if data[i][j+1] in ['-', 'J', '7']:
                    directions.append([0,1])
                if data[i][j-1] in ['-', 'F', 'L']:
                    directions.append([0,-1])

                return np.array([i,j]), np.array(directions)
            j=j+1
        i=i+1

    return "There os no 'S' in data!!"

def solve(data):

    data = [list(x) for x in data]
    data = np.array(data)
    # print()
    # print()
    # print(data)
    # print()

    ## Find start
    start = getStartingData(data)
    # print(start)

    _s = [start[0] for x in start[1]]
    steps = [_s]
    _s = []
    for d in start[1]:
        newPos = start[0] + d
        _s.append(newPos)
    steps.append(_s)
    # print("First step:", steps)

    cornerKeys = list(corners.keys())
    maxSteps = 1000000
    t = 1
    while t < maxSteps:

        _steps = []
        # print()
        for j,d in enumerate(steps[-1]):
            currChar = data[d[0], d[1]]
            # print(j, d, currChar)

            previousInd = -2
            if len(steps) < 2:
                previousInd = -1

            if currChar in cornerKeys:
                # print("corner")
                if d[0] != steps[previousInd][j][0]:
                    # print("y change")
                    # print("\t", corners[currChar])
                    newLoc = d + [0,corners[currChar][1]]
                else:
                    # print("x change")
                    # print("\t", corners[currChar])
                    newLoc = d + [corners[currChar][0],0]
                # newLoc = []
                # continue
            else:
                # print("strait pipe")
                if currChar == '|':
                    # print("vertical")
                    if d[0] < steps[previousInd][j][0]:
                        newLoc = d + [-1,0]
                    else:
                        newLoc = d + [+1,0]

                elif currChar == "-":
                    # print("horizontal")
                    if d[1] > steps[previousInd][j][1]:
                        # print("\tfirst")
                        newLoc = d + [0,1]
                    else:
                    #     print("\tsecond")
                        newLoc = d + [0,-1]

                # print(newLoc)

            _steps.append(newLoc)

        steps.append(_steps)
        # print("steps:", steps)

        # break
        # if t > 2:
        #     break

        t=t+1

        lastStep = steps[-1]
        if (np.diff(np.vstack(lastStep).reshape(len(lastStep),-1),axis=0)==0).all():
            # print("break for equal", t)
            break

    # print()
    # for x in steps:
    #     print(x[0])
    # print()
    # for x in steps:
    #     print(x[1])

    # print(sols)
    print("solution:", t)

    return

data = readFile()
solve(testData11)
solve(testData12)
solve(data)

print()
print("########### PART 2 ###############")
print()

def getPath(data):

    data = [list(x) for x in data]
    data = np.array(data)
    # print()
    # print()
    # print(data)
    # print()

    ## Find start
    start = getStartingData(data)
    # print(start)

    _s = [start[0] for x in start[1]]
    steps = [_s]
    _s = []
    for d in start[1]:
        newPos = start[0] + d
        _s.append(newPos)
    steps.append(_s)
    # print("First step:", steps)

    cornerKeys = list(corners.keys())
    maxSteps = 1000000
    t = 1
    while t < maxSteps:

        _steps = []
        # print()
        for j,d in enumerate(steps[-1]):
            currChar = data[d[0], d[1]]
            # print(j, d, currChar)

            previousInd = -2
            if len(steps) < 2:
                previousInd = -1

            if currChar in cornerKeys:
                # print("corner")
                if d[0] != steps[previousInd][j][0]:
                    # print("y change")
                    # print("\t", corners[currChar])
                    newLoc = d + [0,corners[currChar][1]]
                else:
                    # print("x change")
                    # print("\t", corners[currChar])
                    newLoc = d + [corners[currChar][0],0]
                # newLoc = []
                # continue
            else:
                # print("strait pipe")
                if currChar == '|':
                    # print("vertical")
                    if d[0] < steps[previousInd][j][0]:
                        newLoc = d + [-1,0]
                    else:
                        newLoc = d + [+1,0]

                elif currChar == "-":
                    # print("horizontal")
                    if d[1] > steps[previousInd][j][1]:
                        # print("\tfirst")
                        newLoc = d + [0,1]
                    else:
                    #     print("\tsecond")
                        newLoc = d + [0,-1]

                # print(newLoc)

            _steps.append(newLoc)

        steps.append(_steps)
        # print("steps:", steps)

        # break
        # if t > 2:
        #     break

        t=t+1

        lastStep = steps[-1]
        if t > 2 and (np.diff(np.vstack(lastStep).reshape(len(lastStep),-1),axis=0)==0).all():
            # print("break for equal", t)
            break

    # print()
    # for x in steps:
    #     print(x[0])
    # print()
    # for x in steps:
    #     print(x[1])

    return np.array(steps)

def getNumberOfCrosses(data, point, loop, M):

    nCross = 0
    ray = np.array([[point[0], k] for k in range(point[1]+1,M)])
    # cross = ["L7", "FJ"]

    # if point[0] == 4 and point[1] == 4:
    #     print()
    #     print(point, ray)

    M = len(data[0])
    N = len(ray)
    i = 0
    while i < N:
        r = ray[i]
        # if point[0] == 4 and point[1] == 4:
        #     print("r", r, data[r[0], r[1]], i, N)
        # if np.any(np.all(r == loop, axis=1)):
        if data[r[0], r[1]] in ["|"]:
            # print("ray |", r)
            nCross = nCross + 1

        elif data[r[0], r[1]] in ["L", "F"]:
            # if point[0] == 4 and point[1] == 4:
            #     print("\t", data[r[0], r[1]])
            j=1
            while r[1]+j < M:
                if data[r[0], r[1]+j] != "-":
                    # print("\t", data[r[0], r[1]+j])
                    if data[r[0], r[1]] == "L" and data[r[0], r[1]+j] == "7":
                        # if point[0] == 4 and point[1] == 4:
                        #     print("\tray L7", r)
                        nCross = nCross + 1
                        # i = i + j
                        # break
                    elif data[r[0], r[1]] == "F" and data[r[0], r[1]+j] == "J":
                        # if point[0] == 4 and point[1] == 4:
                        #     print("\tray FJ", r)
                        nCross = nCross + 1
                    i = i + j
                    break
                # else:
                #     print("\t", data[r[0], r[1]+j])
                j=j+1

        i=i+1

    # if point[0] == 3 and point[1] == 14:
    #     print("yay",point, ray, nCross)

    return nCross

def isColinearWithEdge(point, loop):
    # print("is colinear")

    N = len(loop)
    i=0
    while i < N:

        # print(i)
        if i != 0 and loop[i][0] == loop[i-1][0] and point[0] == loop[i][0]:
            # print("'1'")
            return True
        elif i < N-1 and loop[i][0] == loop[i+1][0] and point[0] == loop[i][0]:
            # print("'2'")
            return True

        i=i+1

    return False

def removeColinearLoopPoints(point, loop):

    newLoop = []

    N = len(loop)
    i=0
    while i < N:

        if loop[i][0] == loop[i-1][0] and point[0] == loop[i][0]:
            pass
        elif i < N-1 and loop[i][0] == loop[i+1][0] and point[0] == loop[i][0]:
            pass
        else:
            newLoop.append(loop[i])

        i=i+1

    return newLoop

def cleanData(data, loop):

    N = len(data)
    M = len(data[0])
    i=0
    while i < N:
        j=0
        while j < M:

            if not np.any(np.all([i,j] == loop, axis=1)):
                data[i][j] = "."

            j=j+1
        i=i+1

    return data

def removeStart(data, loop):

    start, steps = getStartingData(data)

    # print(start, steps)
    # print(steps[0]+steps[1])

    p = steps[0]+steps[1]

    if p[0] == 1 and p[1] == 1:
        # print("F")
        y = "F"
    elif p[0] == 1 and p[1] == -1:
        # print("7")
        y = "7"
    elif p[0] == -1 and p[1] == 1:
        # print("L")
        y = "L"
    elif p[0] == -1 and p[1] == -1:
        # print("J")
        y = "J"
    elif steps[0][0] != 0 and steps[1][0] != 0:
        # print("|")
        y = "|"
    elif steps[0][1] != 0 and steps[1][1] != 0:
        # print("-")
        y = "-"
    else:
        # print("so sad")
        y = "S"

    data[start[0]][start[1]] = y

    # print(data[start[0]-2:start[0]+2, start[1]-2: start[1]+2])

    return data

def solve2(data):

    data = [list(x) for x in data]
    data = np.array(data)
    # print(np.shape(data))
    # [print("  ".join(x)) for x in data]

    path  = getPath(data)
    # print(path)

    loop = []
    for step in path:
        loop.append(step[0])
    for step in path[::-1][1:-1]:
        loop.append(step[1])

    loop = np.array(loop)
    # print(loop)
    # print()

    data = cleanData(data, loop)
    # [print("  ".join(x)) for x in data]

    data = removeStart(data, loop)
    # print()
    # [print("  ".join(x)) for x in data]

    n = 0
    N = len(data)
    M = len(data[0])
    i = 0
    while i < N:
        j=0
        while j < M:

            # print([i,j])
            if not np.any(np.all([i,j] == loop, axis=1)):

                nCross = getNumberOfCrosses(data, [i,j], loop, M)
                # print()
                # print([i,j], nCross)

                if nCross != 0 and nCross % 2 == 1:
                    # print("found", [i,j], data[i,j], nCross)
                    n=n+1

                    # [print("  ".join(x)) for x in data]
                    # return

                    # print()
                    # if isColinearWithEdge([i,j], loop) == False:
                    #     print("found", [i,j], data[i,j], nCross)
                    #     n=n+1
                    # else:
                    #     print("colinear", [i,j], data[i,j])
                        # newLoop = removeColinearLoopPoints([i,j], loop)
                        # nCross = getNumberOfCrosses(data, [i,j], newLoop, M)
                        # # print("\t", nCross, newLoop)
                        # if nCross != 0 and nCross % 2 == 1:
                        #     print("found", [i,j], nCross)
                        #     n=n+1
                        # else:
                        #     print("colienar out", [i,j])

            # else:
            #     print("point", [i,j], "is part of loop")

            j=j+1

        i=i+1

    print("solution:", n)

    return

def changeOneNeighbour(data, point, loop, character, N,M):

    if data[point[0], point[1]] != character and not np.any(np.all(point == loop, axis=1)):

        data[point[0], point[1]] = character
        succ = floodFill(data, point, loop, character, N,M)

        if type(succ) == type(True) and succ == False:
            # print("outside", a)
            data[point[0], point[1]] = "."
            return data, False

    else:
        return data, False

    return data, True

def floodFill(data, point, loop, character, N,M, area, n=0):

    # if 0 in point:
    #     return False

    # if a[0] >= N or a[1] >= M or a[0] == 0 or a[1] == 0:
    #     return data, False

    # print(point, np.shape(data))

    if data[point[0], point[1]] != character and not np.any(np.all(point == loop, axis=1)):

        data[point[0], point[1]] = character
        # print("hello", n)

        # if n > 20:
        #     return data, True

        if point[0] > 0:
            a = [point[0]-1, point[1]]
            data, succ = floodFill(data, a, loop, character, N,M, area, n=n+1)
            # if succ == False:
            #     return data, False
            # else:
            area.append(a)

        if point[0] < N-1:
            a = [point[0]+1, point[1]]
            data, succ = floodFill(data, a, loop, character, N,M, area, n=n+1)
            # if succ == False:
            #     return data, False
            # else:
            area.append(a)

        if point[1] > 0:
            a = [point[0], point[1]-1]
            data, succ = floodFill(data, a, loop, character, N,M, area, n=n+1)
            # if succ == False:
            #     return data, False
            # else:
            area.append(a)

        if point[1] < M-1:
            a = [point[0], point[1]+1]
            data, succ = floodFill(data, a, loop, character, N,M, area, n=n+1)
            # if succ == False:
            #     return data, False
            # else:
            area.append(a)

    return data, True

def solve21(data):

    data = [list(x) for x in data]
    data = np.array(data)
    # print(np.shape(data))
    [print("  ".join(x)) for x in data]

    path  = getPath(data)
    # print(path)

    loop = []
    for step in path:
        loop.append(step[0])
    for step in path[::-1][1:-1]:
        loop.append(step[1])

    loop = np.array(loop)
    # print(loop)
    print()

    data = cleanData(data, loop)
    [print("  ".join(x)) for x in data]

    specialChar = "*"

    N = len(data)
    M = len(data[0])
    i = 0
    while i < N:
        j=0
        while j < M:

            area = []

            if not np.any(np.all([i,j] == loop, axis=1)):
                # data[i][j] = specialChar
                data, succ = floodFill(data, [i,j], loop, specialChar, N,M, area, n=0)
                # if succ == False:
                #     for point in area:
                #         data[point[0], point[1]] = "."

                # print()
                # print([i,j])
                # print(area)

            # break

            j=j+1
        # break
        i=i+1

    n = 0
    i = 1
    while i < N-1:
        j=0
        while j < M-1:

            if data[i,j] == specialChar:
                n=n+1

            j=j+1
        i=i+1

    print()
    [print("  ".join(x)) for x in data]

    print("solution:", n)

    return

solve2(testData11)
solve2(testData21)
print(testSol21)
solve2(testData22)
print(testSol22)
solve2(testData23)
print(testSol23)
solve2(testData24)
print(testSol24)
solve2(data)
