def first_unique_char(s: str):
    res = [0] * 26
    ret = -1
    for c in s:
        res[ord(c) - ord("a")] += 1
    print(res)
    for i in range(len(res)):
        if res[i] == 1:
            ret = i
            return ret


print(first_unique_char("leetcode"))
