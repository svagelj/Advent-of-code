import numpy as np
import time
import copy as cp

fileName = "day15_data.txt"
testData = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"
testSol1 = 1320
testSol2 = 145

def readFile():

    data = []

    with open(fileName, "r") as f:
        i=0
        for line in f:
            return line[:-1]
            data.append(line[:-1])
            # line = line[:-1]
            # print(i, line)

            i=i+1

    # print(data[:3])
    # print(data[-3:])

    return data

def HASH0(string):

    _hash = 0
    for char in string:
        _hash = _hash + ord(char)
        _hash = _hash * 17
        _hash = _hash % 256

    return _hash

def solve(data):

    solution = 0

    data = data.split(",")
    # print(data)
    for steps in data:
        h = HASH0(steps)
        # print(steps, h)

        solution = solution + h

    print()
    print("solution:", solution)

    return

data = readFile()
solve(testData)
solve(data)

print()
print("########### PART 2 ###############")
print()


def solve2(data):

    solution = 0

    data = data.split(",")
    # print(data)

    boxes = {}

    for step in data:

        # print()
        # print(step)
        if step[-1] != "-":
            label = step[:-2]
            box = HASH0(label)

            # print("b a:", boxes)
            if box not in boxes.keys():
                boxes[box] = [[label, int(step[-1])]]
            else:
                replaced = False
                i = 0
                while i < len(boxes[box]):
                    if boxes[box][i][0] == label:
                        boxes[box][i][1] = int(step[-1])
                        replaced = True
                        # break
                    i=i+1

                if replaced == False:
                    boxes[box].append([label, int(step[-1])])
            # print("a a:", boxes)

        else:
            label = step[:-1]
            box = HASH0(label)

            ## remove if present
            # print("b r:", boxes)
            if box in boxes.keys():

                # print(boxes[box])
                for _label, _focus in boxes[box]:
                    if _label == label:
                        # print(label)
                        boxes[box].remove([label, _focus])

                if len(boxes[box]) == 0:
                    del boxes[box]
            # print("a r:", boxes)

        # print(step, label, box)
        # print(boxes)

    ## Calculate focusing power
    # print(boxes)
    for box in np.sort(list(boxes.keys())):

        i = 0
        while i < len(boxes[box]):

            focus = (box+1) * (i+1) * boxes[box][i][1]
            # print(box+1, i+1, boxes[box][i][1], "|", focus)

            solution = solution + focus
            i=i+1

    print()
    print("solution:", solution)

    return

solve2(testData)
solve2(data)