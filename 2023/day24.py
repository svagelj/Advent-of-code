import numpy as np
import time
import copy as cp

fileName = "day24_data.txt"
testData1 = ["19, 13, 30 @ -2,  1, -2","18, 19, 22 @ -1, -1, -2","20, 25, 34 @ -2, -2, -4",
             "12, 31, 28 @ -1, -2, -1","20, 19, 15 @  1, -5, -3"]

testSol1 = 2

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

def solve(data, positionMin=7, positionMax=27):

    solution = 0

    # print(data)
    # [print("".join(list(x))) for x in [[ y for y in x] for x in data]]

    N = len(data)
    for i,hail1 in enumerate(data):

        p1, v1 = hail1.split(" @ ")
        p1 = [int(x) for x in p1.split(", ")]
        v1 = [int(x) for x in v1.split(", ")]
        # print()
        # print(i, hail1)
        # print(p1, v1)

        x01, y01, z01 = p1
        u1, v1, w1 = v1

        if i +1 < N:
            for hail2 in data[i+1:]:
                # print("\t", hail2)

                p2, v2 = hail2.split(" @ ")
                x02, y02, z02 = [int(x) for x in p2.split(", ")]
                u2, v2, w2 = [int(x) for x in v2.split(", ")]

                base = v2 - v1*u2/u1

                if base != 0:

                    t2 = ( y01 - y02 + (x02-x01)*v1/u1 ) / base
                    t1 = (x02 - x01 + u2*t2) / u1

                    x1 = round(x01 + u1*t1, 3)
                    y1 = round(y01 + v1*t1, 3)
                    x2 = round(x02 + u2*t2, 3)
                    y2 = round(y02 + v2*t2, 3)
                    # print("\t\t t1, t2:", round(t1,1), round(t2, 1))
                    # print("\t\t pos:", [x1,y1], [x2,y2])

                    timeCon = t1 >= 0 and t2 >= 0
                    positionConX = x1 >= positionMin and x1 <= positionMax and x2 >= positionMin and x2 <= positionMax
                    positionConY= y1 >= positionMin and y1 <= positionMax and y2 >= positionMin and y2 <= positionMax
                    if timeCon and positionConX and positionConY:
                        # print("\t\t yay")
                        solution = solution + 1

                # else:
                #     print("\t\t base is zero")

                # break

        # break

    # print()
    # [print("".join(list(x))) for x in [[ y for y in x] for x in data]]
    print()
    print("solution:", solution)

    return

data = readFile()
solve(testData1)
solve(data, positionMin = 200000000000000, positionMax=400000000000000)

print()
print("########### PART 2 ###############")
print()

