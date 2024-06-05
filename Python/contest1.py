def mostCommonWord(p: str, banned: list[str]) -> str:
    freq = [[] * i for i in range(len(p) + 1)]
    count = {}
    newP = ""
    for i, c in enumerate(p):
        if c in [".", ",", "?", "!", ";"]:
            continue
        newP += c
    pList = list(newP.split())
    for w in pList:
        if w not in banned:
            count[w.lower()] = 1 + count.get(w.lower(), 0)
    for w, c in count.items():
        freq[c].append(w)
    for i in range(len(freq) - 1, -1, -1):
        for w in freq[i]:
            return w


print(mostCommonWord("Bob hit a ball, the hit BALL flew far after it was hit.", ["hit"]))
