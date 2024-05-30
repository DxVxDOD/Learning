def reverse_word(word, ch):
    word = list(word)
    for i in range(len(word)):
        if word[i] == ch:
            word = word[:i][::-1] + word[i:]
    return "".join(word)


print(reverse_word("abcdefd", "d"))
