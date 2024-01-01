import numpy as np
import time
import copy as cp

fileName = "day12_data.txt"
testData = ["???.### 1,1,3",".??..??...?##. 1,1,3","?#?#?#?#?#?#?#? 1,3,1,6",
            "????.#...#... 4,1,1","????.######..#####. 1,6,5","?###???????? 3,2,1"]
testSol = 21

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

def getIndexes(pos):

    n = [[]]
    unknown = []

    i = 0
    while i < len(pos):
        if pos[i] == "#":
            n[-1].append(i)
        elif len(n[-1]) != 0 and pos[i] in [".", "?"]:
            n.append([])

        if pos[i] == "?":
            unknown.append(i)

        i=i+1

    if n[-1] == []:
        del n[-1]

    return n, unknown

def trySetOnePosition(line, position, code, N, level):

    # print(position, n)

    for char in ["#", "."]:

        orgChar = line[position]

        line[position] = char
        n, unknown = getIndexes(line)

        # print("\tchanged", position, char, unknown, "|", level)

        if len(unknown) == 0:

            ## This means that all unknown places are filled
            ## Now we have to check if it's a good solution
            _code = [len(x) for x in n]

            if len(_code) == len(code) and all(_code[i] - code[i] == 0 for i in range(len(code))):
                # print("\tReached the end happy")
                # print("\t", code, _code)
                # print(2*"\t", "".join(line))
                # print("\treset fs1", "".join(line))
                line[position] = orgChar
                # print("\treset fs2", "".join(line))

                return True, N

            else:
                # print("\tReached the end sad")
                # print(2*"\t", code, _code)
                # print(2*"\t", "".join(line))
                # print("\treset ff", "".join(line))
                line[position] = orgChar
                # print("\treset ff", "".join(line))

                # return False, N
        else:
            succ, N = trySetOnePosition(line, unknown[0], code, N, level+1)
            # print("yay", succ, N, "|", level, unknown[0], char)

            if succ == True:
                N = N + 1

            # print("\treset 1", "".join(line))
            line[position] = orgChar
            # print("\treset 2", "".join(line))

    # print(succ)
    return None, N

def solve(data):

    solution = 0

    M = len(data)
    for i, line in enumerate(data):

        # if i % 50 == 0:
        #     print(i, "/", M)
        # line = data[5]
        # line = '???..######..#####. 1,6,5'

        # print()
        # print(line.split())

        pos = list(line.split()[0])
        code = [int(x) for x in line.split()[1].split(",")]

        n, unknown = getIndexes(pos)
        # print("n", n)

        succ, N = trySetOnePosition(pos, unknown[0], code, N=0, level=0)
        # print("final", N)

        solution = solution + N

    print("solution:", solution)

    return

data = readFile()
solve(testData)

t1 = time.time()
# solve(data)
print("elapsed:", round((time.time()-t1)/60., 2), "min")

print()
print("########### PART 2 ###############")
print()

def solve2(data, rng=5):

    solution = 0

    M = len(data)
    for i, line in enumerate(data):

        if i % 50 == 0:
            print(i, "/", M)
        # line = data[1]
        # line = '???..######..#####. 1,6,5'

        # print()
        # print(line.split())

        pos0 = list(line.split()[0])
        code0 = [int(x) for x in line.split()[1].split(",")]
        pos = pos0
        code = code0
        for x in range(rng-1):
            pos = pos + ["?"] + pos0
            code = code + code0

        # print("".join(pos))
        # print(code)
        # return

        # .??..??...?##.?.??..??...?##.

        n, unknown = getIndexes(pos)
        succ, N = trySetOnePosition(pos, unknown[0], code, N=0, level=0)
        # print("n", n)

        # # borderN = 0
        # if rng > 1:
        #     border = []
        #     # _pos = pos + ["?"] + pos

        #     bs = pos[::-1].index("#")
        #     be = pos.index("#")

        #     borderStart = 0
        #     allow = False
        #     for j, char in enumerate(pos[::-1]):
        #         print(j, char)
        #         if char == "#":
        #             allow = True
        #         if char == "." and allow == True:
        #             break

        #     border = pos[-j:n[-1][-1]+1] + pos[::-1][:bs][::-1] + ["?"] + pos[:be] + pos[n[0][0]:n[0][-1]+1]
        #     _n, _unknown = getIndexes(border)

        #     __n = len(code) - len(_n)
        #     _code = [code[-1]] + code[:__n+2]

        #     # print("pos", "".join(_pos))
        #     print("border", "".join(border))
        #     print("_code ", _code)
        #     # return

        #     n, unknown = getIndexes(border)
        #     # print(n, unknown, code, [1, _code, 1])
        #     print("code", code, n, __n,  _code)

        #     succ, borderN = trySetOnePosition(border, unknown[0], _code, N=0, level=0)
        #     print("border N:", N, borderN, N*borderN**rng)

        #     solution = solution + N*borderN**(rng-1)
        # else:
        solution = solution + N
        # break

    print()
    print("solution:", solution, "|", rng)

    return

t1 = time.time()
# solve2(testData, 1)
# solve2(testData, 2)
# solve2(testData, 3)
# solve2(testData, 4)
# solve2(testData, 5)
solve2(data, 5)
print("elapsed:", round((time.time()-t1)/60., 2), "min")