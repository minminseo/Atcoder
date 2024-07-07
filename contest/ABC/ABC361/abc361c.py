"""
結果は不正解。

＜メモ＞
・インデックスの範囲を決定するためにx=n-k-1
・昇順のリストで、なぜa[i+x]-a[i]の範囲の全探索で、答えの最小値の扱いになるかは理解できなかった。

"""

n,k=map(int,input().split())
a=sorted(map(int,input().split()))
x=n-k-1
ans=10**20
for i in range(k+1):
    ans=min(ans,a[i+x]-a[i])
print(ans)