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

def getOverlappingBlocks(state, block, i):

    Mx = np.max(block[:, 0])
    mx = np.min(block[:, 0])
    My = np.max(block[:, 1])
    my = np.min(block[:, 1])
    mz = np.min(block[:, 2])

    inds = []
    for j,block2 in enumerate(state):

        if j==i:
            continue

        elif abs(block2[0, 0] - Mx) < 15 and abs(block2[0, 1] - My) < 15:
            Mx2 = np.max(block2[:, 0])
            mx2 = np.min(block2[:, 0])
            My2 = np.max(block2[:, 1])
            my2 = np.min(block2[:, 1])
            Mz2 = np.max(block2[:,2])

            zCon = Mz2 < mz
            xCon = mx <= Mx2 and mx2 <= Mx
            yCon = my <= My2 and my2 <= My

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

                    # inds2 = getOverlappingBlocks(stacked, newBlock, i)
                    # print("\t", inds2)

                    stacked.append(newBlock)
                    stacked2[i] = {"loc":cp.deepcopy(block), "bellow":None}


        # if i > 12:
        #     break

    # print(10*"-")

    # [print(key, stacked2[key]) for key in stacked2.keys()]
    stacked = np.array(stacked)
    return stacked[stacked[:,0,2].argsort()]

def getTouchingBlocks(state, block, i):

    Mx = np.max(block[:, 0])
    mx = np.min(block[:, 0])
    My = np.max(block[:, 1])
    my = np.min(block[:, 1])
    Mz = np.max(block[:, 2])
    mz = np.min(block[:, 2])

    indsAbove, indsBellow = [], []
    for j,block2 in enumerate(state):

        if i == j:
            continue

        elif abs(block2[0, 2] - Mz) < 15:# or mz - np.max(block2[:, 2]) == 1:

            Mx2 = np.max(block2[:, 0])
            mx2 = np.min(block2[:, 0])
            My2 = np.max(block2[:, 1])
            my2 = np.min(block2[:, 1])
            Mz2 = np.max(block2[:, 2])
            mz2 = np.min(block2[:, 2])

            zConAbove = mz2 - Mz == 1
            zConBellow = mz - Mz2 == 1
            xCon = mx <= Mx2 and mx2 <= Mx
            yCon = my <= My2 and my2 <= My

            if zConAbove == True and xCon == True and yCon == True:
                indsAbove.append(j)
            if zConBellow == True and xCon == True and yCon == True:
                indsBellow.append(j)

    return indsAbove, indsBellow

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

    # print("min", [xMin, yMin, zMin])
    # print("max", [xMax, yMax, zMax])

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
    state = initializeData(data)[:]
    # print(state)
    # plotDebug(state)
    # plotDebugg12(state, "raw")

    print("Stacking")
    stacked = stackBlocks(state)
    # print(stacked)
    # print(len(state), len(stacked))
    # plotDebug(stacked)
    # plotDebugg12(stacked, "Stacked")

    # return
    stacked2 = {}

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

        stacked2[i] = {"loc":block, "bellow":indsB, "above":indsA}

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
                if len(indsB2) > 0:
                   good = good + 1
                else:
                    ## One of the blocks (block2) above is causing the problem, so this block (block) cannot be removed
                    good = -1
                    break

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
    # print(stacked2)
    # [print(key, stacked2[key]) for key in stacked2.keys()]

    # plotDebugg2(gud, bad)

    print("solution:", solution)

    return

data = readFile()
solve(testData1)

# t1 = time.time()
# solve(data)
# print("elapsed", round(time.time()-t1, 2), "s")

print()
print("########### PART 2 ###############")
print()

def isBlockSupported(stacked, ind, collapsed):

    # print(ind, "collapsed", collapsed)

    block2 = stacked[ind]
    indsA2, indsB2 = getTouchingBlocks(stacked, block2, ind)

    # print(20*"-", i)
    # print("\tabove", indsA, indsA2)
    # print("\tb bellow("+str(ind)+")", indsB2)

    for x in cp.deepcopy(indsB2):
        # print("yay", x)
        if x in collapsed:
            # print("yay2", x)
            indsB2.remove(x)

    # print("\ta bellow("+str(ind)+")", indsB2)

    if len(indsB2) == 0:
        return False
    else:
        return True

