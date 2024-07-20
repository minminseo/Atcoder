N, T, P = map(int, input().split())
L = list(map(int, input().split()))

count = 0
for i in L:
    if i >= T:
        count += 1

if count >= P:
    print(0)
else:
    days = 0
    while count < P:
        days += 1
        count = 0
        for i in L:
            if i + days >= T:
                count += 1

    print(days)