N,M=map(int,input().split())
H=list(map(int,input().split()))
cnt = 0

for i in range(N):
    if M > 0:
        M -= H[i]
        cnt += 1
        if  M < 0:
            cnt -= 1
            break
        
print(cnt)
