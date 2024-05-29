arrOfNums = [45, 62, 18, 79, 33, 50, 7, 91, 25, 84, 10, 68, 88, 96, 12, 67, 14, 89, 3, 53]


def partition(nums, lo, hi):
    pivot = nums[hi]
    i = lo - 1

    for j in range(lo, hi):
        if nums[j] <= pivot:
            i += 1
            nums[i], nums[j] = nums[j], nums[i]
    i += 1
    nums[i], nums[hi] = nums[hi], nums[i]
    return i


def quick_sort(arr, low, high):
    if low < high:
        pIndex = partition(arr, low, high)
        quick_sort(arr, low, pIndex - 1)
        quick_sort(arr, pIndex + 1, high)
    return arr


print(quick_sort(arrOfNums, 0, len(arrOfNums) - 1))
