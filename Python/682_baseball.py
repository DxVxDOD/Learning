def baseball(o: list[str]):
    res = []
    for i in range(len(o)):
        if o[i] == "+":
            print(res)
            res.append(res[i] + res[i - 1])
        elif o[i] == "D":
            print(res)
            res.append(res[i] * 2)
        elif o[i] == "C":
            print(res)
            res.pop()
        else:
            print(res)
            res.append(o[i])
    return sum(res)

baseball(["5","2","C","D","+"])
