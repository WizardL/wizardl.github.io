---
title: Codeforce 339D
description: "Codeforce 339D solution explanation (segment) tree"
---

<!-- toc -->

# Codeforce 339D Xenia and Bit Operations
Today we are gonna take a look of problem Codeforce 339D
[Problem Link](https://codeforces.com/problemset/problem/339/D)

> **Problem**
Xenia the beginner programmer has a sequence $a$, consisting of $2^n$ non-negative integers: $a_1$, $a_2$, ..., $a_{2^n}$. Xenia is currently studying bit operations. To better understand how they work, Xenia decided to calculate some value $v$ for $a$.
Namely, it takes several iterations to calculate value $v$. At the first iteration, Xenia writes a new sequence $a_1$ or $a_2$, $a_3$ or $a_4$, ..., $a_{2^{n - 1}}$ or $a_{2^n}$, consisting of $2^{n - 1}$ elements. In other words, she writes down the bit-wise OR of adjacent elements of sequence $a$. At the second iteration, Xenia writes the bitwise exclusive OR of adjacent elements of the sequence obtained after the first iteration. At the third iteration Xenia writes the bitwise OR of the adjacent elements of the sequence obtained after the second iteration. And so on; the operations of bitwise exclusive OR and bitwise OR alternate. In the end, she obtains a sequence consisting of one element, and that element is $v$.
Let's consider an example. Suppose that sequence a = (1, 2, 3, 4). Then let's write down all the transformations (1, 2, 3, 4)  →  (1 or 2 = 3, 3 or 4 = 7)  →  (3 xor 7 = 4). The result is v = 4.
You are given Xenia's initial sequence. But to calculate value $v$ for a given sequence would be too easy, so you are given additional $m$ queries. Each query is a pair of integers $p$, $b$. Query $p$, $b$ means that you need to perform the assignment a $p = b$. After each query, you need to print the new value $v$ for the new sequence $a$.
**Input**
The first line contains two integers $n$ and $m$ ($1 \le n \le 17$, $1 \le m \le 10^5$). The next line contains $2^n$ integers $a_1$, $a_2$, ..., $a_{2^n}$ ($0 \le a_i < 2^30$). Each of the next m lines contains queries. The i-th line contains integers $p_i$, $b_i$ ($1 \le p_i \le 2^n$, $0 \le b_i < 2^30$) — the i-th query.
**Output**
Print $m$ integers — the i-th integer denotes value $v$ for sequence $a$ after the i-th query.

### Idea
If you had heard of the data structure "segment" tree(I quote segment since this name is ambiguous with an existing structure), then you can _possibly_ find that we need to use segment tree.
The tree root stores the number $v$, and the leaves of the tree stores $a[i]$.
The tree we use in this problem doesn't need a query function since we just need the number root->v, $i.e.$ we don't need range query.
And to deal with OR and XOR, it suffices to assign an variable $level$ for each node, when the level is odd, the pull() function($i.e.$ a step in the problem description) do OR, and XOR otherwise.

$p.s.$ I wondered why do we need to use segment tree even if we don't need the query function? I would say, that's because the structure this problem implies $is$ the form of a segment tree, and it incidentally fits with the functions that a segment tree would have.
### Code:
```c++
#include <bits/stdc++.h>
using namespace std;
const int _n = ((1 << 17) + 10);
int n, m, p, b, a[_n];
struct node {
  int v, lvl;
  node *l, *r;
  node(int v, int lvl) : v(v), lvl(lvl) {}
  node(node *l, node *r) : l(l), r(r) { pull(); }
  void pull() {
    lvl = l->lvl + 1;
    if (lvl & 1)
      v = l->v | r->v;
    else
      v = l->v ^ r->v;
  }
};
node *build(int l, int r) {
  if (l + 1 == r) return new node(a[l], 0);
  int mid = (l + r) / 2;
  return new node(build(l, mid), build(mid, r));
}
void modify(node *cur, int pos, int l, int r, int v) {
  if (l + 1 == r) {
    cur->v = v;
    return;
  }
  int mid = (l + r) / 2;
  if (pos < mid)
    modify(cur->l, pos, l, mid, v);
  else
    modify(cur->r, pos, mid, r, v);
  cur->pull();
}
node *rt;
main(void) {
  cin.tie(0);
  ios_base::sync_with_stdio(0);
  cin >> n >> m;
  int nn = 1 << n;
  for (int i = 0; i < nn; i++) cin >> a[i];
  rt = build(0, nn);
  while (m--) {
    cin >> p >> b;
    p--;
    modify(rt, p, 0, nn, b);
    cout << rt->v << '\n';
  }
  return 0;
}
```
[Submission](https://codeforces.com/contest/339/submission/87072697)