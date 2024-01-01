import numpy as np
import time
import copy as cp

fileName = "day22_data.txt"
testData1 = ["0,0,2~2,0,2","0,2,3~2,2,3","0,0,4~0,2,4",
             "2,0,5~2,2,5","0,1,6~2,1,6","1,1,8~1,1,9", "1,0,1~1,2,1"]
testSol1 = 5

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

def initializeData(data):

    _data = []

    for line in data:
        _line = line.split("~")
        # print(_line)

        p1 = [int(x) for x in _line[0].split(",")]
        p2 = [int(x) for x in _line[1].split(",")]

        # print(p1,p2)
        if p1[2] > p2[2]:
            _data.append([p2,p1])
        else:
            _data.append([p1,p2])

    _data = np.array(_data)
    return _data[_data[:,0,2].argsort()]

def getMaxLenght(state):

    maxSize = 0

    for brick in state:

        d = np.max(np.abs(brick[1] - brick[0])) + 1
        maxSize = max(d, maxSize)

    return maxSize

def getOverlappingBlocks(state, block, i):

    Mx = max(block[0,0], block[1,0])
    mx = min(block[0,0], block[1,0])
    My = max(block[0,1], block[1,1])
    my = min(block[0,1], block[1,1])
    # Mz = max(block[0,2], block[1,2])
    mz = min(block[0,2], block[1,2])

    inds = []
    for j,block2 in enumerate(state):

        Mx2 = max(block2[0,0], block2[1,0])
        mx2 = min(block2[0,0], block2[1,0])
        My2 = max(block2[0,1], block2[1,1])
        my2 = min(block2[0,1], block2[1,1])

        zCon = max(block2[0,2], block2[1,2]) < mz
        xCon = (Mx2 <= Mx and Mx2 >= mx) or (mx2 <= Mx and mx2 >= mx) or (Mx <= Mx2 and Mx >= mx2) or (mx <= Mx2 and mx >= mx2)
        yCon = (My2 <= My and My2 >= my) or (my2 <= My and my2 >= my) or (My <= My2 and My >= my2) or (my <= My2 and my >= my2)

        if zCon == True and xCon == True and yCon == True:
            inds.append(j)

    return inds

def stackBlocks(state):

    stacked = []
    stacked2 = {}

    for i,block in enumerate(state):

        # print()
        # print("block", i, [list(x) for x in block])

        if i == 0:
            if block[0,2] or block[1,2] == 1:
                stacked.append(block)
            else:
                newBlock = cp.deepcopy(block)
                mz = min(block[0,2], block[1,2])
                newBlock[0,2] = newBlock[0,2] - mz +1
                newBlock[1,2] = newBlock[1,2] - mz +1

                stacked.append(newBlock)
                stacked2[i] = {"loc":newBlock, "bellow":0}

        else:

            ## Check if this block can be dropped and by how much
            ## i.e. height distance to closest overlaping other block or ground
            if any(block[:,2] == 1):
                ## This block is on the ground
                stacked.append(cp.deepcopy(block))
                stacked2[i] = {"loc":cp.deepcopy(block), "bellow":0}

            else:
                ## This block is in the air

                ## Get indexes of block with same xy but different z
                inds = getOverlappingBlocks(stacked, block, i)
                # print("inds", inds)

                ## This is height of the block - the max amount to drop if there is no blocks bellow it
                mz = min(block[0,2], block[1,2])

                minDy = mz - 1
                for j in inds:
                    if i == j:
                        continue
                    else:
                        block2 = stacked[j]
                        Mz2 = max(block2[0,2], block2[1,2])
                        # mz2 = min(block2[0,2], block2[1,2])

                        # print(Mz, mz2)

                        diff = mz - Mz2 - 1
                        minDy = min(minDy, diff)

                # print("max drop", minDy, minDy == Mz)
                if minDy == 0:
                    stacked.append(cp.deepcopy(block))
                    stacked2[i] = {"loc":cp.deepcopy(block), "bellow":0}
                else:
                    newBlock = cp.deepcopy(block)
                    mz = min(block[0,2], block[1,2])
                    newBlock[0,2] = newBlock[0,2] - minDy
                    newBlock[1,2] = newBlock[1,2] - minDy

                    stacked.append(newBlock)
                    stacked2[i] = {"loc":cp.deepcopy(block), "bellow":"TODO"}


        # if i > 12:
        #     break

    # print(10*"-")

    stacked = np.array(stacked)
    return stacked[stacked[:,0,2].argsort()]

