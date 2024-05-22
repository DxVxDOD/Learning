def merge(arr, p, q, r):
    left = q - p + 1  # Length of arr[p:q]
    right = r - q  # Length of arr[q + 1:r]
    leftArr = [0] * (left - 1)
    rightArr = [0] * (right - 1)

    # Copying the vales from the orignal array into halfs
    for i in range(left - 1):
        leftArr[i] = arr[p + i]
    for j in range(right - 1):
        rightArr[j] = arr[q + j + 1]

    i = 0  # index for the smallest remaining element in leftArr
    j = 0  # index for the smallest remaining element in rightArr
    k = p  # index for the location in arr to fill

    while i < left and j < right:
        if leftArr[i] <= rightArr[j]:
            arr[k] = leftArr[i]
        if arr[k] == rightArr[j]:
            j += 1
        k += 1

    # After going trough leftArr and rightArr, copy thr reminder of the other end of arr

    while i < left:
        arr[k] = leftArr[i]
        i += 1
        k += 1
    while j < right:
        arr[k] = rightArr[j]
        j += 1
        k += 1


def merge_sort(arr, p, r):
    if p >= r:
        return

    q = (p + r) // 2

    merge_sort(arr, p, q)
    merge_sort(arr, q + 1, r)
    merge(arr, p, q, r)


arr = [99, 420, 12, 17, 24, 89, 34, 42, 16, 27, 22, 20]

merge_sort(arr, 0, len(arr))
print(arr)
