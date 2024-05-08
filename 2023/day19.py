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

def doAllConditions(parts, condition, todo, accepted):

    con, dest = condition
    if con[1] == "<":
        # print(condition, con, dest)
        tArray = [max(0, parts[con[0]][0]), min(con[2] - 1, parts[con[0]][1])]
        fArray = [max(con[2], parts[con[0]][0]), max(con[2], parts[con[0]][1])]

        nPart1 = cp.deepcopy(parts)
        nPart1[con[0]] = tArray
        if dest == "A":
            accepted.append(nPart1)
            # return True, variants
            # print("end mid A", nPart1)
        elif dest == "R":
            # print("end mid R", nPart1)
            pass
        else:
            todo[dest] = nPart1

        parts[con[0]] = fArray

        # if curr == "rfg":
        #     print("yay", con[0], curr)
        #     print(todo)
        #     print(tArray, fArray)

    elif con[1] == ">":
        # print(condition, con, dest)
        tArray = [max(con[2]+1, parts[con[0]][0]), max(con[2], parts[con[0]][1])]
        fArray = [max(0, parts[con[0]][0]), min(con[2], parts[con[0]][1])]

        nPart1 = cp.deepcopy(parts)
        nPart1[con[0]] = tArray
        if dest == "A":
            accepted.append(nPart1)
            # print("end mid A", nPart1)
            # return True, variants

        elif dest == "R":
            # print("end mid R", nPart1)
            pass
        else:
            todo[dest] = nPart1

        parts[con[0]] = fArray

        # print("yolo", tArray, con[0], curr)
        # print(data, tArray, fArray)
        # return
    # else:
    #     print("sad")

    return

def processOne(accepted, rules, curr, parts, maxValue=4000, depth=0):

    todo = {}

    # print()
    # print("depth", depth)
    # print("curr:", curr)
    # print("parts:", parts)
    # print("rule:", rules[curr])
    # print("accepted", accepted)

    for condition in rules[curr]:

        ## there is no conditions in current rule - just go to that destination
        if len(condition) == 1:
            # print("last condition")
            dest = condition[0]
            if dest == "A":
                accepted.append(parts)
                # print("end A", part)
                # return True, variants
            elif dest == "R":
                # print("end R", part)
                pass
                # return False, variants
            else:
                todo[dest] = parts
        else:
            doAllConditions(parts, condition, todo, accepted)

    # print(10*"-")
    # print("accepted:")
    # [print(x) for x in accepted]
    # print("TODO:")
    # [print(key, todo[key]) for key in todo.keys()]
    # print(10*"-")

    if depth < 2 or True:
        for key in todo.keys():
            succ, accepted = processOne(accepted, rules, key, todo[key], depth=depth+1)

    # print(succ, variants)
    return None, accepted

def testBoundaries(parts, accepted):

    print("testing")
    for part in parts:
        # print(part)
        i=0
        for acc in accepted:

            ## check all parameters x,m,a,s
            succ = True
            for key in part.keys():
                if not (part[key] >= acc[key][0] and part[key] <= acc[key][1]):
                    succ = False
                    break

            if succ == True:
                print(part, "Accepted", i)
                break

            i=i+1

    return

def solve2(data, maxSteps=100):

    solution = 0

    # print(data)
    # [print("".join(list(x))) for x in [[ str(int(y)) for y in x] for x in data]]

    parts, rules = getPartsAndRules(data)
    # print(parts)
    # [print(str(x).ljust(3), rules[x]) for x in rules.keys()]
    # print()

    curr = "in"
    arr = []

    part = {'x': [1,4000], 'm': [1,4000], 'a': [1,4000], 's': [1,4000]}

    ## proccesing all borders of part number span
    res, accepted = processOne(arr, rules, curr, part)
    # print("accepted:")
    # [print(x) for x in accepted]

    # testBoundaries(parts, accepted)

    ## There is ABSOLUTELY NO NEED to complicate this
    ## just add the number of each accepted path
    ## because they are already mutualy independant and they DO NOT overlap in any way
    for acc in accepted:
        m = 1
        for cat in acc.keys():
            m = m * (acc[cat][1]-acc[cat][0]+1)
        solution = solution + m

    print()
    print("solution:", solution)

    return

solve2(testData)
# print("test sol:", testSol2)
# print(9*" ", 4000**4)

solve2(data)