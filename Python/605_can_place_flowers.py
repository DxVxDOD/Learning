def place_flower(f: list[int], n: int):
    for i in range(len(f)):
        if i == 0 and f[i] == 0 and f[i + 1] == 0:
            f[i] = 1
            n -= 1
        if i == (len(f) - 1) and f[i] == 0 and f[i - 1] == 0:
            f[i] = 1
            n -= 1
        if i > 0 and f[i - 1] == 0 and f[i] == 0 and f[i + 1] == 0 and i < (len(f) - 1):
            f[i] = 1
            n -= 1
    return n < 1


print(place_flower([1, 0, 0, 0, 0, 1], 2))
