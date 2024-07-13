"""
解けなかった。
条件分岐でYesの場合の処理が理解できなかった。
<メモ>
E869120さんの[アルゴリズム・AtCoder のための数学【後編：数学的考察編】](https://qiita.com/e869120/items/1ccb2bdf16890637e767)の
「7-6. 条件を減らした場合を考える（レベル：:star:2）」に考え方が近い。
"""


N = int(input())
L, R = [0] * N, [0] * N
for i in range(N):
    L[i], R[i] = map(int, input().split())

if sum(L) > 0 or sum(R) < 0:
    print("No")
    exit()

X = L.copy()
sumX = sum(X)
for i in range(N):
    d = min(R[i] - L[i], -sumX)
    sumX += d
    X[i] += d

print("Yes")
print(*X)
