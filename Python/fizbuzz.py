def fizzbuzz(num: int):
    res = []
    for i in range(num):
        if i % 3 == 0 and i % 5 == 0:
            res.append("fizz buzz")
        elif i % 3 == 0:
            res.append("fizz")
        elif i % 5 == 0:
            res.append("buzz")
        else:
            res.append(i)
    print(res)


fizzbuzz(100)
