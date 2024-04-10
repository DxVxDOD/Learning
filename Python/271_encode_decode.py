from typing import List


def encode(strs: List[str]) -> str:
    for i in range(len(strs)):
        strs[i] += ":"
    return "".join(strs)


def decode(str: str):
    arr = str.split(":")
    arr.pop()
    return arr


en = encode(["david", "beea", "puffy"])

print(encode(["david", "beea", "puffy"]))
print(decode(en), 3)
