N, L, R = map(int, input().split())

A = []

for i in range(1, N+1):
    A.append(i)

result = []

for i in range(L-1):
    result.append(A[i])


for i in range(R-1, L-2, -1):
    result.append(A[i])


for i in range(R, N):
    result.append(A[i])

print(*result)
# もしくは →　print(' '.join(map(str, result)))