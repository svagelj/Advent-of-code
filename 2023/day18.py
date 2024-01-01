import numpy as np
import time
import copy as cp

fileName = "day18_data.txt"
testData = ["R 6 (#70c710)","D 5 (#0dc571)","L 2 (#5713f0)","D 2 (#d2c081)","R 2 (#59c680)","D 2 (#411b91)",
            "L 5 (#8ceee2)","U 2 (#caa173)","L 1 (#1b58a2)","U 2 (#caa171)","R 2 (#7807d2)","U 3 (#a77fa3)",
            "L 2 (#015232)","U 2 (#7a21e3)"]
testSol1 = 62
testSol2 = 952408144115

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

def walkTheWalk(steps):

    coord = [[0,0]]

    i=0
    while i < len(steps):

        # print(steps[i])
        direction, distance, color = steps[i].split()

        curr = coord[-1]

        if direction == "R":
            d = int(distance)
            for j in range(1,d+1):
                nextPos = [curr[0], curr[1] + j]
                coord.append(nextPos)
        elif direction == "D":
            d = int(distance)
            for j in range(1,d+1):
                nextPos = [curr[0] + j, curr[1]]
                coord.append(nextPos)
        if direction == "L":
            d = int(distance)
            for j in range(1,d+1):
                nextPos = [curr[0], curr[1] - j]
                coord.append(nextPos)
        if direction == "U":
            d = int(distance)
            for j in range(1,d+1):
                nextPos = [curr[0] - j, curr[1]]
                coord.append(nextPos)

        # break

        i=i+1

    return coord[:-1]

def floodPool(data, pos, N,M, level=0):

    i,j = pos
    data[i][j] = 1

    # if level > 2:
    #     return data

    if i < N-1 and data[i+1][j] != 1:
        floodPool(data, [i+1, j], N,M, level=level+1)
    if i > 0 and data[i-1][j] != 1:
        floodPool(data, [i-1, j], N,M, level=level+1)

    if j < M-1 and data[i][j+1] != 1:
        floodPool(data, [i, j+1], N,M, level=level+1)
    if i > 0 and data[i][j-1] != 1:
        floodPool(data, [i, j-1], N,M, level=level+1)

    return data

def isPositionInArray(array, pos):

    for p in array:
        if pos[0] == p[0] and pos[1] == p[1]:
            return True

    return False

def floodPool2(data, pos, N,M, level=0):

    queue = [pos]

    # print(10*"-")
    # print("start", pos)

    k=0
    for p in queue:

        i,j = p
        data[i][j] = 1

        if k % 5000 == 0:
            print("progres", k, "/", M*N, "("+str(round(100.*k/(M*N), 2))+" %)")

        if i < N-1 and data[i+1][j] != 1 and isPositionInArray(queue, [i+1,j]) == False:
            queue.append([i+1, j])
        if i > 0 and data[i-1][j] != 1 and isPositionInArray(queue, [i-1,j]) == False:
            queue.append([i-1, j])

        if j < M-1 and data[i][j+1] != 1 and isPositionInArray(queue, [i,j+1]) == False:
            queue.append([i, j+1])
        if i > 0 and data[i][j-1] != 1 and isPositionInArray(queue, [i,j-1]) == False:
            queue.append([i, j-1])

        k=k+1
        # if k > N:
        #     break

    # print(k)

    # print(10*"-")
    return data

def solve(data, maxSteps=100):

    solution = 0

    # print(data)

    coord = np.array(walkTheWalk(data))
    N = np.max(coord[:,0]) - np.min(coord[:,0]) + 1
    M = np.max(coord[:,1]) - np.min(coord[:,1]) + 1
    dy = abs(np.min(coord[:,0]))
    dx = abs(np.min(coord[:,1]))

    # print()
    # print(coord[:50])
    # print(N,M, np.min(coord[:,0]), np.min(coord[:,1]))

    space = np.zeros((N,M))
    for pos in coord:
        space[pos[0]+dy, pos[1]+dx] = 1

    # print()
    # print(space)
    # [print("".join(list(x))) for x in [[ str(int(y)) for y in x] for x in space]]

    space = floodPool2(space, [dy+1,dx+1], N,M)
    # print()
    # print(space)
    # [print("".join(list(x))) for x in [[ str(int(y)) for y in x] for x in space]]

    solution = np.sum(space)
    print()
    print("solution:", solution, "("+str(round(100*solution/(M*N), 2))+" %)")

    return

