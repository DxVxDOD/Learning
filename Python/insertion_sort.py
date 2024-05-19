arr1 = [1, 5, 11, 15, 33, 22]
arr2 = [1]
arr3 = [1112, 11, 2]


def insertion_sort(arr):
    # right pointer
    i = 1
    while i < len(arr):
        # card in right hand
        k = arr[i]
        # card in left hand, left pointer
        j = i - 1
        # when the last card in the left hand is larger than the card in the right hand swap them
        while j > 0 and arr[j] > k:
            arr[j + 1] = arr[j]
            j = j - 1
        arr[j + 1] = k
        i += 1
# The time complexity of the algorithm is O(N * N) becuase we have 1 nested loop,
# we loop trough it in quadratic time.


insertion_sort(arr1)
insertion_sort(arr2)
insertion_sort(arr3)

print(arr1)
print(arr2)
print(arr3)
