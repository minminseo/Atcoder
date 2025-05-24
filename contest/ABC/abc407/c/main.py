S = input()
need = 0

for i in range(len(S)-1, -1, -1):
    d = int(S[i])
    if d < need:
        t = (need - d + 9) // 10
        d += t * 10
    need = d
print(len(S) + need)