"""
結果は不正解。下のコード間違ってる。
"""

import itertools

N, K = map(int, input().split())
A = list(map(int, input().split()))

combinations = list(itertools.combinations(A, K))
min_list = "inf"


for i in combinations:
    result = max(i) - min(i)
    if result < min_list:
        min_list = result
        
print(result)