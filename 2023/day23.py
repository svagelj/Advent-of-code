import numpy as np
import time
import copy as cp

fileName = "day23_data.txt"
testData1 = ["#.#####################","#.......#########...###","#######.#########.#.###","###.....#.>.>.###.#.###",
             "###v#####.#v#.###.#.###","###.>...#.#.#.....#...#","###v###.#.#.#########.#","###...#.#.#.......#...#",
             "#####.#.#.#######.#.###","#.....#.#.#.......#...#","#.#####.#.#.#########v#","#.#...#...#...###...>.#",
             "#.#.#v#######v###.###v#","#...#.>.#...>.>.#.###.#","#####v#.#.###v#.#.###.#","#.....#...#...#.#.#...#",
             "#.#########.###.#.#.###","#...###...#...#...#.###","###.###.#.###v#####v###","#...#...#.#.>.>.#.>.###",
             "#.###.###.#.###.#.#v###","#.....###...###...#...#","#####################.#"]

testSol1 = 94

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

def getPossibleSteps(data, position, N,M):

    inds = []
    i,j = position

    # print("p 4:", position)

    if i > 0 and data[i-1][j] not in ["#", "v"]:
        inds.append([i-1, j])

    if i < N-1 and data[i+1][j] not in ["#", "^"]:
        inds.append([i+1, j])

    if j > 0 and data[i][j-1] not in ["#", ">"]:
        inds.append([i, j-1])

    if j < M-1 and data[i][j+1] not in ["#", "<"]:
        inds.append([i, j+1])

    return inds

def solve(data, maxSteps=9999):

    solution = 0

    # print(data)
    [print("".join(list(x))) for x in [[ y for y in x] for x in data]]

    N = len(data)
    M = len(data[0])

    start = [0, data[0].index(".")]
    end = [N-1, data[N-1].index(".")]
    # end = [5,5]

    paths = {0:[start]}
    # print(paths)
    bestPath = []
    bestDistance = 0

    i = 0
    while i < maxSteps:

        n = 0
        # hist = cp.deepcopy(paths)
        keys = list(paths.keys())
        for j in keys:
            path = paths[j]

            # print()
            # print(path)

            ## Step downhill and end this path if end is reached
            # print("bb", path)
            ind = path[-1]
            # print(ind)
            h = data[ind[0]][ind[1]]
            # print("\t", ind,h)
            if h == ">":
                path.append([ind[0], ind[1]+1])
            elif h == "<":
                path.append([ind[0], ind[1]-1])
            elif h == "v":
                # print("yay")
                # print(k, paths[k])
                path.append([ind[0]+1, ind[1]])
            elif h == "^":
                path.append([ind[0]-1, ind[1]])

            if path[-1][0] == end[0] and path[-1][1] == end[1]:
                bestPath = cp.deepcopy(path)
                # print("end path")
                bestDistance = max(bestDistance, len(path))
                del paths[j]
                continue
            # print("aa", path)

            steps = getPossibleSteps(data, path[-1], N,M)
            # print("possible b:", steps)

            ## check if step in steps are already been visited
            # print("b", steps)
            for _p in steps:
                if _p in path:
                    steps.remove(_p)
            n = n + len(steps)
            # print("possible a:", steps)

            if len(steps) == 1:
                paths[j].append(steps[0])
            else:
                _p = cp.deepcopy(paths[j])
                for k, step in enumerate(steps):
                    if k == 0:
                        paths[j].append(step)
                    else:
                        _p.append(step)
                        key = np.max(list(paths.keys())) + 1
                        # print("new key:", key)
                        paths[key] = _p

        # for key in paths.keys():
        #     print(paths[key])
        # print(len(paths.keys()))


        # if i > 15:
        #     break

        if n == 0:
            print("no more paths to explore", i)
            break

        i=i+1

    print()
    print("best path")
    solution = bestDistance - 1
    path = [list(x) for x in cp.deepcopy(data)]
    for i,j in bestPath:
        path[i][j] = "o"
    [print("".join(list(x))) for x in [[ y for y in x] for x in path]]

    # print()
    # [print("".join(list(x))) for x in [[ y for y in x] for x in data]]
    print()
    print("solution:", solution)

    return

data = readFile()
solve(testData1)
# solve(data)

print()
print("########### PART 2 ###############")
print()

def getPossibleSteps2(data, position, N,M):

    inds = []
    i,j = position

    # print("p 4:", position)

    if i > 0 and data[i-1][j] != "#":
        inds.append([i-1, j])

    if i < N-1 and data[i+1][j] != "#":
        inds.append([i+1, j])

    if j > 0 and data[i][j-1] != "#":
        inds.append([i, j-1])

    if j < M-1 and data[i][j+1] != "#":
        inds.append([i, j+1])

    return inds

def solve2(data, maxSteps=999999):

    solution = 0

    # print(data)
    [print("".join(list(x))) for x in [[ y for y in x] for x in data]]

    N = len(data)
    M = len(data[0])

    start = [0, data[0].index(".")]
    end = [N-1, data[N-1].index(".")]
    # end = [5,5]

    print("TODO normal Dijkstra  algorithm")
    return

    paths = {0:[start]}
    # print(paths)
    bestPath = []
    bestDistance = 0

    i = 0
    while i < maxSteps:

        n = 0
        # hist = cp.deepcopy(paths)
        keys = list(paths.keys())
        for j in keys:
            path = paths[j]

            print()
            # print(path)

            ## Step downhill and end this path if end is reached
            # print("bb", path)
            # ind = path[-1]
            # print(ind)
            # h = data[ind[0]][ind[1]]
            # print("\t", ind,h)
            # if h == ">":
            #     path.append([ind[0], ind[1]+1])
            # elif h == "<":
            #     path.append([ind[0], ind[1]-1])
            # elif h == "v":
            #     # print("yay")
            #     # print(k, paths[k])
            #     path.append([ind[0]+1, ind[1]])
            # elif h == "^":
            #     path.append([ind[0]-1, ind[1]])

            if path[-1][0] == end[0] and path[-1][1] == end[1]:
                bestPath = cp.deepcopy(path)
                # print("end path")
                bestDistance = max(bestDistance, len(path))
                del paths[j]
                continue
            # print("aa", path)

            steps = getPossibleSteps2(data, path[-1], N,M)
            # print("possible b:", steps)

            ## check if step in steps are already been visited
            # print("b", steps)
            for _p in steps:
                if _p in path:
                    steps.remove(_p)
            n = n + len(steps)
            # print("possible a:", steps)

            if len(steps) == 1:
                paths[j].append(steps[0])
            else:
                _p = cp.deepcopy(paths[j])
                for k, step in enumerate(steps):
                    if k == 0:
                        paths[j].append(step)
                    else:
                        _p.append(step)
                        key = np.max(list(paths.keys())) + 1
                        # print("new key:", key)
                        paths[key] = _p

        for key in paths.keys():
            print(paths[key])
        print(len(paths.keys()))


        if i > 15:
            break

        if n == 0:
            print("no more paths to explore", i)
            break

        i=i+1

    print()
    print("best path")
    solution = bestDistance - 1
    path = [list(x) for x in cp.deepcopy(data)]
    for i,j in bestPath:
        path[i][j] = "o"
    [print("".join(list(x))) for x in [[ y for y in x] for x in path]]

    # print()
    # [print("".join(list(x))) for x in [[ y for y in x] for x in data]]
    print()
    print("solution:", solution)

    return

data = readFile()
solve2(testData1)