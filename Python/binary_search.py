
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


def recursive_binary_search(nums: list, item: int) -> int:
    mid = len(nums) // 2
    if mid == 1 & nums[mid] != item:
        return -1
    if nums[mid] == item:
        return nums[mid]
    if nums[mid] > item:
        return recursive_binary_search(nums[0:mid], item)
    return recursive_binary_search(nums[mid:], item)


random_ints = [3, 7, 10, 12, 14, 18, 25, 33, 45, 50, 53, 62, 67, 68, 79, 84, 88, 89, 91, 96]


print(recursive_binary_search(random_ints, 68))
