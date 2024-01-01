import numpy as np
import time
import copy as cp

fileName = "day13_data.txt"
testData = [["#.##..##.","..#.##.#.","##......#","##......#","..#.##.#.","..##..##.","#.#.##.#."],
              ["#...##..#","#....#..#","..##..###","#####.##.","#####.##.","..##..###","#....#..#"]]
testSol1 = 405
testSol2 = 400

def readFile():

    data = []
    _dat = []

    with open(fileName, "r") as f:
        i=0
        for line in f:
            # print(line)
            if line == "\n":
                data.append(_dat)
                _dat = []
            else:
                _dat.append(line[:-1])
            # line = line[:-1]
            # print(i, line)

            i=i+1
        data.append(_dat)

    # print(data[:3])
    # print(data[-3:])

    return data

def findVerticalSymertry(pattern):

    # afterIndexSymetry = None

    N = len(pattern)
    M = len(pattern[0])

    candidates = []

    i=0
    while i < N:

        _canditates = []
        j=1
        while j < M:
            jMin = min(j, M-j)
            # print(i,j, "|", jMin)

            left = pattern[i][j-jMin:j][::-1]
            right = pattern[i][j:j+jMin]
            # print("\t",left)
            # print("\t",right)

            if all(left[x] == right[x] for x in range(len(left))):
                # print("\tyay")
                _canditates.append(j)
            j=j+1

        # print("\tcand_", _canditates)

        if len(_canditates) == 0:
            # print("so so sad")
            return None

        if i == 0:
            candidates = cp.deepcopy(_canditates)
        else:
            # print("b", _canditates)
            # print(canditates)
            for x in cp.deepcopy(_canditates):
                if x not in candidates:
                    # print("\t_", x, "not in", candidates)
                    _canditates.remove(x)
            # print("a", _canditates)

            # print("b", candidates)
            for xx in cp.deepcopy(candidates):
                # print("\t\tyolo", xx)
                if xx not in _canditates:
                    # print("\t", xx, "not in", _canditates)
                    candidates.remove(xx)
            # print("\t", candidates)

            if len(candidates) == 0:
                # print("so sad", i)
                return None

        # print("\tcand", candidates)

        # canditates = _canditates
        # return
        i=i+1

    # print(canditates)
    return candidates

def solve(data):

    solution = 0
    vert = []
    hor = []

    N = len(data)
    for i,pattern in enumerate(data[::]):
        # pattern = data[-1]

        # [print(" ".join(list(x))) for x in pattern]
        # print()

        pattern2 = cp.deepcopy(pattern)

        patternT = np.array([list(x) for x in pattern])
        patternT = patternT.transpose()
        patternT = [list(x) for x in patternT]

        # [print(" ".join(list(x))) for x in patternT]
        # print()

        horizontal = findVerticalSymertry(patternT)
        # print("horizontal", horizontal)

        # [print(" ".join(list(x))) for x in patternT]
        # print()

        if horizontal != None:
            if len(horizontal) > 1:
                print(i, "horizontal:", horizontal)
            solution = solution + 100*horizontal[0]

            pattern2.insert(horizontal[0], "".join(["-" for x in range(len(pattern2[0]))]))

            if horizontal[0] > len(pattern2) // 2:
                pattern2.insert( -2*(len(pattern2)-horizontal[0])+1, "".join(["=" for x in range(len(pattern2[0]))]))
            else:
                pattern2.insert(2*horizontal[0] - len(pattern2)+1, "".join(["=" for x in range(len(pattern2[0]))]))

        vertical = findVerticalSymertry(pattern)
        # print("vertical", vertical)

        if vertical != None:
            if len(vertical) > 1:
                print(i, "vertical:", vertical)
            solution = solution + vertical[0]

            patternT.insert(vertical[0], "".join(["|" for x in range(len(patternT[0]))]))

            if vertical[0] > len(patternT) // 2:
                patternT.insert( -2*(len(patternT)-vertical[0])+1, "".join(["!" for x in range(len(patternT[0]))]))
            else:
                patternT.insert(2*vertical[0]-len(patternT)+1, "".join(["!" for x in range(len(patternT[0]))]))

            pattern2 = np.array([list(x) for x in patternT])
            pattern2 = pattern2.transpose()
            pattern2 = [list(x) for x in pattern2]

        # [print(" ".join(list(x))) for x in pattern2]
        # print()
        # break
        # input(str(i)+" / "+str(N)+" | yolo "+str(solution))
        # print()

    print("solution:", solution)

    return

