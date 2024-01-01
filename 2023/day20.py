import numpy as np
import time
import copy as cp

fileName = "day20_data.txt"
testData11 = ["broadcaster -> a, b, c","%a -> b","%b -> c","%c -> inv","&inv -> a"]
testData12 = ["broadcaster -> a","%a -> inv, con","&inv -> b","%b -> con","&con -> output"]
testSol1 = 32000000
testSol2 = 11687500

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

    init = {}

    for line in data:

        _line = line.split(" -> ")

        _name, output = _line
        output = output.split(", ")
        # print(_name, output)

        if _name[0] == "%":
            name = _name[1:]
            init[name] = {"output":output, "kind":_name[0]}
            init[name]["state"] = 0
            # init[name]["last"] = 0
        elif _name[0] == "&":
            name = _name[1:]
            init[name] = {"output":output, "kind":_name[0]}
            # init[name]["input"] = []
            init[name]["last"] = None
        else:
            init[_name] = {"output":output}

    # print("after")
    for key in init.keys():
        if key != "broadcaster" and init[key]["kind"] == "&":
            # print(key, init[key])

            for key2 in init.keys():
                # print("\t", key2, init[key2]["output"])
                for out in init[key2]["output"]:
                    if key2 != "broadcaster" and out == key:
                        # print("yay", key2)
                        if "memory" not in init[key].keys():
                            init[key]["memory"] = {key2:0}
                        else:
                            init[key]["memory"][key2] = 0

    return init

def sendPulse(state, origin, pulse):

    if origin == "broadcaster":
        # newOutput = [[to, pulse]]
        newOutput = []
        for output in state[origin]["output"]:
            newOutput.append([output, 0])

        # print("yolo", newOutput)
        return newOutput

    elif state[origin]["kind"] == "%":
        if pulse == 0:
            print("\tb %", state[origin])
            state[origin]["state"] = (state[origin]["state"] + 1)%2
            print("\ta %", state[origin])

            newOutput = []
            for output in state[origin]["output"]:
                newOutput.append([output, state[origin]["state"]])

                if state[output]["kind"] == "&" and origin in state[output]["memory"].keys():
                    print("yay", origin, output, pulse)
                    state[output]["last"] = [origin, pulse]
                    print("yay", state[output]["last"])

            return newOutput

        else:
            for output in state[origin]["output"]:
                if state[output]["kind"] == "&" and origin in state[output]["memory"].keys():
                    print("yay2", origin, output, pulse)
                    state[output]["last"] = [origin, pulse]
                    print("yay2", state[output]["last"])

    elif state[origin]["kind"] == "&":
        # print("\tb &", origin, to, state[origin]["memory"])

        last = state[origin]["last"]
        state[origin]["memory"][last[0]] = last[1]
        # print("\ta &", origin, to, state[origin]["memory"])

        for loc in state[origin]["memory"].keys():
            if state[origin]["memory"][loc] == 0:
                newOutput = []
                for output in state[origin]["output"]:
                    newOutput.append([output, 1])
                return newOutput

        newOutput = []
        for output in state[origin]["output"]:
            newOutput.append([output, 0])
        return newOutput

    return None

def isStateSame(state1, state2):



    return True

def solve(data, nButton=1000):

    solution = 0

    state = initializeData(data)
    [print(key, state[key]) for key in state.keys()]

    # print(data)
    # [print("".join(list(x))) for x in [[ str(int(y)) for y in x] for x in data]]

    nHigh = 0
    nLow = 0

    i=0
    while i < nButton:
        # nLow = nLow + 1
        # print("LOW 0")

        curr = [["broadcaster", 0]]

        k = 0
        while len(curr) != 0:
            print()
            print(50*"-")
            print(curr)

            newOutputs = []

            for pulse in curr:
                print("pulse", pulse)

                if pulse[0] not in state.keys():
                    print("DEAD END")
                    continue

                print("\t", pulse[0], state[pulse[0]])

                if pulse[1] == 0:
                    nLow = nLow + 1
                    print("LOW 1:")#, len(state[pulse[0]]["output"]))
                elif pulse[1] == 1:
                    nHigh = nHigh + 1
                    print("HIGH 1:")#, len(state[pulse[0]]["output"]))

                targets = sendPulse(state, pulse[0], pulse[1])
                if targets != None:
                    newOutputs = newOutputs + targets
                    print("\t", "targets", targets)
                else:
                    print("\t target None")

                # outs = state[pulse[0]]["output"]
                # for output in outs:
                #     if output not in state.keys():
                #         continue

                #     print("\t\tY", output, state[output])

                #     # if not (pulse[0] != "broadcaster" and state[pulse[0]]["kind"] == "%" and pulse[1] == 1):
                #     targets = sendPulse(state, pulse[0], output, pulse[1])
                #     if targets != None:
                #         newOutputs = newOutputs + targets
                #         print("\t", "targets", targets)
                #     else:
                #         print("\t target None")

            print("new", newOutputs)
            curr = cp.deepcopy(newOutputs)

            k=k+1

        print("k", k)
        break
        if i > 0:
            break
        i=i+1

    print()
    print("nLow, nHigh", nLow, nHigh)
    solution = nLow * nHigh
    print("solution:", solution)

    return

data = readFile()
solve(testData11)
print(50*"/")
solve(testData12)
# solve(data)

print()
print("########### PART 2 ###############")
print()
