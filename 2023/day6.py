import numpy as np
import time

fileName = "day5_data.txt"
testTimes =[ 7,  15,   30]
testDistance = [9,  40,  200]
testResult = 288 ## 4*8*9
testTimes2 = 71530
testDistance2 = 940200
testResult2 = 71503

times = [53, 91, 67, 68]
distance = [250, 1330, 1081, 1025]
time2 = 53916768
distance2 = 250133010811025

def solve(time, dist):

    solution = 1
    sols = []

    k=0
    while k < len(time):
        sols.append([])
        i=1
        while i < time[k]:

            d = i * (time[k] - i)
            if d > dist[k]:
                sols[k].append(i)
            i=i+1

        solution = solution * len(sols[k])
        k=k+1



    # print(sols)
    print("solution:", solution)

    return

solve(testTimes, testDistance)
solve(times, distance)

print()
print("########### PART 2 ###############")
print()

def solve2(time, dist):

    solution = 0
    sols = []

    i=1
    while i < time:

        # if i != 0 and i % 1000000 == 0:
        #     print("progress", i, "("+str(round(i/time*100, 2))+" %)")
        #     print("\tcurrent solution =", solution)

        ## d = v * t
        d = i * (time - i)
        if d > dist:
            solution = solution + 1
        i=i+1

    # print(sols)
    print("solution:", solution)

    return

def solve2faster(time, dist):

    solution = 0

    n=0
    i=1
    while i < time:
        d = i * (time - i)
        if d <= dist:
            n = n+1
        else:
            break
        i=i+1

    i=0
    while i < time:
        d = (time-i) * i
        if d <= dist:
            n = n+1
        else:
            break
        i=i+1

    solution = time - n
    print("solution:", solution)

    return

solve2(testTimes2, testDistance2)
solve2faster(testTimes2, testDistance2)
t1 = time.time()
# solve2(time2, distance2)
solve2faster(time2, distance2)
t2 = time.time()
print("time", round(t2-t1, 2))