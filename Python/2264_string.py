def largestGoodInteger(num: str) -> str:
    res = ""
    j = 2
    while j < len(num):
        print(len(res))
        print(res)
        if num[j - 2] == num[j - 1] and num[j - 2] == num[j]:
            if len(res) < 1:
                res = "0"
            if num[j] != "0" and len(res) < 3:
                res = str(max(int(res), int(num[j - 2:j + 1])))
            else:
                res = num[j - 2:j + 1]
        j += 1
    return res


print(largestGoodInteger("000400059"))
