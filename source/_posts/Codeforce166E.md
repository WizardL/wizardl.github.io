---
title: Codeforce 166E
description: "Codeforce 166E dp matrix logarithm time exponential multiply rolling dp"
---

<!-- toc -->

# Codeforce 166E Tetrahedron
Today we are gonna take a look of problem Codeforce 166E
[Problem Link](https://codeforces.com/problemset/problem/166/E)

> **Problem**
You are given a tetrahedron. Let's mark its vertices with letters A, B, C and D correspondingly.
![](https://i.imgur.com/t5TKHXl.png)
An ant is standing in the vertex $D$ of the tetrahedron. The ant is quite active and he wouldn't stay idle. At each moment of time he makes a step from one vertex to another one along some edge of the tetrahedron. The ant just can't stand on one place.
You do not have to do much to solve the problem: your task is to count the number of ways in which the ant can go from the initial vertex $D$ to itself in exactly $n$ steps. In other words, you are asked to find out the number of different cyclic paths with the length of $n$ from vertex $D$ to itself. As the number can be quite large, you should print it modulo $1000000007$ ($10^9 + 7$).
**Input**
The first line contains the only integer $n$ ($1 \le n \le 10^7$) — the required length of the cyclic path.
**Output**
Print the only integer — the required number of ways modulo $1000000007$ ($10^9 + 7$).
### Idea
Thinking about how we calculate this if there's no computer, we will consider the state of the previous vertex. The way we calculate this depends on: Whether the previous vertex is $D$?
So we can derive the dp transfer formula:
$$\begin{cases}
dp[now][0]=dp[previous][1]\\
dp[now][1]=dp[previous][0]\times 3+dp[previous][1]\times 2
\end{cases}$$ Where we denote $dp[i][0]$ as the possibilities at $i$-th step and the last vertex is $D$, and denote $dp[i][1]$ as the possibilities at $i$-th step and the last vertex is not $D$.
And we can find that we only need previous state's data, so we don't need to store all $n$ states' data. We can apply a technique "rolling dp". We denote $i\&1$ as the index of $now$, and $!(i\&1)$ as the index of $previous$, where $i$ denote that we are calculating $i$-th data. Thus we only need to define dp array as $int\ dp[2][2];$
Since $n$ will not exceed $1e7$, the $O(n)$ algorithm is enough.
If $n$ will be $1e9$ or higher, then we need(can) to reduce the complexity to $O(\lg n)$
We can convert the dp transfer formula into matrix arithmetic:
$$
\begin{bmatrix}
dp[now][0]\\
dp[now][1]
\end{bmatrix}=\begin{bmatrix}
0&1\\
3&2
\end{bmatrix}\begin{bmatrix}
dp[previous][0]\\
dp[previous][1]
\end{bmatrix}
$$ And to calculate $n$-th dp data, we just need to calculate $\begin{bmatrix}
0&1\\
3&2
\end{bmatrix}^n\begin{bmatrix}
dp[0\&1][0]\\
dp[0\&1][1]
\end{bmatrix}$
Then by using fast matrix exponent(same as ordinary fast exponential algorithm but change numbers to matrices) we can calculate $\begin{bmatrix}
0&1\\
3&2
\end{bmatrix}^n$ in $O(\lg n)$.
### Code:($O(n)$ solution)
```c++
#include <bits/stdc++.h>
using namespace std;
const int _n = 1e7 + 10, _m = 1e9 + 7;
int n, dp[2][2];
main(void) {
  cin.tie(0);
  ios_base::sync_with_stdio(0);
  cin >> n;
  dp[1][0] = 0, dp[1][1] = 3;
  for (int i = 2; i <= n; i++) {
    dp[i & 1][0] = dp[!(i & 1)][1];
    dp[i & 1][1] = 1ll * dp[!(i & 1)][0] * 3 % _m + 1ll * dp[!(i & 1)][1] * 2 % _m, dp[i & 1][1] %= _m;
  }
  cout << dp[n & 1][0];
  return 0;
}
```
[Submission](https://codeforces.com/contest/166/submission/87251081)