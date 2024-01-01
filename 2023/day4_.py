import numpy as np

fileName = "4_data.txt"

testData = ["Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
            "Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
            "Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
            "Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
            "Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
            "Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"]

def readFile(fileName):

    data = []

    with open(fileName, "r") as f:

        for line in f:
            # print(line[:-1])
            data.append(line[:-1])

            # break

    return data

def solveTestData(data):

    solution = 0

    sol = []
    for i, line in enumerate(data):

        d = line.split(": ")[-1]
        wins, gots = d.split(" | ")
        wins = wins.split()
        gots = gots.split()

        wins = [int(x) for x in wins]
        gots = [int(x) for x in gots]

        print()
        print(i+1)
        print(wins)
        print(gots)

        n = 0
        for x in wins:
            if x in gots:
                n=n+1

        if n > 0:
            w = 2**(n-1)
        else:
            w = 0
        print(n,w)
        sol.append(w)
        # break

    print()
    solution = np.sum(sol)
    print("Solution:", solution)

    return

def solve(data):

    solution = 0

    sol = []
    for i, line in enumerate(data):

        d = line.split(": ")[-1]
        wins, gots = d.split(" | ")
        wins = wins.split()
        gots = gots.split()

        wins = [int(x) for x in wins]
        gots = [int(x) for x in gots]

        n = 0
        for x in wins:
            if x in gots:
                n=n+1

        if n > 0:
            w = 2**(n-1)
        else:
            w = 0
        # print(n,w)
        sol.append(w)
        # break

    solution = np.sum(sol)
    print("Solution:", solution)

    return


data = readFile(fileName)
# solveTestData(testData)
solve(data)

print()
print("$$$$$$$$$$$$$$$ Part II $$$$$$$$$$$$$$$$$$$")
print()

def solve2(data):

    solution = 0

    ns = np.ones(len(data))

    for i, line in enumerate(data):

        d = line.split(": ")[-1]
        wins, gots = d.split(" | ")
        wins = wins.split()
        gots = gots.split()

        wins = [int(x) for x in wins]
        gots = [int(x) for x in gots]

        n = 0
        for x in wins:
            if x in gots:
                n=n+1

        while n > 0:
            m = ns[i]
            while m > 0:
                ns[i+n] = ns[i+n] + 1
                m=m-1
            n=n-1

        # print(n, ns)

        # break

    solution = np.sum(ns)
    print("Solution:", solution)

    return

solve2(testData)
solve2(data)