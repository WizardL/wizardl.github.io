---
title: Codeforce 1352G
description: "Codeforce 1352G math constructive"
---

<!-- toc -->

# Codeforce 1352G Special Permutation
Today we are gonna take a look of problem Codeforce 1352G
[Problem Link](https://codeforces.com/problemset/problem/1352/G)

> **Problem**
A permutation of length n is an array p=[p1,p2,…,pn], which contains every integer from 1 to n (inclusive) and, moreover, each number appears exactly once. For example, p=[3,1,4,2,5] is a permutation of length 5.
For a given number n (n≥2), find a permutation p in which absolute difference (that is, the absolute value of difference) of any two neighboring (adjacent) elements is between 2 and 4, inclusive. Formally, find such permutation p that 2≤|pi−pi+1|≤4 for each i (1≤i<n).
Print any such permutation for the given integer n or determine that it does not exist.
**Input**
The first line contains an integer t (1≤t≤100) — the number of test cases in the input. Then t test cases follow.
Each test case is described by a single line containing an integer n (2≤n≤1000).
**Output**
Print t lines. Print a permutation that meets the given requirements. If there are several such permutations, then print any of them. If no such permutation exists, print -1.

### Idea
Since the difference of odd(even) numbers are $2$, I guess the answer is some arrangement of [odd number sequence][backward even number sequence].
After trying the case $n=10$ and $n=9$, I found a way to construct the sequence. If $n$ is even, swap the largest two odd numbers, if $n$ is odd, swap the largest two even numbers. (It may be more clear by looking at the code).

### Code:
```c++
#include <bits/stdc++.h>
using namespace std;
int t, n;
vector<int> odd, even;
main(void) {
  cin.tie(0);
  ios_base::sync_with_stdio(0);
  cin >> t;
  while (t--) {
    cin >> n;
    if (n <= 3) {
      cout << "-1\n";
      continue;
    }
    odd.clear(), even.clear();
    for (int i = 1; i <= n; i += 2) odd.push_back(i);
    for (int i = 2; i <= n; i += 2) even.push_back(i);
    if (!(n & 1))
      iter_swap(odd.end() - 1, odd.end() - 2);
    else
      iter_swap(even.end() - 1, even.end() - 2);
    for (int i = 0; i < odd.size(); i++) cout << odd[i] << " ";
    for (int i = even.size() - 1; i >= 0; i--) cout << even[i] << " ";
    cout << '\n';
  }
  return 0;
}
```
[Submission](https://codeforces.com/contest/1352/submission/87278856)