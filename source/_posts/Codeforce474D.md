---
title: Codeforce 474D
description: "Codeforce 474D solution explanation"
---

<!-- toc -->

# Codeforce 474D Flowers
Today we are gonna take a look of problem Codeforce 474D
[Problem Link](https://codeforces.com/problemset/problem/474/D)

> **Problem**
> We saw the little game Marmot made for Mole's lunch. Now it's Marmot's dinner time and, as we all know, Marmot eats flowers. At every dinner he eats some red and white flowers. Therefore a dinner can be represented as a sequence of several flowers, some of them white and some of them red.
> But, for a dinner to be tasty, there is a rule: Marmot wants to eat white flowers only in groups of size k.
> Now Marmot wonders in how many ways he can eat between a and b flowers. As the number of ways could be very large, print it modulo 1000000007 (10^9 + 7).
>**Input**
>Input contains several test cases.
>The first line contains two integers t and k (1 ≤ t, k ≤ $10^5$), where t represents the number of test cases.
>The next t lines contain two integers $a_i$ and $b_i$ (1 ≤ $a_i$ ≤ $b_i$ ≤ $10^5$), describing the i-th test.
>**Output**
>Print t lines to the standard output. The i-th line should contain the number of ways in which Marmot can eat between $a_i$ and $b_i$ flowers at dinner modulo 1000000007 (10^9 + 7).
>**Examples**
>**input**
3 2
1 3
2 3
4 4
**output**
6
5
5
**Note**
For K = 2 and length 1 Marmot can eat \(R\).
For K = 2 and length 2 Marmot can eat (RR) and (WW).
For K = 2 and length 3 Marmot can eat (RRR), (RWW) and (WWR).
For K = 2 and length 4 Marmot can eat, for example, (WWWW) or (RWWR), but for example he can't eat (WWWR).

At first glance we may think this problem is about combination math, but if you think further, you will find that you need at least $O(nk)$ time(n denote the max $b_i$ among all $i$), and that's too much.
The solution I found is using dp. You can surprisingly find that the following dp transfer can reduce the complexity to $O(max_i\{b_i\})$.
$$\begin{cases}
dp[0..k-1]=1\\
dp[i]=dp[i-1]+dp[i-k]&i\ge k
\end{cases}$$ Where $dp[i]$ denote the number of combinations of length $i$
For $0\le i\le k-1$, since we can't eat any white flower yet, so the value will be $1$.
For $i\ge k$, We want to add length by $1$ everytime, and there're only two possiblities for the $i$-th flower, red or white. If it's red, it's clear that there're $dp[i-1]$ possibilities, and for white, since we need to put a whole group of $k$ flowers, the trailing $k$ flowers will all be white, and there left $i-k$ spots for us to randomly choose flowers, there're $dp[i-k]$ possibilities.

The crucial part had been done, now it suffice to turn $dp$ array into prefix-sum array, then we can $O(1)$ query any segment.

**Code:**
```c++
#include <bits/stdc++.h>
using namespace std;
const int _m = 1e9 + 7, _n = 1e5 + 10;
int t, k, a, b, dp[_n];
main(void) {
  cin.tie(0);
  ios_base::sync_with_stdio(0);
  cin >> t >> k;
  for (int i = 0; i <= k - 1; i++) dp[i] = 1;
  for (int i = k; i <= 100000; i++) dp[i] = (1ll * dp[i - 1] + 1ll * dp[i - k]) % _m;
  for (int i = 1; i <= 100000; i++) dp[i] = (1ll * dp[i - 1] + 1ll * dp[i]) % _m;
  while (t--) {
    cin >> a >> b;
    cout << ((1ll * dp[b] - 1ll * dp[a - 1]) % _m >= 0 ? (1ll * dp[b] - 1ll * dp[a - 1]) % _m : (1ll * dp[b] - 1ll * dp[a - 1]) % _m + _m) << '\n';
  }
  return 0;
}
```
[Submission](https://codeforces.com/contest/474/submission/86848551)