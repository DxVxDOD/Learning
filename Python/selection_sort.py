def find_smallest(arr):
    small = arr[0]
    small_index = 0
    for i in range(1, len(arr)):
        if arr[i] < small:
            small = arr[i]
            small_index = i
    return small_index


def selection_sor(arr: list) -> list:
    sorted_arr = []
    for i in range(len(arr)):
        smallest = find_smallest(arr)
        sorted_arr.append(arr.pop(smallest))
    return sorted_arr


print(selection_sor([100, 4, 60, 1, 69, 420]))


def factorial(num: int) -> int:
    if num == 1:
        return num
    return num * factorial(num - 1)


print(factorial(6))


def sum_arr(nums: list) -> int:
    if len(nums) == 0:
        return 0
    if len(nums) == 1:
        return nums[0]
    return nums.pop() + sum(nums)


print(sum_arr([1, 2, 3, 5, 5, 30]))


def maximum_num(nums: list) -> int:
    if len(nums) == 2:
        return nums[0] if nums[0] > nums[1] else nums[1]
    sub_max = max(nums[1:])
    return nums[0] if nums[0] > sub_max else sub_max


print(maximum_num([1, 4, 5, 6, 100]))
