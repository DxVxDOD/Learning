def match(pattern: str, string: str):
    sList = list(string.split())

    for i in range(1, len(pattern)):
        print("sList -->", (sList[i] == sList[i - 1]))
        print("pattern -->", (pattern[i] == pattern[i - 1]))
        if (sList[i] == sList[i - 1]) != (pattern[i] == pattern[i - 1]):
            return False
    return True


print(match("abba", "dog cat cat dog"))
print(match("abba", "dog cat cat fish"))