def getTouchingBlocks(state, block, i):

    Mx = max(block[0,0], block[1,0])
    mx = min(block[0,0], block[1,0])
    My = max(block[0,1], block[1,1])
    my = min(block[0,1], block[1,1])
    Mz = max(block[0,2], block[1,2])
    mz = min(block[0,2], block[1,2])

    indsAbove, indsBellow = [], []
    for j,block2 in enumerate(state):

        if i == j:
            continue

        Mx2 = max(block2[0,0], block2[1,0])
        mx2 = min(block2[0,0], block2[1,0])
        My2 = max(block2[0,1], block2[1,1])
        my2 = min(block2[0,1], block2[1,1])
        Mz2 = max(block2[0,2], block2[1,2])
        mz2 = min(block2[0,2], block2[1,2])

        zConAbove = mz2 - Mz in [1,0]
        zConBellow = mz - Mz2 in [1,0]
        xCon = (Mx2 <= Mx and Mx2 >= mx) or (mx2 <= Mx and mx2 >= mx) or (Mx <= Mx2 and Mx >= mx2) or (mx <= Mx2 and mx >= mx2)
        yCon = (My2 <= My and My2 >= my) or (my2 <= My and my2 >= my) or (My <= My2 and My >= my2) or (my <= My2 and my >= my2)

        if zConAbove == True and xCon == True and yCon == True:
            indsAbove.append(j)
        if zConBellow == True and xCon == True and yCon == True:
            indsBellow.append(j)

        # if abs(mz2 - Mz) > 2 and abs(Mz2 - mz) > 2:
        #     if (zConAbove == True or zConBellow == True):
        #         print("sad", [mz, Mz], [mz2, Mz2])
        #         print(i, [list(x) for x in block])
        #         print(j, [list(x) for x in block2])
        #         print(zConAbove == True, zConBellow == True)
        #         raise Exception("TNEhi")
        #     break

    return indsAbove, indsBellow

def getTouchingBlocksAbove(state, block, i, N):

    Mx = max(block[0,0], block[1,0])
    mx = min(block[0,0], block[1,0])
    My = max(block[0,1], block[1,1])
    my = min(block[0,1], block[1,1])
    Mz = max(block[0,2], block[1,2])
    mz = min(block[0,2], block[1,2])

    ## This is no good. Looking only ahead int the list is not the same as above

    inds = []
    j = i + 1
    while True:
        print(j)

        # if j >= N-1:
        #     break

        block2 = state[j]
        Mx2 = max(block2[0,0], block2[1,0])
        mx2 = min(block2[0,0], block2[1,0])
        My2 = max(block2[0,1], block2[1,1])
        my2 = min(block2[0,1], block2[1,1])
        Mz2 = max(block2[0,2], block2[1,2])
        mz2 = min(block2[0,2], block2[1,2])

        zConAbove = mz2 - Mz == 1
        xCon = (Mx2 <= Mx and Mx2 >= mx) or (mx2 <= Mx and mx2 >= mx) or (Mx <= Mx2 and Mx >= mx2) or (mx <= Mx2 and mx >= mx2)
        yCon = (My2 <= My and My2 >= my) or (my2 <= My and my2 >= my) or (My <= My2 and My >= my2) or (my <= My2 and my >= my2)

        if zConAbove == True and xCon == True and yCon == True:
            inds.append(j)

        if abs(mz2 - Mz) > 1 and abs(Mz2 - mz) > 1:
            print(i, [list(x) for x in block])
            print(j, [list(x) for x in block2])
            print(zConAbove == True)

            if zConAbove == True:
                raise Exception("TNEhi above")
            break

        j=j+1

    return inds

def getTouchingBlocksBellow(state, block, i, N):

    Mx = max(block[0,0], block[1,0])
    mx = min(block[0,0], block[1,0])
    My = max(block[0,1], block[1,1])
    my = min(block[0,1], block[1,1])
    Mz = max(block[0,2], block[1,2])
    mz = min(block[0,2], block[1,2])

    ## This is no good. Looking only behind int the list is not the same as bellow

    inds = []
    j = i + 1
    while True:

        # if j >= N-1:
        #     break

        block2 = state[j]
        Mx2 = max(block2[0,0], block2[1,0])
        mx2 = min(block2[0,0], block2[1,0])
        My2 = max(block2[0,1], block2[1,1])
        my2 = min(block2[0,1], block2[1,1])
        Mz2 = max(block2[0,2], block2[1,2])
        mz2 = min(block2[0,2], block2[1,2])

        zConBellow = mz - Mz2 == 1
        xCon = (Mx2 <= Mx and Mx2 >= mx) or (mx2 <= Mx and mx2 >= mx) or (Mx <= Mx2 and Mx >= mx2) or (mx <= Mx2 and mx >= mx2)
        yCon = (My2 <= My and My2 >= my) or (my2 <= My and my2 >= my) or (My <= My2 and My >= my2) or (my <= My2 and my >= my2)

        if zConBellow == True and xCon == True and yCon == True:
            inds.append(j)

        if abs(mz2 - Mz) > 1 and abs(Mz2 - mz) > 1:
            print(i, [list(x) for x in block])
            print(j, [list(x) for x in block2])
            print(zConBellow == True)

            if zConBellow == True:
                raise Exception("TNEhi bellow")
            break

        j=j+1

    return inds

