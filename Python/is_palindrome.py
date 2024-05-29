def is_palindrome(s: str):
    res = ""
    res2 = ""
    for i in range(len(s) - 1, -1, -1):
        if ord(s[i]) >= 48 and ord(s[i]) <= 57 or ord(s[i].lower()) >= ord("a") and ord(s[i].lower()) <= ord("z"):
            res += s[i].lower()
    for i in range(len(s)):
        if ord(s[i]) >= 48 and ord(s[i]) <= 57 or ord(s[i].lower()) >= ord("a") and ord(s[i].lower()) <= ord("z"):
            res2 += s[i].lower()
    print(res)
    print(res2)
    print(res2 == res)


string = "A man, a plan, a canal: Panama"
string2 = "0P"

is_palindrome(string2)
