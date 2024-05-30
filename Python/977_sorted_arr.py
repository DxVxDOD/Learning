def sorted_arr(arr: list[int]):
    l = 0
    while l < len(arr) and arr[l] < 0:
        l += 1

    leftArr = arr[:l]
    rightArr = arr[l:]


    for i, n in enumerate(leftArr):
        leftArr[i] = n * n
    for j, m in enumerate(rightArr):
        rightArr[j] = m * m

    print(leftArr)
    print(rightArr)

sorted_arr([-7,-3,2,3,11])
