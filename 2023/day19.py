import numpy as np
import time
import copy as cp

fileName = "day19_data.txt"
testData = ["px{a<2006:qkq,m>2090:A,rfg}","pv{a>1716:R,A}","lnx{m>1548:A,A}","rfg{s<537:gd,x>2440:R,A}","qs{s>3448:A,lnx}",
            "qkq{x<1416:A,crn}","crn{x>2662:A,R}","in{s<1351:px,qqz}","qqz{s>2770:qs,m<1801:hdj,R}","gd{a>3333:R,R}","hdj{m>838:A,pv}",
            "","{x=787,m=2655,a=1222,s=2876}","{x=1679,m=44,a=2067,s=496}","{x=2036,m=264,a=79,s=2244}","{x=2461,m=1339,a=466,s=291}",
            "{x=2127,m=1623,a=2188,s=1013}"]

testSol1 = 19114
testSol2 = 167409079868000

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

def getPartsAndRules(data):

    parts = []
    rules = {}

    firstHalf = True
    for line in data:

        if line == "":
            firstHalf = False
            continue

        if firstHalf == True:
            key, rs = line.split("{")

            values = []
            for x in rs[:-1].split(","):
                if ":" in x:
                    value = x.split(":")
                    condition = [[value[0][0], value[0][1], int(value[0][2:])], value[1]]
                    values.append(condition)
                else:
                    values.append([x])
            rules[key] = values
        else:
            d = {}
            for values in line[1:-1].split(","):
                key, value = values.split("=")
                d[key] = int(value)
            parts.append(d)

    return parts, rules

def processOneCondition(rule, part):

    for condition in rule:
        if len(condition) == 1:
            # print("last condition")
            return condition[0]
        else:
            con, dest = condition
            # print(condition, con, dest)
            if con[1] == "<" and part[con[0]] < con[2]:
                # print(condition, con, dest)
                # print("yay <", dest)
                return dest
            elif con[1] == ">" and part[con[0]] > con[2]:
                # print(condition, con, dest)
                # print("yay >", dest)
                return dest
            # else:
            #     print("sad")

    print("NONE", rule, part)
    return None

def processOnePart(rules, part, maxIter=100):

    curr = "in"

    i=0
    while i < maxIter:

        curr = processOneCondition(rules[curr], part)
        # print("curr", curr)

        if curr in ["R", "A"]:
            return curr

        i=i+1

    print("NONE", part, i)
    return None

def solve(data, maxSteps=100):

    solution = 0

    # print(data)
    # [print("".join(list(x))) for x in [[ str(int(y)) for y in x] for x in data]]

    parts, rules = getPartsAndRules(data)
    # print(parts)
    # [print(str(x).ljust(3), rules[x]) for x in rules.keys()]
    # print()

    for part in parts:

        # print(part)
        res = processOnePart(rules, part, maxSteps)
        # print()

        if res == "A":
            s = 0
            for key in part.keys():
                s = s + part[key]

            solution = solution + s

    print()
    print("solution:", solution)

    return

data = readFile()
solve(testData)
solve(data)

print()
print("########### PART 2 ###############")
print()

def processOne(variants, rules, curr, part, maxValue=4000):

    todo = {}

    # print()
    # print("curr:", curr)
    # print("part:", part)
    # print("rule:", rules[curr])

    for condition in rules[curr]:
        if len(condition) == 1:
            # print("last condition")
            dest = condition[0]
            if dest == "A":
                variants.append(part)
                # print("end A", part)
                # return True, variants
            elif dest == "R":
                # print("end R", part)
                pass
                # return False, variants
            else:
                todo[dest] = part
        else:
            con, dest = condition
            if con[1] == "<":
                # print(condition, con, dest)
                tArray = [max(0, part[con[0]][0]), min(con[2] - 1, part[con[0]][1])]
                fArray = [max(con[2], part[con[0]][0]), max(con[2], part[con[0]][1])]

                nPart1 = cp.deepcopy(part)
                nPart1[con[0]] = tArray
                if dest == "A":
                    variants.append(nPart1)
                    # return True, variants
                    # print("end mid A", nPart1)
                elif dest == "R":
                    # print("end mid R", nPart1)
                    pass
                else:
                    todo[dest] = nPart1

                part[con[0]] = fArray

                # if curr == "rfg":
                #     print("yay", con[0], curr)
                #     print(todo)
                #     print(tArray, fArray)

            elif con[1] == ">":
                # print(condition, con, dest)
                tArray = [max(con[2]+1, part[con[0]][0]), max(con[2], part[con[0]][1])]
                fArray = [max(0, part[con[0]][0]), min(con[2], part[con[0]][1])]

                nPart1 = cp.deepcopy(part)
                nPart1[con[0]] = tArray
                if dest == "A":
                    variants.append(nPart1)
                    # print("end mid A", nPart1)
                    # return True, variants

                elif dest == "R":
                    # print("end mid R", nPart1)
                    pass
                else:
                    todo[dest] = nPart1

                part[con[0]] = fArray

                # print("yolo", tArray, con[0], curr)
                # print(data, tArray, fArray)
                # return
            # else:
            #     print("sad")


    # print(10*"-")
    # print("TODO:")
    # [print(key, todo[key]) for key in todo.keys()]
    # print(10*"-")
    for key in todo.keys():
        succ, variants = processOne(variants, rules, key, todo[key])

    # print(succ, variants)
    return None, variants

def solve2(data, maxSteps=100):

    solution = 0

    # print(data)
    # [print("".join(list(x))) for x in [[ str(int(y)) for y in x] for x in data]]

    parts, rules = getPartsAndRules(data)
    # print(parts)
    # [print(str(x).ljust(3), rules[x]) for x in rules.keys()]
    # print()

    curr = "in"
    variants = []

    part = {'x': [0,4000], 'm': [0,4000], 'a': [0,4000], 's': [0,4000]}
    union = {'x': [], 'm': [], 'a': [], 's': []}

    i=0
    while i < maxSteps:
        res, variants = processOne(variants, rules, curr, part)

        # print(res)
        break

        i=i+1

    print()
    # print(variants)
    print(len(variants))
    for var in variants:
        print(union)
        print("yolo", var)
        for key in var.keys():

            if len(union[key]) == 0:
                union[key].append(var[key])
                print("yay")
            else:

                N = len(union[key])
                add = []
                inside = False
                i=0
                while i < N:

                    if union[key][i][0] <= var[key][1] and union[key][i][1] >= var[key][0]:
                        print(key, i, "inside")#, union[key][i][0] <= var[key][1], union[key][i][1] >= var[key][0])
                        inside = True
                        if union[key][i][0] > var[key][0]:
                            union[key][i][0] = var[key][0]
                        if union[key][i][1] < var[key][1]:
                            union[key][i][1] = var[key][1]
                    else:
                        print(key, i, "outside")
                        add = var[key]

                    i=i+1

                if inside == False and len(add) != 0:
                    union[key].append(add)

    print("after", union)
    s = 1
    for key in union.keys():
        a = 0
        i=0
        while i < len(union[key]):
            a = a + (union[key][i][1] - union[key][i][0])
            i=i+1
        s = s * a
    solution = solution + s

    print()
    print("solution:", solution)

    return

solve2(testData)
print(9*" ", testSol2)

# solve2(data)