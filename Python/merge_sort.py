def merge_sort(arr):
    if (len(arr) > 1):
        mid = len(arr) // 2
        leftArr = arr[:mid]
        rightArr = arr[mid:]

        merge_sort(leftArr)
        merge_sort(rightArr)
# You keep deviding the array until each one hits the size of 2

        i = j = k = 0

        while i < len(leftArr) and j < len(rightArr):
            if leftArr[i] < rightArr[j]:
                arr[k] = leftArr[i]
                i += 1
            else:
                arr[k] = rightArr[j]
                j += 1
            k += 1

        while i < len(leftArr):
            arr[k] = leftArr[i]
            i += 1
            k += 1
        while j < len(rightArr):
            arr[k] = rightArr[j]
            j += 1
            k += 1


arr = [246917, 99, 420, 12, 17, 24, 89, 34, 42, 16, 27, 22, 20]

print(arr)
merge_sort(arr)
print(arr)
