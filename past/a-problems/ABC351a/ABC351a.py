A = list(map(int, input().split()))
B = list(map(int, input().split()))
Asum = 0
Bsum = 0
for i in range(len(A)):
    Asum += A[i]

for i in range(len(B)):
    Bsum += B[i]
    
print(Asum - Bsum + 1)