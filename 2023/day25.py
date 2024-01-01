import numpy as np
import time
import copy as cp

fileName = "day25_data.txt"
testData1 = ["jqt: rhn xhk nvd","rsh: frs pzl lsr","xhk: hfx","cmg: qnr nvd lhk bvb","rhn: xhk bvb hfx",
             "bvb: xhk hfx","pzl: lsr hfx nvd","qnr: nvd","ntq: jqt hfx bvb xhk","nvd: lhk","lsr: lhk",
             "rzs: qnr cmg lsr rsh","frs: qnr lhk lsr"]
testSol1 = 54

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

def initData(data):

    _data = {}

    for line in data:

        # print(line)

        base, connections = line.split(": ")

        if base not in _data.keys():
            _data[base] = connections.split()
        else:
            [_data[base].append(x) for x in connections.split()]

        for con in connections.split():
            if con not in _data.keys():
                _data[con] = [base]
            else:
                _data[con].append(base)

    # print(_data)

    return _data

def findNext(data, curr, cycle, level=0):

    newCycle = []
    for connection in data[curr]:

        if level == 0:
            print(level, curr)

        if level >= len(data):
            print("yay")
            return data, cycle, True

        nc = cp.deepcopy(cycle)
        nc.append(connection)
        # newCycle.append(nc)

        _d, _cycle, success = findNext(data, connection, nc, level=level+1)
        if success == True:
            newCycle.append(_cycle)

    return data, newCycle, False

def findCycles(data):

    nodes = list(data.keys())

    node = nodes[0]
    # print(node, data[node], [node])
    data, cycle, success = findNext(data, node, [node])

    print("finnit:", cycle, success)

    return

def findCycles2(data, paths, level=0):

    keys = list(paths.keys())
    for key in keys:
        path = paths[key]
        # print("path", path)

        node = path[-1]
        copyPath = cp.deepcopy(path)
        first = False
        for connection in data[node]:

            if connection in path:
                continue

            if first == False:
                paths[key].append(connection)
                first = True
            else:
                newPath = cp.deepcopy(copyPath)
                newPath.append(connection)

                # print("yolo", paths.keys())
                key2 = np.max(list(paths.keys())) + 1
                paths[key2] = newPath

    if level < 3:
        findCycles2(data, paths, level=level+1)

    return paths

def solve(data, positionMin=7, positionMax=27):

    solution = 0

    # print(data)
    # [print("".join(list(x))) for x in [[ y for y in x] for x in data]]

    data2 = initData(data)
    [print(key, data2[key]) for key in data2.keys()]
    print()

    nodes = list(data2.keys())
    paths = findCycles2(data2, {0:[nodes[0]]})
    # [print(key, paths[key]) for key in paths.keys()]
    print("len paths:", len(paths))

    # print()
    # [print("".join(list(x))) for x in [[ y for y in x] for x in data]]
    print()
    print("solution:", solution)

    return

data = readFile()
solve(testData1)

print()
print("########### PART 2 ###############")
print()

