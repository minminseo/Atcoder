S = input().strip()
alphabet = 'abcdefghijklmnopqrstuvwxyz'

for i in range(len(alphabet)):
    
    found = False
    
    for j in range(len(S)):
        if alphabet[i] == S[j]:
            found = True
            break

    if not found:
        print(alphabet[i])
        break
