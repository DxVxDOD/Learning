def binary_search(arr: list[int], needle: int):
    left = 0
    right = len(arr) - 1
    while left < right:
        m = right + left // 2
        if arr[m] == needle:
            return True
        if arr[m] > needle:
            right = m
        else:
            left = m + 1
    return False


print(binary_search([169, 27, 118, 49, 14, 205, 448, 303], 303))
print(binary_search([169, 27, 118, 49, 14, 205, 448, 303], 3033))
