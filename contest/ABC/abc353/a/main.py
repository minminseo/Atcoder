N = int(input())

H = list(map(int, input().split()))
first = H[0]

for i in range(1, N):
    if first < H[i]:
        print(i+1)
        exit()

print(-1)