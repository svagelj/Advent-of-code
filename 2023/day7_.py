import numpy as np
import copy as cp

fileName = "7_data.txt"

testData = ["32T3K 765","T55J5 684","KK677 28","KTJJT 220","QQQJA 483"]
testSolution = 6440
testSolution2 = 5905

strength = ['2', '3', '4', '5', '6', '7', '8', '9', 'T', 'J', 'Q', 'K', 'A']
strength2 = ['J', '2', '3', '4', '5', '6', '7', '8', '9', 'T', 'Q', 'K', 'A']

def readFile():

    data = []

    with open(fileName, "r") as f:

        for line in f:
            # print(line[:-1])
            data.append(line[:-1])

    # print(data[:5])

    return data

def getTypeOfHand(data):

    _data = []
    for i, line in enumerate(data):

        _line = line.split(" ")
        hand = _line[0]
        h = {}
        for c in hand:
            if c not in h.keys():
                h[c] = 1
            else:
                h[c] = h[c] + 1

        keys = h.keys()
        ln = len(keys)
        if ln == 5:
            typ = 0
        elif ln == 4:
            typ = 1
        elif ln == 3:
            if 3 in [h[x] for x in keys]:
                typ = 3
            else:
                typ = 2
        elif ln == 2:
            if 4 in [h[x] for x in keys]:
                typ = 5
            else:
                typ = 4
        else:
            typ = 6

        _data.append([typ, _line[0], _line[1]])

    _data = np.array(_data)
    _data = _data[np.argsort(_data[:,0])]

    return _data[::-1]

def getStronger(hand1, hand2):

    k = 0
    while k < len(hand1):

        st1 = strength.index(hand1[k])
        st2 = strength.index(hand2[k])

        # print(st1, st2)
        if st1 > st2:
            # print(1)
            return 1
        elif st1 < st2:
            # print(2)
            return 2

        k=k+1

    return

def sortBySecondOrder(data):

    # print("########")
    ordered = []
    for i,line in enumerate(data[::-1]):
        # print("org", line)
        # print("ordr", ordered)

        if i == 0:
            ordered.append(line)

        elif i == 1:
            # print("new", ordered[0])
            s1 = getStronger(line[1], ordered[0][1])

            # print(1, s1)
            if s1 == 2:
                ordered.append(line)
            else:
                ordered.insert(-1, line)

        else:
            for j, _line in enumerate(ordered):
                # print("new", _line)

                if j == 0:
                    # print("org", line)
                    # print("new", ordered[0], ordered[1])
                    s1 = getStronger(line[1], ordered[0][1])
                    s2 = getStronger(line[1], ordered[1][1])

                    # print("!", j, s1, s2)
                    if s1 == 1 and s2 == 1:
                        ordered.insert(0, line)
                        break
                    elif s1 != s2:
                        ordered.insert(1, line)
                        break

                else:
                    # print("new", ordered[j-1], ordered[j])
                    s1 = getStronger(line[1], ordered[j][1])
                    s2 = getStronger(line[1], ordered[j-1][1])

                    # print("!!", j, s1, s2)
                    if s1 == 1 and s2 == 1:
                        ordered.insert(j-1, line)
                        break
                    elif s1 != s2:
                        ordered.insert(j, line)
                        break
                    if all(_line == ordered[-1]):
                        # print("XYOLO")
                        ordered.append(line)
                        break
                j=j+1

        # print(np.array(ordered))

    return ordered

def solveTestData(data):

    solution = 0

    _data = getTypeOfHand(data)
    orderedData = []
    print(_data)

    ## Sorting the all hands
    N = len(_data)
    i = 0
    while i < N:

        arr = []
        j=0
        while j+i < N:
            if j != 0 and _data[i+j][0] != _data[i+j-1][0]:
                break
            arr.append(_data[i+j])
            j=j+1

        # print(arr)
        ordrd = sortBySecondOrder(arr)
        # print(ordrd)
        orderedData = orderedData + ordrd

        # break
        i=i+j

    print()
    print(np.array(orderedData))

    ## calculating winnings
    for i,line in enumerate(orderedData[::-1]):

        currScore = (i+1)*int(line[2])
        solution = solution + currScore

    print()
    print(np.array(orderedData))

    print()
    print("Solution:", solution)

    return

def solve(data):

    solution = 0

    _data = getTypeOfHand(data)
    orderedData = []
    # print(_data)

    ## Sorting the all hands
    N = len(_data)
    i = 0
    while i < N:

        arr = []
        j=0
        while j+i < N:
            if j != 0 and _data[i+j][0] != _data[i+j-1][0]:
                break
            arr.append(_data[i+j])
            j=j+1

        # print(arr)
        ordrd = sortBySecondOrder(arr)
        # print(ordrd)
        orderedData = orderedData + ordrd

        # break
        i=i+j

    # print()
    # print(np.array(orderedData))

    ## calculating winnings
    for i,line in enumerate(orderedData[::-1]):

        currScore = (i+1)*int(line[2])
        solution = solution + currScore

    print()
    print("Solution:", solution)

    return


data = readFile()
# solveTestData(testData)
# solve(testData)
solve(data[:])

print()
print("//////////////////////// Part II ////////////////")
print()

def getHandType(handDict):

    keys = handDict.keys()
    values = [handDict[x] for x in keys]
    ln = len(keys)
    if ln == 5:
        typ = 0
    elif ln == 4:
        typ = 1
    elif ln == 3:
        if 3 in values:
            typ = 3
        else:
            typ = 2
    elif ln == 2:
        if 4 in values:
            typ = 5
        else:
            typ = 4
    else:
        typ = 6

    return typ

