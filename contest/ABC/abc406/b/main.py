N, K = map(int, input().split())
A = list(map(int, input().split()))

answer = 1
for i in range(N):
    answer *= A[i]
    if K < len(str(answer)):
        answer = 1

print(answer)

# or

n,k=map(int, input().split())
a=list(map(int, input().split()))
x,y=1,10**k-1
for i in range(n):
    x*=a[i]
    if x>y:
        x=1
print(x)