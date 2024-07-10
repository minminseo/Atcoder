"""
解説

<for i in range(len(S)-1, 0, -1):の部分に関して>
S[len(S)]は存在しないため（文字列のインデックスは0から始まってる）、最後の文字列のインデックスはlen(S)-1になる。

"""

S = input()
first_result = ""
second_result = ""

for i in range(len(S)):
    if S[i] == "|":
        first_result = S[:i]
        break
    
for i in range(len(S)-1, 0, -1):
    if S[i] == "|":
        second_result = S[i+1:]
        break

print(first_result + second_result)