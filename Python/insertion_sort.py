arr1 = [9, 8, 7, 6, 5]
arr4 = [1, 69, 24, 11, 10, 420, 42069, 17, 5, 11, 15, 33, 22]
arr2 = [1]
arr3 = [1112, 11, 2]


def insertion_sort(arr):
    # right pointer
    i = 1
    count = 0
    while i < len(arr):
        count += 1
        # card in right hand
        k = arr[i]
        # card in left hand, left pointer
        j = i - 1
        # when the last card in the left hand is larger than the card in the right hand swap them
        while j >= 0 and arr[j] > k:
            arr[j + 1] = arr[j]
            j = j - 1
            count += 1
        arr[j + 1] = k
        i += 1
    print(count)

# The time complexity of the algorithm is O(N * N) becuase we have 1 nested loop,
# we loop trough it in quadratic time.


insertion_sort(arr1)
# insertion_sort(arr2, count)
# insertion_sort(arr3, count)

print(arr1, len(arr1))
# print(arr2)
# print(arr3)


def reverse_insertion_sort(arr):
    # right pointer
    i = 1
    while i < len(arr):
        # card in right hand
        k = arr[i]
        # card in left hand, left pointer
        j = i - 1
        # when the last card in the left hand is larger than the card in the right hand swap them
        while j >= 0 and arr[j] < k:
            arr[j + 1] = arr[j]
            j = j - 1
        arr[j + 1] = k
        i += 1


reverse_insertion_sort(arr1)
print(arr1)
