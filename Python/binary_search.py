def binary_search(arr: list, item: int) -> int:
    low = 0
    high = len(arr) - 1
    while low <= high:
        mid = (low + high) // 2
        if arr[mid] == item:
            return arr[mid]
        if mid < item:
            low = mid + 1
        else:
            high = mid - 1
    return -1


print(binary_search([1, 3, 5, 7, 9], 9))
