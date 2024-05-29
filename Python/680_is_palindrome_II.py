def is_palindorme(s: str):
    res = ""
    for i in range(len(s) - 1, -1, -1):
        res += s[i]
    if res == s:
        print("res --> ", res)
        return True
    else:
        for i in range(len(res) - 1, -1, -1):
            tmp = list(res)
            tmp.remove(s[i])
            tmp = "".join(tmp)

            print("temp =>", tmp)
            print("res =>", res)
            if tmp == s:
                return True
    return False


print(is_palindorme("abca"))
