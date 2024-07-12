S = input()
if S[0] == S[1]:
  for i in range(len(S)):
    if S[i] != S[0]:
      print(i+1)
      exit()
else:
    if S[0] == S[2]:
      print(2)
      exit()
    else:
      print(1)
      exit()