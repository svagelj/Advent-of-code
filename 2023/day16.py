import numpy as np
import time
import copy as cp

fileName = "day16_data.txt"
testData = [".|...\....","|.-.\.....",".....|-...","........|.","..........",
            ".........\\","..../.\\\..",".-.-/..|..",".|....-|.\\","..//.|...."]

testSol1 = 46
testSol2 = 51

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
    # print(data[-3:])

    return data

def isDirectionInHistory(directions, direction):

    for _dir in directions:
        if _dir[0] == direction[0] and _dir[1] == direction[1]:
            return True

    return False

def isPositionInHistory(history, position):

    # print(history)
    # print(position)

    for pos in history:
        if pos[0] == position[0] and pos[1] == position[1] and pos[2] == position[2] and pos[3] == position[3]:
            return True

    return False

def doOneRay(data, history, dirHistory, rays, rayID, M,N, deletes):

    prevPos, currPos = rays[rayID][-2], rays[rayID][-1]
    direction = currPos[:2] - prevPos[:2]
    curr = data[currPos[0]][currPos[1]]

    # print()
    # print("curr, dir:", currPos, direction)
    # print("\t", curr)#, int(history[currPos[0]][currPos[1]]))

    ## history stuff
    if history[currPos[0]][currPos[1]] == 0:
        history[currPos[0]][currPos[1]] = 1

    if dirHistory[currPos[0]][currPos[1]] == 0:
        dirHistory[currPos[0]][currPos[1]] = [direction]
    elif isDirectionInHistory(dirHistory[currPos[0]][currPos[1]], direction) == True:
        # print("destroy (direction history)")
        del rays[rayID]
        return deletes + 1

    # if isPositionInHistory(rays[rayID][:-1], currPos) == True:
    #     print("destroy (ray history)")
    #     del rays[rayID]
    #     return deletes + 1

    ## Making the next step
    if curr == ".":
        nextStep = currPos[:2] + direction
        # print("nextStep (.) =", nextStep)

        if nextStep[0] >= N or nextStep[1] >= M or any(nextStep < 0):
            # print("\tdestroy")
            del rays[rayID]
            return deletes + 1
        else:
            rays[rayID].append(np.array([nextStep[0], nextStep[1], direction[0], direction[1]]))

    else:
        if (curr == "-" and direction[0] == 0) or (curr == "|" and direction[1] == 0):
            nextStep = currPos[:2] + direction
            # print("nextStep (splitter 0) =", nextStep)

            if nextStep[0] >= N or nextStep[1] >= M or any(nextStep < 0):
                # print("\tdestroy")
                del rays[rayID]
                return deletes + 1
            else:
                rays[rayID].append(np.array([nextStep[0], nextStep[1], direction[0], direction[1]]))

        elif curr == "/":
            nextStep = currPos[:2] + np.array([-direction[1], -direction[0]])
            # print("nextStep (/) =", nextStep)

            if nextStep[0] >= N or nextStep[1] >= M or any(nextStep < 0):
                # print("\tdestroy (/)")
                del rays[rayID]
                return deletes + 1
            else:
                rays[rayID].append(np.array([nextStep[0], nextStep[1], direction[0], direction[1]]))

        elif curr == "\\":
            nextStep = currPos[:2] + np.array([direction[1], direction[0]])
            # print("nextStep (\\) =", nextStep)

            if nextStep[0] >= N or nextStep[1] >= M or any(nextStep < 0):
                # print("\tdestroy")
                del rays[rayID]
                return deletes + 1
            else:
                rays[rayID].append(np.array([nextStep[0], nextStep[1], direction[0], direction[1]]))

        elif (curr == "-" and direction[0] != 0) or (curr == "|" and direction[1] != 0):
            nextStep1 = currPos[:2] + np.array([direction[1], direction[0]])
            # print("nextStep (splitter 11) =", nextStep1)

            copyRay = cp.deepcopy(rays[rayID])

            old = False
            if nextStep1[0] >= N or nextStep1[1] >= M or any(nextStep1 < 0):
                # print("\tdestroy")
                deletes = deletes + 1
            else:
                # print("\tcontinue new direction")
                rays[rayID].append(np.array([nextStep1[0], nextStep1[1], direction[0], direction[1]]))
                old = True

            nextStep2 = currPos[:2] + np.array([-direction[1], -direction[0]])
            # print("nextStep (splitter 12) =", nextStep2)

            if nextStep2[0] >= N or nextStep2[1] >= M or any(nextStep2 < 0):
                # print("\tdestroy")
                deletes = deletes + 1
            else:
                if old == False:
                    # print("\tcontinue new direction")
                    rays[rayID].append(np.array([nextStep2[0], nextStep2[1], direction[0], direction[1]]))
                    old = True
                else:
                    maxID = np.max(list(rays.keys()))
                    # print("\tcreate new ray", maxID)
                    rays[maxID+1] = copyRay
                    rays[maxID+1].append(np.array([nextStep2[0], nextStep2[1], direction[0], direction[1]]))

    return deletes

