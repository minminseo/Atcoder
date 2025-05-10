N, M = map(int, input().split())
A = list(map(int, input().split()))

for k in range(N + 1):
    current_length = N - k
    ok = True
    
    for target in range(1, M + 1):
        found = False

        for j in range(current_length):

            if A[j] == target:
                found = True
                break

        if not found:
            ok = False
            break

    if not ok:
        print(k)
        break
