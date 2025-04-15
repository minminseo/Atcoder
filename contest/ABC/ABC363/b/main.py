N, T, P = map(int, input().split())
L = list(map(int, input().split()))
days = 0

while True:
    count = 0
    for i in range(N):
        if L[i] >= T:
            count += 1
    if count >= P:
        print(days)
        break
    else:
        for i in range(len(L)):
            L[i] += 1
        days += 1
        
        
