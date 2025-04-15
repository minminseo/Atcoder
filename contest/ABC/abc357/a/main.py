
N,M=map(int,input().split())
H=list(map(int,input().split()))

ans=0

for i in H:
    M-=i
    if M<0:
        break
    ans+=1
    
print(ans)
