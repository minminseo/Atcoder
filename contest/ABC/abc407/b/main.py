
X, Y = map(int, input().split())

count = 0
for a in range(1, 7):
    for b in range(1, 7):
        
        if a + b >= X or abs(a - b) >= Y:
            count += 1
            
result = count / 36
print(result)
