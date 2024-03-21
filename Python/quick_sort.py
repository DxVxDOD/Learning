def quick_sort(arr: list) -> list:
    # base case
    if len(arr) < 2:
        return arr
    # recursive case
    pivot = arr[0]
    less = [i for i in arr[1:] if i <= pivot]
    greater = [i for i in arr[1:] if i > pivot]
    return quick_sort(less) + [pivot] + quick_sort(greater)


print(quick_sort([45, 62, 18, 79, 33, 50, 7, 91, 25, 84, 10, 68, 88, 96, 12, 67, 14, 89, 3, 53]))
