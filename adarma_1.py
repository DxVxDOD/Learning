arr = [
    [0, 0, 0],
    [0, 0, 0],
    [0, 0, 0],
]

m = 9

while m > 0:

    turn = 1

    for a in arr:
        print(a)

    place = int(input("Pick place: "))

    if place <= 3:
        arr[0][place - 1] = turn
        turn *= -1
    if place > 3 and place <= 6:
        arr[1][place - 3] = turn
        turn *= -1
    if place > 6 and place <= 9:
        arr[1][place - 6] = turn
        turn *= -1
    m -= 1
