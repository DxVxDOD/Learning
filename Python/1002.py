def commonChars(words: list[str]) -> list[str]:
    res = []
    for i in range(len(words[0])):
        count = 0
        for w in words:
            if words[0][i] in w:
                count += 1
        if count % 3 == 0:
            res.append(words[0][i])
        print(count)
    return res


print(commonChars(["bella","label","roller"]))
