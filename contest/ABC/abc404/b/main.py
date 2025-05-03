N = int(input())

S = []
for _ in range(N):
    S.append(input().rstrip())

T = []
for _ in range(N):
    T.append(input().rstrip())

best = N * N

current = S[:]

for k in range(4):

    mismatch = 0
    for i in range(N):
        for j in range(N):
            if current[i][j] != T[i][j]:
                mismatch += 1

    total = k + mismatch

    if total < best:
        best = total

    Spinned = []

    for i in range(N):
        new_row = ''
        for j in range(N):
            new_row += current[N-1-j][i]

        Spinned.append(new_row)

    current = Spinned

print(best)