data = readFile()
solve(testData)

# import sys
# sys.setrecursionlimit(1000000)
# solve(data)

print()
print("########### PART 2 ###############")
print()

def walkTheWalk2(steps):

    coord = [[0,0]]

    dirs = {"0":"R", "1":"D", "2":"L","3":"U"}

    i=0
    while i < len(steps):

        # print(steps[i])
        direction, distance, color = steps[i].split()

        color = color[2:-1]
        distance = int(color[:5], 16)
        direction = dirs[color[-1]]

        # print(color, distance, direction)
        # break

        curr = coord[-1]

        if direction == "R":
            d = int(distance)
            nextPos = [curr[0], curr[1] + d]
            coord.append(nextPos)
        elif direction == "D":
            d = int(distance)
            nextPos = [curr[0] + d, curr[1]]
            coord.append(nextPos)
        if direction == "L":
            d = int(distance)
            nextPos = [curr[0], curr[1] - d]
            coord.append(nextPos)
        if direction == "U":
            d = int(distance)
            nextPos = [curr[0] - d, curr[1]]
            coord.append(nextPos)

        # break

        i=i+1

    return coord[:]

def solve2(data, maxSteps=100):

    solution = 0

    # data = ["R 2 s", "D 5 3", "R 5 3", "D 5 23", "L 5 3", "D 2 2", "R 5 2", "D 1 2", "L 7 3", "U 13 2"]
    # data = ["R 5 4", "D 5 3", "R 5 3", "D 5 23", "L 10 3", "U 10 2"]
    # data = ["R 5 4", "D 5 3", "L 5 3", "U 5 2"]
    # print(data)

    coord = np.array(walkTheWalk2(data))
    N = np.max(coord[:,0]) - np.min(coord[:,0]) + 1
    M = np.max(coord[:,1]) - np.min(coord[:,1]) + 1
    dy = abs(np.min(coord[:,0]))
    dx = abs(np.min(coord[:,1]))

    # print()
    # print(coord[:50])
    # print(N,M, np.min(coord[:,0]), np.min(coord[:,1]))

    ###
    ### https://en.wikipedia.org/wiki/Pick's_theorem
    ### A = i + b/2 -1   =>   i + b = A + b/2 + 1
    ### A - normal area, b - boundary, i - interier points
    ## we need i + b because that is the number of interior and boundary points
    ###
    A = 0
    line = 0
    # print()
    i=1
    while i < len(coord):
        w = coord[i-1,1] - coord[i,1]
        h = (coord[i,0] + coord[i-1,0])/2
        A = A + w*h

        line = line + np.max(np.abs(coord[i] - coord[i-1]))

        i=i+1

    # print(A, line)
    # print(A - 1 + line/2.)

    solution = A + 1 + line/2

    # coord = np.array(walkTheWalk(data))
    # space = np.zeros((N,M))
    # for pos in coord:
    #     space[pos[0], pos[1]] = 1

    # print()
    # # print(space)
    # [print(" ".join(list(x))) for x in [[ str(int(y)) for y in x] for x in space]]

    # space = floodPool2(space, [dy+1,dx+1], N,M)
    # # print()
    # # print(space)
    # [print(" ".join(list(x))) for x in [[ str(int(y)) for y in x] for x in space]]

    # solution = np.sum(space)
    print()
    print("solution:", solution)#, "|", np.sum(space))

    return

solve2(testData)
solve2(data)