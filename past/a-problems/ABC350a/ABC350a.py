S = input()
contest = 0

for i in range(len(S)):
    if i == 3:
        contest += int(S[i]) * 100
    elif i == 4:
        contest += int(S[i]) * 10
    elif i == 5:
        contest += int(S[i])
        
if 0 < contest < 316 or 316 < contest < 350:
    print("Yes")
else:
    print("No")
        