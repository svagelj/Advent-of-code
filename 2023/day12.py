import numpy as np
import time
import copy as cp
import random

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

def getCode(line):

    n = []
    for char in line:
        if char == "#":
            if len(n) == 0:
                n = [1]
            else:
                n[-1] = n[-1] + 1
        elif len(n) != 0 and n[-1] != 0:
            n.append(0)

    if len(n) != 0 and n[-1] == 0:
        del n[-1]
    return n

def isCurrentCodePossible(code, currentCode):

    if len(currentCode) == 0:
        return True
    elif len(currentCode) > len(code):
        return False
    else:

        N = len(currentCode)
        i = 0
        while i < N-1:
        # for i,num in enumerate(currentCode):
            if currentCode[i] != code[i]:
                return False

            i=i+1

        if currentCode[-1] > code[N-1]:# or N == len(code):
            return False


    return True

def trySetOnePosition2(line, position, code, N, ind, level):

    # if level == 0 or True:
    #     # [print(x) for x in [line, position, code, N, level]]
    #     print("\ncurrent", line, position, code, N, ind, level)

    ## Ending condition
    if ind == len(position):
        _code = getCode(line)
        if len(_code) == len(code) and all(_code[i] - code[i] == 0 for i in range(len(code))):
            N = N + 1
            # print("succ, N:", True, N)
            return True, N
        else:
            # print("succ, N:", False, N)
            return False, N

    pos = position[ind]
    orgChar = line[pos]

    maxInd = pos + 1
    for char in line[maxInd:]:
        if char != "#":
            break
        maxInd = maxInd + 1
    # print("\t maxInd:", maxInd)

    chars = [".", "#"]
    random.shuffle(chars)
    for char in chars:

        line[pos] = char
        # print("\t", pos, char)
        # print("\t", "".join(line))

        ## check if current string is possible for solution
        # n, unknown = getIndexes(line[:maxInd])
        # print("\t curr", line[:maxInd])
        # print("\t", n)
        # print("\t", unknown)

        currentCode = getCode(line[:maxInd])
        isPossible = isCurrentCodePossible(code, currentCode)
        # print("\tcode", currentCode, isPossible)

        # input("continue?")

        if isPossible == True:
            succ, N = trySetOnePosition2(line, position, code, N, ind+1, level+1)
        # else:
        #     print("\tthis one is not possible - skipping")

        line[pos] = orgChar

    # print(succ)
    return None, N

def stripDots(position):

    newPos = []
    for char in position:
        if len(newPos) == 0 or not (char == "." and newPos[-1] == "."):
            newPos.append(char)

    return "".join(newPos)

def solved(line, code):

    # print()
    # print(line, code)

    ## key represents index where is the start of the search for current group number
    ## value represents a counter of possible arrangements
    ## at the start the value must be 1 because this will be added each time a succesfull arangement is found
    ##      key must be 0 because this is the first index
    positions = {0: 1}
    # print("\t starting positions", positions)

    ## loop over each group number
    for i, nGroup in enumerate(code):

        new_positions = {}

        for start, counter in positions.items():

            ## Length of the line MINUS sum of leftover codes PLUS lenth of array of leftover codes
            ## WHY is this lenght as it is?
            ## This is the amount of line to include in search for current group number (or interval)
            ## it has to be long enough to include all posible combinations of this group number
            ## and at the same time leave enough space for the rest of the group numbers
            ## it is ended if start is '#'
            length = len(line) - sum(code[i + 1:]) + len(code[i + 1:])
            # print("\t", i, nGroup, "|", start,"to", length)
            # print("\t", len(line), sum(code[i + 1:]), len(code[i + 1:]))

            for n in range(start, length):
                # print(2*"\t", n,"|", n + nGroup, line[n:n + nGroup])

                ## we are looking at interval [n, n+nGroup]
                ## i.e. from n to length of '#' there has to be according to 'code'
                ## so on this interval we expect NO '.'

                intervalEnd = n + nGroup

                ## the whole interval must be inside the line and without '.'
                if intervalEnd - 1 < len(line) and '.' not in line[n:intervalEnd]:

                    ## if we are looking at last code group AND there are no '#' from this interval onwards
                    ##      that means we reached the end of the string and all group numbers are done
                    ## or any other code group AND this interval is within the whole line AND '#' is not the first char after this interval
                    ##      this means that the number group ended cleanly, with '.' just after this group

                    if (i == len(code) - 1 and '#' not in line[intervalEnd:]) or \
                        (i < len(code) - 1 and intervalEnd < len(line) and '#' != line[intervalEnd]):

                        ## increase the counter for current group number
                        ## value of key is used as index position to start the next search

                        # print(3*"\t", "yay", n, intervalEnd)

                        if intervalEnd + 1 in new_positions:
                            new_positions[intervalEnd + 1] = new_positions[intervalEnd + 1] + counter
                        else:
                            new_positions[intervalEnd + 1] = counter

                        # print(new_positions)

                ## stop this group search if first element of interval [n, n+nGroup] is '#'
                ## because that means that the next group number has started, so we end the current one
                if line[n] == '#':
                    # print(4*"\t", "break")
                    break

            # raise Exception("trol")

        ## if this dictionary is empty, there are no solutions
        ## if there are solutions for this group number, the next search start is behind the end of this group number
        positions = new_positions
        # print("\t new positions", positions)


    N =  sum(positions.values())

    return N

def solve2(data, rng=5):

    solution = 0

    M = len(data)
    for i, line in enumerate(data[::]):

        # print()

        # if i % 10 == 0:
        #     print(i, "/", M)
        # line = data[1]
        # line = '???..######..#####. 1,6,5'

        # print(line.split())

        pos0 = line.split()[0]
        pos0 = stripDots(pos0)

        code0 = [int(x) for x in line.split()[1].split(",")]
        pos = list("?".join(rng*[pos0]))
        code = rng*code0

        n, unknown = getIndexes(pos)

        # print("---")
        # print("".join(pos))
        # print(code, len(unknown))
        # return

        # succ, N = trySetOnePosition2(pos, unknown, code, N=0, ind=0, level=0)
        N = solved("".join(pos), code)
        # print("N", N)

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
solve2(testData, 5)
solve2(data, 5)
print("elapsed:", round((time.time()-t1)/60., 2), "min")