data = readFile()
solve(testData)
solve(data)

# t1 = time.time()
# print("elapsed:", round((time.time()-t1)/60., 2), "min")

print()
print("########### PART 2 ###############")
print()

def findVerticalSymertry2(pattern):

    # print()
    # [print(" ".join(list(x))) for x in pattern]

    N = len(pattern)
    M = len(pattern[0])

    candidates = {}

    i=0
    while i < N:

        j=1
        while j < M:
            jMin = min(j, M-j)
            # print(i,j, "|", jMin)

            left = pattern[i][j-jMin:j][::-1]
            right = pattern[i][j:j+jMin]
            # print("\t",left)
            # print("\t",right)

            n = 0
            for x in range(len(left)):
                if left[x] != right[x]:
                    n=n+1

            # if n == 1:
            #     print(i,j)
            #     print("\t",left)
            #     print("\t",right)

            if j not in candidates.keys():
                candidates[j] = n
            else:
                candidates[j] = candidates[j] + n

            # if all(left[x] == right[x] for x in range(len(left))):
            #     # print("\tyay")
            #     _canditates.append(j)
            j=j+1

        # print("\tcand_", _canditates)

        # canditates = _canditates
        # return
        i=i+1

    for key in cp.deepcopy(candidates).keys():
        if candidates[key] != 1:
            del candidates[key]

    # print(list(candidates.keys()))
    # print("\tcand", candidates.keys())

    if len(candidates.keys()) == 0:
        return None

    return list(candidates.keys())

def solve2(data):

    solution = 0

    N = len(data)
    for i,pattern in enumerate(data[::]):
        # pattern = data[-1]

        # [print(" ".join(list(x))) for x in pattern]
        # print()

        pattern2 = cp.deepcopy(pattern)

        patternT = np.array([list(x) for x in pattern])
        patternT = patternT.transpose()
        patternT = [list(x) for x in patternT]

        # [print(" ".join(list(x))) for x in patternT]
        # print()

        horizontal = findVerticalSymertry2(patternT)
        # print("horizontal", horizontal)

        # [print(" ".join(list(x))) for x in patternT]
        # print()

        if horizontal != None:
            if len(horizontal) > 1:
                print(i, "horizontal:", horizontal)
            solution = solution + 100*horizontal[0]

            pattern2.insert(horizontal[0], "".join(["-" for x in range(len(pattern2[0]))]))

            if horizontal[0] > len(pattern2) // 2:
                pattern2.insert( -2*(len(pattern2)-horizontal[0])+1, "".join(["=" for x in range(len(pattern2[0]))]))
            else:
                pattern2.insert(2*horizontal[0] - len(pattern2)+1, "".join(["=" for x in range(len(pattern2[0]))]))


        # break

        vertical = findVerticalSymertry2(pattern)
        # print("vertical", vertical)

        if vertical != None:
            if len(vertical) > 1:
                print(i, "vertical:", vertical)
            solution = solution + vertical[0]

            patternT.insert(vertical[0], "".join(["|" for x in range(len(patternT[0]))]))

            if vertical[0] > len(patternT) // 2:
                patternT.insert( -2*(len(patternT)-vertical[0])+1, "".join(["!" for x in range(len(patternT[0]))]))
            else:
                patternT.insert(2*vertical[0]-len(patternT)+1, "".join(["!" for x in range(len(patternT[0]))]))

            pattern2 = np.array([list(x) for x in patternT])
            pattern2 = pattern2.transpose()
            pattern2 = [list(x) for x in pattern2]

        # [print(" ".join(list(x))) for x in pattern2]
        # print()
        # break
        # input(str(i)+" / "+str(N)+" | yolo "+str(solution))
        # print()

    print("solution:", solution)

    return

solve2(testData)
solve2(data)