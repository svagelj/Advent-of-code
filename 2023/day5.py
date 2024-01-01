import numpy as np
import time

fileName = "day5_data.txt"
testData = """seeds: 79 14 55 13

seed-to-soil map:
50 98 2
52 50 48

soil-to-fertilizer map:
0 15 37
37 52 2
39 0 15

fertilizer-to-water map:
49 53 8
0 11 42
42 0 7
57 7 4

water-to-light map:
88 18 7
18 25 70

light-to-temperature map:
45 77 23
81 45 19
68 64 13

temperature-to-humidity map:
0 69 1
1 0 69

humidity-to-location map:
60 56 37
56 93 4
"""

testResult = 35

def decodeTestData(data):

    dat = []
    _map = []
    mapName = None

    lines = data.split("\n")
    for i, line in enumerate(lines):
        # print(line)

        if i == 0:
            dat.append([int(x) for x in line.split(": ")[1].split()])

        elif i != 1 and  "" == line:
            dat.append(_map)
            _map = []

            # print("yolo")

        elif "map:" in line:
            mapName = line.split()[0]

        elif line != "" and "map:" not in line:
            _line = line.split()
            _line = [int(x) for x in _line]
            # print(_line)
            # src = np.arange(_line[1], _line[1]+_line[2])
            # dst = np.arange(_line[0], _line[0]+ _line[2])
            src = [_line[1], _line[1]+_line[2]-1]
            dst = [_line[0], _line[0]+_line[2]-1]

            # print("src:", src)
            # print("dst:", dst)

            _map.append([mapName, list(src), list(dst)])


        # if i > 10:
        #     break


    # print(dat)

    return dat

def readFile():

    data = []
    _map = []
    mapName = None

    with open(fileName, "r") as f:
        i=0
        for line in f:
            # data.append(line[:-1])
            line = line[:-1]
            # print(i, line)

            if i == 0:
                data.append([int(x) for x in line.split(": ")[1].split()])
                # print(data)

            elif i != 1 and  "" == line:
                data.append(_map)
                # print(_map[0])
                _map = []
                # print(i, line)

            elif "map:" in line:
                mapName = line.split()[0]

            elif line != "" and "map:" not in line:
                _line = line.split()
                _line = [int(x) for x in _line]
                # print(_line)
                # src = np.arange(_line[1], _line[1]+_line[2])
                # dst = np.arange(_line[0], _line[0]+ _line[2])
                src = [_line[1], _line[1]+_line[2]-1]
                dst = [_line[0], _line[0]+_line[2]-1]

                # print("src:", src)
                # print("dst:", dst)

                _map.append([mapName, list(src), list(dst)])

            i=i+1

        # print(_map)

    return data

def findLocation(seed, data):

    find = seed
    for i, mp in enumerate(data[1:]):
        # if i == 0:
        #     continue

        # print(i, mp)
        for mps in mp:
            # print(mp)
            mName, src, dst = mps
            # print(mName)
            # print("src", src)
            # print("dst", dst)
            # if find in src:
            #     ind = src.index(find)
            #     find = dst[ind]
            if find >= src[0] and find <= src[1]:
                ind = find - src[0]
                find = dst[0] + ind
                break

        # print(i,"now looking for number:", find)

        # break

    return find

def solve(data):

    solution = 0

    seeds = data[0]
    sols = []

    for seed in seeds:
        # print("\nseed:", seed)

        location = findLocation(seed, data)

        # print(seed, "goes to", find)
        sols.append(location)
        # print()
        # break

    # print()
    # print(seeds)
    # print(sols)
    solution = np.min(sols)
    print("solution:", solution)

    return

td = decodeTestData(testData)
# solve(td)

data = readFile()
# print(data)
solve(data)

print()
print("########### PART 2 ###############")
print()

def reverseMappings(data):

    reverse = []
    for tr in data[:]:
        reverse.append([])
        for line in tr:
            # print(line)
            reverse[-1].append([line[0], line[2], line[1]])

    return reverse[::-1]

def seedInStartingPack(seed, allSeeds):

    for seedRange in allSeeds:
        # allSeeds.append([seeds[i], seeds[i]+seeds[i+1]-1])
        if seed <= seedRange[1] and seed >= seedRange[0]:
            return True
    return False

def findLocation2(seed, data):

    find = seed

    ## Loop over all transformations all the way from seed - soil - ... - location
    for mp in data:
        ## loop over each row of mappings in one transformation - ex. seeds-to-soil
        for mName, src, dst in mp:
            if find >= src[0] and find <= src[1]:
                find = find + dst[0] - src[0]
                break

    return find

def solve2(data):

    solution = 99999999999999999

    seeds = data[0]
    allSeeds = []
    i=0
    N = 0
    while i < len(seeds):
        allSeeds.append([seeds[i], seeds[i]+seeds[i+1]-1])
        N = N + seeds[i+1]
        i=i+2

    i=0
    k=0
    for sp in allSeeds[::-1]:
        print(sp)

        if k in []:
            k=k+1
            i = i + allSeeds[k][1] - allSeeds[k][0]
            continue

        for seed in range(sp[0], sp[1]+1):
            # print("\nseed:", seed)

            if i != 0 and i % 500000 == 0:
                print("progress", i, "/", N, "("+str(round(i/N*100, 2))+" %)")
                print("\t",k,"current solution =", solution)

            loc = findLocation2(seed, data[1:])

            # _seeds = np.arange(_seed, min(_seed+n, sp[1])+1, 1)
            # # print(_seeds)
            # locs = findLocation2(_seeds, data)

            # print(seed, "goes to", find)
            # sols.append(location)
            # solution = min(solution, location)
            solution = min(solution, loc)
            # print()
            # break

            # seed = seed + 1
            i=i+1
        k=k+1

    print()
    # print(sols)
    # solution = np.min(sols)
    print("solution:", solution)

    return

def solve2reverse(data):

    reverseMap = reverseMappings(data[1:])
    # return

    # print(data[1:])
    # print()
    # print(reverseMap)

    # return

    seeds = data[0]
    allSeeds = []
    i=0
    N = 0
    while i < len(seeds):
        allSeeds.append([seeds[i], seeds[i]+seeds[i+1]-1])
        N = N + seeds[i+1]
        i=i+2

    maxIter = 1000000000
    i=10000000
    while i < maxIter:

        if i != 0 and i % 1000000 == 0:
            print("progress", i, "/", maxIter, "("+str(round(i/maxIter*100, 2))+" %)")
            print("\tcurrent solution =", i)

        ## We are finding seed from location with the same function as locatin,
        ## but reversing transformation to get the seed
        loc = findLocation2(i, reverseMap)
        if seedInStartingPack(loc, allSeeds) == True:
            print("Solution:", i)
            return

        i=i+1

    print("Didn't find a solution")

t1 = time.time()
# data = readFile()
# solve2(td)
# solve2(data)

# solve2reverse(td)
solve2reverse(data)

dt = time.time() - t1
print("elapsed", round(dt/60, 2), "min")