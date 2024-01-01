import numpy as np

fileName = "day1_data.txt"
testData = ["1abc2","pqr3stu8vwx","a1b2c3d4e5f","treb7uchet"]
testSolv = [12, 38, 15, 77]
testSolution = 142

testData2 = ["two1nine", "eightwothree", "abcone2threexyz", "xtwone3four", "4nineeightseven2", "zoneight234", "7pqrstsixteen"]
testSolv2 = [29, 83, 13, 24, 42, 14, 76]
testSolution2 = 281

def readFile():

    data = []

    with open(fileName, "r") as f:
        for line in f:
            data.append(line.split()[0])

    return data

def testSolve(data):

    print(data)

    sol = []
    for i, line in enumerate(data):

        s = []
        for char in line:
            try:
                z = int(char)
                s.append(z)
            except ValueError:
                pass

        sol.append(s[0]*10 + s[-1])

    print(sol)
    print(testSolv)

    print(np.sum(sol), np.sum(testSolv))

    return

def solve(data):

    sol = []
    for line in data:
        s = []
        for char in line:
            try:
                z = int(char)
                s.append(z)
            except ValueError:
                pass

        sol.append(s[0]*10 + s[-1])

    print("solution:", np.sum(sol))

    return

testSolve(testData)
data = readFile()
solve(data)

print()
print("########### PART 2 ###############")
print()

def testSolve2(data):

    print(data)
    words = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]

    sol = []
    for i, line in enumerate(data):

        ## Findung actual numbers
        s = []
        for j, char in enumerate(line):
            try:
                z = int(char)
                s.append([j,z])
            except ValueError:
                pass

        ## Finding numbers as words
        wMin = 9999999
        wMax = -9999999
        wordMin = "---"
        wordMax = "---"
        for word in words:
            m = line.find(word)
            M = line.rfind(word)

            if m != -1 and m < wMin:
                wMin = m
                wordMin = word
            if M != -1 and M > wMax:
                wMax = M
                wordMax = word

        if len(s) != 0:
            # sol.append(s[0]*10 + s[-1])
            if s[0][0] > wMin:
                nMin = words.index(wordMin) + 1
            else:
                nMin = s[0][1]

            if s[-1][0] < wMax:
                nMax = words.index(wordMax) + 1
            else:
                nMax = s[-1][1]
        else:
            nMin = words.index(wordMin) + 1
            nMax = words.index(wordMax) + 1

        sol.append(nMin*10 + nMax)

    print(sol)
    print(testSolv2)

    print(np.sum(sol), np.sum(testSolv2))

    return

def solve2(data):

    words = ["one", "two", "three", "four", "five", "six", "seven", "eight", "nine"]

    sol = []
    for i, line in enumerate(data):

        ## Findung actual numbers
        s = []
        for j, char in enumerate(line):
            try:
                z = int(char)
                s.append([j,z])
            except ValueError:
                pass

        ## Finding numbers as words
        wMin = 9999999
        wMax = -9999999
        wordMin = "---"
        wordMax = "---"
        for word in words:
            m = line.find(word)
            M = line.rfind(word)

            if m != -1 and m < wMin:
                wMin = m
                wordMin = word
            if M != -1 and M > wMax:
                wMax = M
                wordMax = word

        if len(s) != 0:
            # sol.append(s[0]*10 + s[-1])
            if s[0][0] > wMin:
                nMin = words.index(wordMin) + 1
            else:
                nMin = s[0][1]

            if s[-1][0] < wMax:
                nMax = words.index(wordMax) + 1
            else:
                nMax = s[-1][1]
        else:
            nMin = words.index(wordMin) + 1
            nMax = words.index(wordMax) + 1

        sol.append(nMin*10 + nMax)

    print("solution 2:", np.sum(sol))

    return

testSolve2(testData2)
solve2(data)