def getBestJokerHand(handDict):

    # print("before", handDict)
    keys = handDict.keys()

    maxKey = None
    maxValue = 0
    for x in keys:
        if x != "J" and maxValue < handDict[x]:
            maxValue = handDict[x]
            maxKey = x

    # print(maxKey, maxValue)
    if maxKey != None and "J" in keys:
        handDict[maxKey] = handDict[maxKey] + handDict["J"]

        # while handDict["J"] > 0:

        #     bestHand = cp.deepcopy(handDict)

        #     ## Do one permutiation
        #     for key in list(keys):
        #         if key == "J":
        #             continue
        #         handDict[key] = handDict[key] + 1



        #     handDict["J"] = handDict["J"] - 1

        del handDict["J"]

    # print("after ", handDict)

    return handDict

def getTypeOfHand2(data):

    _data = []
    for i, line in enumerate(data):

        _line = line.split(" ")
        hand = _line[0]

        h = {}
        for c in hand:
            if c not in h.keys():
                h[c] = 1
            else:
                h[c] = h[c] + 1

        h = getBestJokerHand(h)

        typ = getHandType(h)

        _data.append([typ, _line[0], _line[1]])

    _data = np.array(_data)
    _data = _data[np.argsort(_data[:,0])]

    return _data[::-1]

def getStronger2(hand1, hand2):

    k = 0
    while k < len(hand1):

        st1 = strength2.index(hand1[k])
        st2 = strength2.index(hand2[k])

        # print(st1, st2)
        if st1 > st2:
            # print(1)
            return 1
        elif st1 < st2:
            # print(2)
            return 2

        k=k+1

    return

def sortBySecondOrder2(data):

    # print("########")
    ordered = []
    for i,line in enumerate(data[::-1]):
        # print("org", line)
        # print("ordr", ordered)

        if i == 0:
            ordered.append(line)

        elif i == 1:
            # print("new", ordered[0])
            s1 = getStronger2(line[1], ordered[0][1])

            # print(1, s1)
            if s1 == 2:
                ordered.append(line)
            else:
                ordered.insert(-1, line)

        else:
            for j, _line in enumerate(ordered):
                # print("new", _line)

                if j == 0:
                    # print("org", line)
                    # print("new", ordered[0], ordered[1])
                    s1 = getStronger2(line[1], ordered[0][1])
                    s2 = getStronger2(line[1], ordered[1][1])

                    # print("!", j, s1, s2)
                    if s1 == 1 and s2 == 1:
                        ordered.insert(0, line)
                        break
                    elif s1 != s2:
                        ordered.insert(1, line)
                        break

                else:
                    # print("new", ordered[j-1], ordered[j])
                    s1 = getStronger2(line[1], ordered[j][1])
                    s2 = getStronger2(line[1], ordered[j-1][1])

                    # print("!!", j, s1, s2)
                    if s1 == 1 and s2 == 1:
                        ordered.insert(j-1, line)
                        break
                    elif s1 != s2:
                        ordered.insert(j, line)
                        break
                    if all(_line == ordered[-1]):
                        # print("XYOLO")
                        ordered.append(line)
                        break
                j=j+1

        # print(np.array(ordered))

    return ordered

def solve2(data):

    solution = 0

    _data = getTypeOfHand2(data)
    orderedData = []
    # print(_data)

    # return

    ## Sorting the all hands
    N = len(_data)
    i = 0
    while i < N:

        arr = []
        j=0
        while j+i < N:
            if j != 0 and _data[i+j][0] != _data[i+j-1][0]:
                break
            arr.append(_data[i+j])
            j=j+1

        # print(arr)
        ordrd = sortBySecondOrder2(arr)
        # print(ordrd)
        orderedData = orderedData + ordrd

        # break
        i=i+j

    # print()
    # print(np.array(orderedData[:5]))
    print(len(data), len(orderedData))

    ## calculating winnings
    for i,line in enumerate(orderedData[::-1]):
        currScore = (i+1)*int(line[2])
        solution = solution + currScore

    print("Solution:", solution)

    return orderedData

################

# def eval0(line):
#     hand, bid = line.split()
#     hand = hand.translate(str.maketrans('TJQKA', face))
#     best = max(type0(hand.replace('0', r)) for r in hand)
#     return best, hand, int(bid)

# def type0(hand):
#     return sorted(map(hand.count, hand), reverse=True)

# for face in 'ABCDE', 'A0CDE':
#     print(sum(rank * bid for rank, (*_, bid) in enumerate(sorted(map(eval0, open(fileName))), start=1)))

# face = "A0CDE"
# arr = []
# for x in sorted(map(eval0, open(fileName))):
#     arr.append(list(x))

################

# testJokers = ["J265J 1", "J8Q2J 2", "J9J2T 3", "JT94J 4"]
# getTypeOfHand2(testJokers)
# print()
solve2(testData)

data2 = readFile()
debug = solve2(data2[:])

# debug = np.array(debug[::-1])
# i = 0
# while i < len(debug):

#     # print(debug[i][-1], arr[i][-1])

#     if int(debug[i][-1]) != arr[i][-1]:
#         print(i, [debug[i][-1], arr[i][-1]])

#         print()
#         print(debug[i-2:i+8])
#         [print(x) for x in arr[i-2:i+3]]
#         break

#     i=i+1