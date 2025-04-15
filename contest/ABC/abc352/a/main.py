N, X, Y, Z = map(int, input().split())

if X < Y:
    for i in range(X, Y+1):
        if i == Z:
            print("Yes")
            exit()
else:
    for i in range(Y, X+1):
        if i == Z:
            print("Yes")
            exit()

print("No")
