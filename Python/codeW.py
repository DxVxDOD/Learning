def order_weight(strng):
    # your code
    def bs(arr, target):
        l = 0
        prox = float("inf")
        h = len(arr)
        while l < h:
            mid = (l + h) // 2
            close = if arr[mid] - target < 0 : return mid + 1 else return mid
            prox = min(prox, )
            print(prox)
            sum_mid = 0
            for y in arr[mid]:
                sum_mid += int(y)
            if sum_mid == target:
                return mid
            if sum_mid > target:
                h = mid
            else:
                l = mid + 1
        return prox
    strs = strng.split()
    res = []
    for n in strs:
        nums = list(n)
        sum_str = 0
        for m in nums:
            sum_str += int(m)
        if len(res) == 0:
            res.append(n)
            continue
        if len(res) == 1:
            sum_l = 0
            for x in res[-1]:
                sum_l += int(x)
            if sum_l < sum_str:
                res.append(n)
            else:
                res.insert(0, n)
            continue
        print(res)
        index = bs(res, sum_str)
        res.insert(index + 1, n)
        print(index)
        print(res)
        print("end")
    return " ".join(res)


print(order_weight("103 123 4444 99 2000"))
