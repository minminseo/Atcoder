"""
別解
"""

S = input()

for i in set(S):
    if S.count(i) == 1:
        print(S.find(i) + 1)