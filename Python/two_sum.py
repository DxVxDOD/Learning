def two_sum(arr, target):
    seen = {}
    for i in range(len(arr)):
        diff = target - arr[i]
        if diff in seen:
            return [seen[diff], i]
        seen[arr[i]] = i


arr = [5, 121, 11, 12, 4, 1, 3, 6, 17]

print(two_sum(arr, 5))