def solve(data, maxSteps=100):

    solution = 0

    # [print("".join(list(x))) for x in data]

    N=len(data)
    M=len(data[0])

    history = np.zeros([N,M])
    dirHistory = np.zeros([N,M])
    dirHistory = [list(x) for x in dirHistory]
    rays = {0:[np.array([0,-1, 0,1]), np.array([0,0, 0,1])]}
    deletes = 0

    n = 0
    k=0
    while k < maxSteps:

        # if k % 10 == 0:
        #     print("progress", k, "/", maxSteps, "||", len(rays.keys()), n, solution)

        # print()
        keys = cp.deepcopy(list(rays.keys()))
        for rayID in keys:
            deletes = doOneRay(data, history, dirHistory, rays, rayID, M, N, deletes)

        # print("rays")
        # [print([list(y) for y in rays[x]]) for x in rays.keys()]

        sol = np.count_nonzero(history)

        if sol != solution:
            solution = sol
            n=0
        else:
            n=n+1

        if n > 100 or len(rays.keys()) == 0:
            # print("break", n, M,N, "|", k)
            # print("rays", len(rays.keys()))
            # print("delets", deletes)
            break

        # print()
        # [print("  "+"  ".join(list(x))) for x in data]
        # print()
        # print(history)

        # if k > 50:
        #     break

        k=k+1

    # # print()
    # [print("".join(xx)) for xx in [[str(int(y)) for y in x] for x in history]]

    print()
    print("solution:", solution)

    return

data = readFile()
solve(testData, 1000000)
solve(data, 100000)

print()
print("########### PART 2 ###############")
print()

def solveOneStartingPoint(data, rays, maxSteps=100):

    solution = 0

    # [print("".join(list(x))) for x in data]

    N=len(data)
    M=len(data[0])

    history = np.zeros([N,M])
    dirHistory = np.zeros([N,M])
    dirHistory = [list(x) for x in dirHistory]
    deletes = 0

    n = 0
    k=0
    while k < maxSteps:

        # if k % 10 == 0:
        #     print("progress", k, "/", maxSteps, "||", len(rays.keys()), n, solution)

        # print()
        keys = cp.deepcopy(list(rays.keys()))
        for rayID in keys:
            deletes = doOneRay(data, history, dirHistory, rays, rayID, M, N, deletes)

        # print("rays")
        # [print([list(y) for y in rays[x]]) for x in rays.keys()]

        sol = np.count_nonzero(history)

        if sol != solution:
            solution = sol
            n=0
        else:
            n=n+1

        if n > 100 or len(rays.keys()) == 0:
            # print("break", n, M,N, "|", k)
            # print("rays", len(rays.keys()))
            # print("delets", deletes)
            break

        # print()
        # [print("  "+"  ".join(list(x))) for x in data]
        # print()
        # print(history)

        # if k > 50:
        #     break

        k=k+1

    # print()
    # [print(" ".join(xx)) for xx in [[str(int(y)) for y in x] for x in history]]

    return solution, k

def solve2(data, maxSteps=1000):

    # [print(" ".join(list(x))) for x in data]

    N=len(data)
    M=len(data[0])
    # dirHistory = np.zeros([N,M])
    # dirHistory = [list(x) for x in dirHistory]

    ## start from top/bottom
    bestVert = [0]
    i=0
    while i < M:
        # print()
        rays = {0:[np.array([-1,i, 1,0]), np.array([0,i, 1,0])]}
        sol1, k1 = solveOneStartingPoint(data, rays, maxSteps=maxSteps)

        # print()
        rays = {0:[np.array([N,i, -1,0]), np.array([N-1,i, -1,0])]}
        sol2, k2 = solveOneStartingPoint(data, rays, maxSteps=maxSteps)

        if sol1 > bestVert[0]:
            bestVert = [sol1, i, 1,0, k1]
        if sol2  > bestVert[0]:
            bestVert = [sol2, i, -1,0, k2]

        # print()
        # [print(" ".join(list(x))) for x in data]

        # print(i, sol1, sol2, bestVert)
        # return

        # break
        i=i+1

    print("done vertical")

    ## start from left/right
    bestHor = [0]
    i=0
    while i < N:
        # print()
        rays = {0:[np.array([i,-1, 0,1]), np.array([i,0, 0,1])]}
        sol1, k1 = solveOneStartingPoint(data, rays, maxSteps=maxSteps)

        # print()
        rays = {0:[np.array([i,M, 0,-1]), np.array([i,M-1, 0,-1])]}
        sol2, k2 = solveOneStartingPoint(data, rays, maxSteps=maxSteps)

        if sol1 > bestHor[0]:
            bestHor = [sol1, i, 0,1, k1]
        if sol2  > bestHor[0]:
            bestHor = [sol2, i, 0,-1, k2]

        # print()
        # [print(" ".join(list(x))) for x in data]

        # if i==0:
        # print(i, sol1, sol2, bestHor)

        # break
        i=i+1

    # print()
    # print(bestVert)
    # print(bestHor)
    print("solution:", max(bestHor[0], bestVert[0]))

    return

solve2(testData)

t1 = time.time()
solve2(data, maxSteps=1000000000000)
print("elapsed", round((time.time()-t1), 2), "s")