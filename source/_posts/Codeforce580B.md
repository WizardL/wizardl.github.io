---
title: Codeforce 580B
description: "Codeforce 580B solution explanation binary search"
---

<!-- toc -->

# Codeforce 1370C Kefa and Company
Today we are gonna take a look of problem Codeforce 580B
[Problem Link](https://codeforces.com/problemset/problem/580/B)

> **Problem**
Kefa wants to celebrate his first big salary by going to restaurant. However, he needs company.
Kefa has $n$ friends, each friend will agree to go to the restaurant if Kefa asks. Each friend is characterized by the amount of money he has and the friendship factor in respect to Kefa. The parrot doesn't want any friend to feel poor compared to somebody else in the company (Kefa doesn't count). A friend feels poor if in the company there is someone who has at least d units of money more than he does. Also, Kefa wants the total friendship factor of the members of the company to be maximum. Help him invite an optimal company!
**Input**
The first line of the input contains two space-separated integers, n and d ($1 \le n \le 10^5$, $1\le d\le10^9$) — the number of Kefa's friends and the minimum difference between the amount of money in order to feel poor, respectively.
Next n lines contain the descriptions of Kefa's friends, the (i + 1)-th line contains the description of the i-th friend of type $m_i$, $s_i$ ($0 \le m_i, s_i \le 10^9$) — the amount of money and the friendship factor, respectively.
**Output**
Print the maximum total friendship factor that can be reached.

### Idea
Since for two friends that are able to be in the same company, the another firend that have money between them must also be able to be in the same company, so we would find that we might need to sort the array with respect to money.
Now we want to find that which _contiguous_ sub-array have the maximum firendship factor. We might want to try out all possible contiguous sub-arrays, but the complexity will be $O(n^2)$, that's too much.
Further observe, we would find that we only need to check the sub-arrays that _cannot_ be further expanded(with one fixed index, say, the right hand side), $i.e.$ we may thought of binary search, to search the first occurance of friend that has money more than $firend[i].money-d$, where we iterate index $i$ from $n$ to $1$.
### Code:
```c++
#include <bits/stdc++.h>
#define ll long long
using namespace std;
const int _n = 1e5 + 10;
int n, d, s, m;
ll maxx = -1, pre[_n];
pair<int, int> a[_n];
bool cmp(int val, pair<int, int> a) { return a.first > val; }
main(void) {
  cin.tie(0);
  ios_base::sync_with_stdio(0);
  cin >> n >> d;
  for (int i = 1; i <= n; i++) {
    cin >> s >> m;
    a[i] = {s, m};
  }
  sort(a + 1, a + n + 1), pre[1] = a[1].second, pre[0] = 0;
  for (int i = 2; i <= n; i++) pre[i] = pre[i - 1] + a[i].second;
  for (int i = n; i >= 1; i--) {
    int l = upper_bound(a + 1, a + 1 + i, a[i].first - d, cmp) - a;
    maxx = max(maxx, pre[i] - pre[l - 1]);
  }
  cout << maxx << '\n';
  return 0;
}
```
Note that $upper\_bound$ is based on binary-search, and to search in type $pair<int,int>$, we need to define the compare function for $upper\_bound$, the parameter types are defined in documents.
[Submission](https://codeforces.com/contest/580/submission/86955857)