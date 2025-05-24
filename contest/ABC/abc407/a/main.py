A, B = map(int, input().split())

r1 = A // B
r2 = A % B

if r2 * 2 >= B:
    print(r1 + 1)
else:
    print(r1)
