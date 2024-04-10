testArray = [1, 1, 1, 2, 2, 3, 4, 4, 4, 4, 4]


def test(nums, k):
    res = []
    freq = {}
    for n in nums:
        freq[n] = 1 + freq.get(n, 0)
    i = 0
    print(freq.keys())
    print(freq.values())
    print(freq.items())

    for n, c in freq.items():
        :wq

    return res


print(test(testArray, 2))