def collapseOneBlock(stacked, i, indsA, collapsed, memo):
    ## NOT WORKING

    # print(i, "collapsed", collapsed)

    queue = []

    for ind in indsA:

        # print("\ta", indsB2)
        # print("\t"+10*"-")

        if isBlockSupported(stacked, ind, collapsed) == False:
            ## block2 has no other support
            # print("\t", i, "is bad block because", ind)
            # print(stacked)

            collapsed.add(ind)
            # collapsed = collapseOneBlock(stacked, ind, indsA2, collapsed, memo)
            queue.append(ind)
        #     print("Not supported")
        # else:
        #     print("supported")

    # print("\tqueue:", queue, collapsed, memo)
    for ind in queue:
        if ind not in memo.keys():
            # print("\t", "calculate", ind)
            indsA2, indsB2 = getTouchingBlocks(stacked, stacked[ind], ind)
            collapsed = collapseOneBlock(stacked, ind, indsA2, collapsed, memo)
        else:
            # print("\t", "memo", ind)
            collapsed.update(memo[ind])


    # if i in collapsed:
    #     collapsed.remove(i)
    # memo[i] = collapsed
    return collapsed

def procesOne(stacked, j, collapsed, queue):

    indsA, indsB = getTouchingBlocks(stacked, stacked[j], j)
    # print("\t", j, indsA)

    for ind in indsA:
        # print("\t",j, ind,"calc")
        # print("\ta", indsB2)
        # print("\t"+10*"-")

        if isBlockSupported(stacked, ind, collapsed) == False:
            ## block2 has no other support
            # print("\t", j, "is bad block because", ind)
            # print(stacked)
            # print("\t\t", ind, "bad")

            collapsed.add(ind)
            if ind not in queue:
                queue.append(ind)

    return collapsed, queue

def collapseOneBlock2(stacked, i, memo):

    # print(i, "collapsed", collapsed)

    queue = [i]
    collapsed = {i}

    for j in queue:

        ## TODO fix -> MEMO IS NOT WORKING
        if True and j in memo:
            collapsed.update(memo[j])
            # collapsed = collapsed.union(memo[j])
            # print("memo:", j, memo[j], collapsed)
            continue

        collapsed, queue = procesOne(stacked, j, collapsed, queue)


    collapsed.remove(i)
    return collapsed

def getNumberOfCollapsingBlocks(stackedInfo):

    n = 0

    for key in stackedInfo.keys():
        n = n + stackedInfo[key]["nCollapsing"]

    return n

def solve2(data):

    solution = 0

    print("Initialization")
    state = initializeData(data)[:]
    # plotDebug(state)
    # plotDebugg12(state, "raw")

    print("Stacking")
    stacked = stackBlocks(state)
    # plotDebug(stacked)
    # plotDebugg12(stacked, "Stacked")

    info = {}

    N = len(stacked)
    print("Checking each block:", N)
    gud, bad = [], []

    ## Go from the top to bottom so that memo makes sense
    memo = {}
    i = N - 1
    while i >= 0:
        block = stacked[i]

        if i % 10 == 0:
            print("progress", str(i).rjust(4), N, "|", solution)

        # print()
        # print(i, [list(x) for x in block])
        indsA, indsB = getTouchingBlocks(stacked, block, i)

        ## count how many block would collapse if current one is removed - one by one
        ## Use memorized data first to avoid unnecessary recursion0
        if len(indsA) == 0:
            gud.append(block)
        else:

            # print("\tb memo", memo)
            # print("\tb collapsed", set([]))
            collapsed = collapseOneBlock2(cp.deepcopy(stacked), i, memo)
            # print("\ta collapsed", collapsed)
            if len(collapsed) != 0:
                memo[i] = collapsed
                # print("update memo:", i, memo[i])
            # print("\ta memo", memo)
            # print("\ta coll", collapsed)

            solution = solution + len(collapsed)

        # break
        # if i > 0:
        #     break
        i=i-1

    # print(stacked2)
    [print(key, memo[key]) for key in list(memo.keys())[-30:]]

    print("solution:", solution)

    return

data = readFile()
solve2(testData1)
print()

t1 = time.time()
solve2(data)
print("elapsed", round(time.time()-t1, 2), "s")