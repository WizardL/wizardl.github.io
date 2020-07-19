---
title: Codeforce 699D
description: "Codeforce 699D solution explanation undirected tree find detect cycle"
---

<!-- toc -->

# Codeforce 699D Fix a Tree
Today we are gonna take a look of problem Codeforce 699D
[Problem Link](https://codeforces.com/problemset/problem/699/D)

> **Problem**
A tree is an undirected connected graph without cycles.
Let's consider a rooted undirected tree with n vertices, numbered 1 through n. There are many ways to represent such a tree. One way is to create an array with n integers $p_1$,$p_2$, ..., $p_n$, where $p_i$ denotes a parent of vertex $i$ (here, for convenience a root is considered its own parent).
![](https://i.imgur.com/gh0W1ms.png)
Given a sequence $p_1$, $p_2$, ..., $p_n$, one is able to restore a tree:
>1. There must be exactly one index $r$ that $p_r=r$. A vertex $r$ is a root of the tree.
>2. For all other $n - 1$ vertices $i$, there is an edge between vertex $i$ and vertex $p_i$.
A sequence $p_1$, $p_2$, ..., $p_n$ is called valid if the described procedure generates some (any) rooted tree. For example, for $n = 3$ sequences $(1,2,2)$, $(2,3,1)$ and $(2,1,3)$ are not valid.
>
>You are given a sequence $a_1$, $a_2$, ..., $a_n$, not necessarily valid. Your task is to change the minimum number of elements, in order to get a valid sequence. Print the minimum number of changes and an example of a valid sequence after that number of changes. If there are many valid sequences achievable in the minimum number of changes, print any of them.
**Input**
The first line of the input contains an integer $n$ ($2 ≤ n ≤ 200 000$) — the number of vertices in the tree.
The second line contains $n$ integers $a_1$, $a_2$, ..., $a_n$ ($1 ≤ a_i ≤ n$).
**Output**
In the first line print the minimum number of elements to change, in order to get a valid sequence.
In the second line, print any valid sequence possible to get from ($a_1$, $a_2$, ..., $a_n$) in the minimum number of changes. If there are many such sequences, any of them will be accepted.

Thinking the structure of the graph, I found that it suffice to identitfy all Cycles(Rings in code) and all self-looping node(Roots in code), then we do the following:
1. Choose arbitary node in the cycle and make it a self-looping node
2. Eliminate self-looping node until there's only one(I chain all Roots)

The problems left here are:
1. How to find self-looping node?
2. How to identitfy cycles in undirected graph?

The first problem is reletivly easy, all node with $a[i]==i$
To resolve the second problem, we need to use another structure: disjoint-set.
Note that in this problem, one node only aggresively connect to one node, so the structure of cycles here is very simple.(think about it)
Then we can loop among all nodes, and if $a[i]!=i$, we union the two nodes. If $a[i]!=i$ but the two nodes are already in the same group, that means there's a cycle.
**Code:**
```c++
#include <bits/stdc++.h>
#define rank dkngpsngs
using namespace std;
const int _n = 2e5 + 10;
int n, a[_n], cnt = 0, ans[_n];
vector<int> Rings, Roots;

//start of structure disjoint-set
int parent[_n], rank[_n];
inline void dsinit(int n) {
  for (int i = 0; i < n; i++)
    parent[i] = i;
  memset(rank, 0, sizeof rank);
}
inline int dsfind(int e) { return parent[e] == e ? e : parent[e] = dsfind(parent[e]); }
inline void dsunion(int s1, int s2) {
  if (rank[s1] < rank[s2]) swap(s1, s2);
  parent[s2] = s1;
  if (rank[s1] == rank[s2]) rank[s1]++;
}
//end of disjoint set

main(void) {
  cin.tie(0);
  ios_base::sync_with_stdio(0);
  cin >> n;
  dsinit(n + 1);
  for (int i = 1; i <= n; i++) {
    cin >> a[i];
    ans[i] = a[i];
  }
  for (int i = 1; i <= n; i++) {
    if (a[i] == i)
      Roots.push_back(i);
    else {
      if (dsfind(a[i]) != dsfind(i))
        dsunion(dsfind(a[i]), dsfind(i));
      else
        Rings.push_back(i);
    }
  }
  for (int i = 0; i < (int)Rings.size(); i++) {
    ans[Rings[i]] = Rings[i];
    Roots.push_back(Rings[i]);
  }
  for (int i = 1; i < (int)Roots.size(); i++) ans[Roots[i]] = Roots[i - 1];
  for (int i = 1; i <= n; i++)
    if (ans[i] != a[i]) cnt++;
  cout << cnt << '\n';
  for (int i = 1; i <= n; i++) {
    cout << ans[i] << " ";
  }
  cout << '\n';
  return 0;
}
```
[Submission](https://codeforces.com/contest/699/submission/86767521)