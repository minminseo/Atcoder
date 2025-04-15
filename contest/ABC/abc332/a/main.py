N, S, K = map(int, input().split())

sum = 0
for i in range(N):
    P, Q = map(int, input().split())
    sum += P*Q
    
if S <= sum:
    print(sum)
else:
    print(sum + K)