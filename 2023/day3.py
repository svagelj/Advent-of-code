import numpy as np

fileName = "day3_data.txt"
testData = ["467..114..","...*......","..35..633.","......#...","617*......",".....+.58.","..592.....","......755.","...$.*....",".664.598.."]
testResult = 4361

notSymbols = ["1", "2", "3", "4", "5", "6", "7", "8", "9", "0", "."]
digits = ["1", "2", "3", "4", "5", "6", "7", "8", "9", "0"]

def readFile():

    data = []

    with open(fileName, "r") as f:
        i=0
        for line in f:
            data.append(line[:-1])
            # _line = line[:-1]

            # if i > 0:
            #     break
            # i = i+1

    # print(data[:])

    return data

def findCloseNumbers(data, i,j):

    # if i == 0:
    #     pass
    # elif i == len(data):
    #     pass
    # else:
    #     print(data[i-1])
    #     print(data[i])
    #     print(data[i+1])

    #     print()
    #     print(i,j, data[i][j])

    #     if j == 0:
    #         pass
    #     if j == len(data[i]):
    #         pass
    #     else:
    #         neighbours = [data[i][j-1], data[i][j+1],
    #                       data[i-1][j], data[i+1][j],
    #                       data[i+1][j+1], data[i+1][j-1],
    #                       data[i-1][j+1], data[i-1][j-1]]

    #         print(neighbours)

    numbers = []
    n = -1
    while n <= 1:
        if i+n >= len(data) or i-n < 0:
            n=n+1
            continue

        m=-1
        while m <= 1:
            if (n==0 and m==0) or j+m >= len(data[i]) or j-m < 0:
                m=m+1
                continue

            # print("sad", i,j, n,m, data[i][j])
            if data[i+n][j+m] in digits:
                # print(n,m, data[i+n][j+m])

                k = 0
                while j+m+k < len(data[i]) and data[i+n][j+m+k] in digits:
                    # print("yolo", k, data[i+n][j+m+k])
                    k=k+1
                kMax = k

                k = 1
                while j+m-k >= 0 and data[i+n][j+m-k] in digits:
                    # print("yol2", k, data[i+n][j+m-k])
                    k=k+1
                kMin = k-1

                jMin, jMax = j+m-kMin, j+m+kMax

                # print(kMin, kMax, jMin, jMax)
                number = data[i+n][jMin:jMax]
                numbers.append(int(number))

                # print("before", data[i+n])
                data[i+n] = list(data[i+n])
                data[i+n][jMin:jMax] = len(number)*"."
                data[i+n] = "".join(data[i+n])
                # print("after", data[i+n])

                # print("yay", number)
                # print()
                # return

            m=m+1
        n=n+1

    return numbers

def solve(data):

    solution = 0

    numbers = []
    for i,line in enumerate(data):
        # print(line)
        for j,char in enumerate(line):
            if char not in notSymbols:
                nums = findCloseNumbers(data, i,j)
                numbers = numbers + nums
                # print(numbers)
                # return
    solution = np.sum(numbers)

    print()
    print("Solution =", solution)

    return

data = readFile()
# solve(testData)
solve(data)

print()
print("########### PART 2 ###############")
print()

def solve2(data):

    solution = 0

    numbers = []
    for i,line in enumerate(data):
        # print(line)
        for j,char in enumerate(line):
            if char == "*":
                nums = findCloseNumbers(data, i,j)
                if len(nums) == 2:
                    # print(nums, nums[0]*nums[1])
                    numbers = numbers + [nums[0]*nums[1]]
                # print(numbers)
                # return
    solution = np.sum(numbers)

    print()
    print("Solution =", solution)

    return

data = readFile()
# solve2(testData)
solve2(data)