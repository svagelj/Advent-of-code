import numpy as np

fileName = "day2_data.txt"
testData = ["Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
            "Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
            "Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
            "Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
            "Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green"]

## 12 red cubes, 13 green cubes, and 14 blue cubes
myColors = [12, 13, 14]

def readFile():

    data = []

    with open(fileName, "r") as f:
        for line in f:
            # data.append(line)
            _line = line[:-1]

            game = "".join(_line.split(":")[1:])[1:]
            # print(game)
            sets = game.split("; ")
            # print(sets)

            data.append([])
            for s in sets:
                nRed, nGreen, nBlue = 0, 0, 0
                # print("\t",s)
                colors = s.split(", ")
                # print(2*"\t", colors)
                for color in colors:
                    # print(3*"\t", color.split())
                    n = int(color.split()[0])
                    c = color.split()[1]

                    if c == "red":
                        nRed = nRed + n
                    if c == "green":
                        nGreen = nGreen + n
                    if c == "blue":
                        nBlue = nBlue + n
                data[-1].append([nRed, nGreen, nBlue])

    # print(data[:2])

    return data

def solve(data):

    solution = 0

    for i, game in enumerate(data):

        # print(i, game, type(i))

        possible = True
        for s in game:
            # print("\t", s)

            for j,n in enumerate(s):
                if n > myColors[j]:
                    possible = False
                    break

        if possible == True:
            solution = solution + i+1

        # if i>0:
        #     break

    print("Solution =", solution)

    return

data = readFile()
solve(data)

print()
print("########### PART 2 ###############")
print()

def decodeTestData(testData):

    data = []

    for line in testData:
        # print(line)
        _line = line

        game = "".join(_line.split(":")[1:])[1:]
        # print(game)
        sets = game.split("; ")
        # print(sets)

        data.append([])
        for s in sets:
            nRed, nGreen, nBlue = 0, 0, 0
            # print("\t",s)
            colors = s.split(", ")
            # print(2*"\t", colors)
            for color in colors:
                # print(3*"\t", color.split())
                n = int(color.split()[0])
                c = color.split()[1]

                if c == "red":
                    nRed = nRed + n
                if c == "green":
                    nGreen = nGreen + n
                if c == "blue":
                    nBlue = nBlue + n
            data[-1].append([nRed, nGreen, nBlue])

    return data

def solve2(data):

    solution = 0

    for i, game in enumerate(data):

        # print()
        # print(i, game)

        maxs = [0, 0, 0]
        for s in game:
            # print("\t", s)

            maxs[0] = max(maxs[0], s[0])
            maxs[1] = max(maxs[1], s[1])
            maxs[2] = max(maxs[2], s[2])

        power = 1
        for v in maxs:
            power = power*v
        # print(maxs, power)
        solution = solution + power


    # print()
    print("Solution =", solution)

    return

td = decodeTestData(testData)
solve2(data)