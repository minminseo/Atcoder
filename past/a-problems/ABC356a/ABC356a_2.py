N, L, R =map(int,input().split())

A = []

for i in range(1,N+1):
    A.append(i)

A = A[:L-1] + A[L-1:R][::-1] + A[R:]

print(*A)