def plotDebug(state):

    import matplotlib.pyplot as plt

    fig = plt.figure()
    ax = fig.add_subplot(111, projection='3d')

    for block in state:
        xs = [block[0,0], block[1,0]]
        ys = [block[0,1], block[1,1]]
        zs = [block[0,2], block[1,2]]
        ax.plot(xs, ys,zs=zs)

    ax.set_xlabel("x")
    ax.set_ylabel("y")
    ax.set_zlabel("z")

    return

def plotDebugg2(good, bad):

    import matplotlib.pyplot as plt

    fig = plt.figure()
    ax = fig.add_subplot(111, projection='3d')

    for block in good:
        xs = [block[0,0], block[1,0]]
        ys = [block[0,1], block[1,1]]
        zs = [block[0,2], block[1,2]]
        ax.plot(xs, ys,zs=zs, c="g")

    for block in bad:
        xs = [block[0,0], block[1,0]]
        ys = [block[0,1], block[1,1]]
        zs = [block[0,2], block[1,2]]
        ax.plot(xs, ys,zs=zs, c="r")

    ax.set_xlabel("x")
    ax.set_ylabel("y")
    ax.set_zlabel("z")

    return

def plotDebugg12(state, title=""):

    import matplotlib.pyplot as plt

    xMin = np.min(state[:, :,0])
    xMax = np.max(state[:, :,0])
    yMin = np.min(state[:, :,1])
    yMax = np.max(state[:, :,1])
    zMin = np.min(state[:, :,2])
    zMax = np.max(state[:, :,2])

    print("min", [xMin, yMin, zMin])
    print("max", [xMax, yMax, zMax])

    xz = np.zeros((zMax-zMin+1, xMax-xMin+1))*np.nan
    yz = np.zeros((zMax-zMin+1, yMax-yMin+1))*np.nan

    n = 1
    for block in state:
        # print(block)

        x0, x1 = (block[0, 0], block[1,0])
        y0, y1 = (block[0, 1], block[1,1])
        z0, z1 = (block[0, 2], block[1,2])

        xz[z0-1:z1,x0:x1+1] = n
        yz[z0-1:z1,y0:y1+1] = n

        n=n+1

    plt.figure()

    plt.subplot(121)
    plt.imshow(xz, origin="lower")
    plt.xlabel("x")
    plt.ylabel("z")

    plt.subplot(122)
    plt.imshow(yz, origin="lower")
    plt.xlabel("y")
    plt.ylabel("z")

    plt.suptitle(title)

    return

def solve(data):

    solution = 0

    # print(data)
    # [print("".join(list(x))) for x in [[ y for y in x] for x in data]]

    print("Initialization")
    state = initializeData(data)[:100]
    # print(state)
    # plotDebug(state)
    plotDebugg12(state, "raw")

    print("Stacking")
    stacked = stackBlocks(state)
    # print(stacked)
    # print(len(state), len(stacked))
    # plotDebug(stacked)
    plotDebugg12(stacked, "Stacked")

    return

    print("Checking each block")
    N = len(state)
    gud, bad = [], []
    for i,block in enumerate(stacked):

        # print()
        # print(i, [list(x) for x in block])
        indsA, indsB = getTouchingBlocks(stacked, block, i)
        # print(indsA, indsB)
        # indsA2 = getTouchingBlocksAbove(state, block, i, N)
        # indsB2 = getTouchingBlocksBellow(state, block, i, N)
        # print(indsA2, indsB2)

        if len(indsA) == 0:
            gud.append(block)
            solution = solution + 1
        else:
            good = 0
            for ind in indsA:
                block2 = stacked[ind]
                indsA2, indsB2 = getTouchingBlocks(stacked, block2, ind)
                # print("\t", ind, indsB2)
                if i in indsB2:
                    indsB2.remove(i)
                if len(indsB2) != 0:
                   good = good + 1

            if good > 0:
                gud.append(block)
                solution = solution + 1
            else:
                bad.append(block)

        # print("inds:", indsA, indsB)

        # break
        # if i > 0:
        #     break
    print("len", len(stacked), len(gud)+len(bad))

    # plotDebugg2(gud, bad)

    print("solution:", solution)

    return

data = readFile()
solve(testData1)
# solve(data)

print()
print("########### PART 2 ###############")
print()

