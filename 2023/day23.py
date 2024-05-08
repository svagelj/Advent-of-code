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
testSol2 = 154

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

def getCrossroads(data, N,M, extraNodes=None):

    ## "cross road location":[possible direction from crossroad]
    crossroads = {}

    i=0
    while i < N:
        j=0
        while j < M:
            if data[i][j] != "#":
                directions = getPossibleSteps2(data, [i,j], N, M)
                if len(directions) > 2:
                    crossroads[(i,j)] = directions
                    # return crossroads
            j=j+1
        i=i+1

    if extraNodes != None:
        for node in extraNodes:
            directions = getPossibleSteps2(data, list(node), N, M)
            crossroads[tuple(node)] = directions

    return crossroads

def stepForwardUntilCrossroad(data, path, N,M):

    # print("path")
    # print(path)
    possibleSteps = getPossibleSteps2(data, path[-1], N, M)
    # print(possibleSteps)

    nextSteps = []

    for step in possibleSteps:
        if step not in path or False:
            nextSteps.append(step)
    # nextSteps = possibleSteps

    if len(nextSteps) < 1:
        return path, nextSteps
    elif len(nextSteps) > 1:
        # print("yay return")
        return path, nextSteps
    else:
        newPath = path + [nextSteps[0]]
        # print("bbb", path, newPath)
        finalPath, nextSteps = stepForwardUntilCrossroad(data, newPath, N, M)

    return finalPath, nextSteps

def connectCrossroads(data, crossroads, N,M):

    ## "cross road location":[(connected crossrads, distance to croassroads)]
    tree = {}

    for crossroad in crossroads:

        # print("yaay", crossroad)
        for direction in crossroads[crossroad]:
            # print(crossroad, direction)

            ## step forward until we get to crossroad
            d = 1
            if tuple(direction) in crossroads.keys():
                ## Some crossroad is next to the current one i.e. two touching croassroads
                if crossroad not in tree:
                    tree[crossroad] == [(direction, d)]
                else:
                    tree[crossroad].append((direction, d))
            else:
                ## actualy step forward until we get to crossroad
                path = [list(crossroad), direction]
                nextPath, nextSteps = stepForwardUntilCrossroad(data, path, N, M)

                if nextPath != None and nextSteps != None:
                    dist = len(nextPath) - 1
                    # print(nextPath, nextSteps)
                    if crossroad not in tree:
                        tree[crossroad] = [(nextPath[-1], dist)]#, nextSteps)]
                    else:
                        tree[crossroad].append((nextPath[-1], dist))#, nextSteps))

    return tree

def bestRouteInTree(tree, start, end, visited):

    # start = tuple(start)
    # end = tuple(end)
    if start == end:
        ## This is added to the current distance at the end - it must be 0 for end point
        return [0]

    visited.add(start)  ## to prevent cycles
    distances = []

    # print("start:", start, tree[start])
    for crossroad in tree[start]:
        # print("yolo", crossroad)
        nextCrossroad, currDistance = crossroad ## Here is date about connected crossrads from current one
        nextCrossroad = tuple(nextCrossroad)

        if nextCrossroad not in visited:
            newDistances = bestRouteInTree(tree, nextCrossroad, end, visited)

            ## Total distance is current distance (from start to the next crossrod)
            ## and new distance (from all connected crossroads from start and onwards)
            for nextDistance in newDistances:
                distances.append(currDistance + nextDistance)

    ## so that just current path cannot visit current crossroad
    visited.remove(start)

    return distances

def printCrossrads(data, crossroads):

    data = [[ y for y in x] for x in data]

    for node in crossroads:
        i,j = node
        data[i][j] = str(len(crossroads[node]))

    print()
    print(4*" "+"".join([str(x).ljust(3) for x in np.arange(len(data[0]))]))
    print(4*" "+(3*len(data[0])-2)*"-")
    [print(str(i).rjust(2)+" |"+"  ".join(list(x))) for i,x in enumerate(data)]

    return

def printCrossradsTree(data, tree):

    data = [[ y for y in x] for x in data]

    for crossroad in tree:
        i,j = crossroad
        data[i][j] = str(len(tree[crossroad]))

    print()
    print(4*" "+"".join([str(x).ljust(3) for x in np.arange(len(data[0]))]))
    print(4*" "+(3*len(data[0])-2)*"-")
    [print(str(i).rjust(2)+" |"+"  ".join(list(x))) for i,x in enumerate(data)]

    return

def solve2(data, maxSteps=999999):

    solution = 0

    # print(data)
    # [print("".join(list(x))) for x in [[ y for y in x] for x in data]]


    N = len(data)
    M = len(data[0])
    start = [0, data[0].index(".")]
    end = [N-1, data[N-1].index(".")]
    # end = [5,5]

    # bestPath, bestDistance = bestWalk(data, start, end)

    crossroads = getCrossroads(data, N,M, [start, end])
    # print("crossroads:")
    # [print(key, value) for key, value in crossroads.items()]
    # printCrossrads(data, crossroads)

    tree = connectCrossroads(data, crossroads, N,M)
    # print("tree:")
    # [print(key, value) for key, value in tree.items()]
    # printCrossradsTree(data, tree)

    print()
    # print("start, end", start, end)
    solved = bestRouteInTree(tree, tuple(start), tuple(end), set())
    print("solved paths:", len(solved))

    solution = np.max(solved)
    # print("test solution:", testSol2)

    # print(bestDistance)
    # print(bestPath)
    # return

    # print()
    # print("best path")
    # solution = bestDistance - 1

    # ## Printout the path
    # path = [list(x) for x in cp.deepcopy(data)]
    # for i,j in bestPath:
    #     path[i][j] = "o"
    # [print("".join(list(x))) for x in [[ y for y in x] for x in path]]

    # print()
    # [print("".join(list(x))) for x in [[ y for y in x] for x in data]]
    print()
    print("solution:", solution)

    return

data = readFile()
t1 = time.time()
solve2(testData1)
print("elapsed", round(time.time()-t1, 3), "s")

t1 = time.time()
solve2(data)
print("elapsed", round((time.time()-t1)/60., 2), "min")