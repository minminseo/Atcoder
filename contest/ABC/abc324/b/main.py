N = int(input())

while N % 2 == 0:
    N = N // 2

while N % 3 == 0:
    N = N // 3

if N == 1:
    print("Yes")
else:
    print("No")